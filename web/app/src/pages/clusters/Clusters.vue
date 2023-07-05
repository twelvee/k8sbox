<template>
  <default-layout>
    <div class="flex w-full h-full">
      <div class="flex flex-col w-full">
        <div class="flex justify-between items-center border-y-[1px] border-white/5">
          <div class="text-xl pl-10 py-4">
            Clusters
          </div>
          <div>
            <router-link to="/clusters/connect" class="btn btn-sm mr-2">Connect new cluster</router-link>
          </div>
        </div>
        <div class="pl-4 mt-4">
          <div class="overflow-y-auto h-screen">
            <div class="w-full grid grid-cols-4 gap-6" v-if="getClusters">

              <div v-for="[key, cluster] in getClusters">
                <router-link :to="'/clusters/'+cluster.Name" class="block rounded-lg border-[1px] border-white/10 p-4 hover:border-primary/50 cursor-pointer">
                  <div class="flex items-center">
                    <div class="m-0 p-0">
                      <h3 class="text-lg font-medium text-white">{{ cluster.Name }}</h3>
                      <div class="text-xs">Connected: {{ format(new Date(cluster.CreatedAt * 1000), 'yyyy-MM-dd') }}</div>
                    </div>
                  </div>

                  <ul class="mt-4 space-y-2">
                    <li>
                      <a
                          href="#"
                          class="block h-full text-white/70 hover:text-white rounded-lg border border-white/5 p-4 hover:border-primary/20"
                      >
                        <strong class="font-medium">Last deployment</strong>

                        <p class="mt-1 text-xs font-medium">
                          Unknown
                        </p>
                      </a>
                    </li>
                  </ul>
                </router-link>
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
import {format} from "date-fns";
</script>
<script lang="ts">
import {mapGetters} from "vuex";

export default {
  computed: {
    ...mapGetters(['getClusters'])
  },
  async mounted() {
    if (this.getClusters.size == 0) {
      await this.$store.dispatch('getClusters')
    }
  }
}
</script>