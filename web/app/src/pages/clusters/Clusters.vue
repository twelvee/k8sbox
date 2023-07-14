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

              <div v-for="cluster in clusters">
                <router-link :to="'/clusters/'+cluster.Name" class="block rounded-lg border-[1px] border-white/10 p-4 hover:border-primary/50 cursor-pointer">
                  <div class="flex items-center">
                    <div class="m-0 p-0">
                      <h3 class="text-lg font-medium text-white">{{ cluster.Name }}</h3>
                      <div class="text-xs">Connected: {{ format(new Date(cluster.CreatedAt * 1000), 'yyyy-MM-dd') }}</div>
                    </div>
                  </div>

                  <ul class="mt-4 space-y-2">
                    <li v-if="cluster.LastDeployment">
                      <router-link
                          :to="'/environments/'+cluster.LastDeployment.Name"
                          class="block h-full text-white/70 hover:text-white rounded-lg border border-white/5 p-4 hover:border-primary/20"
                      >
                        <strong class="font-medium">Last deployment</strong>

                        <p class="mt-1 text-xs font-medium" v-if="!cluster.LastDeployment">
                          Unknown
                        </p>
                        <p class="mt-1 text-xs font-medium" v-else>
                          {{cluster.LastDeploymentTime}} by {{cluster.LastDeployedUser.Name}}
                        </p>
                      </router-link>
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
import {isProxy, toRaw} from "vue";
import {differenceInMinutes} from "date-fns";

export default {
  computed: {
    ...mapGetters(['getClusters', 'getEnvironments', 'getUsers']),
    clusters() {
      if (this.getClusters.size === 0) {
        return null
      }
      const self = this
      return Array.from(this.getClusters).map((cl) => {
        if (isProxy(cl[1])) {
          cl[1] = toRaw(cl[1])
        }
        let lastEnvironment = null;
        this.getEnvironments.forEach((env) => {
          if (lastEnvironment == null &&  env.ClusterName == cl[1].Name) {
            lastEnvironment = env
            return
          }
          if (env.CreatedAt > lastEnvironment.CreatedAt && env.ClusterName == cl[1].Name) {
            lastEnvironment = env
          }
        })
        if (lastEnvironment != null) {
          let diffInTime = differenceInMinutes(new Date(), new Date(lastEnvironment.CreatedAt * 1000))
          if (diffInTime <= 60) {
            cl[1].LastDeploymentTime = diffInTime + 'm'
          } else if (diffInTime > 60 && diffInTime <= 60*24) {
            cl[1].LastDeploymentTime = Math.round(diffInTime/60) + 'h'
          } else {
            cl[1].LastDeploymentTime = Math.round(diffInTime/(60*24)) + 'd'
          }
          cl[1].LastDeployedUser = self.$store.getters.getUserById(lastEnvironment.UserID)
        }
        cl[1].LastDeployment = lastEnvironment

        return cl[1]
      })
    }
  },
  async mounted() {
    if (this.getClusters.size == 0) {
      await this.$store.dispatch('getClusters')
    }
    if (this.getEnvironments.size == 0) {
      await this.$store.dispatch('getEnvironments')
    }
    if (this.getUsers.size == 0) {
      await this.$store.dispatch('getUsers')
    }
  }
}
</script>