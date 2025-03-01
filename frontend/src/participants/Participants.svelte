<script>
    import TableComponent from "../table/TableComponent.svelte";
    import { FindParticipants } from "../../wailsjs/go/participants/ParticipantsController.js";
    import { toastError, toastSuccess } from "../toast/toastStore.js";
    import { onMount } from "svelte";

    const PARTICIPANTS_OVERVIEW = "participants_overview";
    const ADD_NEW_PARTICIPANT = "new_participant";
    const UPDATE_PARTICIPANT = "update_participant";
    const MEMBER_CARD_OVERVIEW = "member_card_overview";

    let componentRender = PARTICIPANTS_OVERVIEW;

    let participantsPage;
    let findParticipantsParams;

    function findParticipants(filter, pagination) {
        FindParticipants(filter, pagination ?? { currentPage: 1, pageSize: 5 })
            .then((page) => (participantsPage = page))
            .catch((err) => {
                console.error(err);
                toastError();
            });
    }

    function displayDetails(index) {
        console.log(index);
        componentRender = UPDATE_PARTICIPANT;
    }

    function newParticipant() {
        componentRender = ADD_NEW_PARTICIPANT;
    }

    function displayMemberCardOverview(index) {
        componentRender = MEMBER_CARD_OVERVIEW;
    }

    function archiveCourse(index) {
        console.log(index);
    }

    function onPaginationFilterChanged(filter, pagination) {
        findParticipants(filter, pagination);
    }

    onMount(() => findParticipants());
</script>

{#if componentRender === PARTICIPANTS_OVERVIEW}
    <TableComponent
        tableHeader={"Participants overview"}
        total={participantsPage?.total ?? 0}
        mainFilters={[
            {
                key: "withActiveCard",
                label: "With active member cards",
                type: "checkbox",
            }
        ]}
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
            {
                key: "group",
                header: "group",
                filterbar: true,
            },
        ]}
        rows={participantsPage?.data?.map((participant) => {
            return {
                name: participant.name,
                surname: participant.surname,
                group: participant.group,
            };
        })}
        actions={[
            {
                title: "edit",
                icon: "edit",
                onClick: displayDetails,
            },
            {
                title: "card",
                icon: "card",
                onclick: displayMemberCardOverview,
            },
            {
                title: "trash",
                icon: "trash",
                requireConfirmation: true,
                onClick: archiveCourse,
            },
        ]}
        headerActions={[
            {
                title: "plus",
                icon: "plus",
                onClick: newParticipant,
            },
        ]}
        {onPaginationFilterChanged}
    />
{/if}
{#if componentRender === ADD_NEW_PARTICIPANT}{/if}

{#if componentRender === UPDATE_PARTICIPANT}{/if}

{#if componentRender === MEMBER_CARD_OVERVIEW}{/if}
