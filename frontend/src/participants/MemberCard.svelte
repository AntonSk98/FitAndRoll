<script>
    import Modal from "../common/Modal.svelte";
    import TableComponent from "../common/TableComponent.svelte";
    import { toastError, toastSuccess } from "../toast/toastStore";
    import {
        FindAllMemberCards,
        IssueNewMemberCard,
        UndoIssueNewMemberCard,
    } from "../../wailsjs/go/participants/MemberCardHandler";

    import { LoadMemberCardCourseHistory } from "../../wailsjs/go/membercardattendance/MemberCardAttendanceHandler";
    import BackButton from "../common/BackButton.svelte";
    import { i18n } from "../common/i18n";

    import {LogError} from "../../wailsjs/runtime/runtime"

    export let member;
    export let onBackButtonClicked;

    let displayMemberCardParticipationHistory = false;
    let displayNewCardConfirmationDialog = false;

    let memberCards = [];

    let memberCardParticipationHistory = [];
    let selectedMemberCardId;

    $: if (
        !displayMemberCardParticipationHistory &&
        !displayNewCardConfirmationDialog
    ) {
        findAllMemberCards();
    }

    function findAllMemberCards() {
        FindAllMemberCards(member.id)
            .then((cards) => (memberCards = cards || []))
            .catch((error) => {
                LogError(`Error occurred while fetching all member cards: ${error}`)
                toastError();
            });
    }

    function deleteEmptyMemberCard() {
        UndoIssueNewMemberCard(member.id, selectedMemberCardId)
            .then(() => {
                findAllMemberCards();
                displayMemberCardParticipationHistory = false;
                toastSuccess();
            })
            .catch((error) => {
                LogError(`Failed to delete an empty member card. Member: ${member}. Card id: ${selectedMemberCardId}. Error: ${error}`)
                toastError();
            });
    }

    function onMemberCardClicked(index) {
        selectedMemberCardId = memberCards[index]?.id;
        LoadMemberCardCourseHistory(member.id, selectedMemberCardId)
            .then((history) => {
                memberCardParticipationHistory = history;
                displayMemberCardParticipationHistory = true;
            })
            .catch((err) => {
                LogError(`Failed to fetch member card participation history. Member: ${member}. Card id: ${selectedMemberCardId}. Error: ${err}`)
                toastError();
            });
    }

    function issueNewMemberCard() {
        IssueNewMemberCard(member.id)
            .then(() => {
                displayNewCardConfirmationDialog = false;
                findAllMemberCards();
                toastSuccess();
            })
            .catch((error) => {
                LogError(`Failed to issue a new member card: Member: ${member}. Error: ${error}`)
                toastError();
            });
    }
</script>

<div class="member-card-header">
    <div class="flex items-center gap-3">
        <span class="holder">{member.name} {member.surname}</span>
        <BackButton {onBackButtonClicked} />
    </div>
    <button
        class="primary"
        on:click={() => (displayNewCardConfirmationDialog = true)}
        >{i18n("participants.memberCard.newCard")}</button
    >
</div>

{#if displayNewCardConfirmationDialog}
    <Modal
        onModalConfirmed={issueNewMemberCard}
        onModalCanceled={() => (displayNewCardConfirmationDialog = false)}
    >
        <div slot="body">
            <div>
                {i18n("participants.memberCard.issuePrefix")}
                <b class="text-[var(--primary-color)]"
                    >{member.name} {member.surname}</b
                >
                {i18n("participants.memberCard.issuePostfix")}.
            </div>
        </div>
        <div slot="confirm">
            {i18n("participants.memberCard.issueConfirmed")}
        </div>
        <div slot="cancel">{i18n("participants.memberCard.issueCanceled")}</div>
    </Modal>
{/if}

{#if displayMemberCardParticipationHistory}
    <Modal
        onModalConfirmed={() => {
            if (memberCardParticipationHistory?.length > 0) {
                displayMemberCardParticipationHistory = false;
                return;
            }
            deleteEmptyMemberCard();
        }}
        onModalCanceled={!memberCardParticipationHistory ||
        memberCardParticipationHistory?.length === 0
            ? () => (displayMemberCardParticipationHistory = false)
            : null}
    >
        <svg
            slot="icon"
            xmlns="http://www.w3.org/2000/svg"
            fill="none"
            viewBox="0 0 24 24"
            stroke-width="2"
            stroke="currentColor"
        >
            <path
                stroke-linecap="round"
                stroke-linejoin="round"
                d="M12 6v6h4.5m4.5 0a9 9 0 1 1-18 0 9 9 0 0 1 18 0Z"
            />
        </svg>

        <div slot="body" class="text-left">
            {#if memberCardParticipationHistory?.length > 0}
                <div class="member-card-overview-header">
                    {i18n("participants.memberCard.cardHistory")}
                </div>
                <ul class="px-3 py-2 flex flex-col gap-3">
                    {#each memberCardParticipationHistory as memberCardParticipationEntry}
                        <li class="p-2 bg-[var(--secondary-color)] rounded-xl">
                            <span class="text-[var(--primary-color)] text-lg"
                                >📌</span
                            >
                            <span class="font-medium"
                                >{memberCardParticipationEntry.course}</span
                            >
                            <span
                                class="font-bold text-xl text-[var(--primary-color)]"
                                >|</span
                            >
                            <span class="text-gray-600 text-sm"
                                >{memberCardParticipationEntry.attendedAt}</span
                            >
                        </li>
                    {/each}
                </ul>
            {:else}
                <div class="font-bold">
                    {i18n("participants.memberCard.removeEmptyCardHeader")}
                </div>
                <div>
                    {i18n("participants.memberCard.removeEmptyCard")}.
                </div>
            {/if}
        </div>
        <div slot="confirm">
            {#if memberCardParticipationHistory?.length > 0}
                {i18n("participants.memberCard.ok")}
            {:else}
                {i18n("participants.memberCard.removeEmptyCardConfirm")}
            {/if}
        </div>
        <div slot="cancel">{i18n("participants.memberCard.close")}</div>
    </Modal>
{/if}

<TableComponent
    tableHeader={i18n("participants.memberCard.header")}
    total="0"
    onRowClicked={onMemberCardClicked}
    columns={[
        {
            key: "index",
            header: i18n("participants.memberCard.card"),
        },
        {
            key: "issued",
            header: i18n("participants.memberCard.issuedAt"),
        },
        {
            key: "expired",
            header: i18n("participants.memberCard.expiredAt"),
        },
    ]}
    rows={memberCards.map((card, index) => {
        return {
            index: `#${index + 1}`,
            issued: card.issuedAt,
            expired: card.expiredAt || "X",
        };
    })}
/>

<style>
    .member-card-header {
        max-width: var(--container-max-width);
        margin: 0 auto;
        margin-bottom: -1rem;
    }

    .holder {
        font-size: 2.5rem;
        font-weight: bold;
        color: var(--secondary-color-dark);
        cursor: default;
        letter-spacing: 0.2rem;
    }

    button {
        cursor: pointer;
    }

    .primary {
        font-weight: bold;
        border-radius: 1rem;
        transition: transform 0.5s ease-in-out, background-color 0.5s ease-in-out ;
        padding: 0.5rem 2rem;
        background-color: var(--primary-color);
        color: white;
    }

    .primary:hover {
        transform: scale(1.02);
        background-color: var(--primary-color-dark);
    }

    .member-card-overview-header {
        font-weight: bold;
        font-size: 1.5rem;
        color: var(--primary-color);
    }
</style>
