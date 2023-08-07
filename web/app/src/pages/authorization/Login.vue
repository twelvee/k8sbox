<template>
  <div class="hero min-h-screen bg-midnight">
    <div class="hero-content flex-col lg:flex-row-reverse">
      <div class="card flex-shrink-0 w-full max-w-sm shadow-2xl bg-base-100">
        <div class="card-header text-center">
          <h1 class="h1 text-2xl pt-4">Welcome to Boxie!</h1>
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
              <span class="label-text">Password</span>
            </label>
            <input type="password" v-model="password" placeholder="" class="input input-bordered"/>
          </div>
          <div v-if="errors" class="text-error">
            {{errors}}
          </div>
          <div class="form-control mt-6">
            <button class="btn btn-primary" @click="login">Login</button>
            <router-link to="/login/redeem" class="btn mt-3 btn-sm">Sign up with an invitation code</router-link>
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
      password: '',
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
    async login() {
      try {
        if (this.password.length <= 5 || this.email.length === 0) {
          this.errors = 'Email and password fields are required. Password should be longer than 6 characters.'
          return
        }
        if (!this.validateEmail()) {
          this.errors = 'Email field should be actual email address.'
          return
        }

        const data = await this.$store.dispatch('login', {
          email: this.email,
          password: this.password,
        })
        if (!data) {
          this.errors = 'Login failed'
          return
        }
        location.reload()
      } catch (e) {
        this.errors = e
      }
    }
  }
}
</script>