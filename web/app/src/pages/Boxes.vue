<template>
  <default-layout>
    <div class="flex w-full h-full">
      <div class="flex flex-col w-full">
        <div class="flex justify-between items-center border-y-[1px] border-white/5">
          <div class="text-xl pl-10 py-4">
            Boxes
          </div>
          <div>
            <a href="#" class="underline mr-4">How to work with boxes</a>
            <button class="btn btn-sm mr-2">Create new box</button>
          </div>
        </div>
        <div class="pl-4 mt-4 mr-2">
          <div class="overflow-y-auto h-screen">
            <div class="w-full grid grid-cols-3 gap-6">
              <div>
                <article class="rounded-lg border-[1px] border-white/10 p-4">
                  <div class="flex items-center justify-between mb-2">
                    <h3 class="text-lg font-medium text-white">Analytics</h3>
                    <div class="dropdown dropdown-end">
                      <label tabindex="0" class="text-lg hover:cursor-pointer">
                        <EllipsisVerticalIcon class="h-5 w-5" />
                      </label>
                      <div tabindex="0" class="dropdown-content z-[1] p-2 bg-midnight rounded-md w-40">
                        <a class="hover:cursor-pointer block p-2 hover:bg-white/10 w-full rounded-md text-md flex items-center">
                          <div><EyeIcon class="h-4 w-4 mr-2" /></div>
                          <div>View</div>
                        </a>
                        <a class="hover:cursor-pointer block p-2 hover:bg-white/10 w-full rounded-md text-md flex items-center">
                          <div><Cog6ToothIcon class="h-4 w-4 mr-2" /></div>
                          <div>Manage</div>
                        </a>
                      </div>
                    </div>
                  </div>
                  <div class="items-start">
                    <div class="badge">cronjob-test</div>
                    <div class="badge badge-outline-primary">deployment-test</div>
                    <div class="badge badge-outline-primary">nginx-ingress</div>
                    <div class="badge badge-outline-primary">nice-service</div>
                    <div class="badge">cronjob-test</div>
                    <div class="badge badge-outline-primary">deployment-test</div>
                    <div class="badge badge-outline-primary">nginx-ingress</div>
                    <div class="badge badge-outline-primary">nice-service</div>
                    <div class="badge">cronjob-test</div>
                    <div class="badge badge-outline-primary">deployment-test</div>
                    <div class="badge badge-outline-primary">nginx-ingress</div>
                    <div class="badge badge-outline-primary">nice-service</div>
                  </div>
                </article>
              </div>
              <div>
                <article class="rounded-lg border-[1px] border-white/10 p-4">
                  <div class="flex items-center justify-between mb-2">
                    <h3 class="text-lg font-medium text-white">GraphQL</h3>
                    <div class="dropdown dropdown-end">
                      <label tabindex="0" class="text-lg hover:cursor-pointer">
                        <EllipsisVerticalIcon class="h-5 w-5" />
                      </label>
                      <div tabindex="0" class="dropdown-content z-[1] p-2 bg-midnight rounded-md w-40">
                        <a class="hover:cursor-pointer block p-2 hover:bg-white/10 w-full rounded-md text-md flex items-center">
                          <div><EyeIcon class="h-4 w-4 mr-2" /></div>
                          <div>View</div>
                        </a>
                        <a class="hover:cursor-pointer block p-2 hover:bg-white/10 w-full rounded-md text-md flex items-center">
                          <div><Cog6ToothIcon class="h-4 w-4 mr-2" /></div>
                          <div>Manage</div>
                        </a>
                      </div>
                    </div>
                  </div>

                  <div class="items-start">
                    <div class="badge">cronjob-test</div>
                  </div>
                </article>
              </div>
              <div>
                <article class="rounded-lg border-[1px] border-white/10 p-4">
                  <div class="flex items-center justify-between mb-2">
                    <h3 class="text-lg font-medium text-white">GraphQL Books</h3>
                    <div class="dropdown dropdown-end">
                      <label tabindex="0" class="text-lg hover:cursor-pointer">
                        <EllipsisVerticalIcon class="h-5 w-5" />
                      </label>
                      <div tabindex="0" class="dropdown-content z-[1] p-2 bg-midnight rounded-md w-40">
                        <a class="hover:cursor-pointer block p-2 hover:bg-white/10 w-full rounded-md text-md flex items-center">
                          <div><EyeIcon class="h-4 w-4 mr-2" /></div>
                          <div>View</div>
                        </a>
                        <a class="hover:cursor-pointer block p-2 hover:bg-white/10 w-full rounded-md text-md flex items-center">
                          <div><Cog6ToothIcon class="h-4 w-4 mr-2" /></div>
                          <div>Manage</div>
                        </a>
                      </div>
                    </div>
                  </div>
                  <div class="items-start">
                    <div class="badge">cronjob-test</div>
                    <div class="badge badge-outline-primary">deployment-test</div>
                    <div class="badge badge-outline-primary">nginx-ingress</div>
                    <div class="badge badge-outline-primary">nice-service</div>
                    <div class="badge">cronjob-test</div>
                    <div class="badge badge-outline-primary">deployment-test</div>
                    <div class="badge badge-outline-primary">nginx-ingress</div>
                  </div>
                </article>
              </div>
              <div>
                <article class="rounded-lg border-[1px] border-white/10 p-4">
                  <div class="flex items-center justify-between mb-2">
                    <h3 class="text-lg font-medium text-white">External API</h3>
                    <div class="dropdown dropdown-end">
                      <label tabindex="0" class="text-lg hover:cursor-pointer">
                        <EllipsisVerticalIcon class="h-5 w-5" />
                      </label>
                      <div tabindex="0" class="dropdown-content z-[1] p-2 bg-midnight rounded-md w-40">
                        <a class="hover:cursor-pointer block p-2 hover:bg-white/10 w-full rounded-md text-md flex items-center">
                          <div><EyeIcon class="h-4 w-4 mr-2" /></div>
                          <div>View</div>
                        </a>
                        <a class="hover:cursor-pointer block p-2 hover:bg-white/10 w-full rounded-md text-md flex items-center">
                          <div><Cog6ToothIcon class="h-4 w-4 mr-2" /></div>
                          <div>Manage</div>
                        </a>
                      </div>
                    </div>
                  </div>
                  <div class="items-start">
                    <div class="badge">cronjob-test</div>
                    <div class="badge badge-outline-primary">deployment-test</div>
                    <div class="badge badge-outline-primary">nginx-ingress</div>
                    <div class="badge badge-outline-primary">nice-service</div>
                    <div class="badge">cronjob-test</div>
                    <div class="badge badge-outline-primary">deployment-test</div>
                    <div class="badge badge-outline-primary">nginx-ingress</div>
                  </div>
                </article>
              </div>
              <div>
                <article class="rounded-lg border-[1px] border-white/10 p-4">
                  <div class="flex items-center justify-between mb-2">
                    <h3 class="text-lg font-medium text-white">Backoffice</h3>
                    <div class="dropdown dropdown-end">
                      <label tabindex="0" class="text-lg hover:cursor-pointer">
                        <EllipsisVerticalIcon class="h-5 w-5" />
                      </label>
                      <div tabindex="0" class="dropdown-content z-[1] p-2 bg-midnight rounded-md w-40">
                        <a class="hover:cursor-pointer block p-2 hover:bg-white/10 w-full rounded-md text-md flex items-center">
                          <div><EyeIcon class="h-4 w-4 mr-2" /></div>
                          <div>View</div>
                        </a>
                        <a class="hover:cursor-pointer block p-2 hover:bg-white/10 w-full rounded-md text-md flex items-center">
                          <div><Cog6ToothIcon class="h-4 w-4 mr-2" /></div>
                          <div>Manage</div>
                        </a>
                      </div>
                    </div>
                  </div>
                  <div class="items-start">
                    <div class="badge">cronjob-test</div>
                    <div class="badge badge-outline-primary">deployment-test</div>
                    <div class="badge badge-outline-primary">nginx-ingress</div>
                    <div class="badge badge-outline-primary">nice-service</div>
                    <div class="badge">cronjob-test</div>
                    <div class="badge badge-outline-primary">deployment-test</div>
                    <div class="badge badge-outline-primary">nginx-ingress</div>
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
import {EllipsisVerticalIcon, Cog6ToothIcon, EyeIcon} from '@heroicons/vue/24/solid'

</script>