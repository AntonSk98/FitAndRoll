<script>
    import Modal from "../common/Modal.svelte";
    import PaginationComponent from "../../pagination/PaginationComponent.svelte";
    import { space } from "svelte/internal";

    export let tableHeader;
    export let total;

    export let columns = [];
    export let rows = [];
    export let actions = [];
    export let headerActions = [];
    export let onPaginationFilterChanged = (filter, pagination) => {};

    let filter = {};
    let paginationRef;

    let pendingAction;
    let showModal = false;
    function displayModal() {
        showModal = true;
    }
    function onModalConfirmed() {
        if (pendingAction) {
            pendingAction();
        }
        showModal = false;
    }
    function onModalCanceled() {
        showModal = false;
    }

    const icons = {
        edit: `<?xml version="1.0"?><svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="2" stroke="currentColor" class="w-5 h-5"><path stroke-linecap="round" stroke-linejoin="round" d="m16.862 4.487 1.687-1.688a1.875 1.875 0 1 1 2.652 2.652L10.582 16.07a4.5 4.5 0 0 1-1.897 1.13L6 18l.8-2.685a4.5 4.5 0 0 1 1.13-1.897l8.932-8.931Zm0 0L19.5 7.125M18 14v4.75A2.25 2.25 0 0 1 15.75 21H5.25A2.25 2.25 0 0 1 3 18.75V8.25A2.25 2.25 0 0 1 5.25 6H10"/></svg>`,
        info: `<?xml version="1.0"?><svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="2" stroke="currentColor" class="w-5 h-5"><path stroke-linecap="round" stroke-linejoin="round" d="M17.982 18.725A7.488 7.488 0 0 0 12 15.75a7.488 7.488 0 0 0-5.982 2.975m11.963 0a9 9 0 1 0-11.963 0m11.963 0A8.966 8.966 0 0 1 12 21a8.966 8.966 0 0 1-5.982-2.275M15 9.75a3 3 0 1 1-6 0 3 3 0 0 1 6 0Z"/></svg>`,
        calender: `<?xml version="1.0"?><svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="2" stroke="currentColor" class="w-5 h-5"><path stroke-linecap="round" stroke-linejoin="round" d="M6.75 2.994v2.25m10.5-2.25v2.25m-14.252 13.5V7.491a2.25 2.25 0 0 1 2.25-2.25h13.5a2.25 2.25 0 0 1 2.25 2.25v11.251m-18 0a2.25 2.25 0 0 0 2.25 2.25h13.5a2.25 2.25 0 0 0 2.25-2.25m-18 0v-7.5a2.25 2.25 0 0 1 2.25-2.25h13.5a2.25 2.25 0 0 1 2.25 2.25v7.5m-6.75-6h2.25m-9 2.25h4.5m.002-2.25h.005v.006H12v-.006Zm-.001 4.5h.006v.006h-.006v-.005Zm-2.25.001h.005v.006H9.75v-.006Zm-2.25 0h.005v.005h-.006v-.005Zm6.75-2.247h.005v.005h-.005v-.005Zm0 2.247h.006v.006h-.006v-.006Zm2.25-2.248h.006V15H16.5v-.005Z"/></svg>`,
        trash: `<?xml version="1.0"?><svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="2" stroke="currentColor" class="w-5 h-5"><path stroke-linecap="round" stroke-linejoin="round" d="m14.74 9-.346 9m-4.788 0L9.26 9m9.968-3.21c.342.052.682.107 1.022.166m-1.022-.165L18.16 19.673a2.25 2.25 0 0 1-2.244 2.077H8.084a2.25 2.25 0 0 1-2.244-2.077L4.772 5.79m14.456 0a48.108 48.108 0 0 0-3.478-.397m-12 .562c.34-.059.68-.114 1.022-.165m0 0a48.11 48.11 0 0 1 3.478-.397m7.5 0v-.916c0-1.18-.91-2.164-2.09-2.201a51.964 51.964 0 0 0-3.32 0c-1.18.037-2.09 1.022-2.09 2.201v.916m7.5 0a48.667 48.667 0 0 0-7.5 0"/></svg>`,
        plus: `<?xml version="1.0"?><svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="2" stroke="currentColor" class="w-5 h-5"><path stroke-linecap="round" stroke-linejoin="round" d="M12 4.5v15m7.5-7.5h-15"/></svg>`,
    };

    function onPagination(newPagination) {
        onPaginationFilterChanged(filter, newPagination);
    }

    function onFilter(key, value) {
        if (!value) {
            delete filter[key];
            onPaginationFilterChanged(filter, paginationRef?.resetPagination());
            return;
        }
        filter[key] = value;

        onPaginationFilterChanged(filter, paginationRef?.resetPagination());
    }
