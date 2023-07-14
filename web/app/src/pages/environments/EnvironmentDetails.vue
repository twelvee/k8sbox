<template>
  <default-layout>
    <div class="flex w-full">
      <div class="flex flex-col w-full">
        <div class="flex justify-between items-center border-y-[1px] border-white/5">
          <div class="text-xl pl-10 py-4">
            Environment details
          </div>
        </div>
        <div class="pl-4 mt-4 h-screen overflow-y-auto" v-if="environment">
          <div class="rounded-lg border-[1px] p-4 border-primary/30 mr-2">
            <div class="flex justify-between">
              <div class="flex flex-col relative overflow-x-hidden">
                <div class="text-xl">{{ environment.Name }}</div>
                <div class="text-sm text-white/50"> {{ environment.Namespace }}</div>
                <div class="!overflow-x-scroll flex max-w-full justify-start gap-2 mt-2 mr-36">
                  <div class="min-w-[150px] px-4 py-2 rounded-md border-[1px] border-white/20"
                       v-for="box in environment.Boxes.slice(0,5)">{{ box }}
                  </div>
                </div>
                <div class="text-sm text-white/30" v-if="environment.Boxes.length > 5"> and
                  {{ environment.Boxes.length - 5 }} more boxes
                </div>
              </div>
              <div class="w-[20%] flex flex-col justify-start text-end">
                <router-link :to="'/users/'+environment.UserID" class="flex justify-end gap-2 items-center">
                  <div class="avatar placeholder">
                    <div class="bg-neutral-focus text-neutral-content rounded-full w-12 h-12">
                      <span class="text-lg">{{ environment.User.AvatarInitials }}</span>
                    </div>
                  </div>
                  <div>
                    <div class="font-bold truncate">{{ environment.User.Name }}</div>
                  </div>
                </router-link>
                <div class="text-white/30 -mt-2">
                  Created {{ environment.CreatedFromNow }} ago
                </div>
                <div class="text-white/30 -mt-2">
                  on
                  <router-link class="underline hover:text-primary" :to="'/clusters/'+environment.ClusterName">
                    {{ environment.ClusterName }}
                  </router-link>
                </div>
              </div>

            </div>
            <div class="flex justify-between mt-4 border-t-[1px] border-white/10 pt-4 items-center">
              <div class="text-white/20">
                <span class="badge">{{ environment.StatusText }}</span>
              </div>
              <div class="flex gap-2">
                <button class="btn btn-secondary btn-disabled btn-sm">Re-deploy</button>
                <button class="btn btn-secondary btn-disabled btn-sm" disabled>Logs</button>
                <router-link :to="'/environments/'+environment.Name+'/delete'" class="btn btn-error btn-sm">Delete</router-link>
              </div>
            </div>
          </div>
          <div class="flex flex-col mt-4 min-h-screen pb-[200px]">
            <h3 class="text-xl">Applications</h3>
            <div class="flex flex-col" v-for="app in environment.EnvironmentApplications">
              <h3 class="text-lg">{{ app.Name }}</h3>
              <div class="tabs">
                <a class="tab tab-active" v-if="tabs[app.Name].description">Details</a>
                <a class="tab" @click="tabDescription(app.Name)" v-else>Details</a>
                <a class="tab tab-active" v-if="tabs[app.Name].chart">Chart</a>
                <a class="tab" @click="tabChart(app.Name)" v-else>Chart</a>
                <a class="tab tab-active" v-if="tabs[app.Name].runtime">Runtime data</a>
                <a class="tab" @click="tabRuntime(app.Name)" v-else>Runtime data</a>
              </div>
              <div class="text-white/70" v-show="tabs[app.Name].description">
                <div class="flex flex-col p-4 border-[1px] border-white/10 rounded-md">
                  <div>
                    Kind: <span class="font-semibold">{{ app.RuntimeData.Kind }}</span>, Name: <span
                      class="font-semibold">{{ app.RuntimeData.Name }}</span>
                  </div>
                  <div v-if="app.RuntimeData.Kind === 'Ingress'">
                    <div class="text-md">Host: <span
                        class="font-semibold">{{ app.RuntimeData.Data[0].spec.rules[0].host }}</span></div>
                  </div>
                  <div v-else-if="app.RuntimeData.Kind === 'Deployment'">
                    <div class="text-md font-semibold">Containers</div>
                    <div class="flex gap-4 text-md"
                         v-for="container in app.RuntimeData.Data[0].spec.template.spec.containers">
                      <div class=" p-2 border-[1px] border-white/10 rounded-md">{{ container.image }}</div>
                    </div>
                    <div class="text-md font-semibold">Conditions</div>
                    <div class="text-md" v-for="condition in app.RuntimeData.Data[0].status.conditions">
                      {{ condition.lastUpdateTime }} | {{ condition.reason }}: {{ condition.message }}
                    </div>
                  </div>
                  <div v-else-if="app.RuntimeData.Kind === 'Service'">
                    <div class="text-md">Cluster IP: <span
                        class="font-semibold">{{ app.RuntimeData.Data[0].spec.clusterIP }}</span></div>
                    <div class="text-md font-semibold">Ports</div>
                    <div class="text-md" v-for="port in app.RuntimeData.Data[0].spec.ports">
                      Name: {{ port.name }}, Protocol: {{ port.protocol }}, Port: {{ port.port }}
                    </div>
                  </div>
                  <div v-else>
                    <div class="text-sm text-white/20">Short description for this kind of application is not developed
                      yet.. To be done.
                    </div>
                  </div>
                </div>
              </div>
              <div class="text-white/70" v-show="tabs[app.Name].chart">
                <Codemirror
                    v-model:value="app.Chart"
                    :options="cmOptions"
                    placeholder="dot-env style environment variables"
                    class="rounded-md w-full border-[1px] border-white/10 !text-[16px] !not-italic !font-sans"
                    :height="400"
                />
              </div>
              <div class="text-white/70" v-show="tabs[app.Name].runtime">
                <Codemirror
                    v-model:value="app.runtimeYaml"
                    :options="cmOptions"
                    placeholder="dot-env style environment variables"
                    class="rounded-md w-full border-[1px] border-white/10 !text-[16px] !not-italic !font-sans"
                    :height="400"
                />
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
import {parse, stringify} from 'yaml'
// placeholder
import "codemirror/addon/display/placeholder.js";

