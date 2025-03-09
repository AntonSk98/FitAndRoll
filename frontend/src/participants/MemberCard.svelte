<script>
    import Modal from "../common/Modal.svelte";
    import TableComponent from "../common/TableComponent.svelte";
    import { toastError, toastSuccess } from "../toast/toastStore";
    import {
        FindAllMemberCards,
        IssueNewMemberCard,
        UndoIssueNewMemberCard,
    } from "../../wailsjs/go/participants/MemberCardController";

    import {
        LoadMemberCardCourseHistory
    } from "../../wailsjs/go/membercardattendance/MemberCardAttendanceController"
    import BackButton from "../common/BackButton.svelte";

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
                console.error(error);
                toastError();
            });
    }

    function deleteEmptyMemberCard() {
        UndoIssueNewMemberCard(member.id, selectedMemberCardId)
            .then(() => {
                findAllMemberCards();
                displayMemberCardParticipationHistory = false;
            })
            .catch((error) => {
                console.error(error);
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
                console.error(err);
                toastError();
            });
    }

    function issueNewMemberCard() {
        IssueNewMemberCard(member.id)
            .then(() => {
                displayNewCardConfirmationDialog = false;
                findAllMemberCards();
            })
            .catch((error) => {
                console.error(error);
                toastError();
            });
    }
</script>

<div class="member-card-header">
    <div class="flex justify-between items-center gap-3">
        <span class="holder">{member.name} {member.surname}</span>
        <BackButton {onBackButtonClicked} />
    </div>
    <button
        class="primary"
        on:click={() => (displayNewCardConfirmationDialog = true)}
        >New Card</button
    >
</div>

{#if displayNewCardConfirmationDialog}
    <Modal
        onModalConfirmed={issueNewMemberCard}
        onModalCanceled={() => (displayNewCardConfirmationDialog = false)}
    >
        <div slot="body">
            <div>
                A new member card is about to be issued to
                <b>{member.name} {member.surname}</b>.
            </div>
            <div>Please confirm.</div>
        </div>
        <div slot="confirm">Issue new card</div>
        <div slot="cancel">Cancel</div>
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
                <div class="member-card-overview-header">History</div>
                <ul class="px-3 py-2 flex flex-col gap-3">
                    {#each memberCardParticipationHistory as memberCardParticipationEntry}
                        <li
                            class="p-2 bg-[var(--secondary-color)] rounded-xl"
                        >
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
                <div class="font-bold">Unused member card</div>
                <div>
                    Since this card has not been used yet you can remove this
                    member card
                </div>
            {/if}
        </div>
        <div slot="confirm">
            {#if memberCardParticipationHistory?.length > 0}
                OK
            {:else}
                Remove member card
            {/if}
        </div>
        <div slot="cancel">Close</div>
    </Modal>
{/if}

<TableComponent
    tableHeader={"Member card history"}
    total="0"
    onRowClicked={onMemberCardClicked}
    columns={[
        {
            key: "index",
            header: "card",
        },
        {
            key: "issued",
            header: "issued",
        },
        {
            key: "expired",
            header: "expired",
        },
    ]}
    rows={memberCards.map((card, index) => {
        return {
            index: `#${index}`,
            issued: card.issuedAt,
            expired: card.expiredAt || "X",
        };
    })}
/>

<style>
    .member-card-header {
        max-width: var(--container-max-width);
        margin: 0 auto;
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
        border-radius: 8px;
        transition: all 0.3s ease-in-out;
        padding: 0.5rem 6rem;
        background-color: var(--primary-color);
        color: white;
        border: none;
    }

    .primary:hover {
        background-color: var(--primary-color-dark);
        transform: scale(1.05);
        box-shadow: 0px 4px 10px rgba(16, 185, 129, 0.5);
    }

    .member-card-overview-header {
        font-weight: bold;
        font-size: 1.5rem;
        color: var(--primary-color);
    }
</style>
