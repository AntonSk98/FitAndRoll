<script>
    import { toastError, toastSuccess } from "../toast/toastStore";
    import FormComponent from "../common/Form.svelte";
    import {
        CreateUpdateParticipant,
        FindParticipantDetails,
    } from "../../wailsjs/go/participants/ParticipantsHandler";
    import { onMount } from "svelte";
    import { i18n } from "../common/i18n";

    import { LogError } from "../../wailsjs/runtime/runtime";

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
            header: !participant
                ? i18n("participants.form.new")
                : i18n("participants.form.update"),
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
                {
                    key: "phone",
                    type: "text",
                    value: participant?.phone,
                    display: i18n("participants.phone"),
                },
                {
                    key: "email",
                    type: "text",
                    value: participant?.email,
                    display: i18n("participants.email"),
                },
                {
                    key: "birthday",
                    type: "date",
                    value: participant?.birthday,
                    display: i18n("participants.birthday"),
                },
                {
                    key: "createdAt",
                    type: "text",
                    value: participant?.createdAt,
                    display: i18n("participants.createdAt"),
                    readonly: true,
                },
                {
                    key: "privacyPolicy",
                    type: "checkbox",
                    value: participant?.privacyPolicy,
                    display: `${i18n("participants.privacyPolicy")}${participant?.privacyPolicyAcceptedAt ? ` | ${participant.privacyPolicyAcceptedAt}` : ""}`,
                },
                {
                    key: "address",
                    type: "textarea",
                    value: participant?.address,
                    display: i18n("participants.address"),
                },
                {
                    key: "parents",
                    type: "textarea",
                    value: participant?.parents,
                    display: i18n("participants.parents"),
                },
                {
                    key: "notes",
                    type: "textarea",
                    value: participant?.notes,
                    display: i18n("participants.notes"),
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
        createUpdateCourse(form);
    }

    function findParticipantDetails(id) {
        FindParticipantDetails(id)
            .then((details) => (participant = details))
            .then(() => initForm())
            .catch((error) => {
                LogError(
                    `Error while fetching participant details by id ${id}. Error: ${error}`,
                );
                toastError();
            });
    }

    function createUpdateCourse(form) {
        console.log("Form submitted", form);
        CreateUpdateParticipant(form)
            .then(() => {
                toastSuccess();
                backToOverview();
            })
            .catch((err) => {
                LogError(
                    `Error while persisting participant details. Form: ${form}. Error: ${err}`,
                );
                toastError();
            });
    }
</script>

<FormComponent formConfig={formGroup} />
