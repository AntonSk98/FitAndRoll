<script>
  import DateTimePicker from "../common/DateTimePicker.svelte";
  import { i18n } from "./i18n";

  export let formConfig;
  const form = {};
  let isFormValid = false;

  $: if (formConfig) {
    formConfig.fields.forEach((field) => {
      const formField = {
        value: field?.value ?? "",
        valid: field?.validation?.function(field?.value ?? "") ?? true,
      };
      form[field.key] = formField;
    });
    validateForm();
  }

  function onInput(field, value) {
    form[field.key] = {
      value: value,
      valid: field?.validation?.function(value) ?? true,
    };
    validateForm();
  }

  function validateForm() {
    isFormValid = Object.values(form).every((formEntry) => formEntry.valid);
  }

  function submitForm() {
    formConfig?.actions?.submit(
      Object.fromEntries(
        Object.entries(form).map(([key, { value }]) => [key, value]),
      ),
    );
  }

  function onDateTimeSet(field, dateTIme) {
    const currentDateTimeList = form[field.key]?.["value"] ?? [];
    onInput(field, [...new Set([...currentDateTimeList, dateTIme])]);
  }

  function removeDateTimeEntry(field, toBeRemovedDateTimeEntry) {
    onInput(
      field,
      form[field.key]?.["value"]?.filter(
        (dateTimeEntry) => dateTimeEntry !== toBeRemovedDateTimeEntry,
      ),
    );
  }

  function resolveDay(day) {
    return i18n(`dateTimePicker.options.${day}`);
  }
</script>

<div class="form-container">
  <h1 class="form-header">{formConfig?.header}</h1>

  <form class="form-content">
    {#each formConfig?.fields || [] as field}
      <div class="form-group">
        {#if field.type !== "list"}
          <input
            id={field.key}
            type={field.type}
            value={form?.[field.key]?.value}
            class={`form-input ${
              field.validation
                ? form?.[field.key]?.["valid"]
                  ? "input-valid"
                  : "input-invalid"
                : ""
            }`}
            placeholder={`${field.display}${field.validation ? " *" : ""}`}
            on:input={(event) => onInput(field, event.target.value)}
          />
          {#if !form?.[field.key]?.["valid"]}
            <p class="error-message">{field?.validation?.message}</p>
          {/if}
        {:else}
          <div class="list-input">
            <div>
              <div class="list-input-header">{field.display}</div>
              {#if !form?.[field.key]?.["valid"]}
                <div class="error-message">{field?.validation?.message}</div>
              {/if}
            </div>
            {#if field.subtype === "datetime"}
              <div>
                {#each form?.[field.key]?.["value"] || [] as dateTimeEntry}
                  <ul>
                    <li
                      class="datetime-entry"
                      on:click={() => removeDateTimeEntry(field, dateTimeEntry)}
                    >
                      {`${resolveDay(dateTimeEntry.day)} | ${dateTimeEntry.time}`}
                    </li>
                  </ul>
                {/each}
              </div>
              <DateTimePicker
                datePrefix={i18n("form.dateTimePicker.datePrefix")}
                timePrefix={i18n("form.dateTimePicker.timePrefix")}
                setOnButton={true}
                setButtonText={i18n("form.dateTimePicker.addEntryButtonText")}
                onDateTimeSet={(dateTime) => onDateTimeSet(field, dateTime)}
              />
            {/if}
          </div>
        {/if}
      </div>
    {/each}

    <div class="form-actions">
      <button
        type="button"
        on:click={formConfig?.actions?.cancel}
        class="btn-cancel"
      >
        {i18n("form.cancel")}
      </button>
      <button
        type="button"
        disabled={!isFormValid}
        on:click={submitForm}
        class="btn-submit {isFormValid ? '' : 'btn-disabled'}"
      >
        {i18n("form.save")}
      </button>
    </div>
  </form>
</div>

<style>
  .form-container {
    max-width: 700px;
    margin: 3rem auto;
    background-color: white;
    box-shadow: 0 6px 30px rgba(0, 0, 0, 0.1);
    border-radius: 20px;
    padding: 40px;
  }

  .form-header {
    color: var(--primary-color);
    font-size: 2rem;
    font-weight: bold;
    margin-bottom: 28px;
  }

  .form-content {
    display: flex;
    flex-direction: column;
    gap: 1rem;
  }

  .form-group {
    display: flex;
    flex-direction: column;
    gap: 0.1rem;
  }

  .form-input {
    width: 100%;
    padding: 0.7rem;
    border-radius: 1rem;
    border: 2px solid #e0e0e0;
    transition:
      border-color 0.3s,
      box-shadow 0.3s;
    outline: none;
  }

  .input-valid {
    border-color: var(--primary-color);
  }

  .input-invalid {
    border-color: var(--error-color);
  }

  .error-message {
    color: var(--error-color);
    font-size: 0.9rem;
  }

  .form-actions {
    display: flex;
    justify-content: flex-end;
    gap: 16px;
    flex-wrap: wrap;
  }

  .btn-cancel {
    padding: 12px 28px;
    border-radius: 14px;
    background-color: var(--secondary-color);
    color: var(--text-color);
    border: none;
    cursor: pointer;
    transition:
      background-color 0.3s,
      transform 0.2s;
  }

  .btn-cancel:hover {
    background-color: var(--secondary-color-darker);
    transform: translateY(-2px);
  }

  .btn-submit {
    padding: 12px 28px;
    border-radius: 14px;
    border: none;
    color: white;
    font-weight: 700;
    background-color: var(--primary-color);
    cursor: pointer;
    transition:
      background-color 0.3s,
      transform 0.2s,
      opacity 0.3s;
  }

  .btn-submit:hover {
    background-color: var(--primary-color-dark);
    transform: translateY(-2px);
  }

  .btn-submit:disabled,
  .btn-disabled {
    background-color: var(--primary-color-light);
    cursor: not-allowed;
  }

  .list-input {
    display: flex;
    flex-direction: column;
    align-items: start;
    gap: 0.5rem;
    background-color: var(--secondary-color-light);
    padding: 1rem;
    border-radius: 1rem;
  }

  .list-input-header {
    color: var(--text-color);
    font-size: 1.2rem;
    font-weight: bolder;
  }

  .datetime-entry {
    color: var(--text-color);
    cursor: pointer;
    transition: all 0.2s;
  }

  .datetime-entry:hover {
    color: var(--error-color);
    text-decoration: line-through;
    transform: scale(1.05);
  }
</style>
