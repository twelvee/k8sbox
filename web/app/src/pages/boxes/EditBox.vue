<template>
  <default-layout>
    <div class="flex w-full">
      <div class="flex flex-col w-full">
        <div class="flex justify-between items-center border-y-[1px] border-white/5">
          <div class="text-xl pl-10 py-4">
            Edit box
          </div>
          <div v-if="box">
            <router-link :to="'/boxes/'+name+'/delete'" class="btn btn-sm mr-2 btn-error">delete</router-link>
          </div>
        </div>
        <div class="pl-4 mt-4 h-screen overflow-y-auto" v-if="box">
          <div class="h-auto pb-[200px] !overflow-y-auto">
            <div class="w-full p-6 pt-2">
              <div>
                <div>
                  <select v-model="type" required class="bg-midnight border-[1px] border-white/10 px-4 py-2 rounded-md w-full max-w-xs">
                    <option value="helm" selected>Helm</option>
                  </select>
                </div>
                <div class="mt-2">
                  <input type="text" placeholder="Name"
                         disabled
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
                  <Codemirror
                      v-model:value="chart"
                      :options="cmOptions"
                      placeholder="helm chart"
                      class="rounded-md w-full border-[1px] border-white/10 !text-[16px] !not-italic !font-sans"
                      :height="400"
                  />
                </div>
                <div class="mt-2">
                  <Codemirror
                      v-model:value="values"
                      :options="cmOptions"
                      placeholder="helm values"
                      class="rounded-md w-full border-[1px] border-white/10 !text-[16px] !not-italic !font-sans"
                      :height="400"
                  />
                </div>
                <div class="mt-2 text-lg">
                  Applications
                </div>

                <div class="border-b-[1px] border-white/20" v-for="(app, i) in applications">
                  <div class="mt-2">
                    <button class="link link-error" @click="deleteApplication(i)">Delete application</button>
                  </div>
                  <div class="mt-2">
                    <input type="text" placeholder="Name"
                           v-model="app.Name"
                           required
                           class="bg-midnight border-[1px] border-white/10 px-4 py-2 rounded-md w-full max-w-xs"/>
                  </div>
                  <div class="mt-2">
                    <Codemirror
                        v-model:value="app.Chart"
                        :options="cmOptions"
                        placeholder="application chart"
                        class="rounded-md w-full border-[1px] border-white/10 !text-[16px] !not-italic !font-sans"
                        :height="400"
                    />
                  </div>
                </div>
                <div class="mt-2">
                  <button class="btn btn-secondary" @click="addApplication">Add application</button>
                </div>

                <div class="mt-4">
                  <button class="btn btn-primary btn-sm" @click="updateBox">Update box</button>
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
      box: {},
      name: this.box?.Name,
      type: this.box?.Type,
      chart: this.box?.Chart,
      values: this.box?.Values,
      namespace: this.box?.Namespace,
      applications: this.box?.Applications,
      errors: '',
      cmOptions: {
        mode: "text/yaml", // Language mode
        theme: "material-palenight", // Theme
      },
    }
  },
  async mounted() {
    this.box = await this.$store.getters.getBoxByName(this.$route.params.name)
    if (!this.box) {
      await this.$store.dispatch('getBoxes')
    }
    this.box = await this.$store.getters.getBoxByName(this.$route.params.name)
    console.log(this.box)
    this.name = this.box.Name
    this.namespace = this.box.Namespace
    this.chart = this.box.Chart
    this.values = this.box.Values
    this.type = this.box.Type
    this.applications = this.box.Applications
  },
  methods: {
    validateName(name) {
      if (name.length > 253 || name.length <= 0) {
        return false
      }
      let reg = new RegExp(/[a-z][a-z0-9-.]{0,253}[a-z]$/, 'gm')
      return reg.test(name)
    },
    addApplication() {
      this.applications.push({
        Name: '',
        Chart: '',
      })
    },
    deleteApplication(i) {
      if (this.applications.length == 1) {
        this.applications = [];
        return
      }
      this.applications = this.applications.splice(i-1, 1);
    },
    async updateBox() {
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
      if (this.applications.length === 0) {
        this.errors = `At least one application should be set.`
        return
      }

      if (this.chart.length === 0) {
        this.errors = `Box chart is required.`
        return
      }

      for (let i = 0; i<this.applications.length; i++) {
        if (!this.applications[i].Chart || this.applications[i].Chart.length === 0) {
          this.errors = 'Application chart is required.'
          return
        }
        if (!this.applications[i].Name || this.applications[i].Name.length === 0) {
          this.errors = 'Application name is required.'
          return
        }
      }
      try {
        const response = await this.$store.dispatch(
            'updateBox',
            {
              Name: this.name,
              Namespace: this.namespace,
              Type: this.type,
              Chart: this.chart,
              Values: this.values,
              Applications: this.applications
            }
        )

        if (response?.error) {
          this.errors = response.error
          return
        }
        location.href="/boxes"

      } catch (e) {
        this.errors = e
      }
    }
  }
}
</script>