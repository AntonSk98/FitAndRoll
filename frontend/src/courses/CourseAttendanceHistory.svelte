<script>
    import { onMount } from "svelte";
    import { toastError, toastSuccess } from "../toast/toastStore.js";
    import {
        FindCourseAttendanceHistory,
        UndoCourseAttendance,
    } from "../../wailsjs/go/courseattendance/CourseAttendanceHandler.js";
    import TableComponent from "../common/TableComponent.svelte";
    import Modal from "../common/Modal.svelte";
    import { i18n } from "../common/i18n.js";

    export let courseId = null;
    export let onComponentDestroyed;

    let courseAttendanceHistoryPage;
    let showUndoParticipationConfirmationModal = false;
    let selectedParticipationHistoryEntry;
    let tableRef;

    onMount(() => loadCourseAttendanceHistory());

    function getMainFilters() {
        const commonFilters = [
            {
                key: "excludeArchivedCourse",
                label: i18n(
                    "courseHistory.commonFilters.excludeArchivedCourse",
                ),
                type: "checkbox",
            },
            {
                key: "excludeWithMemberCard",
                label: i18n(
                    "courseHistory.commonFilters.excludeWithMemberCard",
                ),
                type: "checkbox",
            },
            {
                key: "excludeNoMemberCard",
                label: i18n("courseHistory.commonFilters.excludeNoMemberCard"),
                type: "checkbox",
            },

            {
                key: "attendedRange",
                label: i18n("courseHistory.commonFilters.attendedRange"),
                type: "date",
            },
            {
                key: "excludeTrialAttendance",
                label: i18n(
                    "courseHistory.commonFilters.excludeTrialAttendance",
                ),
                type: "checkbox",
            },
        ];

        return courseId
            ? commonFilters.filter(
                  (filter) => filter.key !== "excludeArchivedCourse",
              )
            : commonFilters;
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
                loadCourseAttendanceHistory(tableRef?.getFilter());
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
                    html: `<span class='text-[var(--primary-color)] font-bold'>${i18n("courseHistory.attendanceType.withCard")}</span>`,
                };
            case "TRIAL_ATTENDANCE":
                return {
                    html: `<span class='text-[var(--warning-color)] font-bold'>${i18n("courseHistory.attendanceType.trial")}</span>`,
                };
            case "WITHOUT_MEMBER_CARD":
                return {
                    html: `<span class='text-[var(--error-color)] font-bold'>${i18n("courseHistory.attendanceType.noCard")}</span>`,
                };
            default:
                return "";
        }
    }

    function getActions() {
        if (courseId) {
            return [];
        }

        return [
            {
                title: i18n("courseHistory.removeHistoryEntry"),
                icon: "handRaised",
                onClick: undoParticipation,
            },
        ];
    }
</script>

<TableComponent
    bind:this={tableRef}
    tableHeader={i18n("courseHistory.header")}
    total={courseAttendanceHistoryPage?.total ?? 0}
    mainFilters={getMainFilters()}
    columns={[
        {
            key: "fullname",
            header: i18n("courseHistory.columns.fullname"),
            filterbar: true,
        },
        {
            key: "course",
            header: i18n("courseHistory.columns.course"),
            filterbar: true,
        },
        {
            key: "attendedAt",
            header: i18n("courseHistory.columns.attendedAt"),
        },
        {
            key: "attendanceType",
            header: i18n("courseHistory.columns.attendedAs"),
        },
    ]}
    actions={getActions()}
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
            title: i18n("courseHistory.back"),
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
                        {i18n("courseHistory.modal.withCard.header")}?
                    </div>
                    <div>
                        <span class="font-bold text-[var(--primary-color)]"
                            >{selectedParticipationHistoryEntry.fullname}</span
                        >
                        {i18n("courseHistory.modal.withCard.bodyP1")}
                        <span class="font-bold text-[var(--primary-color)]"
                            >{selectedParticipationHistoryEntry.course}</span
                        >
                        {i18n("courseHistory.modal.withCard.bodyP2")}.
                    </div>
                    <div>
                        {i18n("courseHistory.modal.withCard.footer")}.
                    </div>
                </div>
            {:else}
                <div>
                    {i18n("courseHistory.modal.generic")}?
                </div>
            {/if}
        </div>
    </Modal>
{/if}
