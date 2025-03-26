<script>
    import Modal from "../common/Modal.svelte";
    import { toastError, toastSuccess } from "../toast/toastStore.js";
    import { AttendCourse } from "../../wailsjs/go/membercardattendance/MemberCardAttendanceHandler";
    import { i18n } from "../common/i18n";

    export let attendanceType;
    export let selectedParticipant;
    export let selectedCourse;
    export let activeMemberCards;
    export let onDestroy;

    function attendCourseWithMemberCard(memberCard) {
        abstractAttendCourse(memberCard);
    }

    function trialAttend() {
        abstractAttendCourse();
    }

    function attendCourseWIthoutMemberCard() {
        abstractAttendCourse();
    }

    function abstractAttendCourse(memberCard) {
        const attendCourseCommand = {
            memberCard: memberCard?.id ?? activeMemberCards[0]?.id,
            course: selectedCourse?.id,
            participant: selectedParticipant?.id,
            attendanceType: attendanceType,
        };

        AttendCourse(attendCourseCommand)
            .then((__) => {
                toastSuccess();
                onDestroy();
            })
            .catch((err) => {
                console.error(err);
                toastError();
            });
    }
</script>

{#if attendanceType === "WITH_MEMBER_CARD"}
    <Modal
        onModalCanceled={onDestroy}
        onModalConfirmed={activeMemberCards?.length === 1
            ? () => attendCourseWithMemberCard()
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
                d="M7.5 8.25h9m-9 3H12m-9.75 1.51c0 1.6 1.123 2.994 2.707 3.227 1.129.166 2.27.293 3.423.379.35.026.67.21.865.501L12 21l2.755-4.133a1.14 1.14 0 0 1 .865-.501 48.172 48.172 0 0 0 3.423-.379c1.584-.233 2.707-1.626 2.707-3.228V6.741c0-1.602-1.123-2.995-2.707-3.228A48.394 48.394 0 0 0 12 3c-2.392 0-4.744.175-7.043.513C3.373 3.746 2.25 5.14 2.25 6.741v6.018Z"
            />
        </svg>

        <div slot="body">
            {#if activeMemberCards?.length === 1}
                <div class="text-justify">
                    <div class="text-[var(--primary-color)] font-bold text-2xl">
                        {i18n("courseParticipant.modal.card")}
                    </div>
                    <div class="flex flex-col my-1">
                        <div>
                            {i18n("courseParticipant.modal.issuedFor")}:
                            <span class="font-semibold"
                                >{selectedParticipant.name}
                                {selectedParticipant.surname}</span
                            >
                        </div>
                        <div>
                            {i18n("courseParticipant.modal.issuedOn")}:
                            <span class="font-semibold"
                                >{activeMemberCards[0].issuedAt}</span
                            >
                        </div>
                        <div>
                            {i18n("courseParticipant.modal.remainingSlots")}:
                            <span class="font-semibold"
                                >{activeMemberCards[0].capacity}</span
                            >
                        </div>
                    </div>
                </div>
            {:else if activeMemberCards?.length > 1}
                <div class="text-justify">
                    <div>
                        <span
                            class="font-bold text-[var(--primary-color)] text-lg"
                        >
                            {selectedParticipant.name}
                            {selectedParticipant.surname}
                        </span>
                        {i18n("courseParticipant.modal.severalCards")}.
                    </div>
                    <div>
                        {i18n("courseParticipant.modal.selectCardPrompt")}.
                    </div>
                    <div class="flex flex-col gap-4 my-2">
                        {#each activeMemberCards as activeMemberCard, index}
                            <div
                                on:click={() =>
                                    attendCourseWithMemberCard(
                                        activeMemberCard,
                                    )}
                                class="text-sm flex flex-col cursor-pointer bg-[var(--primary-color)] hover:bg-[var(--primary-color-dark)] text-white p-2 rounded-md duration-700 ease-in-out transform hover:scale-[1.02]"
                            >
                                <span class="text-lg font-bold"
                                    >{i18n("courseParticipant.modal.card")} #{index +
                                        1}</span
                                >
                                <span
                                    >{i18n(
                                        "courseParticipant.modal.remainingSlots",
                                    )}:
                                    <span class="font-semibold"
                                        >{activeMemberCard.capacity}</span
                                    ></span
                                >
                                <span
                                    >{i18n("courseParticipant.modal.issuedOn")}:
                                    <span class="font-semibold"
                                        >{activeMemberCard.issuedAt}</span
                                    ></span
                                >
                            </div>
                        {/each}
                    </div>
                </div>
            {:else}
                {i18n("courseParticipant.modal.noActiveCard")}
            {/if}
        </div>

        <div slot="confirm">
            {i18n("courseParticipant.modal.attendConfirm")}
        </div>
        <div slot="cancel">{i18n("courseParticipant.modal.attendCancel")}</div>
    </Modal>
{/if}

{#if attendanceType === "TRIAL_ATTENDANCE"}
    <Modal onModalCanceled={onDestroy} onModalConfirmed={trialAttend}>
        <div slot="body">
            {i18n("courseParticipant.modal.confirmTrial1")}
            <span class="font-bold text-[var(--primary-color)]"
                >{selectedParticipant.name} {selectedParticipant.surname}</span
            >
            {i18n("courseParticipant.modal.confirmTrial2")}
            <span class="font-bold text-[var(--primary-color)]"
                >{selectedCourse.name}</span
            >
            {i18n("courseParticipant.modal.confirmTrial3")}.
        </div>
        <div slot="confirm">{i18n("courseParticipant.modal.trialConfirm")}</div>
        <div slot="cancel">{i18n("courseParticipant.modal.attendCancel")}</div>
    </Modal>
{/if}

{#if attendanceType === "WITHOUT_MEMBER_CARD"}
    <Modal
        onModalCanceled={onDestroy}
        onModalConfirmed={attendCourseWIthoutMemberCard}
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
                d="M12 9v3.75m-9.303 3.376c-.866 1.5.217 3.374 1.948 3.374h14.71c1.73 0 2.813-1.874 1.948-3.374L13.949 3.378c-.866-1.5-3.032-1.5-3.898 0L2.697 16.126ZM12 15.75h.007v.008H12v-.008Z"
            />
        </svg>

        <div slot="body">
            {i18n("courseParticipant.modal.confirmTrial1")}
            <span class="font-bold text-[var(--primary-color)]"
                >{selectedParticipant.name} {selectedParticipant.surname}</span
            >
            {i18n("courseParticipant.modal.confirmTrial4")}
            <span class="font-bold text-[var(--primary-color)]"
                >{selectedCourse.name}</span
            >
            {i18n("courseParticipant.modal.confirmTrial5")}.
        </div>
        <div slot="confirm">
            {i18n("courseParticipant.modal.noCardConfirm")}
        </div>
        <div slot="cancel">{i18n("courseParticipant.modal.attendCancel")}</div>
    </Modal>
{/if}
