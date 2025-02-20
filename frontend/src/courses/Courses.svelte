<script>
    import { t } from "svelte-i18n";
    import { onMount } from "svelte";
    import { FindCourses } from "../../wailsjs/go/courses/CourseController.js";
    import { ComponentControl } from "./componentControl.js";
    import TableComponent from "../table/TableComponent.svelte";
    import { addToast } from "../toast/toastStore.js";

    let coursePage;

    let componentControl = new ComponentControl();

    async function loadCourses(filter, pagination) {
        try {
            coursePage = await FindCourses(
                filter,
                pagination ?? { currentPage: 1, pageSize: 5 },
            );
            console.log(coursePage);
        } catch (error) {
            console.error("Error loading courses:", error);
            addToast({ type: "error" });
        }
    }

    function onPaginationFilterChanged(filter, pagination) {
        loadCourses(filter, pagination);
    }

    onMount(loadCourses);

    function displayDetails(index) {
        console.log(coursePage.data[index]);
    }

    function displayCourseParticipants(index) {
        console.log(index);
    }

    function displayParticipationHistory(index) {
        console.log(index);
    }

    function displayCreateCourse() {
        console.log("hi");
    }

    function archiveCourse(index) {
        console.log(index);
    }
</script>

{#if componentControl.showCourseOverview}
    <TableComponent
        tableHeader={$t("courses.table.header")}
        total={coursePage?.total ?? 0}
        columns={[
            { key: "name", header: $t("courses.table.columns.course"), filterbar: true },
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
                onClick: displayCreateCourse,
            },
        ]}
        {onPaginationFilterChanged}
    />
{/if}
