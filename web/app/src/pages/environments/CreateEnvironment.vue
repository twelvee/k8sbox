<template>
  <default-layout>
    <div class="flex w-full">
      <div class="flex flex-col w-full">
        <div class="flex justify-between items-center border-y-[1px] border-white/5">
          <div class="text-xl pl-10 py-4">
            Run new environment
          </div>
        </div>
        <div class="pl-4 mt-4 h-screen overflow-y-auto">
          <div class="h-auto pb-[200px] !overflow-y-auto">
            <div class="w-full p-6 pt-2">
              <div>
                <div>
                  <select v-model="cluster.name" required
                          class="bg-midnight border-[1px] border-white/10 px-4 py-2 rounded-md w-full max-w-xs">
                    <option value="" selected disabled>Select cluster</option>
                    <option v-for="[k, cl] in getClusters" :value="k">{{ k }}</option>
                  </select>
                </div>
                <div class="mt-2">
                  <input type="text" placeholder="Name"
                         v-model="name"
                         required
                         class="bg-midnight border-[1px] border-white/10 px-4 py-2 rounded-md w-full max-w-xs"/>
                </div>
                <div class="mt-2">
                  <input type="text" placeholder="Namespace (optional)"
                         v-model="namespace"
                         required
                         class="bg-midnight border-[1px] border-white/10 px-4 py-2 rounded-md w-full max-w-xs"/>
                </div>

                <div class="mt-2">
                  <div class="text-lg">Select boxes</div>
                  <div class="form-control grid grid-cols-5 gap-4 mt-2">
                    <div v-for="[bk, box] in getBoxes">
                      <label v-if="boxes.includes(bk)"
                             class="flex justify-center truncate label cursor-pointer p-4 border-[1px] border-primary/70 rounded-md"
                             @click="selectBox(bk)">
                        {{ bk }}
                      </label>
                      <label v-else
                             class="flex justify-center truncate label cursor-pointer p-4 border-[1px] border-white/10 hover:border-primary/40 rounded-md"
                             @click="selectBox(bk)">
                        {{ bk }}
                      </label>
                    </div>
                  </div>
                </div>

                <div class="mt-4">
                  <Codemirror
                      v-model:value="variables"
                      :options="cmOptions"
                      placeholder="dot-env style environment variables"
                      class="rounded-md w-full border-[1px] border-white/10 !text-[16px] !not-italic !font-sans"
                      :height="400"
                  />
                </div>

                <div class="mt-4">
                  <button class="btn btn-primary btn-sm" @click="runEnvironment">Run environment</button>
                </div>
                <div v-if="errors" class="text-error mt-4 max-w-[800px]">
                  {{ errors }}
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </default-layout>
</template>
<script setup lang="ts">
import DefaultLayout from "../../components/layouts/DefaultLayout.vue";
</script>
<script lang="ts">
// placeholder
import "codemirror/addon/display/placeholder.js";

// language
import "codemirror/mode/yaml/yaml.js";
// placeholder
import "codemirror/addon/display/placeholder.js";
// theme
import "codemirror/theme/material-palenight.css";
import {mapGetters} from "vuex";

export default {
  data() {
    return {
      name: '',
      cluster: {
        name: ''
      },
      namespace: '',
      boxes: [],
      variables: '',
      errors: '',
      cmOptions: {
        mode: "text/yaml", // Language mode
        theme: "material-palenight", // Theme
      },
    }
  },
  async mounted() {
    if (this.getClusters.size === 0) {
      await this.$store.dispatch('getClusters')
    }
    if (this.getBoxes.size === 0) {
      await this.$store.dispatch('getBoxes')
    }
  },
  computed: {
    ...mapGetters(['getClusters', 'getBoxes']),
  },
  methods: {
    findVariables() {
      const variablesRegExp = new RegExp(/\$\{[A-Z0-9_][^ :\#\<%\/\$\{\}]+\}|\$[A-Z0-9_][^ :\#\<%\/\$\{\}]+\b/gm)
      let boxVariables = []
      this.boxes.forEach((selectedBox) => {
        if (!this.getBoxes.has(selectedBox)) {
          return
        }
        const b = this.getBoxes.get(selectedBox)
        const chartVariables = b.Chart.match(variablesRegExp)
        const valuesVariables = b.Values.match(variablesRegExp)
        let appsVariables = []
        for (let i = 0; i < b.Applications.length; i++) {
          const appVars = b.Applications[i].Chart.match(variablesRegExp)
          if (appVars != null) {
            appsVariables.push(appVars.values())
          }
        }
        if (chartVariables) {
          boxVariables.push(...chartVariables)
        }
        if (valuesVariables) {
          boxVariables.push(...valuesVariables)
        }
        if (appsVariables && appsVariables.length > 0) {
          boxVariables.push(...appsVariables)
        }
      })
      for (let i = 0; i < boxVariables.length; i++) {
        boxVariables[i] = boxVariables[i].replace('{', '').replace('}', '').replace('$', '')
      }
      let variablesArray = [...new Set(boxVariables)]
      this.variables = ''
      for (let i = 0; i < variablesArray.length; i++) {
        this.variables += variablesArray[i] + '=\r\n'
      }

    },
    selectBox(key) {
      if (this.boxes.includes(key)) {
        const index = this.boxes.indexOf(key);
        this.boxes.splice(index, 1)
      } else {
        this.boxes.push(key)
      }
      this.findVariables()
    },
    validateName(name) {
      if (name.length > 253 || name.length <= 0) {
        return false
      }
      let reg = new RegExp(/[a-z][a-z0-9-.]{0,253}[a-z]$/, 'gm')
      return reg.test(name)
    },
    async runEnvironment() {
      this.errors = ''
      if (!this.validateName(this.name)) {
        this.errors = `
        Name fields must
        contain no more than 253 characters,
        contain only lowercase alphanumeric characters, '-' or '.',
        start with an alphanumeric character,
        end with an alphanumeric character
        `
        return
      }
      if (this.namespace.length > 0 && !this.validateName(this.namespace)) {
        this.errors = `
        Namespace fields must
        contain no more than 253 characters,
        contain only lowercase alphanumeric characters, '-' or '.',
        start with an alphanumeric character,
        end with an alphanumeric character
        `
        return
      }
      if (this.boxes.length === 0) {
        this.errors = `At least one box should be set.`
        return
      }
      if (this.cluster.name === '') {
        this.errors = 'Cluster field is required.'
      }
      try {
        let variables = '{'
        const variableLines = this.variables.split('\n')
        for (let i = 0; i< variableLines.length; i++) {
          if (variableLines[i].length == 0) {
            continue;
          }
          if (variableLines[i].match(/(^[A-Z0-9_]+)(\=)(.*\n(?=[A-Z])|.*$)/gm)) {
            const prepareMap = variableLines[i].split('=')
            variables += '"'+prepareMap[0] + '": "' + prepareMap[1] + '",'
          }
        }
        if (variables.endsWith(',')) {
          variables = variables.slice(0,-1)
        }
        variables+='}'

        let boxesWithApps = [];
        for(let i = 0; i<this.boxes.length; i++) {
          boxesWithApps.push(this.getBoxes.get(this.boxes[i]))
        }
        const response = await this.$store.dispatch(
            'createEnvironment',
            {
              Name: this.name,
              Namespace: this.namespace,
              ClusterName: this.cluster.name,
              Boxes: boxesWithApps,
              VariablesMap: JSON.parse(variables),
            }
        )

        if (response?.error) {
          this.errors = response.error
          return
        }

        location.href="/environments/"+this.name

      } catch (e) {
        this.errors = e
      }
    }
  }
}
</script>