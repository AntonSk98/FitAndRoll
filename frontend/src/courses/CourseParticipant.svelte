<script>
    import BackButton from "../common/BackButton.svelte";
    import { FindParticipants } from "../../wailsjs/go/participants/ParticipantsController.js";
    import { toastError, toastSuccess } from "../toast/toastStore.js";
    import { onDestroy, onMount } from "svelte";
    import TableComponent from "../table/TableComponent.svelte";
    import { FindActiveMemberCards } from "../../wailsjs/go/participants/MemberCardCourseParticipationController";
    import VisitCourseWithMemberCard from "./VisitCourseWithMemberCard.svelte";

    export let selectedCourse;
    export let returnToCourseOverview;

    let participants = [];
    let total;

    let selectedParticipant;
    let activeMemberCards = [];
    let showOnVisitMemberCardModal;

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
        findParticipants(filter, pagination);
    }

    function visitWithMemberCard(index) {
        FindActiveMemberCards(participants[index]?.id)
            .then((memberCards) => {
                selectedParticipant = participants[index];
                activeMemberCards = memberCards;
                showOnVisitMemberCardModal = true;
            })
            .catch((err) => {
                console.error(err);
                toastError();
            });
    }

    function visitAsTrial(index) {
        console.log("trial visit!", participants[index]);
    }

    function visitWithoutMemberCard(index) {
        console.log("visit without member card!", participants[index]);
    }

    function onVisitMemberCardModalDestroyed() {
        activeMemberCards = [];
        showOnVisitMemberCardModal = false;
    }
</script>

<div class="header">
    <div>
        <div class="title">{selectedCourse.name}</div>
        <div class="description">
            Here should be some description. Here should be some description.
            Here should be some description. Here should be some description.
            Here should be some description. Here should be some description.
            Here should be some description. Here should be some description.
            Here should be some description. Here should be some
            description.Here should be some description. Here should be some
            description.Here should be some description. Here should be some
            description.Here should be some description.Here should be some
            description.Here should be some description.
        </div>
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

{#if showOnVisitMemberCardModal}
    <VisitCourseWithMemberCard
        {selectedParticipant}
        {selectedCourse}
        {activeMemberCards}
        onDestroy={onVisitMemberCardModalDestroyed}
    />
{/if}

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
