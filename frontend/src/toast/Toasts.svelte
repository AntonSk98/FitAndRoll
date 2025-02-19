<script>
    import Toast from "./Toast.svelte";
    import { dismissToast, toasts } from "./toastStore";
</script>

{#if $toasts}
    <section>
        {#each $toasts as toast (toast.id)}
            <Toast type={toast.type} on:dismiss={() => dismissToast(toast.id)}
                >{toast.message}
            </Toast>
        {/each}
    </section>
{/if}

<style>
    section {
        position: absolute;
        top: 0;
        right: 0;
    }
</style>
