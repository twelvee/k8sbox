<template>
  <default-layout>
    <div class="flex w-full">
      <div class="flex flex-col w-full">
        <div class="flex justify-between items-center border-y-[1px] border-white/5">
          <div class="text-xl pl-10 py-4">
            Edit cluster
          </div>
          <div v-if="name.length > 0">
            <router-link :to="'/clusters/'+name+'/delete'" class="btn btn-sm mr-2 btn-error">delete</router-link>
          </div>
        </div>
        <div class="pl-4 mt-4" v-if="name.length > 0">
          <div class="h-screen overflow-y-auto">
            <div class="w-full p-6 pt-2">
              <div>
                <div>
                  <input type="text" placeholder="Cluster name"
                         v-model="name"
                         disabled
                         required
                         class="bg-midnight border-[1px] border-white/10 px-4 py-2 rounded-md w-full max-w-xs"/>
                </div>
                <div class="mt-2">
                  <Codemirror
                      v-model:value="kubeconfig"
                      :options="cmOptions"
                      placeholder="kubeconfig"
                      class="rounded-md w-full border-[1px] border-white/10 !text-[16px] !not-italic !font-sans"
                      :height="400"
                  />
                </div>
                <div class="mt-4">
                  <button class="mr-2 btn btn-secondary btn-sm" @click="testConnection">Test connection</button>
                  <button class="btn btn-primary btn-sm" @click="updateCluster">Update cluster</button>
                </div>
                <div v-if="available" class="text-success mt-4 max-w-[400px]">
                  Cluster is available!
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
export default {
  data() {
    return {
      name: '',
      kubeconfig: '',
      errors: '',
      available: false,
      cmOptions: {
        mode: "text/yaml", // Language mode
        theme: "material-palenight", // Theme
      },
    }
  },
  async mounted() {
    let cluster = this.$store.getters.getClusterByName(this.$route.params.name)
    if (!cluster) {
       await this.$store.dispatch('getClusters')
    }
    cluster = this.$store.getters.getClusterByName(this.$route.params.name)
    if (!cluster) {
      location.href="/clusters"
      return
    }
    this.name = cluster.Name
    this.kubeconfig = cluster.Kubeconfig
    this.errors = ''
    this.available = false

  },
  methods: {
    validateName() {
      if (this.name.length > 253 || this.name.length <= 0) {
        return false
      }
      let reg = new RegExp(/[a-z][a-z0-9-.]{0,253}[a-z]$/, 'gm')
      return reg.test(this.name)
    },
    async updateCluster() {
      if (!this.validateName()) {
        this.errors = `
        Name field must
        contain no more than 253 characters,
        contain only lowercase alphanumeric characters, '-' or '.',
        start with an alphanumeric character,
        end with an alphanumeric character
        `
        this.available = false
        return
      }
      try {
        const response = await this.$store.dispatch('updateCluster', {Name: this.name, Kubeconfig: this.kubeconfig})

        if (response?.error) {
          this.errors = response.error
          return
        }
        location.href="/clusters"

      } catch (e) {
        this.errors = e
      }
    },
    async testConnection() {
      try {
        const response = await this.$store.dispatch('testClusterConnection', {Name: this.name, Kubeconfig: this.kubeconfig})

        if (response?.error) {
          this.errors = response.error
          return
        }
        if (response.available) {
          this.errors = ''
          this.available = response.available
        }

      } catch (e) {
        this.errors = e
      }
    }
  }
}
</script>