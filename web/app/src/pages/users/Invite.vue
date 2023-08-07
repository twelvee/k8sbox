<template>
  <default-layout>
    <div class="flex w-full h-full">
      <div class="flex flex-col w-full">
        <div class="flex justify-between items-center border-y-[1px] border-white/5">
          <div class="text-xl pl-10 py-4">
            Invite new user
          </div>
        </div>
        <div class="pl-4 mt-4">
          <div class="overflow-y-auto h-screen">
            <div class="w-full p-6 pt-2">
              <div class="alert alert-midnight mb-4" v-if="code">
                <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24"
                     class="stroke-info shrink-0 w-6 h-6">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                        d="M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"></path>
                </svg>
                <div class="flex flex-col">
                  <h3 class="font-semibold">New code just generated!</h3>
                  <div>
                    <span class="badge">{{ code }}</span>
                  </div>
                  <div class="text-xs"> Once the page is refreshed, the code cannot
                    be retrieved again.
                  </div>
                </div>
                <button class="btn btn-sm btn-primary" @click="copyCode">Copy</button>
              </div>
              <div>
                <div>
                  <input type="email" placeholder="Email"
                         v-model="email"
                         required
                         class="bg-midnight border-[1px] border-white/10 px-4 py-2 rounded-md w-full max-w-xs"/>
                </div>
                <div class="mt-2">
                  <input type="text" placeholder="Name"
                         v-model="name"
                         required
                         class="bg-midnight border-[1px] border-white/10 px-4 py-2 rounded-md w-full max-w-xs"/>
                </div>
                <div class="mt-4">
                  <button class="btn btn-primary btn-sm" @click="createInvite">Create invite</button>
                </div>
                <div v-if="errors" class="alert alert-error mt-4 max-w-[400px]">
                  <div class="flex flex-col">
                    <div class="font-semibold">Failed</div>
                    <div class="ml-2">{{ errors }}</div>
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
export default {
  data() {
    return {
      name: '',
      email: '',
      errors: '',
      code: ''
    }
  },
  methods: {
    async copyCode() {
      try {
        await navigator.clipboard.writeText(this.code);
        this.code = 'COPIED'
      } catch($e) {
        this.errors = 'Cannot copy the code.'
      }
    },
    validateEmail() {
      return String(this.email)
          .toLowerCase()
          .match(
              /^(([^<>()[\]\\.,;:\s@\"]+(\.[^<>()[\]\\.,;:\s@\"]+)*)|(\".+\"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/
          )
    },
    async createInvite() {
      if (!this.validateEmail()) {
        this.errors = 'Email field should be actual email address.'
        return
      }
      try {
        const response = await this.$store.dispatch('createInvite', {email: this.email, name: this.name})
        if (response.invite_code.length > 0) {
          this.code = response.invite_code
          this.name = ''
          this.email = ''
          this.errors = ''
          return
        }

        if (response?.error) {
          this.errors = response.error
          return
        }

      } catch (e) {
        this.errors = e
      }
    }
  }
}
</script>