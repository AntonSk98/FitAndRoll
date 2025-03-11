<script>
    import { onMount } from "svelte";
    import { toastError, toastSuccess } from "../toast/toastStore.js";
    import { FindCourseAttendanceHistory } from "../../wailsjs/go/courseattendance/CourseAttendanceController.js";
    import TableComponent from "../common/TableComponent.svelte";

    export let courseId = null;
    export let onComponentDestroyed;

    let courseAttendanceHistoryPage;

    onMount(() => loadCourseAttendanceHistory());

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

    function mapAttendanceType(attendanceType) {
        switch (attendanceType) {
            case "WITH_MEMBER_CARD":
                return {
                    html: `<span class='text-[var(--primary-color)] font-bold'>with member card</span>`
                };
            case "TRIAL_ATTENDANCE":
            return {
                    html: `<span class='text-[var(--warning-color)] font-bold'>trial attendance</span>`
                };
            case "WITHOUT_MEMBER_CARD":
            return {
                    html: `<span class='text-[var(--error-color)] font-bold'>without member card</span>`
                };
            default:
                return '';
        }
    }
</script>

<TableComponent
    tableHeader="Course Attendance History Overview"
    total={courseAttendanceHistoryPage?.total ?? 0}
    mainFilters={[
        {
            key: "excludeArchivedCourse",
            label: "Hide archived courses",
            type: "checkbox",
        },
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
    ]}
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
            dynamicContent: true
        },
    ]}
    rows={courseAttendanceHistoryPage?.data?.map((entry) => {
        return {
            ...entry,
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
