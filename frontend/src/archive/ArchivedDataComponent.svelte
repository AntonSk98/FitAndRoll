<script>
    import { onMount } from "svelte";
    import TableComponent from "../common/TableComponent.svelte";
    import { toastError, toastSuccess } from "../toast/toastStore.js";
    import Modal from "../common/Modal.svelte";

    export let header;
    export let loadArchivedDataPromise;
    export let unarchiveEntryPromise;

    export let onComponentDestroyed;

    let archivedDataPage;
    let selectedArchivedEntry;

    onMount(() => loadArchivedData());

    function loadArchivedData(filter, pagination) {
        loadArchivedDataPromise(
            filter,
            pagination ?? { currentPage: 1, pageSize: 5 },
        )
            .then((page) => (archivedDataPage = page))
            .catch((err) => {
                console.error("Error loading archived data:", err);
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
                loadArchivedData();
                selectedArchivedEntry = null;
                toastSuccess();
            })
            .catch((err) => {
                console.error(err);
                toastError();
            });
    }
</script>

<TableComponent
    tableHeader="Archived participants overview"
    total={archivedDataPage?.total ?? 0}
    columns={[
        {
            key: "name",
            header: "Name",
            filterbar: true,
        },
        {
            key: "archivedAt",
            header: "Archived at",
        },
    ]}
    actions={[
        {
                title: 'Unarchive',
                icon: "lockOpen",
                onClick: unarchiveEntry
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
            title: "close",
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
        <div>Are you sure you would like to unarchive <span class="font-bold text-[var(--primary-color)]">{selectedArchivedEntry?.name}</span></div>
    </div>
</Modal>
{/if}