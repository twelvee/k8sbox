<script>
    import {Avatar, popup} from '@skeletonlabs/skeleton';

    const popupFeatured = {
        // Represents the type of event that opens/closed the popup
        event: 'hover',
        // Matches the data-popup value on your popup element
        target: 'popupFeatured',
        // Defines which side of your trigger the popup will appear
        placement: 'bottom',
    };
</script>
<div class="container mx-auto mt-4">
    <div class="flex justify-between">
        <h1 class="h1 inline-block">Environments</h1>
        <a href="#" class="btn variant-ghost-primary font-semibold">Run environment</a>
    </div>

    <div class="mt-6 flex justify-between content-center">
        <h1 class="h3">2 active environments</h1>
    </div>
    <div class="grid grid-cols-4 mt-4 gap-4">
        <div class="card">
            <header class="card-header relative">
                <div>
                    <Avatar
                            initials="AA"
                            border="border-3 border-surface-300-600-token hover:!border-primary-500"
                            cursor="cursor-pointer"
                            class="w-[44px] absolute "
                    />
                </div>
                <div class="font-semibold ml-[60px] truncate">
                    ts-plus-test-sdf
                </div>
            </header>
            <section class="p-4 pt-2">
                <div class="block truncate -mt-3 mb-2 ml-[60px]">
                    <span class="badge variant-ghost-primary h-4 mr-1">Minikube</span>
                    <span class="badge variant-filled-tertiary h-4 mr-1">pending</span>
                </div>
                <div class="font-thin">4 boxes</div>
                <div class="block truncate" use:popup={popupFeatured}>
                    <span class="badge variant-ghost-surface">graphql-books</span>
                    <span class="badge variant-ghost-surface">content</span>
                    <span class="badge variant-ghost-surface">analytics</span>
                    <span class="badge variant-ghost-surface">backend-graphql</span>
                </div>

            </section>
            <footer class="card-footer">
                <a href="#" class="btn btn-sm variant-outline-secondary">Details</a>
                <a href="#" class="btn btn-sm variant-outline-warning">Edit</a>
                <a href="#" class="btn btn-sm variant-outline-error">Delete</a>
            </footer>

            <div class="card p-4 w-72 shadow-xl" data-popup="popupFeatured">
                <span class="badge variant-ghost-surface">graphql-books</span>
                <span class="badge variant-ghost-surface">content</span>
                <span class="badge variant-ghost-surface">analytics</span>
                <span class="badge variant-ghost-surface">backend-graphql</span>
            </div>
        </div>
    </div>
</div>