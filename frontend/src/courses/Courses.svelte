<script>
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
            console.log(coursePage)
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
        tableHeader="Course Overview"
        total={coursePage?.total ?? 0}
        columns={[
            { key: "name", header: "course", filterbar: true },
            {
                key: "schedule",
                header: "schedule",
            },
        ]}
        rows={coursePage?.data?.map(course => {
            return {
                name: course.name,
                schedules: course?.schedules?.length
                    ? course.schedules.map(
                          (schedule) => `${schedule.day} (${schedule.time})`,
                      )
                    : "No schedules available",
            };
        })}
        actions={[
            {
                title: "Details",
                icon: "edit",
                onClick: displayDetails,
            },
            {
                title: "Participants",
                icon: "info",
                onClick: displayCourseParticipants,
            },
            {
                title: "Participation History",
                icon: "calender",
                onClick: displayParticipationHistory,
            },
            {
                title: "Archive",
                icon: "trash",
                requireConfirmation: true,
                onClick: archiveCourse,
            },
        ]}
        headerActions={[
            {
                title: "New Course",
                icon: "plus",
                onClick: displayCreateCourse,
            },
        ]}
        {onPaginationFilterChanged}
    />
{/if}
