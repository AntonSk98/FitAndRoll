<script>
    import { i18n } from "../common/i18n";
    import {
        FindCourses,
        ArchiveCourse,
    } from "../../wailsjs/go/courses/CourseHandler.js";
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
                toastSuccess();
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

    function mapDay(day) {
        const days = {
            monday: "courses.table.days.monday",
            tuesday: "courses.table.days.tuesday",
            wednesday: "courses.table.days.wednesday",
            thursday: "courses.table.days.thursday",
            friday: "courses.table.days.friday",
            saturday: "courses.table.days.saturday",
            sunday: "courses.table.days.sunday",
        };

        return i18n(days[day]);
    }
</script>

{#if componentControl.showCourseOverview}
    <TableComponent
        tableHeader={i18n("courses.table.header")}
        total={coursePage?.total ?? 0}
        columns={[
            {
                key: "name",
                header: i18n("courses.table.columns.course"),
                filterbar: true,
            },
            {
                key: "schedule",
                header: i18n("courses.table.columns.schedule"),
            },
        ]}
        rows={coursePage?.data?.map((course) => {
            return {
                name: course.name,
                schedules: course?.schedules?.length
                    ? course.schedules.map(
                          (schedule) =>
                              `${mapDay(schedule.day)} (${schedule.time})`,
                      )
                    : i18n("course.test"),
            };
        })}
        actions={[
            {
                title: i18n("courses.table.actions.details"),
                icon: "edit",
                onClick: displayDetails,
            },
            {
                title: i18n("courses.table.actions.participants"),
                icon: "info",
                onClick: displayCourseParticipants,
            },
            {
                title: i18n("courses.table.actions.history"),
                icon: "calender",
                onClick: displayParticipationHistory,
            },
            {
                title: i18n("courses.table.actions.archive"),
                icon: "trash",
                requireConfirmation: true,
                onClick: archiveCourse,
            },
        ]}
        headerActions={[
            {
                title: i18n("courses.table.actions.newCourse"),
                icon: "plus",
                onClick: () =>
                    (componentControl =
                        componentControl.showDefineNewCourseComponent()),
            },
            {
                title: i18n(
                    "courses.table.actions.overallParticipationHistory",
                ),
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
