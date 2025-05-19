<script>
    import { i18n } from "./i18n";
    import { onDestroy, onMount } from "svelte";

    import Modal from "./Modal.svelte";
    import PaginationComponent from "./PaginationComponent.svelte";
    import DateRangePicker from "./DateRangePicker.svelte";

    import {IS_HOTKEY_DOWN_FUNCTION, IS_HOTKEY_UP_FUNCTION} from "./scale_hotkeys";

    export let tableHeader = null;
    export let total;

    export let mainFilters = [];
    export let columns = [];
    export let rows = [];
    export let actions = [];
    export let headerActions = [];
    export let onPaginationFilterChanged = (filter, pagination) => {};
    export let onRowClicked = null;

    let filter = {};
    let paginationRef;

    let pendingAction;
    let showModal = false;
    let onlyDateMainFilter = false;
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
        card: `<svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="2" stroke="currentColor" class="w-5 h-5"><path stroke-linecap="round" stroke-linejoin="round" d="M2.25 8.25h19.5M2.25 9h19.5m-16.5 5.25h6m-6 2.25h3m-3.75 3h15a2.25 2.25 0 0 0 2.25-2.25V6.75A2.25 2.25 0 0 0 19.5 4.5h-15a2.25 2.25 0 0 0-2.25 2.25v10.5A2.25 2.25 0 0 0 4.5 19.5Z" /></svg>`,
        checkCircle: `<svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="2" stroke="currentColor" class="w-5 h-5"><path stroke-linecap="round" stroke-linejoin="round" d="M9 12.75 11.25 15 15 9.75M21 12a9 9 0 1 1-18 0 9 9 0 0 1 18 0Z" /></svg>`,
        exclamationCircle: `<svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="2" stroke="currentColor" class="w-5 h-5"><path stroke-linecap="round" stroke-linejoin="round" d="M12 9v3.75m9-.75a9 9 0 1 1-18 0 9 9 0 0 1 18 0Zm-9 3.75h.008v.008H12v-.008Z" /></svg>`,
        xCircle: `<svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="2" stroke="currentColor" class="w-5 h-5"><path stroke-linecap="round" stroke-linejoin="round" d="m9.75 9.75 4.5 4.5m0-4.5-4.5 4.5M21 12a9 9 0 1 1-18 0 9 9 0 0 1 18 0Z" /></svg>`,
        userGroup: `<svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="2" stroke="currentColor" class="w-5 h-5"><path stroke-linecap="round" stroke-linejoin="round" d="M18 18.72a9.094 9.094 0 0 0 3.741-.479 3 3 0 0 0-4.682-2.72m.94 3.198.001.031c0 .225-.012.447-.037.666A11.944 11.944 0 0 1 12 21c-2.17 0-4.207-.576-5.963-1.584A6.062 6.062 0 0 1 6 18.719m12 0a5.971 5.971 0 0 0-.941-3.197m0 0A5.995 5.995 0 0 0 12 12.75a5.995 5.995 0 0 0-5.058 2.772m0 0a3 3 0 0 0-4.681 2.72 8.986 8.986 0 0 0 3.74.477m.94-3.197a5.971 5.971 0 0 0-.94 3.197M15 6.75a3 3 0 1 1-6 0 3 3 0 0 1 6 0Zm6 3a2.25 2.25 0 1 1-4.5 0 2.25 2.25 0 0 1 4.5 0Zm-13.5 0a2.25 2.25 0 1 1-4.5 0 2.25 2.25 0 0 1 4.5 0Z" /></svg>`,
        xMark: `<svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="2" stroke="currentColor" class="w-5 h-5"><path stroke-linecap="round" stroke-linejoin="round" d="M6 18 18 6M6 6l12 12" /></svg>`,
        handRaised: `<svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="2" stroke="currentColor" class="w-5 h-5"><path stroke-linecap="round" stroke-linejoin="round" d="M10.05 4.575a1.575 1.575 0 1 0-3.15 0v3m3.15-3v-1.5a1.575 1.575 0 0 1 3.15 0v1.5m-3.15 0 .075 5.925m3.075.75V4.575m0 0a1.575 1.575 0 0 1 3.15 0V15M6.9 7.575a1.575 1.575 0 1 0-3.15 0v8.175a6.75 6.75 0 0 0 6.75 6.75h2.018a5.25 5.25 0 0 0 3.712-1.538l1.732-1.732a5.25 5.25 0 0 0 1.538-3.712l.003-2.024a.668.668 0 0 1 .198-.471 1.575 1.575 0 1 0-2.228-2.228 3.818 3.818 0 0 0-1.12 2.687M6.9 7.575V12m6.27 4.318A4.49 4.49 0 0 1 16.35 15m.002 0h-.002" /></svg>`,
        lockOpen: `<svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="2" stroke="currentColor" class="w-5 h-5"><path stroke-linecap="round" stroke-linejoin="round" d="M13.5 10.5V6.75a4.5 4.5 0 1 1 9 0v3.75M3.75 21.75h10.5a2.25 2.25 0 0 0 2.25-2.25v-6.75a2.25 2.25 0 0 0-2.25-2.25H3.75a2.25 2.25 0 0 0-2.25 2.25v6.75a2.25 2.25 0 0 0 2.25 2.25Z" /></svg>`,
        documentArrowUp: `<svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="2" stroke="currentColor" class="w-5 h-5"><path stroke-linecap="round" stroke-linejoin="round" d="M19.5 14.25v-2.625a3.375 3.375 0 0 0-3.375-3.375h-1.5A1.125 1.125 0 0 1 13.5 7.125v-1.5a3.375 3.375 0 0 0-3.375-3.375H8.25m6.75 12-3-3m0 0-3 3m3-3v6m-1.5-15H5.625c-.621 0-1.125.504-1.125 1.125v17.25c0 .621.504 1.125 1.125 1.125h12.75c.621 0 1.125-.504 1.125-1.125V11.25a9 9 0 0 0-9-9Z" /></svg>`
    };

    function onPagination(newPagination) {
        onPaginationFilterChanged(filter, newPagination);
        resizeTableresetTableScrollIfNeeded();
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

    function setOnlyDateMainFilter() {
        onlyDateMainFilter = mainFilters?.every(
            (filter) => filter.type === "date",
        );
    }

    function resizeTableresetTableScrollIfNeeded() {
        const tableContainer = document.querySelector(".table-container");
        if (
            tableContainer &&
            tableContainer.scrollWidth >= tableContainer.clientWidth
        ) {
            tableContainer.scrollLeft = 0;
        }
    }

    export function getFilter() {
        return filter;
    }

    onMount(() => {
        setOnlyDateMainFilter();
        window.onresize = () => resizeTableresetTableScrollIfNeeded();
        window.addEventListener("keydown", ($event) => {
            if (
                IS_HOTKEY_UP_FUNCTION($event) ||
                IS_HOTKEY_DOWN_FUNCTION($event)
            ) {
                resizeTableresetTableScrollIfNeeded();
            }
        });
    });

    onDestroy(() => {
        window.onresize = null;
        window.removeEventListener(
            "keydown",
            resizeTableresetTableScrollIfNeeded,
        );
    });
</script>

<div class="table-outer-container">
    <div class="flex gap-x-3.5 items-center">
        {#if tableHeader}
            <h1 class="table-header">{tableHeader}</h1>
        {/if}
        <div class="flex gap-1.5">
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
    <div class="main-filters mb-2 mt-2">
        <div class="flex gap-x-10 items-end justify-between flex-wrap">
            {#if !onlyDateMainFilter}
                <div class="flex flex-col">
                    {#each mainFilters as mainFilter}
                        {#if mainFilter.type === "checkbox"}
                            <div
                                class="flex items-center gap-1 text-[var(--secondary-color-dark)] font-semibold tracking-wide text-base"
                            >
                                <input
                                    type="checkbox"
                                    class="w-5 h-5 cursor-pointer transition-all duration-300"
                                    on:input={(event) =>
                                        onFilter(
                                            mainFilter.key,
                                            event?.target?.checked,
                                        )}
                                />
                                <label>
                                    <span>{mainFilter.label}</span>
                                </label>
                            </div>
                        {/if}
                    {/each}
                </div>
            {/if}

            <!-- Right side (Date Pickers) -->
            <div class="flex flex-col">
                {#each mainFilters as mainFilter}
                    {#if mainFilter.type === "date"}
                        <div
                            class="flex items-end gap-2 text-[var(--secondary-color-dark)] italic font-semibold tracking-wide"
                        >
                            <span>{mainFilter.label}:</span>
                            <DateRangePicker
                                align={!onlyDateMainFilter ? "right" : "left"}
                                onDateRangePicked={(date) =>
                                    onFilter(mainFilter.key, date)}
                            />
                        </div>
                    {/if}
                {/each}
            </div>
        </div>
    </div>
    <div class="table-container shadow-lg">
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
                                            placeholder={i18n(
                                                "table.filter.search",
                                            )}
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
                        <th>{i18n("table.columns.action")}</th>
                    {/if}
                </tr>
            </thead>
            <tbody>
                {#if rows}
                    {#each rows as row, index}
                        <tr
                            class={onRowClicked ? "!cursor-pointer" : ""}
                            on:click={() => onRowClicked?.(index)}
                        >
                            {#each Object.keys(row) as rowKey}
                                {#if Array.isArray(row[rowKey])}
                                    <td>
                                        {#each row[rowKey] as subRow}
                                            <div>
                                                {subRow}
                                            </div>
                                        {/each}
                                    </td>
                                {:else if row[rowKey]?.["html"]}
                                    <td>{@html row[rowKey]?.["html"]}</td>
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
                    <tr
                        ><td
                            colspan={columns?.length + 1}
                            class="no-data text-left!"
                            >{i18n("table.noData")}</td
                        ></tr
                    >
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
        color: var(--primary-color);
        cursor: default;
    }

    .table-outer-container {
        max-width: var(--container-max-width);
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
        transition: background-color 0.3s ease-in-out;
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
        transition:
            background-color 0.5s ease,
            transform 0.5s ease,
            box-shadow 0.5s ease;
    }

    .header-action-button:hover {
        background-color: var(--primary-color-dark);
        transform: translateY(-2px) scale(1.02);
    }

    table {
        width: 100%;
        font-size: 0.875rem;
        color: var(--text-color-darker);
    }

    thead {
        font-size: 1.125rem;
        font-weight: bold;
        text-transform: uppercase;
        background-color: var(--secondary-color);
    }

    thead tr th {
        padding-left: 1.5rem;
        padding-right: 1.5rem;
        padding-top: 0.75rem;
        padding-bottom: 0.75rem;
    }

    tbody tr {
        border-bottom: 1px solid var(--secondary-color-darker);
        cursor: default;
    }

    tbody tr:hover {
        background-color: var(--secondary-color-light);
    }

    tbody tr td {
        padding-left: 1.5rem;
        padding-right: 1.5rem;
        padding-top: 1rem;
        padding-bottom: 1rem;
        text-align: center;
        font-weight: 500;
        color: var(--text-color);
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
        color: var(--text-color-darker);
        border: 1px solid var(--secondary-color-darker);
        border-radius: 0.5rem;
        background-color: var(--secondary-color-light);
        transition:
            border-color 0.3s ease-in-out,
            box-shadow 0.2s ease-in-out;
    }

    .searchable-column .search-input:focus {
        outline: none;
        border-color: var(--primary-color-dark);
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
