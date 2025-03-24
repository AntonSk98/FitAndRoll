<script>
    import { onMount } from "svelte";
    import { toastError, toastSuccess } from "../toast/toastStore.js";
    import {
        ParticipantStatistics,
        CourseStatistics,
    } from "../../wailsjs/go/statistics/StatisticsHandler";
    import TableComponent from "../common/TableComponent.svelte";

    let statisticsPage;
    let displayCourseStatistics = false;

    const PARTICIPANTS_STATISTICS = "PARTICIPANTS_STATISTICS";
    const COURSE_STATISTICS = "COURSE_STATISTICS";

    $: if (displayCourseStatistics) {
        loadCourseStatistics();
    } else {
        loadParticipantStatistics();
    }

    function loadCourseStatistics(filter, pagination) {
        CourseStatistics(filter, pagination ?? { currentPage: 1, pageSize: 5 })
            .then((page) => (statisticsPage = page))
            .catch((err) => {
                console.error(err);
                toastError();
            });
    }

    function loadParticipantStatistics(filter, pagination) {
        ParticipantStatistics(
            filter,
            pagination ?? { currentPage: 1, pageSize: 5 },
        )
            .then((page) => (statisticsPage = page))
            .catch((err) => {
                console.error(err);
                toastError();
            });
    }

    function onPaginationFilterChanged(filter, pagination, type) {
        if (PARTICIPANTS_STATISTICS === type) {
            loadParticipantStatistics(filter, pagination);
        }

        if (COURSE_STATISTICS === type) {
            loadCourseStatistics(filter, pagination);
        }
    }
</script>

{#if !displayCourseStatistics}
    <TableComponent
        tableHeader="Participants statistics"
        total={statisticsPage?.total ?? 0}
        mainFilters={[
            {
                key: "attendedRange",
                label: "Attended range",
                type: "date",
            },
        ]}
        headerActions={[
            {
                title: "Course statistics",
                icon: "calender",
                onClick: () => displayCourseStatistics = true,
            },
        ]}
        columns={[
            {
                key: "name",
                header: "Fullname",
                filterbar: true,
            },
            {
                key: "withMemberCard",
                header: "With member card",
            },
            {
                key: "trialAttendance",
                header: "Trial attendance",
            },
            {
                key: "withoutMemberCard",
                header: "Without member card",
            },
        ]}
        rows={statisticsPage?.data}
        onPaginationFilterChanged={(filter, pagination) =>
            onPaginationFilterChanged(
                filter,
                pagination,
                PARTICIPANTS_STATISTICS,
            )}
    />
{/if}

{#if displayCourseStatistics}
    <TableComponent
        tableHeader="Course statistics"
        total={statisticsPage?.total ?? 0}
        mainFilters={[
            {
                key: "attendedRange",
                label: "Attended range",
                type: "date",
            },
        ]}
        headerActions={[
            {
                title: "close",
                icon: "xMark",
                onClick: () => displayCourseStatistics = false,
            },
        ]}
        columns={[
            {
                key: "name",
                header: "Course name",
                filterbar: true,
            },
            {
                key: "withMemberCard",
                header: "With member card",
            },
            {
                key: "trialAttendance",
                header: "Trial attendance",
            },
            {
                key: "withoutMemberCard",
                header: "Without member card",
            },
        ]}
        rows={statisticsPage?.data}
        onPaginationFilterChanged={(filter, pagination) =>
            onPaginationFilterChanged(
                filter,
                pagination,
                COURSE_STATISTICS,
            )}
    />
{/if}
