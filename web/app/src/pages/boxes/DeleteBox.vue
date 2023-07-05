<template>
  <default-layout>
    <div class="flex w-full">
      <div class="flex flex-col w-full">
        <div class="flex justify-between items-center border-y-[1px] border-white/5">
          <div class="text-xl pl-10 py-4">
            Delete Box
          </div>

        </div>
        <div class="pl-4 mt-4" v-if="name.length > 0">
          <div class="h-screen overflow-y-auto">
            <div class="w-full p-6 pt-2">
              <div>
                <div>
                  <input type="text" placeholder="Box name"
                         v-model="name"
                         disabled
                         required
                         class="bg-midnight border-[1px] border-white/10 px-4 py-2 rounded-md w-full max-w-xs"/>
                </div>
                <div class="mt-4">
                  <div>
                    Are you sure that you want to delete this box completely from boxie?
                  </div>
                  <button class="btn btn-error btn-sm" @click="deleteBox">Yes, delete it</button>
                  <router-link :to="'/boxes/'+name+'/edit'" class="btn btn-primary btn-sm ml-2">Nevermind</router-link>
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
      errors: '',
      cmOptions: {
        readOnly: true,
        mode: "text/yaml", // Language mode
        theme: "material-palenight", // Theme
      },
    }
  },
  async mounted() {
    let box = this.$store.getters.getBoxByName(this.$route.params.name)
    if (!box) {
       await this.$store.dispatch('getBoxes')
    }
    box = this.$store.getters.getBoxByName(this.$route.params.name)
    if (!box) {
      location.href="/boxes"
      return
    }
    this.name = box.Name
    this.errors = ''
  },
  methods: {
    async deleteBox() {
      try {
        const response = await this.$store.dispatch('deleteBox', {Name: this.name})

        if (response?.error) {
          this.errors = response.error
          return
        }
        location.href="/boxes"

      } catch (e) {
        this.errors = e
      }
    },
  }
}
</script>