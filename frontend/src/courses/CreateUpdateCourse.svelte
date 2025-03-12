<script>
    import { ComponentControl } from "./componentControl";
    import FormComponent from "../common/Form.svelte";
    import {
        FindCourseDetails,
        NewCourse,
        UpdateCourse
    } from "../../wailsjs/go/courses/CourseHandler.js";
    import { toastError, toastSuccess } from "../toast/toastStore.js";
    import { onMount } from "svelte";
    import Form from "../common/Form.svelte";

    export let courseId = null;
    export let componentControl = new ComponentControl();

    let formConfig;
    let courseDetails;

    onMount(async () => {
        if (courseId) {
            await findCourseDetails(courseId);
        } else {
            defineForm(null);
        }
    });

    async function findCourseDetails(courseId) {
        try {
            courseDetails = await FindCourseDetails(courseId);
            defineForm(courseDetails);
        } catch (error) {
            console.error("Error loading course details:", error);
            toastError();
        }
    }

    function onCourseFormCancelled() {
        componentControl = componentControl.resetComponentControl();
    }

    async function onCourseFormSubmit(form) {
        try {
            courseId ? await UpdateCourse({id: courseId, ...form}) : await NewCourse(form);
            toastSuccess()
            componentControl = componentControl.resetComponentControl();
        } catch(error) {
            console.error("Error occurred while updating|saving the course details:", error);
            toastError();
        }
    }

    function defineForm(course) {
        formConfig = {
            header: !courseDetails ? "Create Course" : "Update Course",
            fields: [
                {
                    key: "name",
                    type: "text",
                    value: course?.name,
                    display: "Name",
                    validation: {
                        message: "Must contain at least 2 characters",
                        function: (value) => value.length >= 2,
                    },
                },
                {
                    key: "description",
                    type: "text",
                    value: course?.description,
                    display: "Beschreibung",
                },
                {
                    key: "schedules",
                    type: "list",
                    value: course?.schedules,
                    subtype: "datetime",
                    display: "Schedules",
                    validation: {
                        message: "Must contain at least one schedule entry",
                        function: (values) => values?.length > 0,
                    },
                },
            ],
            actions: {
                cancel: onCourseFormCancelled,
                submit: onCourseFormSubmit,
            },
        };
    }
</script>

<FormComponent {formConfig} />
