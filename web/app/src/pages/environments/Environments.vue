<template>
  <default-layout>
    <div class="flex w-full h-full">
      <div class="flex flex-col w-full">
        <div class="flex justify-between items-center border-y-[1px] border-white/5">
          <div class="text-xl pl-10 py-4">
            Environments {{environment}}
          </div>
          <div>
            <a class="underline mr-4">How to work with environments</a>
            <router-link to="/environments/create" class="btn btn-sm mr-2">Run</router-link>
          </div>
        </div>
        <div class="pl-4 mt-4 mr-2">
          <div class="overflow-y-auto h-screen">
            <div class="w-full grid grid-cols-2 gap-6" v-if="getEnvironments?.size > 0">
              <div v-for="environment in environments">
                <router-link :to="'/environments/'+environment.Name"
                    class="block rounded-lg border-[1px] border-white/10 p-4 hover:border-primary transition cursor-pointer">
                  <div class="flex items-center mb-4">
                    <span class="badge badge-success inline mr-2"></span>
                    <div class="m-0 p-0">
                      <h3 class="text-lg font-medium text-white">{{ environment.Name }}</h3>
                      <div class="text-sm">Running on {{ environment.ClusterName }}</div>
                    </div>
                  </div>
                  <div class="flex gap-2 items-center">
                    <div
                        href="#"
                        class="block text-white/70 hover:text-white rounded-lg border border-white/5 p-4 hover:border-primary/20"
                        v-for="app in environment.BoxesInUse.slice(0,3)"
                    >
                      <strong class="font-medium">{{ app.BoxName }}</strong>
                    </div>
                  </div>
                  <div class="text-white/50 pl-4 py-2" v-if="environment.BoxesInUse.length > 3"> ... and
                    {{ environment.BoxesInUse.length - 3 }} more boxes
                  </div>
                  <div class="text-white/70 pt-4">
                    <div class="avatar placeholder">
                      <div class="bg-neutral-focus text-neutral-content rounded-full w-8 mr-2">
                        <span class="text-sm">SZ</span>
                      </div>
                    </div>
                    <router-link :to="'/users/'+environment.UserID" class="underline hover:text-primary cursor-pointer">{{environment.User.Name}}</router-link>
                    {{ environment.CreatedFromNow }} ago on cluster
                    <router-link :to="'/clusters/'+environment.ClusterName" class="underline hover:text-primary cursor-pointer">{{environment.ClusterName }}</router-link>
                  </div>
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
</script>
<script lang="ts">
import {isProxy, toRaw} from 'vue';
import {mapGetters} from "vuex";
import {format, differenceInMinutes} from 'date-fns'

export default {
  async mounted() {
    if (this.getEnvironments.size === 0) {
      await this.$store.dispatch('getEnvironments')
    }

    if (this.getUsers.size === 0) {
      await this.$store.dispatch('getUsers')
    }
  },
  computed: {
    ...mapGetters(['getEnvironments', 'getUsers']),
    environments() {
      if (this.getEnvironments.size === 0) {
        return null
      }
      const self = this
      let inArray = []
      return Array.from(this.getEnvironments).map((env) => {
        if (isProxy(env[1])) {
          env[1] = toRaw(env[1])
        }
        let diffInTime = differenceInMinutes(new Date(), new Date(env[1].CreatedAt * 1000))
        if (diffInTime <= 60) {
          env[1].CreatedFromNow = diffInTime + 'm'
        } else if (diffInTime > 60 && diffInTime <= 60*24) {
          env[1].CreatedFromNow = Math.round(diffInTime/60) + 'h'
        } else {
          env[1].CreatedFromNow = Math.round(diffInTime/(60*24)) + 'd'
        }
        env[1].User = self.$store.getters.getUserById(env[1].UserID)
        env[1].BoxesInUse = env[1].EnvironmentApplications.filter((value, index, array) => {
          if (inArray.includes(value.BoxName)) {
            return null
          } else {
            inArray.push(value.BoxName)
          }
          return inArray.includes(value.BoxName)
        })
        return env[1]
      })
    }
  }
}
</script>