<script>
    import BackButton from "../common/BackButton.svelte";
    import { FindParticipants } from "../../wailsjs/go/participants/ParticipantsHandler";
    import { toastError, toastSuccess } from "../toast/toastStore.js";
    import { onMount } from "svelte";
    import TableComponent from "../common/TableComponent.svelte";
    import { FindActiveMemberCards } from "../../wailsjs/go/membercardattendance/MemberCardAttendanceHandler";
    import AttendCourseModal from "./AttendCourseModal.svelte";
    import { i18n } from "../common/i18n";

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
                attendCourseType = "WITH_MEMBER_CARD";
            })
            .catch((err) => {
                console.error(err);
                toastError();
            });
    }

    function trialAttend(index) {
        selectedParticipant = participants[index];
        attendCourseType = "TRIAL_ATTENDANCE";
    }

    function attendWithoutMemberCard(index) {
        selectedParticipant = participants[index];
        attendCourseType = "WITHOUT_MEMBER_CARD";
    }

    function onAttendCourseModalDestroyed() {
        activeMemberCards = [];
        attendCourseType = null;
    }
</script>

<TableComponent
    tableHeader = {`${i18n('courseParticipant.header')} '${selectedCourse?.name}'`}
    total={total ?? 0}
    columns={[
        {
            key: "name",
            header: i18n("courseParticipant.name"),
            filterbar: true,
        },
        {
            key: "surname",
            header: i18n("courseParticipant.surname"),
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
            title: i18n("courseParticipant.withMemberCard"),
            icon: "checkCircle",
            onClick: attendWithMemberCard,
        },
        {
            title: i18n("courseParticipant.trialAttendance"),
            icon: "exclamationCircle",
            onClick: trialAttend,
        },
        {
            title: i18n("courseParticipant.withoutMemberCard"),
            icon: "xCircle",
            onClick: attendWithoutMemberCard,
        },
    ]}
     headerActions={[
        {
            title: i18n("courseParticipant.close"),
            icon: "xMark",
            onClick: () => returnToCourseOverview(),
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
