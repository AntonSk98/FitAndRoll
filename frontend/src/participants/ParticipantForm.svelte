<script>
    import { toastError, toastSuccess } from "../toast/toastStore";
    import FormComponent from "../common/Form.svelte";
    import {
        CreateUpdateParticipant,
        FindParticipantDetails,
    } from "../../wailsjs/go/participants/ParticipantsHandler";
    import { onMount } from "svelte";

    export let participantId = null;
    export let backToOverview;

    let formGroup;
    let participant;

    onMount(() => {
        if (participantId) {
            findParticipantDetails(participantId);
            return;
        }

        initForm();
    });

    function initForm() {
        formGroup = {
            header: !participant ? "New Participant" : "Update Participant ",
            fields: [
                {
                    key: "name",
                    type: "text",
                    value: participant?.name,
                    display: "Name",
                    validation: {
                        message: "Must contain at least 2 characters",
                        function: (value) => value.length >= 2,
                    },
                },
                {
                    key: "surname",
                    type: "text",
                    value: participant?.surname,
                    display: "Surname",
                    validation: {
                        message: "Must contain at least 2 characters",
                        function: (value) => value.length >= 2,
                    },
                },
                {
                    key: "group",
                    type: "text",
                    value: participant?.group,
                    display: "Group",
                    validation: {
                        message: "Must contain at least 2 characters",
                        function: (value) => value?.length > 2,
                    },
                },
            ],
            actions: {
                cancel: backToOverview,
                submit: onSubmitForm,
            },
        };
    }

    function onSubmitForm(form) {
        if (participantId) {
            form.id = participantId;
        }
        createUpdateCourse(form)
    }

    function findParticipantDetails(id) {
        FindParticipantDetails(id)
            .then(details => participant = details)
            .then(() => initForm())
            .catch(error => {
                console.error(`Error while fetching participant details. Error ${error.message}`)
                toastError();
            });
    }

    function createUpdateCourse(form) {
        CreateUpdateParticipant(form)
        .then(() => {
            toastSuccess();
            backToOverview();
        })
        .catch(err => {
            console.error('Error while registering a customer', err);
            toastError();
        })
    }
</script>

<FormComponent formConfig={formGroup} />
