<template>
  <div class="flex relative h-screen overflow-hidden">
    <div class="w-1/4 border-r-[1px] border-white/5 h-screen bg-midnight">
      <MainMenu/>
    </div>
    <div class="bg-midnight-light text-white w-full min-h-screen">
      <div class="h-20 flex justify-end p-6">
        <span class="badge p-3">
          <a href="https://github.com/twelvee/boxie">
          <img src="/assets/github.png" alt="github page" class="w-12"/>
          </a>
        </span>
      </div>
      <div>
      <slot/>
      </div>
    </div>
  </div>
</template>
<script setup lang="ts">
import MainMenu from "../navigations/MainMenu.vue";
</script>