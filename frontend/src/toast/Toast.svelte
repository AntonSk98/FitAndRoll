<script>
    import { createEventDispatcher } from "svelte";
    import { fade, scale } from "svelte/transition";
    const dispatch = createEventDispatcher();

    export let type = "error";

    const icons = {
        success: `<?xml version="1.0"?><svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="2.5" stroke="currentColor" class="size-6"><path stroke-linecap="round" stroke-linejoin="round" d="M9 12.75 11.25 15 15 9.75M21 12a9 9 0 1 1-18 0 9 9 0 0 1 18 0Z"/></svg>`,
        error: `<?xml version="1.0"?><svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="2.5" stroke="currentColor" class="size-6 icon"><path stroke-linecap="round" stroke-linejoin="round" d="M12 9v3.75m-9.303 3.376c-.866 1.5.217 3.374 1.948 3.374h14.71c1.73 0 2.813-1.874 1.948-3.374L13.949 3.378c-.866-1.5-3.032-1.5-3.898 0L2.697 16.126ZM12 15.75h.007v.008H12v-.008Z"/></svg>`
    };
</script>

<!-- svelte-ignore a11y-click-events-have-key-events -->
<div class="notification {type}" on:click={() => dispatch("dismiss")} transition:fade>
    <div class="icon">{@html icons[type] ?? ""}</div>
    <div class="content"><slot /></div>
</div>

<style>
    .notification {
        display: flex;
        align-items: center;
        gap: 0.5rem;
        padding: 1rem 1.4rem;
        border-radius: 1rem;
        width: 18rem;
        margin-top: 0.6rem;
        margin-right: 0.6rem;
        box-shadow: 0 10px 30px rgba(0, 0, 0, 0.1);
        cursor: pointer;
        transition: transform 0.2s ease, box-shadow 0.2s ease;
    }

    .notification:hover {
        transform: translateY(-4px) scale(1.03);
        box-shadow: 0 12px 35px rgba(0, 0, 0, 0.15);
    }

    .icon {
        color: white;
        font-weight: bold;
        line-height: 1;
        opacity: 0.9;
    }

    .content {
        font-weight: bold;
        color: white;
        font-size: 0.95rem;
        line-height: 1.5;
        opacity: 0.9;
    }

    .success {
        background: linear-gradient(135deg, var(--primary-color) 0%, var(--primary-color-dark) 100%);
    }

    .error {
        background: linear-gradient(135deg, #EF4444 0%, #B91C1C 100%);
    }
</style>
