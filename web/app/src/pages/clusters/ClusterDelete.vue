<template>
  <default-layout>
    <div class="flex w-full">
      <div class="flex flex-col w-full">
        <div class="flex justify-between items-center border-y-[1px] border-white/5">
          <div class="text-xl pl-10 py-4">
            Delete Cluster
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
                  <div>
                    Are you sure that you want to delete this kubernetes cluster completely from boxie? All current running environments will be lost.
                  </div>
                  <button class="btn btn-error btn-sm" @click="deleteCluster">Yes, delete it</button>
                  <router-link :to="'/clusters/'+name" class="btn btn-primary btn-sm ml-2">Nevermind</router-link>
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
      cmOptions: {
        readOnly: true,
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
    async deleteCluster() {
      try {
        const response = await this.$store.dispatch('deleteCluster', {Name: this.name})

        if (response?.error) {
          this.errors = response.error
          return
        }
        location.href="/clusters"

      } catch (e) {
        this.errors = e
      }
    },
  }
}
</script>