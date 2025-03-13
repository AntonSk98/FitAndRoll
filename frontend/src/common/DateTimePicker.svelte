<script>
    import { t } from "svelte-i18n";

    let hours;
    let minutes;
    let selectedDay;

    export let datePrefix;
    export let timePrefix;
    export let onDateTimeSet;
    export let setOnButton;
    export let setButtonText;

    function updateHours() {
        hours = String(Math.min(parseInt(hours) || 0, 23)).padStart(2, "0");
        if (!setOnButton) {
            triggerDateTimeSet();
        }
    }

    function updateMinutes() {
        minutes = String(Math.min(parseInt(minutes) || 0, 59)).padStart(2, "0");
        if (!setOnButton) {
            triggerDateTimeSet();
        }
    }

    function updateDay() {
        if (!setOnButton) {
            triggerDateTimeSet();
        }
    }

    function triggerDateTimeSet() {
        if (onDateTimeSet && selectedDay && hours && minutes) {
            onDateTimeSet({
                day: selectedDay,
                time: `${hours}:${minutes}`,
            });
        }
    }
</script>

<div class="datetime-picker">
    <div class="date-picker">
        {#if datePrefix}
            <label for="day-select">{datePrefix}:</label>
        {/if}
        <select id="day-select" bind:value={selectedDay} on:change={updateDay}>
            <option value="" disabled selected>{$t("dateTimePicker.options.none")}</option>
            <option value="monday">{$t("dateTimePicker.options.monday")}</option>
            <option value="tuesday">{$t("dateTimePicker.options.tuesday")}</option>
            <option value="wednesday">{$t("dateTimePicker.options.wednesday")}</option>
            <option value="thursday">{$t("dateTimePicker.options.thursday")}</option>
            <option value="friday">{$t("dateTimePicker.options.friday")}</option>
            <option value="saturday">{$t("dateTimePicker.options.saturday")}</option>
            <option value="sunday">{$t("dateTimePicker.options.sunday")}</option>
        </select>
    </div>

    <div class="time-picker">
        {#if timePrefix}
            <label for="day-select">{timePrefix}:</label>
        {/if}
        <input
            type="text"
            placeholder="HH"
            bind:value={hours}
            on:input={updateHours}
        />
        :
        <input
            type="text"
            placeholder="MM"
            bind:value={minutes}
            on:input={updateMinutes}
        />
    </div>

    {#if setOnButton}
        <button
            type="button"
            class="new-element-button"
            disabled={!(selectedDay && hours && minutes)}
            on:click={triggerDateTimeSet}
        >
            {setButtonText || "+"}
        </button>
    {/if}
</div>

<style>
    .datetime-picker {
        display: flex;
        flex-direction: column;
        gap: 0.5rem;
        border-left: 3px var(--primary-color-dark) solid;
        padding-left: 1rem;
    }

    .date-picker,
    .time-picket {
        display: flex;
        align-items: center;
        gap: 0.5rem;
        flex-wrap: wrap;
    }

    .time-picker input {
        max-width: 2.5rem;
        text-align: center;
    }

    select,
    input[type="text"] {
        padding: 0.2rem;
        border: 1px solid #ccc;
        outline: none;
    }

    input[type="text"] {
        border-radius: 0.5rem;
    }

    .new-element-button {
        color: white;
        background-color: var(--primary-color);
        border-radius: 2rem;
        padding: 0.4rem 0.7rem;
        cursor: pointer;
        width: auto;
        align-self: flex-start;
    }

    .new-element-button:hover {
        background-color: var(--primary-color-dark);
        transition: 0.5s ease;
    }

    .new-element-button:disabled {
        background-color: var(--primary-color-light);
        cursor: not-allowed;
    }

    .new-element-button:disabled:hover {
        background-color: var(--primary-color-light);
        cursor: not-allowed;
    }
</style>
