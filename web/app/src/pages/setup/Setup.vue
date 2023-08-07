<template>
  <div class="hero min-h-screen bg-midnight">
    <div class="hero-content flex-col lg:flex-row-reverse">
      <div class="card flex-shrink-0 w-full max-w-sm shadow-2xl bg-base-100">
        <div class="card-header text-center">
          <h1 class="h1 text-2xl pt-4">Run boxie, run!</h1>
          <div class="text-sm">
            The last step before you start is to create your first account
          </div>
        </div>
        <div class="card-body w-[370px]">
          <div class="form-control">
            <label class="label">
              <span class="label-text">Email</span>
            </label>
            <input type="email" v-model="email" placeholder="" class="input input-bordered"/>
          </div>
          <div class="form-control">
            <label class="label">
              <span class="label-text">Full name</span>
            </label>
            <input type="text" v-model="name" placeholder="" class="input input-bordered"/>
          </div>
          <div v-if="errors" class="text-error">
            {{errors}}
          </div>
          <div class="form-control mt-6">
            <button class="btn btn-primary" @click="createFirstAccount">Create an account</button>
          </div>
        </div>
      </div>
    </div>
  </div>

</template>
<script lang="ts">
export default {
  data() {
    return {
      email: '',
      name: '',
      code: '',
      errors: ''
    }
  },
  methods: {
    validateEmail() {
      return String(this.email)
          .toLowerCase()
          .match(
              /^(([^<>()[\]\\.,;:\s@\"]+(\.[^<>()[\]\\.,;:\s@\"]+)*)|(\".+\"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/
          )
    },
    async createFirstAccount() {
      try {
        if (!this.validateEmail()) {
          this.errors = 'Email field should be actual email address.'
          return
        }

        const data = await this.$store.dispatch('createFirstInvite', {
          email: this.email,
          name: this.name,
        })
        if (!data) {
          this.errors = 'Failed by unknown reason'
          return
        }
        if (data.invite_code.length > 0) {
          this.code = data.invite_code
          this.name = ''
          this.email = ''
          this.errors = ''
          location.href='/login/redeem?code='+this.code
          return
        }
      } catch (e) {
        this.errors = e
      }
    }
  }
}
</script>