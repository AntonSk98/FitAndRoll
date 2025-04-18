<script>
    import { onMount } from "svelte";
    import TableComponent from "../common/TableComponent.svelte";
    import { toastError, toastSuccess } from "../toast/toastStore.js";
    import Modal from "../common/Modal.svelte";
    import { i18n } from "../common/i18n";
    import {LogError} from "../../wailsjs/runtime/runtime"

    export let header;
    export let loadArchivedDataPromise;
    export let unarchiveEntryPromise;

    export let onComponentDestroyed;

    let archivedDataPage;
    let selectedArchivedEntry;
    let tableRef;

    onMount(() => loadArchivedData());

    function loadArchivedData(filter, pagination) {
        loadArchivedDataPromise(
            filter,
            pagination ?? { currentPage: 1, pageSize: 5 },
        )
            .then((page) => (archivedDataPage = page))
            .catch((err) => {
                LogError(`Error loading archived data. Filters: ${filter}. Pagination: ${pagination}. Error: ${err}`);
                toastError();
            });
    }

    function onPaginationFilterChanged(filter, pagination) {
        loadArchivedData(filter, pagination);
    }

    function unarchiveEntry(index) {
        if (!selectedArchivedEntry) {
            selectedArchivedEntry = archivedDataPage?.data[index];
            return;
        }

        unarchiveEntryPromise(selectedArchivedEntry?.id)
            .then(() => {
                loadArchivedData(tableRef?.getFilter());
                selectedArchivedEntry = null;
                toastSuccess();
            })
            .catch((err) => {
                LogError(`Failed to unarchive an entry. Entry details: ${selectedArchivedEntry}. Error: ${err}`);
                toastError();
            });
    }
</script>

<TableComponent
    bind:this={tableRef}
    tableHeader={header}
    total={archivedDataPage?.total ?? 0}
    columns={[
        {
            key: "name",
            header: i18n("archive.table.columns.name"),
            filterbar: true,
        },
        {
            key: "archivedAt",
            header: i18n("archive.table.columns.archivedAt"),
        },
    ]}
    actions={[
        {
            title: i18n("archive.table.actions.unarchive"),
            icon: "lockOpen",
            onClick: unarchiveEntry,
        },
    ]}
    rows={archivedDataPage?.data?.map((entry) => {
        return {
            name: entry.name,
            archivedAt: entry.archivedAt,
        };
    })}
    headerActions={[
        {
            title: i18n("close"),
            icon: "xMark",
            onClick: () => onComponentDestroyed(),
        },
    ]}
    {onPaginationFilterChanged}
/>

{#if selectedArchivedEntry}
    <Modal
        onModalCanceled={() => (selectedArchivedEntry = null)}
        onModalConfirmed={unarchiveEntry}
    >
        <div slot="body">
            <div>
                {i18n("archive.table.actions.unarchivePrePrompt")}
                <span class="font-bold text-[var(--primary-color)]"
                    >{selectedArchivedEntry?.name}</span
                >
                {i18n("archive.table.actions.unarchivePostPrompt")}
            </div>
        </div>
    </Modal>
{/if}