</script>

<div class="table-outer-container">
    <div class="flex flex-wrap gap-3.5 items-baseline">
        <h1 class="table-header">{tableHeader}</h1>
        <div>
            {#each headerActions as headerAction}
                <button
                    title={headerAction.title}
                    class="header-action-button"
                    on:click={headerAction.onClick}
                >
                    {#if icons[headerAction.icon]}{@html icons[
                            headerAction.icon
                        ]}{:else}{headerAction.title}{/if}
                </button>
            {/each}
        </div>
    </div>
    <div class="table-container">
        <table>
            <thead>
                <tr>
                    {#each columns as column}
                        {#if column.filterbar}
                            <th>
                                <div class="searchable-column">
                                    <span>{column.header}</span>
                                    <div class="search-wrapper">
                                        <!-- Search Icon -->
                                        <div class="search-icon">
                                            <svg
                                                aria-hidden="true"
                                                xmlns="http://www.w3.org/2000/svg"
                                                fill="none"
                                                viewBox="0 0 20 20"
                                            >
                                                <path
                                                    stroke="currentColor"
                                                    stroke-linecap="round"
                                                    stroke-linejoin="round"
                                                    stroke-width="3"
                                                    d="m19 19-4-4m0-7A7 7 0 1 1 1 8a7 7 0 0 1 14 0Z"
                                                />
                                            </svg>
                                        </div>
                                        <input
                                            type="search"
                                            id="search"
                                            class="search-input"
                                            placeholder="Search"
                                            on:input={(event) =>
                                                onFilter(
                                                    column.key,
                                                    event?.target?.value,
                                                )}
                                        />
                                    </div>
                                </div>
                            </th>
                        {:else}
                            <th>{column.header}</th>
                        {/if}
                    {/each}
                    {#if actions && actions.length}
                        <th>Actions</th>
                    {/if}
                </tr>
            </thead>
            <tbody>
                {#if rows}
                    {#each rows as row, index}
                        <tr>
                            {#each Object.keys(row) as rowKey}
                                {#if Array.isArray(row[rowKey])}
                                    <td>
                                        {#each row[rowKey] as subRow}
                                            <div>
                                                {subRow}
                                            </div>
                                        {/each}
                                    </td>
                                {:else}
                                    <td>{row[rowKey]}</td>
                                {/if}
                            {/each}
                            {#if actions && actions.length}
                                <td>
                                    {#each actions as action}
                                        <button
                                            title={action.title}
                                            class="table-action-button"
                                            on:click={() => {
                                                if (
                                                    action.requireConfirmation
                                                ) {
                                                    displayModal();
                                                    pendingAction = () =>
                                                        action.onClick(index);
                                                    return;
                                                }
                                                action.onClick(index);
                                            }}
                                        >
                                            {#if icons[action.icon]}{@html icons[
                                                    action.icon
                                                ]}{:else}{action.title}{/if}
                                        </button>
                                    {/each}
                                </td>
                            {/if}
                        </tr>
                    {/each}
                {:else}
                    <div class="no-data">No data...</div>
                {/if}
            </tbody>
        </table>
    </div>
    {#if total > 0}
        <PaginationComponent
            {total}
            onPaginationChanged={(pagination) => onPagination(pagination)}
            bind:this={paginationRef}
        />
    {/if}
</div>

{#if showModal}
    <Modal {onModalConfirmed} {onModalCanceled} />
{/if}

<style>
    .table-header {
        font-size: 1.8rem;
        font-weight: bold;
        margin-bottom: 1rem;
        color: var(--primary-color);
        cursor: default;
    }

    .table-outer-container {
        max-width: 1536px;
        margin: 2rem auto;
    }

    .table-container {
        position: relative;
        overflow-x: auto;
        border-radius: 0.5rem;
        cursor: default;
    }

    .table-action-button {
        cursor: pointer;
        padding: 0.5rem;
        margin-left: 0.1rem;
        margin-right: 0.1rem;
        color: white;
        background-color: var(--primary-color);
        border-radius: 0.375rem;
        transition: background-color 0.2s ease-in-out;
    }

    .table-action-button:hover,
    .header-action-button:hover {
        background-color: var(--primary-color-dark);
    }

    .header-action-button {
        cursor: pointer;
        padding: 0.6rem 0.6rem;
        color: white;
        background-color: var(--primary-color);
        border-radius: 1rem;
        box-shadow: 0 4px 12px rgba(16, 185, 129, 0.3);
        transition:
            background-color 0.3s ease,
            transform 0.2s ease,
            box-shadow 0.3s ease;
    }

    .header-action-button:hover {
        background-color: var(--primary-color-dark);
        transform: translateY(-2px) scale(1.02);
        box-shadow: 0 6px 16px rgba(4, 120, 87, 0.4);
    }

    table {
        /* table-layout: fixed; */
        width: 100%;
        font-size: 0.875rem;
        color: #374151;
        box-shadow: 0 10px 10px rgba(0, 0, 0, 0.1);
    }

    thead {
        font-size: 1.125rem;
        font-weight: bold;
        text-transform: uppercase;
        background-color: #f3f4f6;
    }

    thead tr th {
        padding-left: 1.5rem;
        padding-right: 1.5rem;
        padding-top: 0.75rem;
        padding-bottom: 0.75rem;
    }

    tbody tr {
        border-bottom: 1px solid #e5e7eb;
        cursor: default;
    }

    tbody tr:hover {
        background-color: #f9fafb;
    }

    tbody tr td {
        padding-left: 1.5rem;
        padding-right: 1.5rem;
        padding-top: 1rem;
        padding-bottom: 1rem;
        text-align: center;
        font-weight: 500;
        color: #4b5563;
        white-space: nowrap;
    }

    .searchable-column {
        display: flex;
        flex-direction: column;
        align-items: center;
    }

    .searchable-column .search-wrapper {
        position: relative;
    }

    .searchable-column .search-icon {
        position: absolute;
        top: 0;
        bottom: 0;
        left: 0.75rem;
        display: flex;
        align-items: center;
        pointer-events: none;
    }

    .searchable-column .search-icon svg {
        width: 1rem;
        height: 1rem;
        color: var(--primary-color-dark);
        transition: color 0.3s ease-in-out;
    }

    .searchable-column .search-input {
        width: auto;
        max-width: 7.5rem;
        padding-left: 2.3rem;
        padding-right: 0.75rem;
        padding-top: 0.375rem;
        padding-bottom: 0.375rem;
        font-size: 0.875rem;
        color: #374151;
        border: 1px solid #e5e7eb;
        border-radius: 0.5rem;
        background-color: #f9fafb;
        transition:
            border-color 0.3s ease-in-out,
            box-shadow 0.2s ease-in-out;
    }

    .searchable-column .search-input:focus {
        outline: none;
        border-color: var(--primary-color-dark);
        box-shadow: 0 0 0 2px rgba(4, 120, 87, 0.3);
    }

    .searchable-column .search-input::placeholder {
        letter-spacing: 0.05em;
    }

    .no-data {
        margin: 0.7rem 2rem;
        font-weight: bold;
        font-size: 1rem;
        letter-spacing: 0.1rem;
    }
</style>
