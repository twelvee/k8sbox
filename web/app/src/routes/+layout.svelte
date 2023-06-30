<script lang='ts'>
    // The ordering of these imports is critical to your app working properly
    import '@skeletonlabs/skeleton/themes/theme-crimson.css';
    // If you have source.organizeImports set to true in VSCode, then it will auto change this ordering
    import '@skeletonlabs/skeleton/styles/skeleton.css';
    // Most of your app wide CSS should be put in this file
    import '../app.postcss';
    import {AppBar, AppShell} from '@skeletonlabs/skeleton';
    import {page} from '$app/stores';
    import logo from '$img/logo.svg'
    import { computePosition, autoUpdate, offset, shift, flip, arrow } from '@floating-ui/dom';

    import { storePopup } from '@skeletonlabs/skeleton';
    storePopup.set({ computePosition, autoUpdate, offset, shift, flip, arrow });

</script>

<!-- App Shell -->
<AppShell>
    <svelte:fragment slot="header">
        <!-- App Bar -->
        <AppBar>
            <svelte:fragment slot="lead">
                <img src={logo} width="30px" class="inline"/>
                <strong class="text-xl uppercase ml-[10px]">
                    Boxie
                </strong>
            </svelte:fragment>
            <div class="w-full flex justify-center gap-6">
                {#if $page.route.id === "/app"}
                    <a href="/app" class="text-primary-500">Dashboard</a>
                {:else}
                    <a href="/app" class="hover:text-primary-400">Dashboard</a>
                {/if}
                {#if $page.route.id === "/app/tokens"}
                    <a href="/app/tokens" class="text-primary-500">Tokens</a>
                {:else}
                    <a href="/app/tokens" class="hover:text-primary-400">Tokens</a>
                {/if}
                {#if $page.route.id === "/app/clusters"}
                    <a href="/app/clusters" class="text-primary-500">Clusters</a>
                {:else}
                    <a href="/app/clusters" class="hover:text-primary-400">Clusters</a>
                {/if}
                {#if $page.route.id === "/app/boxes"}
                    <a href="/app/boxes" class="text-primary-500">Boxes</a>
                {:else}
                    <a href="/app/boxes" class="hover:text-primary-400">Boxes</a>
                {/if}
                {#if $page.route.id === "/app/environments"}
                    <a href="/app/environments" class="text-primary-500">Environments</a>
                {:else}
                    <a href="/app/environments" class="hover:text-primary-400">Environments</a>
                {/if}
            </div>
            <svelte:fragment slot="trail">
                <a
                        class="btn btn-sm variant-ghost-surface"
                        href="https://github.com/twelvee/boxie"
                        target="_blank"
                        rel="noreferrer"
                >
                    GitHub
                </a>
            </svelte:fragment>
        </AppBar>
    </svelte:fragment>
    <!-- Page Route Content -->
    <slot/>
</AppShell>
