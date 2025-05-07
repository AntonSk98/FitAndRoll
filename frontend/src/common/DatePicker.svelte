<script>
    import { DatePicker } from "@svelte-plugins/datepicker";
    import { format } from "date-fns";
    import { i18n } from "./i18n";

    let isDatePickerOpen = false;
    let selectedDateDisplay = "";

    export let onDatePicked;
    export let placeholder;
    export let selectorClass;
    export let selectedDateValue = undefined;


    $: if (selectedDateValue) {
        const selectedDate = new Date(selectedDateValue);
        selectedDateDisplay = format(selectedDate, "yyyy-MM-dd");
        onDatePicked(selectedDateValue);
    }

    function onDateChange(date) {
        if (!date) return "";

        const selectedDate = new Date(date.startDate);

        selectedDateValue = selectedDate.toISOString();
    }
</script>

<div class="w-fit">
    <DatePicker
        bind:isOpen={isDatePickerOpen}
        startDate={selectedDateValue}
        monthLabels={i18n("datePicker.months")}
        dowLabels={i18n("datePicker.days")}
        align={"left"}
        defaultYear={1998}
        defaultMonth={2}
        includeFont={false}
        onDateChange={(date) => onDateChange(date)}
    >
        <div
            class={`${selectorClass} cursor-pointer transition-all duration-500 hover:scale-105 text-base`}
            on:click={() => {
                isDatePickerOpen = !isDatePickerOpen;
            }}
        >
            {#if selectedDateDisplay}
                {selectedDateDisplay}
            {:else}
                <svg
                    xmlns="http://www.w3.org/2000/svg"
                    viewBox="0 0 24 24"
                    fill="currentColor"
                    class="w-6 h-6 text-[var(--text-color-darker)]"
                >
                    <path
                        d="M12.75 12.75a.75.75 0 1 1-1.5 0 .75.75 0 0 1 1.5 0ZM7.5 15.75a.75.75 0 1 0 0-1.5.75.75 0 0 0 0 1.5ZM8.25 17.25a.75.75 0 1 1-1.5 0 .75.75 0 0 1 1.5 0ZM9.75 15.75a.75.75 0 1 0 0-1.5.75.75 0 0 0 0 1.5ZM10.5 17.25a.75.75 0 1 1-1.5 0 .75.75 0 0 1 1.5 0ZM12 15.75a.75.75 0 1 0 0-1.5.75.75 0 0 0 0 1.5ZM12.75 17.25a.75.75 0 1 1-1.5 0 .75.75 0 0 1 1.5 0ZM14.25 15.75a.75.75 0 1 0 0-1.5.75.75 0 0 0 0 1.5ZM15 17.25a.75.75 0 1 1-1.5 0 .75.75 0 0 1 1.5 0ZM16.5 15.75a.75.75 0 1 0 0-1.5.75.75 0 0 0 0 1.5ZM15 12.75a.75.75 0 1 1-1.5 0 .75.75 0 0 1 1.5 0ZM16.5 13.5a.75.75 0 1 0 0-1.5.75.75 0 0 0 0 1.5Z"
                    />
                    <path
                        fill-rule="evenodd"
                        d="M6.75 2.25A.75.75 0 0 1 7.5 3v1.5h9V3A.75.75 0 0 1 18 3v1.5h.75a3 3 0 0 1 3 3v11.25a3 3 0 0 1-3 3H5.25a3 3 0 0 1-3-3V7.5a3 3 0 0 1 3-3H6V3a.75.75 0 0 1 .75-.75Zm13.5 9a1.5 1.5 0 0 0-1.5-1.5H5.25a1.5 1.5 0 0 0-1.5 1.5v7.5a1.5 1.5 0 0 0 1.5 1.5h13.5a1.5 1.5 0 0 0 1.5-1.5v-7.5Z"
                        clip-rule="evenodd"
                    />
                </svg>
            {/if}
        </div>
    </DatePicker>
</div>
