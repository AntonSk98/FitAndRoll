<script>
    import { t } from "svelte-i18n";

    export let activePage;
</script>

<nav class="mb-4">
    <ul class="flex justify-start gap-1.5 flex-wrap">
        {#each [$t("navigation.courses"), $t("navigation.participants"), $t("navigation.archive")] as page}
            <li>
                <button
                    on:click={() => activePage.set(page)}
                    class="p-3
                        text-md
                        font-semibold
                        text-gray-700
                        rounded-full
                        border-2
                        border-transparent
                        transition
                        duration-300
                        hover:bg-[var(--primary-color-light)] hover:text-white
                    {$activePage === page
                        ? 'bg-[var(--primary-color-dark)] text-white shadow-xl shadow-[var(--primary-color-dark)]'
                        : ''}"
                >
                    {page}
                </button>
            </li>
        {/each}
    </ul>
</nav>

<style>
</style>
