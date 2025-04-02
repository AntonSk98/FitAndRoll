<script>
    import { writable } from "svelte/store";
    import { i18n, getLocale, setLocale } from "./i18n";
    import { onMount } from "svelte";

    export let activePage;
    export let localeSet;

    let pageSection;

    function getPageSection() {
        return {
            courses: i18n("navigation.courses"),
            participants: i18n("navigation.participants"),
            statistics: i18n("navigation.statistics"),
            archive: i18n("navigation.archive"),
        };
    }

    $: if (localeSet) {
        pageSection = getPageSection();
    }
</script>

<nav class="mb-4 flex gap-3 flex-wrap items-center">
    <ul class="flex justify-start gap-2 flex-wrap">
        {#each Object.keys(pageSection) as page}
            <li>
                <button
                    on:click={() => activePage.set(page)}
                    class="p-3
                    text-md
                    font-semibold
                    text-[var(--text-color-darker)]
                    rounded-full
                    border-2
                    border-transparent
                    transition
                    duration-700
                    hover:scale-[1.08]
                    hover:bg-[var(--primary-color-light)] hover:text-white
                    cursor-pointer
                {$activePage === page
                        ? 'bg-[var(--primary-color-dark)] text-white shadow-xl shadow-[var(--primary-color-dark)]'
                        : ''}"
                >
                    {pageSection[page]}
                </button>
            </li>
        {/each}
    </ul>
</nav>
