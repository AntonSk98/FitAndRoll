<script>
    import BackButton from "../common/BackButton.svelte";
    import { FindParticipants } from "../../wailsjs/go/participants/ParticipantsController.js";
    import { toastError, toastSuccess } from "../toast/toastStore.js";
    import { onMount } from "svelte";
    import TableComponent from "../common/TableComponent.svelte";
    import { FindActiveMemberCards } from "../../wailsjs/go/membercardattendance/MemberCardAttendanceController";
    import AttendCourseModal from "./AttendCourseModal.svelte";

    export let selectedCourse;
    export let returnToCourseOverview;

    let participants = [];
    let total;

    let selectedParticipant;
    let activeMemberCards = [];

    let attendCourseType;

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

    function attendWithMemberCard(index) {
        FindActiveMemberCards(participants[index]?.id)
            .then((memberCards) => {
                selectedParticipant = participants[index];
                activeMemberCards = memberCards;
                attendCourseType = 'WITH_MEMBER_CARD';
            })
            .catch((err) => {
                console.error(err);
                toastError();
            });
    }

    function trialAttend(index) {
        selectedParticipant = participants[index];
        attendCourseType = 'TRIAL_ATTENDANCE';
    }

    function attendWithoutMemberCard(index) {
        selectedParticipant = participants[index];
        attendCourseType = 'WITHOUT_MEMBER_CARD';
    }

    function onAttendCourseModalDestroyed() {
        activeMemberCards = [];
        attendCourseType = null;
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
            onClick: attendWithMemberCard,
        },
        {
            title: "trial attend",
            icon: "exclamationCircle",
            onClick: trialAttend,
        },
        {
            title: "attend without member card",
            icon: "xCircle",
            onClick: attendWithoutMemberCard,
        },
    ]}
    {onPaginationFilterChanged}
/>

{#if attendCourseType}
    <AttendCourseModal
        attendanceType={attendCourseType}
        {selectedParticipant}
        {selectedCourse}
        {activeMemberCards}
        onDestroy={onAttendCourseModalDestroyed}
    />
{/if}

<style>
    .header {
        display: flex;
        align-items: end;
        gap: 2vw;
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
