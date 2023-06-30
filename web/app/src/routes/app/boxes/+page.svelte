<script>
    import {popup} from '@skeletonlabs/skeleton';

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
        <h1 class="h1 inline-block">Boxes</h1>
        <a href="#" class="btn variant-ghost-primary font-semibold">Create Box</a>
    </div>

    <div class="mt-6 flex justify-between content-center">
        <h1 class="h3">12 boxes</h1>
    </div>
    <div class="grid grid-cols-4 mt-4 gap-4">
        <div class="card">
            <header class="card-header font-semibold truncate">Analytics</header>
            <section class="p-4 pt-2">
                <div class="font-thin -mt-2 mb-2">12 applications</div>
                <div class="block truncate" use:popup={popupFeatured}>
                    <span class="badge variant-ghost-surface">api</span>
                    <span class="badge variant-ghost-surface">configmap</span>
                    <span class="badge variant-ghost-surface">deployment-update-user-information</span>
                    <span class="badge variant-ghost-surface">cronjob-create-app</span>
                </div>

            </section>
            <footer class="card-footer">
                <a href="#" class="btn btn-sm variant-outline-secondary">Details</a>
                <a href="#" class="btn btn-sm variant-outline-warning">Edit</a>
                <a href="#" class="btn btn-sm variant-outline-error">Delete</a>
            </footer>

            <div class="card p-4 w-72 shadow-xl" data-popup="popupFeatured">
                <span class="badge variant-ghost-surface">api</span>
                <span class="badge variant-ghost-surface">configmap</span>
                <span class="badge variant-ghost-surface">deployment-update-user-information</span>
                <span class="badge variant-ghost-surface">cronjob-create-app</span>
                <span class="badge variant-ghost-surface">api</span>
                <span class="badge variant-ghost-surface">configmap</span>
                <span class="badge variant-ghost-surface">deployment-update-user-information</span>
                <span class="badge variant-ghost-surface">cronjob-create-app</span>
                <span class="badge variant-ghost-surface">api</span>
                <span class="badge variant-ghost-surface">configmap</span>
                <span class="badge variant-ghost-surface">deployment-update-user-information</span>
                <span class="badge variant-ghost-surface">cronjob-create-app</span>
            </div>
        </div>
    </div>
</div>