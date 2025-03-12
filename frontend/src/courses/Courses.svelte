<script>
    import { t } from "svelte-i18n";
    import {
        FindCourses,
        ArchiveCourse,
    } from "../../wailsjs/go/courses/CourseController.js";
    import { ComponentControl } from "./componentControl.js";
    import TableComponent from "../common/TableComponent.svelte";
    import { toastError, toastSuccess } from "../toast/toastStore.js";
    import CreateUpdateCourse from "./CreateUpdateCourse.svelte";
    import CourseParticipant from "./CourseParticipant.svelte";
    import CourseAttendanceHistory from "./CourseAttendanceHistory.svelte";

    let coursePage;

    let componentControl = new ComponentControl();
    let selectedCourse;

    $: if (componentControl?.showCourseOverview) {
        loadCourses();
    }

    async function loadCourses(filter, pagination) {
        try {
            coursePage = await FindCourses(
                filter,
                pagination ?? { currentPage: 1, pageSize: 5 },
            );
        } catch (error) {
            console.error("Error loading courses:", error);
            toastError();
        }
    }

    function onPaginationFilterChanged(filter, pagination) {
        loadCourses(filter, pagination);
    }

    function displayDetails(index) {
        selectedCourse = toSelectedCourse(index);
        componentControl = componentControl.showUpdateCourseComponent();
    }

    function displayCourseParticipants(index) {
        selectedCourse = toSelectedCourse(index);
        componentControl = componentControl.showCourseParticipantComponent();
    }

    function displayParticipationHistory(index) {
        selectedCourse = toSelectedCourse(index);
        componentControl.showCourseParticipantsHistoryComponent();
    }

    function archiveCourse(index) {
        ArchiveCourse(coursePage.data[index]?.id)
            .then(() => {
                toastSuccess("Course was successfully archived");
                loadCourses();
            })
            .catch((err) => {
                console.error("Error while archiving a course: ", err);
                toastError();
            });
    }

    function toSelectedCourse(index) {
        return coursePage.data[index];
    }
</script>

{#if componentControl.showCourseOverview}
    <TableComponent
        tableHeader={$t("courses.table.header")}
        total={coursePage?.total ?? 0}
        columns={[
            {
                key: "name",
                header: $t("courses.table.columns.course"),
                filterbar: true,
            },
            {
                key: "schedule",
                header: $t("courses.table.columns.schedule"),
            },
        ]}
        rows={coursePage?.data?.map((course) => {
            return {
                name: course.name,
                schedules: course?.schedules?.length
                    ? course.schedules.map(
                          (schedule) => `${schedule.day} (${schedule.time})`,
                      )
                    : $t("course.test"),
            };
        })}
        actions={[
            {
                title: $t("courses.table.actions.details"),
                icon: "edit",
                onClick: displayDetails,
            },
            {
                title: $t("courses.table.actions.participants"),
                icon: "info",
                onClick: displayCourseParticipants,
            },
            {
                title: $t("courses.table.actions.history"),
                icon: "calender",
                onClick: displayParticipationHistory,
            },
            {
                title: $t("courses.table.actions.archive"),
                icon: "trash",
                requireConfirmation: true,
                onClick: archiveCourse,
            },
        ]}
        headerActions={[
            {
                title: $t("courses.table.actions.newCourse"),
                icon: "plus",
                onClick: () =>
                    (componentControl =
                        componentControl.showDefineNewCourseComponent()),
            },
            {
                title: "Course attendance history",
                icon: "userGroup",
                onClick: () =>
                    (componentControl =
                        componentControl.showAllAttendanceHistoryComponent()),
            },
        ]}
        {onPaginationFilterChanged}
    />
{/if}

{#if componentControl.defineNewCourseComponent}
    <CreateUpdateCourse bind:componentControl />
{/if}

{#if componentControl.updateCourseComponent}
    <CreateUpdateCourse bind:componentControl courseId={selectedCourse?.id} />
{/if}

{#if componentControl.courseParticipantsComponent}
    <CourseParticipant
        {selectedCourse}
        returnToCourseOverview={() =>
            (componentControl = componentControl.resetComponentControl())}
    />
{/if}

{#if componentControl.courseParticipantsHistoryComponent}
    <CourseAttendanceHistory
        courseId={selectedCourse?.id}
        onComponentDestroyed={() =>
            (componentControl = componentControl.resetComponentControl())}
    />
{/if}

{#if componentControl.allAttendanceHistoryComponent}
    <CourseAttendanceHistory
        onComponentDestroyed={() =>
            (componentControl = componentControl.resetComponentControl())}
    />
{/if}
