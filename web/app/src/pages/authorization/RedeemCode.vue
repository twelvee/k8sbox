<template>
  <div class="hero min-h-screen bg-midnight">
    <div class="hero-content flex-col lg:flex-row-reverse">
      <div class="card flex-shrink-0 w-full max-w-sm shadow-2xl bg-base-100">
        <div class="card-header text-center">
          <h1 class="h1 text-2xl pt-4">Welcome to Boxie!</h1>
          <div class="badge" v-if="urlCode">{{urlCode}}</div>
        </div>
        <div class="card-body">
          <div class="form-control">
            <label class="label">
              <span class="label-text">Code</span>
            </label>
            <input type="text" v-model="code" placeholder="" class="input input-bordered" />
          </div>
          <div class="form-control">
            <label class="label">
              <span class="label-text">Password</span>
            </label>
            <input type="password" v-model="password" placeholder="" class="input input-bordered" />
          </div>
          <div class="form-control">
            <label class="label">
              <span class="label-text">Confirm password</span>
            </label>
            <input type="password" v-model="password_confirmation" placeholder="" class="input input-bordered" />
          </div>
          <div v-if="errors" class="text-error">
            {{errors}}
          </div>
          <div class="form-control mt-6">
            <button @click="redeemCode" class="btn btn-primary">Create account</button>
            <router-link to="/login" class="btn mt-3 btn-sm">Sign in with email and password</router-link>
          </div>
        </div>
      </div>
    </div>
  </div>

</template>
<script setup lang="ts">
</script>

<script lang="ts">
export default {
  data() {
    return {
      code: this.$route.query.code,
      urlCode: this.$route.query.code,
      password: '',
      password_confirmation: '',
      errors: ''
    }
  },
  mounted() {
    if (this.code != this.urlCode) {
      this.code = this.urlCode
    }
  },
  methods: {
    async redeemCode() {
      try {
        const user = await this.$store.dispatch('redeemCode', {
          code: this.code,
          password: this.password,
          password_confirmation: this.password_confirmation
        })
        if (user?.error) {
          this.errors = user?.error
        }
        location.reload()
      } catch (e) {
        this.errors = e
      }
    }
  }
}
</script>