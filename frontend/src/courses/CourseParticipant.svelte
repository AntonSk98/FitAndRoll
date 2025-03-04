<script>
    import BackButton from "../common/BackButton.svelte";
    import { FindParticipants } from "../../wailsjs/go/participants/ParticipantsController.js";
    import { toastError, toastSuccess } from "../toast/toastStore.js";
    import { onMount } from "svelte";
    import TableComponent from "../table/TableComponent.svelte";

    export let selectedCourse;
    export let returnToCourseOverview;

    let participants = [];
    let total;

    onMount(() => findParticipants());

    function findParticipants(filter, pagination) {
        FindParticipants(filter, pagination ?? { currentPage: 1, pageSize: 5 })
            .then((page) => {
                participants = page.data;
                total = page.total;
            })
            .catch((err) => {
                console.error(err);
                toastError();
            });
    }

    function onPaginationFilterChanged(filter, pagination) {
        console.log(filter);
        findParticipants(filter, pagination);
    }

    function visitWithMemberCard(index) {
        console.log('member card visit!', participants[index]);
    }

    function visitAsTrial(index) {
        console.log('trial visit!', participants[index]);
    }

    function visitWithoutMemberCard(index) {
        console.log('visit without member card!', participants[index]);

    }
</script>

<div class="header">
    <div>
        <div class="title">{selectedCourse.name}</div>
        <div class="description">Here should be some description. Here should be some description. Here should be some description. Here should be some description. Here should be some description. Here should be some description. Here should be some description. Here should be some description. Here should be some description. Here should be some description.Here should be some description. Here should be some description.Here should be some description. Here should be some description.Here should be some description.Here should be some description.Here should be some description.</div>
    </div>
    <BackButton onBackButtonClicked={returnToCourseOverview} />
</div>

<TableComponent
    total={total ?? 0}
    columns={[
        {
            key: "name",
            header: "name",
            filterbar: true,
        },
        {
            key: "surname",
            header: "surname",
            filterbar: true,
        },
    ]}
    rows={participants?.map((participant) => {
        return {
            name: participant.name,
            surname: participant.surname,
        };
    })}
    actions={[
        {
            title: "attend with member card",
            icon: "checkCircle",
            requireConfirmation: true,
            onClick: visitWithMemberCard,
        },
        {
            title: "trial attend",
            icon: "exclamationCircle",
            requireConfirmation: true,
            onClick: visitAsTrial,
        },
        {
            title: "attend without member card",
            icon: "xCircle",
            requireConfirmation: true,
            onClick: visitWithoutMemberCard,
        },
    ]}
    {onPaginationFilterChanged}
/>

<style>
    .header {
        display: flex;
        align-items: end;
        gap: 5rem;
        justify-content: space-between;
        max-width: var(--container-max-width);
        margin: 0 auto;
        cursor: default;
    }

    .title {
        font-size: 2.5rem;
        font-weight: bold;
        color: var(--secondary-color-dark);
    }

    .description {
        font-weight: bolder;
        color: var(--text-color);
        letter-spacing: 0.1rem;
        text-align: justify;
    }
</style>
