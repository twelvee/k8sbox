<template>
  <default-layout>
    <div class="flex w-full">
      <div class="flex flex-col w-full">
        <div class="flex justify-between items-center border-y-[1px] border-white/5">
          <div class="text-xl pl-10 py-4">
            Delete environment
          </div>
        </div>
        <div class="pl-4 mt-4 h-screen overflow-y-auto">
          <div class="pl-4 mt-4" v-if="environment">
            <div class="h-screen overflow-y-auto">
              <div class="w-full p-6 pt-2">
                <div>
                  <div>
                    <input type="text" placeholder="Environment name"
                           v-model="environment.Name"
                           disabled
                           required
                           class="bg-midnight border-[1px] border-white/10 px-4 py-2 rounded-md w-full max-w-xs"/>
                  </div>
                  <div class="mt-4">
                    <div>
                      Are you sure that you want to delete this environment completely?
                    </div>
                    <button class="btn btn-error btn-sm" @click="deleteEnvironment">Yes, delete it</button>
                    <router-link :to="'/environments/'+environment.Name" class="btn btn-primary btn-sm ml-2">Nevermind</router-link>
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
    </div>
  </default-layout>
</template>
<script setup lang="ts">
import DefaultLayout from "../../components/layouts/DefaultLayout.vue";
</script>
<script lang="ts">
import {mapGetters} from "vuex";

export default {
  data() {
    return {
      environment: null,
      errors: '',
    }
  },
  async mounted() {
    if (this.getEnvironments.size === 0) {
      await this.$store.dispatch('getEnvironments')
    }
    this.environment = this.$store.getters.getEnvironmentByName(this.$route.params.name)
  },
  computed: {
    ...mapGetters(['getEnvironments']),
  },
  methods: {

    async deleteEnvironment() {

      try {

        const response = await this.$store.dispatch(
            'deleteEnvironment',
            {
              Name: this.environment.Name,
            }
        )

        if (response?.error) {
          this.errors = response.error
          return
        }

        location.href="/environments"

      } catch (e) {
        this.errors = e
      }
    }
  }
}
</script>