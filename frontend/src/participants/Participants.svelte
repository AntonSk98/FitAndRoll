<script>
    import TableComponent from "../table/TableComponent.svelte";
    import {
        FindParticipants,
        ArchiveParticipant,
    } from "../../wailsjs/go/participants/ParticipantsController.js";
    import { toastError, toastSuccess } from "../toast/toastStore.js";
    import ParticipantForm from "./ParticipantForm.svelte";
    import MemberCard from "./MemberCard.svelte";

    const PARTICIPANTS_OVERVIEW = "participants_overview";
    const ADD_NEW_PARTICIPANT = "new_participant";
    const UPDATE_PARTICIPANT = "update_participant";
    const MEMBER_CARD_OVERVIEW = "member_card_overview";

    let componentRender = PARTICIPANTS_OVERVIEW;

    let participantsPage;
    let selectedParticipant;

    $: if (componentRender === PARTICIPANTS_OVERVIEW) {
        findParticipants();
    }

    function findParticipants(filter, pagination) {
        FindParticipants(filter, pagination ?? { currentPage: 1, pageSize: 5 })
            .then((page) => (participantsPage = page))
            .catch((err) => {
                console.error(err);
                toastError();
            });
    }

    function displayDetails(index) {
        selectedParticipant = toSelectedUser(index);
        componentRender = UPDATE_PARTICIPANT;
    }

    function newParticipant() {
        componentRender = ADD_NEW_PARTICIPANT;
    }

    function displayMemberCardOverview(index) {
        selectedParticipant = toSelectedUser(index);
        componentRender = MEMBER_CARD_OVERVIEW;
    }

    function archiveParticipant(index) {
        selectedParticipant = toSelectedUser(index);
        ArchiveParticipant(selectedParticipant?.id)
            .then(() => {
                toastSuccess();
                findParticipants();
            })
            .catch((error) => {
                console.error(error.message);
                toastError();
            });
    }

    function onPaginationFilterChanged(filter, pagination) {
        findParticipants(filter, pagination);
    }

    function toSelectedUser(index) {
        const participant = participantsPage.data[index];
        if (!participant) {
            toastError();
            throw new Error(
                `Invalid participant: No participant found at index ${index}`,
            );
            return;
        }

        return participant;
    }
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
            },
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
                onClick: displayMemberCardOverview,
            },
            {
                title: "trash",
                icon: "trash",
                requireConfirmation: true,
                onClick: archiveParticipant,
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
{#if componentRender === ADD_NEW_PARTICIPANT}
    <ParticipantForm
        backToOverview={() => (componentRender = PARTICIPANTS_OVERVIEW)}
    />
{/if}

{#if componentRender === UPDATE_PARTICIPANT}
    <ParticipantForm
        participantId={selectedParticipant?.id}
        backToOverview={() => (componentRender = PARTICIPANTS_OVERVIEW)}
    />
{/if}

{#if componentRender === MEMBER_CARD_OVERVIEW}
    <MemberCard
        member={{
            id: selectedParticipant.id,
            name: selectedParticipant.name,
            surname: selectedParticipant.surname,
        }}
        onBackButtonClicked={() => (componentRender = PARTICIPANTS_OVERVIEW)}
    />
{/if}
