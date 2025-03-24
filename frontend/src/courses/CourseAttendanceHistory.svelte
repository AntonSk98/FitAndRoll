<script>
    import { onMount } from "svelte";
    import { toastError, toastSuccess } from "../toast/toastStore.js";
    import {
        FindCourseAttendanceHistory,
        UndoCourseAttendance,
    } from "../../wailsjs/go/courseattendance/CourseAttendanceHandler.js";
    import TableComponent from "../common/TableComponent.svelte";
    import Modal from "../common/Modal.svelte";

    export let courseId = null;
    export let onComponentDestroyed;

    let courseAttendanceHistoryPage;
    let showUndoParticipationConfirmationModal = false;
    let selectedParticipationHistoryEntry;

    onMount(() => loadCourseAttendanceHistory());

    function getMainFilters() {
        const commonFilters = [
            {
                key: "excludeTrialAttendance",
                label: "Hide trial attendances",
                type: "checkbox",
            },
            {
                key: "excludeWithMemberCard",
                label: "Hide with member card attendances",
                type: "checkbox",
            },
            {
                key: "excludeNoMemberCard",
                label: "Hide without member card attendances",
                type: "checkbox",
            },
            {
                key: "attendedRange",
                label: "Attended range",
                type: "date",
            },
        ];

        return courseId
            ? commonFilters
            : [
                  {
                      key: "excludeArchivedCourse",
                      label: "Hide archived courses",
                      type: "checkbox",
                  },
                  ...commonFilters,
              ];
    }

    async function loadCourseAttendanceHistory(filter, pagination) {
        try {
            courseAttendanceHistoryPage = await FindCourseAttendanceHistory(
                courseId ? { courseId: courseId, ...filter } : filter,
                pagination ?? { currentPage: 1, pageSize: 5 },
            );
        } catch (error) {
            console.error("Error loading course attendance history:", error);
            toastError();
        }
    }

    function onPaginationFilterChanged(filter, pagination) {
        loadCourseAttendanceHistory(filter, pagination);
    }

    function undoParticipation(index) {
        if (!showUndoParticipationConfirmationModal) {
            showUndoParticipationConfirmationModal = true;
            selectedParticipationHistoryEntry =
                courseAttendanceHistoryPage?.data[index];
            return;
        }

        UndoCourseAttendance(selectedParticipationHistoryEntry?.id)
            .then(() => {
                toastSuccess();
                loadCourseAttendanceHistory();
            })
            .catch((err) => {
                console.error(err);
                toastError();
            });
        showUndoParticipationConfirmationModal =
            !showUndoParticipationConfirmationModal;
    }

    function mapAttendanceType(attendanceType) {
        switch (attendanceType) {
            case "WITH_MEMBER_CARD":
                return {
                    html: `<span class='text-[var(--primary-color)] font-bold'>with member card</span>`,
                };
            case "TRIAL_ATTENDANCE":
                return {
                    html: `<span class='text-[var(--warning-color)] font-bold'>trial attendance</span>`,
                };
            case "WITHOUT_MEMBER_CARD":
                return {
                    html: `<span class='text-[var(--error-color)] font-bold'>without member card</span>`,
                };
            default:
                return "";
        }
    }
</script>

<TableComponent
    tableHeader="Course Attendance History Overview"
    total={courseAttendanceHistoryPage?.total ?? 0}
    mainFilters={getMainFilters()}
    columns={[
        {
            key: "fullname",
            header: "Fullname",
            filterbar: true,
        },
        {
            key: "course",
            header: "Course",
            filterbar: true,
        },
        {
            key: "attendedAt",
            header: "Attended at",
        },
        {
            key: "attendanceType",
            header: "Attended as",
        },
    ]}
    actions={[
        {
            title: "undo participation",
            icon: "handRaised",
            onClick: undoParticipation,
        },
    ]}
    rows={courseAttendanceHistoryPage?.data?.map((entry) => {
        return {
            fullname: entry.fullname,
            course: entry.course,
            attendedAt: entry.attendedAt,
            attendanceType: mapAttendanceType(entry?.attendanceType),
        };
    })}
    headerActions={[
        {
            title: "close",
            icon: "xMark",
            onClick: () => onComponentDestroyed(),
        },
    ]}
    {onPaginationFilterChanged}
/>

{#if showUndoParticipationConfirmationModal}
    <Modal
        onModalConfirmed={undoParticipation}
        onModalCanceled={() => (showUndoParticipationConfirmationModal = false)}
    >
        <div slot="body">
            {#if selectedParticipationHistoryEntry.attendanceType === "WITH_MEMBER_CARD"}
                <div class="text-[var(--text-color)] text-justify">
                    <div class="font-bold">
                        Are you sure you want to delete this attendance entry?
                    </div>
                    <div>
                        <span class="font-bold text-[var(--primary-color)]"
                            >{selectedParticipationHistoryEntry.fullname}</span
                        >
                        attended
                        <span class="font-bold text-[var(--primary-color)]"
                            >{selectedParticipationHistoryEntry.course}</span
                        > using a member card.
                    </div>
                    <div>
                        Deleting this entry will restore one available slot to
                        the member card.
                    </div>
                </div>
            {:else}
                <div>
                    Are you sure you want to remove this attendance entry?
                </div>
            {/if}
        </div>
    </Modal>
{/if}
