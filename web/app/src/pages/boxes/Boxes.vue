<template>
  <default-layout>
    <div class="flex w-full h-full">
      <div class="flex flex-col w-full">
        <div class="flex justify-between items-center border-y-[1px] border-white/5">
          <div class="text-xl pl-10 py-4">
            Boxes
          </div>
          <div>
            <a href="#" class="underline mr-4">How to work with boxes</a>
            <router-link to="/boxes/create" class="btn btn-sm mr-2">Create new box</router-link>
          </div>
        </div>
        <div class="pl-4 mt-4 mr-2">
          <div class="overflow-y-auto h-screen">
            <div class="w-full grid grid-cols-3 gap-6" v-if="getBoxes">
              <div v-for="[key, box] in getBoxes">
                <article class="rounded-lg border-[1px] border-white/10 p-4">
                  <div class="flex items-center justify-between mb-2">
                    <h3 class="text-lg font-medium text-white">{{ box.Name }}</h3>
                    <div class="dropdown dropdown-end">
                      <label tabindex="0" class="text-lg hover:cursor-pointer">
                        <EllipsisVerticalIcon class="h-5 w-5" />
                      </label>
                      <div tabindex="0" class="dropdown-content z-[1] p-2 bg-midnight rounded-md w-40">
                        <router-link :to="'/boxes/'+box.Name+'/edit'" class="hover:cursor-pointer block p-2 hover:bg-white/10 w-full rounded-md text-md flex items-center">
                          <div><Cog6ToothIcon class="h-4 w-4 mr-2" /></div>
                          <div>Manage</div>
                        </router-link>
                      </div>
                    </div>
                  </div>
                  <div class="items-start">
                    <div class="badge" v-for="app in box.Applications">{{ app.Name }}</div>
                  </div>
                </article>
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
import {EllipsisVerticalIcon, Cog6ToothIcon, EyeIcon} from '@heroicons/vue/24/solid'

</script>

<script lang="ts">
import {mapGetters} from "vuex";

export default {
  async mounted() {
    if (this.getBoxes.size === 0) {
      await this.$store.dispatch('getBoxes')
    }
  },
  computed: {
    ...mapGetters(['getBoxes'])
  }
}
</script>