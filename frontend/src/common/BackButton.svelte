<script>
    export let onBackButtonClicked;
</script>

<button class="back" on:click={onBackButtonClicked}>
    <svg
        xmlns="http://www.w3.org/2000/svg"
        fill="none"
        viewBox="0 0 24 24"
        stroke-width="2.5"
        stroke="currentColor"
        class="w-5 h-5"
    >
        <path
            stroke-linecap="round"
            stroke-linejoin="round"
            d="M12 9.75 14.25 12m0 0 2.25 2.25M14.25 12l2.25-2.25M14.25 12 12 14.25m-2.58 4.92-6.374-6.375a1.125 1.125 0 0 1 0-1.59L9.42 4.83c.21-.211.497-.33.795-.33H19.5a2.25 2.25 0 0 1 2.25 2.25v10.5a2.25 2.25 0 0 1-2.25 2.25h-9.284c-.298 0-.585-.119-.795-.33Z"
        />
    </svg>
</button>

<style>
    .back {
        transition: all 0.5s ease-in-out;
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
