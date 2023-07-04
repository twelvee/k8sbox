<template>
  <div></div>
</template>
<script lang="ts">
import Cookies from 'js-cookie'
export default {
  mounted() {
    Cookies.remove('x-auth-token')
    location.reload()
  }
}
</script>