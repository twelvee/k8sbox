<template>
  <default-layout>
    <div class="flex w-full h-full">
      <div class="flex flex-col w-full">
        <div class="flex justify-between items-center border-y-[1px] border-white/5">
          <div class="text-xl pl-10 py-4">
            Environments
          </div>
          <div>
            <a href="#" class="underline mr-4">How to work with environments</a>
            <button class="btn btn-sm mr-2">Run</button>
          </div>
        </div>
        <div class="pl-4 mt-4 mr-2">
          <div class="overflow-y-auto h-screen">
            <div class="w-full grid grid-cols-2 gap-6">
              <div class="">
                <article class="rounded-lg border-[1px] border-white/10 p-4 hover:border-primary transition cursor-pointer">
                  <div class="flex items-center mb-4">
                    <span class="badge badge-success inline mr-2"></span>
                    <div class="m-0 p-0">
                      <h3 class="text-lg font-medium text-white">test-environment-1</h3>
                    </div>
                  </div>
                  <div class="flex gap-2 items-center">
                    <div>
                      <a
                          href="#"
                          class="block text-white/70 hover:text-white rounded-lg border border-white/5 p-4 hover:border-primary/20"
                      >
                        <strong class="font-medium">Analytics</strong>
                        <p class="mt-1 text-xs font-medium">
                          11 applications
                        </p>
                      </a>
                    </div>
                    <div>
                      <a
                          href="#"
                          class="block text-white/70 hover:text-white rounded-lg border border-white/5 p-4 hover:border-primary/20"
                      >
                        <strong class="font-medium">Backoffice</strong>

                        <p class="mt-1 text-xs font-medium">
                          4 applications
                        </p>
                      </a>
                    </div>
                    <div>
                      <a
                          href="#"
                          class="block h-full text-white/70 hover:text-white rounded-lg border border-white/5 p-4 hover:border-primary/20"
                      >
                        <strong class="font-medium">External Api</strong>

                        <p class="mt-1 text-xs font-medium">
                          21 applications
                        </p>
                      </a>
                    </div>
                  </div>
                  <div class="text-white/50 pl-4 py-2"> ... and 3 more boxes</div>
                  <div class="text-white/70 pt-4">
                    <div class="avatar placeholder">
                      <div class="bg-neutral-focus text-neutral-content rounded-full w-8 mr-2">
                        <span class="text-sm">SZ</span>
                      </div>
                    </div>
                    <span class="underline hover:text-primary cursor-pointer">Sergey Zyryanov</span> 1 hour ago on cluster <span class="underline hover:text-primary cursor-pointer">Minikube</span>
                  </div>
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