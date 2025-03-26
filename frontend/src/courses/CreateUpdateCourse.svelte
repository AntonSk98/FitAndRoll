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
    import { i18n } from "../common/i18n";

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
            header: !courseDetails ? i18n("courseForm.create") : i18n("courseForm.update"),
            fields: [
                {
                    key: "name",
                    type: "text",
                    value: course?.name,
                    display: i18n("courseForm.name"),
                    validation: {
                        message: i18n("courseForm.nameValidation"),
                        function: (value) => value?.trim()?.length > 0,
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
                    display: i18n("courseForm.schedule"),
                    validation: {
                        message: i18n("courseForm.scheduleValidation"),
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
