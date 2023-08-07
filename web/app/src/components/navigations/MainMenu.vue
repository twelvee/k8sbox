<template>
  <div class="text-moonlight h-screen flex flex-col justify-between">
    <div>
    <div class="p-4 flex">
      <span class="block text-2xl font-semibold uppercase">Boxie</span>
    </div>
    <div class="p-4">
      <nav class="grid grid-cols-1 gap-1">
        <router-link to="/" v-if="$route.name.startsWith('Dashboard')" class="p-2 flex bg-white/10 rounded-md text-white">
          <span>
            <ChartBarIcon class="h-4 w-4 mt-1"/>
          </span>
          <span class="ml-3">Dashboard</span>
        </router-link>
        <router-link to="/" v-else class="p-2 flex hover:bg-white/10 rounded-md hover:text-white">
          <span>
            <ChartBarIcon class="h-4 w-4 mt-1"/>
          </span>
          <span class="ml-3">Dashboard</span>
        </router-link>

        <router-link to="/users" v-if="$route.name.startsWith('Users')" class="p-2 flex bg-white/10 rounded-md text-white">
          <span>
            <UsersIcon class="h-4 w-4 mt-1"/>
          </span>
          <span class="ml-3">Users</span>
        </router-link>
        <router-link to="/users" v-else class="p-2 flex hover:bg-white/10 rounded-md hover:text-white">
          <span>
            <UsersIcon class="h-4 w-4 mt-1"/>
          </span>
          <span class="ml-3">Users</span>
        </router-link>

        <router-link to="/clusters" v-if="$route.name.startsWith('Clusters')"
                     class="p-2 flex bg-white/10 rounded-md text-white">
          <span>
            <ServerStackIcon class="h-4 w-4 mt-1"/>
          </span>
          <span class="ml-3">Clusters</span>
        </router-link>
        <router-link to="/clusters" v-else class="p-2 flex hover:bg-white/10 rounded-md hover:text-white">
          <span>
            <ServerStackIcon class="h-4 w-4 mt-1"/>
          </span>
          <span class="ml-3">Clusters</span>
        </router-link>

        <router-link to="/boxes" v-if="$route.name.startsWith('Boxes')" class="p-2 flex bg-white/10 rounded-md text-white">
          <span>
            <CubeIcon class="h-4 w-4 mt-1"/>
          </span>
          <span class="ml-3">Boxes</span>
        </router-link>
        <router-link to="/boxes" v-else class="p-2 flex hover:bg-white/10 rounded-md hover:text-white">
          <span>
            <CubeIcon class="h-4 w-4 mt-1"/>
          </span>
          <span class="ml-3">Boxes</span>
        </router-link>

        <router-link to="/environments" v-if="$route.name.startsWith('Environments')"
                     class="p-2 flex bg-white/10 rounded-md text-white">
          <span>
            <Squares2X2Icon class="h-4 w-4 mt-1"/>
          </span>
          <span class="ml-3">Environments</span>
        </router-link>
        <router-link to="/environments" v-else class="p-2 flex hover:bg-white/10 rounded-md hover:text-white">
          <span>
            <Squares2X2Icon class="h-4 w-4 mt-1"/>
          </span>
          <span class="ml-3">Environments</span>
        </router-link>

      </nav>
    </div>
    </div>
    <div class="p-4">
      <div class="flex items-center p-4 hover:border-white/20 rounded-md border-[1px] border-white/10">
        <div class="avatar placeholder mr-4">
          <div class="bg-neutral-focus text-neutral-content rounded-full w-10">
            <span class="text-lg">{{ getUser.AvatarInitials }}</span>
          </div>
        </div>
        <div>
          <div>
        {{ getUser.Name }}
          </div>
          <div class="text-xs">
            {{ getUser.Email }}
          </div>
        </div>
      </div>
      <div class="text-sm flex justify-between mt-2">
        <div class="text-white/30">Boxie v2.0.0</div>
        <router-link to="/logout" class="link">Logout</router-link>
      </div>
    </div>
  </div>
</template>
<script setup lang="ts">
import {ChartBarIcon, CubeIcon, ServerStackIcon, Squares2X2Icon, UsersIcon} from '@heroicons/vue/24/solid'
import {mapGetters} from 'vuex'
</script>

<script lang="ts">
import {mapGetters} from "vuex";

export default {
  computed: {
    ...mapGetters(['getUser'])
  }
}
</script>