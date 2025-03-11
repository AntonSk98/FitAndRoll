<script>
    import { t } from "svelte-i18n";

    export let total;
    export let onPaginationChanged = (currentPagination) => {};

    export let pageSizeOptions = [5, 10, 25];
    $: pagination = {
        currentPage: 1,
        pageSize: pageSizeOptions[0],
        totalItems: total,
    };

    $: isFirstPage = pagination.currentPage === 1;
    $: isLastPage =
        pagination.currentPage >=
        Math.ceil(pagination.totalItems / pagination.pageSize);

    function previousPage() {
        if (isFirstPage) {
            return;
        }

        pagination.currentPage--;
        onPaginationChanged(pagination);
    }

    function nextPage() {
        if (isLastPage) {
            return;
        }

        pagination.currentPage++;
        onPaginationChanged(pagination);
    }

    function setPageSize(pageSize) {
        pagination.pageSize = pageSize;
        pagination.currentPage = 1;
        onPaginationChanged(pagination);
    }

    export function resetPagination() {
        pagination.currentPage = 1;
        pagination.pageSize = pageSizeOptions[0];
        return pagination;
    }
</script>

<div class="pagination-container">
    <div class="flex flex-col gap-1 items-start">
        <div>
            <button
                class="pagination-button"
                disabled={isFirstPage}
                on:click={previousPage}
            >
                <svg
                    xmlns="http://www.w3.org/2000/svg"
                    viewBox="0 0 24 24"
                    fill="currentColor"
                    class="w-4 h-4"
                >
                    <path
                        fill-rule="evenodd"
                        d="M11.03 3.97a.75.75 0 0 1 0 1.06l-6.22 6.22H21a.75.75 0 0 1 0 1.5H4.81l6.22 6.22a.75.75 0 1 1-1.06 1.06l-7.5-7.5a.75.75 0 0 1 0-1.06l7.5-7.5a.75.75 0 0 1 1.06 0Z"
                        clip-rule="evenodd"
                    />
                </svg>
            </button>

            <span class="pagination-info">
                <strong>{pagination.currentPage}</strong>
                {$t("common.pagination.of")}
                <strong
                    >{Math.ceil(
                        pagination.totalItems / pagination.pageSize,
                    )}</strong
                >
            </span>

            <button
                class="pagination-button"
                disabled={isLastPage}
                on:click={nextPage}
            >
                <svg
                    xmlns="http://www.w3.org/2000/svg"
                    viewBox="0 0 24 24"
                    fill="currentColor"
                    class="w-4 h-4"
                >
                    <path
                        fill-rule="evenodd"
                        d="M12.97 3.97a.75.75 0 0 1 1.06 0l7.5 7.5a.75.75 0 0 1 0 1.06l-7.5 7.5a.75.75 0 1 1-1.06-1.06l6.22-6.22H3a.75.75 0 0 1 0-1.5h16.19l-6.22-6.22a.75.75 0 0 1 0-1.06Z"
                        clip-rule="evenodd"
                    />
                </svg>
            </button>
        </div>
        <div class="text-[var(--text-color)] font-bold text-sm cursor-default">Total: {total}</div>
    </div>

    <div class="page-size-container">
        <span>{$t("common.pagination.size")}</span>
        {#each pageSizeOptions as pageSizeOption}
            <button
                class="page-size-button {pagination.pageSize === pageSizeOption
                    ? 'active'
                    : ''}"
                on:click={() => setPageSize(pageSizeOption)}
                >{pageSizeOption}</button
            >
        {/each}
    </div>
</div>

<style>
    .pagination-container {
        display: flex;
        align-items: center;
        gap: 1rem;
        justify-content: space-between;
        margin-top: 1.5rem;
        flex-wrap: wrap;
    }

    .pagination-button {
        cursor: pointer;
        padding: 0.5rem;
        border-radius: 0.375rem;
        background-color: var(--primary-color);
        color: white;
        border: none;
        transition:
            background-color 0.2s ease-in-out,
            box-shadow 0.2s ease-in-out;
        box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
    }

    .pagination-button:hover {
        background-color: var(--primary-color-dark);
    }

    .pagination-button:disabled {
        opacity: 0.5;
        cursor: not-allowed;
        box-shadow: none;
    }

    .pagination-info {
        font-size: 0.875rem;
        color: var(--text-color);
        font-weight: 500;
    }

    .page-size-container {
        display: flex;
        align-items: center;
        gap: 0.5rem;
    }

    .page-size-button {
        padding: 0.5rem 0.5rem;
        border-radius: 2rem;
        font-size: 0.875rem;
        font-weight: 500;
        min-width: 2.5rem;
        background-color: var(--secondary-color);
        color: var(--text-color);
        border: 1px solid var(--secondary-color-darker);
        cursor: pointer;
        transition: background-color 0.2s ease-in-out;
    }

    .page-size-button:hover {
        background-color: var(--primary-color);
        color: white;
        border-color: var(--primary-color);
    }

    .page-size-button.active {
        background-color: var(--primary-color-dark);
        color: white;
        border-color: var(--primary-color-dark);
    }
</style>