// language
import "codemirror/mode/yaml/yaml.js";
// placeholder
import "codemirror/addon/display/placeholder.js";
// theme
import "codemirror/theme/material-palenight.css";
import {mapGetters} from "vuex";
import environments from "./Environments.vue";
import {differenceInMinutes} from "date-fns";

export default {
  data() {
    return {
      environmentData: null,
      errors: '',
      tabs: [],
      cmOptions: {
        mode: "text/yaml", // Language mode
        theme: "material-palenight", // Theme
        readOnly: true,
      },
    }
  },
  async mounted() {
    if (this.getEnvironments.size === 0) {
      await this.$store.dispatch('getEnvironments')
    }
    if (this.getUsers.size === 0) {
      await this.$store.dispatch('getUsers')
    }
    if (!this.environmentData) {
      this.environmentData = this.$store.getters.getEnvironmentByName(this.$route.params.name)
    }
  },
  computed: {
    ...mapGetters(['getEnvironments', 'getUsers']),
    environment() {
      if (!this.environmentData) {
        return null
      }
      let env = this.environmentData
      if (env.Status == 0) {
        env.StatusText = 'Pending'
      } else if (env.Status == 1) {
        env.StatusText = 'Installing'
      } else if (env.Status == 2) {
        env.StatusText = 'Running'
      } else if (env.Status == 3) {
        env.StatusText = 'Failed'
      } else {
        env.StatusText = 'Deleted'
      }
      if (!env.Boxes) {
        env.Boxes = []
        for (let i = 0; i < this.environmentData.EnvironmentApplications.length; i++) {
          this.tabs[this.environmentData.EnvironmentApplications[i].Name] = {
            chart: false,
            description: true,
            runtime: false
          }
          env.EnvironmentApplications[i].runtimeYaml = stringify(this.environmentData.EnvironmentApplications[i].RuntimeData.Data)

          if (env.Boxes.length === 0) {

            env.Boxes.push(this.environmentData.EnvironmentApplications[i].BoxName)
            continue
          }
          if (!env.Boxes.includes(this.environmentData.EnvironmentApplications[i].BoxName)) {
            env.Boxes.push(this.environmentData.EnvironmentApplications[i].BoxName)
          }
        }
      }
      if (!env.User) {
        env.User = this.$store.getters.getUserById(env.UserID)
      }

      let diffInTime = differenceInMinutes(new Date(), new Date(env.CreatedAt * 1000))
      if (diffInTime <= 60) {
        env.CreatedFromNow = diffInTime + 'm'
      } else if (diffInTime > 60 && diffInTime <= 60 * 24) {
        env.CreatedFromNow = Math.round(diffInTime / 60) + 'h'
      } else {
        env.CreatedFromNow = Math.round(diffInTime / (60 * 24)) + 'd'
      }
      return env
    }
  },
  methods: {
    tabChart(name) {
      this.tabs[name].chart = true
      this.tabs[name].description = false
      this.tabs[name].runtime = false
    },
    tabDescription(name) {
      this.tabs[name].chart = false
      this.tabs[name].description = true
      this.tabs[name].runtime = false
    },
    tabRuntime(name) {
      this.tabs[name].chart = false
      this.tabs[name].description = false
      this.tabs[name].runtime = true
    }
  }
}
</script>