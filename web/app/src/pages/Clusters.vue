<template>
  <default-layout>
    <div class="flex w-full h-full">
      <div class="flex flex-col w-full">
        <div class="flex justify-between items-center border-y-[1px] border-white/5">
          <div class="text-xl pl-10 py-4">
            Clusters
          </div>
          <div>
            <button class="btn btn-sm mr-2">Connect new cluster</button>
          </div>
        </div>
        <div class="pl-4 mt-4">
          <div class="overflow-y-auto h-screen">
            <div class="w-full grid grid-cols-4 gap-6">
              <div>
                <article class="rounded-lg border-[1px] border-white/10 p-4">
                      <div class="flex items-center">
                        <span class="badge badge-success inline mr-2"></span>
                        <div class="m-0 p-0">
                        <h3 class="text-lg font-medium text-white">Minikube</h3>

                        </div>
                      </div>
                  <div class="text-xs mt-2">IP: 127.0.0.1</div>
                  <div class="text-xs">Connected: 12.11.2023</div>

                  <ul class="mt-4 space-y-2">
                    <li>
                      <a
                          href="#"
                          class="block h-full text-white/70 hover:text-white rounded-lg border border-white/5 p-4 hover:border-primary/20"
                      >
                        <strong class="font-medium">Last deployment</strong>

                        <p class="mt-1 text-xs font-medium">
                          1h ago by Sergey Zyryanov
                        </p>
                      </a>
                    </li>
                  </ul>
                </article>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </default-layout>
</template>
<script setup lang="ts">
import DefaultLayout from "../components/layouts/DefaultLayout.vue";
</script>