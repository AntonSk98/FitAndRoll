<script>
    import { onMount } from "svelte";
    import { FindCourses } from "../../wailsjs/go/courses/CourseController.js";
    import { ComponentControl } from "./componentControl.js";
    import TableComponent from "../table/TableComponent.svelte";
    import { addToast } from "../toast/toastStore.js";

    let coursePage;

    let componentControl = new ComponentControl();

    async function loadCourses(pagination) {
        try {
            coursePage = await FindCourses(
                pagination?.currentPage ?? 1,
                pagination?.pageSize ?? 5,
            );
        } catch (error) {
            console.error("Error loading courses:", error);
        }
    }

    function onPaginationFilterChanged(filter, pagination) {
        console.log({
            ...filter,
            ...pagination,
        });
        loadCourses(pagination);
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

    function triggerToast() {
        addToast({ message: "This is an info toast ", type: "error" });
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
                filterbar: true,
            },
        ]}
        rows={coursePage?.data.map((course) => {
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

<button class="ml-[400px]" on:click={triggerToast}>Show Info Toast</button>
