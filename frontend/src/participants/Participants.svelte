<script>
    import TableComponent from "../common/TableComponent.svelte";
    import {
        FindParticipants,
        ArchiveParticipant,
    } from "../../wailsjs/go/participants/ParticipantsHandler";
    import { ImportParticipants } from "../../wailsjs/go/import_/ImportParticipantsHandler";
    import { toastError, toastSuccess } from "../toast/toastStore.js";
    import ParticipantForm from "./ParticipantForm.svelte";
    import MemberCard from "./MemberCard.svelte";
    import { i18n } from "../common/i18n";

    import { LogError } from "../../wailsjs/runtime/runtime";

    const PARTICIPANTS_OVERVIEW = "participants_overview";
    const ADD_NEW_PARTICIPANT = "new_participant";
    const UPDATE_PARTICIPANT = "update_participant";
    const MEMBER_CARD_OVERVIEW = "member_card_overview";

    let componentRender = PARTICIPANTS_OVERVIEW;

    let participantsPage;
    let selectedParticipant;
    let tableRef;

    $: if (componentRender === PARTICIPANTS_OVERVIEW) {
        findParticipants();
    }

    function findParticipants(filter, pagination) {
        FindParticipants(filter, pagination ?? { currentPage: 1, pageSize: 5 })
            .then((page) => (participantsPage = page))
            .catch((err) => {
                LogError(
                    `Error while fetching courses. Filter: ${filter}. Pagination: ${pagination}. Error: ${err}`,
                );
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
                findParticipants(tableRef?.getFilter());
            })
            .catch((error) => {
                LogError(
                    `Error while archiving a participant. Participant details: ${selectedParticipant}. Error: ${error}`,
                );
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

    function importParticipants() {
        const CODE_TO_MESSAGE = {
            SELECT_FILE_ERROR: i18n("participants.import.selectFileError"),
            OPEN_FILE_ERROR: i18n("participants.import.openFileError"),
            NO_SHEET_ERROR: i18n("participants.import.noSheetError"),
            GET_ROWS_ERROR: i18n("participants.import.getRowsError"),
            REQUIRED_COLUMNS_ERROR: i18n(
                "participants.import.requiredColumnsError",
            ),
            REQUIRED_PARAMETER_MISSING_ERROR: i18n(
                "participants.import.requiredParameterMissingError",
            ),
            TRAINING_START_PARSE_ERROR: i18n(
                "participants.import.trainingStartParseError",
            ),
            BIRTHDAY_PARSE_ERROR: i18n(
                "participants.import.birthdayParseError",
            ),
            PRIVACY_POLICY_ACCEPTED_AT_PARSE_ERROR: i18n(
                "participants.import.privacyPolicyAcceptedAtParseError",
            ),
            PRIVACY_POLICY_ACCEPTED_AT_REQUIRED_ERROR: i18n(
                "participants.import.privacyPolicyAcceptedAtRequiredError",
            ),
            NO_DATA_ERROR: i18n("participants.import.noDataError"),
            DEFAULT_ERROR: i18n("participants.import.defaultError"),
        };

        ImportParticipants()
            .then(() => {
                toastSuccess();
                findParticipants(tableRef?.getFilter());
            })
            .catch((error) => {
                LogError(`Error while importing participants. Error: ${error}`);
                const jsonError = JSON.parse(error);
                toastError(
                    `${jsonError?.id ? `${jsonError.id}: ` : ""}${CODE_TO_MESSAGE[jsonError?.code] ?? CODE_TO_MESSAGE.DEFAULT_ERROR}`,
                );
            });
    }
</script>

{#if componentRender === PARTICIPANTS_OVERVIEW}
    <TableComponent
        bind:this={tableRef}
        tableHeader={i18n("participants.header")}
        total={participantsPage?.total ?? 0}
        mainFilters={[
            {
                key: "withActiveCard",
                label: i18n("participants.filters.withActiveCard"),
                type: "checkbox",
            },
        ]}
        columns={[
            {
                key: "name",
                header: i18n("participants.name"),
                filterbar: true,
            },
            {
                key: "surname",
                header: i18n("participants.surname"),
                filterbar: true,
            },
            {
                key: "group",
                header: i18n("participants.group"),
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
                title: i18n("participants.actions.edit"),
                icon: "edit",
                onClick: displayDetails,
            },
            {
                title: i18n("participants.actions.card"),
                icon: "card",
                onClick: displayMemberCardOverview,
            },
            {
                title: i18n("participants.actions.archive"),
                icon: "trash",
                requireConfirmation: true,
                onClick: archiveParticipant,
            },
        ]}
        headerActions={[
            {
                title: i18n("participants.actions.create"),
                icon: "plus",
                onClick: newParticipant,
            },
            {
                title: i18n("participants.actions.import"),
                icon: "documentArrowUp",
                onClick: importParticipants,
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
