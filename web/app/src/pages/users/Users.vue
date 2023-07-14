<template>
  <default-layout>
    <div class="flex w-full h-full">
      <div class="flex flex-col w-full">
        <div class="flex justify-between items-center border-y-[1px] border-white/5">
          <div class="text-xl pl-10 py-4">
            Users
          </div>
          <div>
            <router-link to="/users/invite" class="btn btn-sm mr-2">Invite new user</router-link>
          </div>
        </div>
        <div class="pl-4">
          <div class="overflow-y-auto h-screen">
            <div class="w-full">
              <div class="flex justify-between items-center p-4 border-b-[1px] border-white/5"  v-for="[key, user] in getUsers">
                <div class="w-full flex space-x-4">
                  <div class="avatar placeholder">
                    <div class="bg-neutral-focus text-neutral-content rounded-full w-12">
                      <span class="text-lg">{{ user.AvatarInitials }}</span>
                    </div>
                  </div>
                  <div>
                    <div class="font-bold">{{ user.Name }}</div>
                    <div class="text-sm opacity-50">{{ user.Email }}</div>
                  </div>
                </div>
                <div class="">
                  <router-link :to="'/users/'+user.ID" class="btn btn-ghost btn-xs">details</router-link>
                </div>
              </div>

            </div>
          </div>
        </div>
      </div>
      <div class="flex flex-col bg-midnight w-2/3">
        <div class="flex justify-between items-center border-y-[1px] border-l-[1px] border-white/5">
          <div class="text-lg p-4">
            Activity feed
          </div>
          <div>
            <button class="btn btn-sm mr-2 btn-ghost text-white/50 hover:text-white">View all</button>
          </div>
        </div>
        <div class="overflow-y-auto">

          <div class="p-4 border-b-[1px] border-l-[1px] border-white/5">
            <!--<div class="flex justify-between">
              <div>
                <div class="avatar placeholder">
                  <div class="bg-neutral-focus text-neutral-content rounded-full w-8 h-8 mr-1">
                    <span class="text-sm">SZ</span>
                  </div>
                </div>
                Sergey Zyryanov
              </div>
              <div class="text-sm font-thin text-white/50"> 1h</div>
            </div>
            <div class="text-white/50 pt-1 font-thin">Created <span
                class="font-normal underline"> environment #1123</span> on cluster <span class="font-normal underline">Minikube</span>
            </div>-->
            Work in progress..
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
import {mapGetters} from "vuex";

export default {
  async mounted() {
    if (this.getUsers.size === 0) {
      await this.$store.dispatch('getUsers')
    }
  },
  computed: {
    ...mapGetters(['getUsers'])
  }
}
</script>