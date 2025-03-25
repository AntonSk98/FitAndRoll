<script>
    import { toastError, toastSuccess } from "../toast/toastStore.js";
    import {
        ParticipantStatistics,
        CourseStatistics,
    } from "../../wailsjs/go/statistics/StatisticsHandler";
    import TableComponent from "../common/TableComponent.svelte";
    import { i18n } from "../common/i18n.js";

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
        tableHeader= {i18n("statistics.participants.header")}
        total={statisticsPage?.total ?? 0}
        mainFilters={[
            {
                key: "attendedRange",
                label: i18n("statistics.attendedRange"),
                type: "date",
            },
        ]}
        headerActions={[
            {
                title: i18n("statistics.courses.header"),
                icon: "calender",
                onClick: () => displayCourseStatistics = true,
            },
        ]}
        columns={[
            {
                key: "name",
                header: i18n("statistics.participants.fullname"),
                filterbar: true,
            },
            {
                key: "withMemberCard",
                header: i18n("statistics.participants.withMemberCard")
            },
            {
                key: "trialAttendance",
                header: i18n("statistics.participants.trialAttendance"),
            },
            {
                key: "withoutMemberCard",
                header: i18n("statistics.participants.withoutMemberCard"),
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
        tableHeader={i18n("statistics.courses.header")}
        total={statisticsPage?.total ?? 0}
        mainFilters={[
            {
                key: "attendedRange",
                label: i18n("statistics.attendedRange"),
                type: "date",
            },
        ]}
        headerActions={[
            {
                title: i18n("statistics.close"),
                icon: "xMark",
                onClick: () => displayCourseStatistics = false,
            },
        ]}
        columns={[
            {
                key: "name",
                header: i18n("statistics.courses.name"),
                filterbar: true,
            },
            {
                key: "withMemberCard",
                header: i18n("statistics.courses.withMemberCard"),
            },
            {
                key: "trialAttendance",
                header: i18n("statistics.courses.trialAttendance"),
            },
            {
                key: "withoutMemberCard",
                header: i18n("statistics.courses.withoutMemberCard"),
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
