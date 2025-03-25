<script>
    import { toastError, toastSuccess } from "../toast/toastStore";
    import FormComponent from "../common/Form.svelte";
    import {
        CreateUpdateParticipant,
        FindParticipantDetails,
    } from "../../wailsjs/go/participants/ParticipantsHandler";
    import { onMount } from "svelte";
    import { i18n } from "../common/i18n";

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
            header: !participant ? i18n("participants.form.new") : i18n("participants.form.update"),
            fields: [
                {
                    key: "name",
                    type: "text",
                    value: participant?.name,
                    display: i18n("participants.name"),
                    validation: {
                        message: i18n("participants.form.notEmptyValidation"),
                        function: (value) => value?.trim()?.length > 0,
                    },
                },
                {
                    key: "surname",
                    type: "text",
                    value: participant?.surname,
                    display: i18n("participants.surname"),
                    validation: {
                        message: i18n("participants.form.notEmptyValidation"),
                        function: (value) => value?.trim()?.length > 0,
                    },
                },
                {
                    key: "group",
                    type: "text",
                    value: participant?.group,
                    display: i18n("participants.group"),
                    validation: {
                        message: i18n("participants.form.notEmptyValidation"),
                        function: (value) => value?.trim()?.length > 0,
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
