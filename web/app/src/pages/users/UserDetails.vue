<template>
  <default-layout>
    <div class="flex w-full h-full">
      <div class="flex flex-col w-full"  v-if="user">
        <div class="flex justify-between items-center border-y-[1px] border-white/5">
          <div class="text-xl pl-10 py-4">
            User details
          </div>
        </div>
        <div class="pl-4 mt-4">
          <div class="overflow-y-auto h-screen">
            <div class="w-full p-6 pt-2">
              <div>
                <div class="overflow-x-auto">
                  <table class="table">
                    <tbody>
                    <tr>
                      <th>ID</th>
                      <th>{{ user.ID }}</th>
                    </tr>
                    <tr>
                      <th>Name</th>
                      <th>{{ user.Name }}</th>
                    </tr>
                    <tr>
                      <th>Email</th>
                      <th>{{ user.Email }}</th>
                    </tr>
                    <tr>
                      <th>Created At</th>
                      <th>{{ user.CreatedAt }}</th>
                    </tr>
                    </tbody>
                  </table>
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
import { format } from 'date-fns'
export default {
  data() {
    return {
      user: null
    }
  },
  computed: {
    ...mapGetters(['getUsers']),
  },
  async mounted() {
    if (this.getUsers.size === 0) {
      await this.$store.dispatch('getUsers')
    }
    this.user = this.$store.getters.getUserById(this.$route.params.id)
    this.user.CreatedAt = format(new Date(this.user.CreatedAt * 1000), 'yyyy-MM-dd')
  }
}
</script>