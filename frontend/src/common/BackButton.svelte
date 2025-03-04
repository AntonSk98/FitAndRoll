<script>
    export let onBackButtonClicked;
</script>

<button class="back" on:click={onBackButtonClicked}>
    <svg
        xmlns="http://www.w3.org/2000/svg"
        class="w-5 h-5"
        fill="none"
        viewBox="0 0 24 24"
        stroke="currentColor"
    >
        <path
            stroke-linecap="round"
            stroke-linejoin="round"
            stroke-width="2"
            d="M15 19l-7-7 7-7"
        />
    </svg>
</button>

<style>
    .back {
        transition: all 0.3s ease-in-out;
        background-color: var(--primary-color);
        padding: 0.8rem;
        border-radius: 1rem;
        color: white;
        cursor: pointer;
    }

    .back:hover {
        background-color: var(--secondary-color-dark);
        transform: scale(1.08);
    }
</style>
