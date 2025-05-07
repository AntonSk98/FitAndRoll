<script>
    export let onDateRangePicked;
    import { DatePicker } from "@svelte-plugins/datepicker";
    import { format } from "date-fns";
    import { i18n } from "./i18n";

    export let align = "right";

    let isOpenFrom = false;
    let isOpenTo = false;

    let from;
    let to;

    let formattedFrom = "";
    let formattedTo = "";

    const toggleDatePicker = (dateType) => {
        if (dateType === "from") {
            isOpenFrom = !isOpenFrom;
        } else if (dateType === "to") {
            isOpenTo = !isOpenTo;
        }
    };

    const handleDateChange = (date, dateType) => {
        if (!date) return "";

        const selectedDate = new Date(date.startDate);

        if (dateType === "from") {
            from = selectedDate.toISOString();
            formattedFrom = format(selectedDate, "yyyy-MM-dd");
        }
        if (dateType === "to") {
            selectedDate.setHours(23);
            selectedDate.setMinutes(59);
            selectedDate.setSeconds(59);
            selectedDate.setMilliseconds(999);
            to = selectedDate.toISOString();
            formattedTo = format(selectedDate, "yyyy-MM-dd");
        }

        // Emit the updated date range with both dates as ISO strings
        onDateRangePicked({ from, to });
    };

    const clearDate = (dateType) => {
        if (dateType === "from") {
            from = null;
            formattedFrom = "";
        } else if (dateType === "to") {
            to = null;
            formattedTo = "";
        }

        onDateRangePicked({ from, to });
    };
</script>

<div class="flex gap-1 w-fit">
    <DatePicker
        bind:isOpen={isOpenFrom}
        startDate={from}
        monthLabels={i18n("datePicker.months")}
        dowLabels={i18n("datePicker.days")}
        {align}
        includeFont={false}
        onDateChange={(date) => handleDateChange(date, "from")}
    >
        <div class="input-container">
            {#if formattedFrom}
                <button class="clear-btn" on:click={() => clearDate("from")}
                    >&times;</button
                >
            {/if}
            <input
                type="text"
                placeholder={i18n("datePicker.from")}
                bind:value={formattedFrom}
                on:click={() => toggleDatePicker("from")}
                readonly
            />
        </div>
    </DatePicker>

    <DatePicker
        bind:isOpen={isOpenTo}
        startDate={to}
        {align}
        includeFont={false}
        enableFutureDates={true}
        monthLabels={i18n("datePicker.months")}
        dowLabels={i18n("datePicker.days")}
        onDateChange={(date) => handleDateChange(date, "to")}
    >
        <div class="input-container">
            {#if formattedTo}
                <button class="clear-btn" on:click={() => clearDate("to")}
                    >&times;</button
                >
            {/if}
            <input
                type="text"
                placeholder={i18n("datePicker.to")}
                bind:value={formattedTo}
                on:click={() => toggleDatePicker("to")}
                readonly
            />
        </div>
    </DatePicker>
</div>

<style>
    .input-container {
        cursor: pointer;
        display: inline-block;
        padding: 0.2rem;
        font-size: 0.875rem;
        border: 2px solid var(--secondary-color-darker);
        border-radius: 0.4rem;
        transition:
            border-color 0.3s ease-in-out,
            box-shadow 0.2s ease-in-out;
    }

    input {
        cursor: pointer;
        border: none;
        outline: none;
        font-weight: bold;
        color: var(--text-color);
        max-width: 6rem;
        min-height: 1.5rem;
    }

    .clear-btn {
        cursor: pointer;
        background: var(--primary-color);
        border: none;
        color: white;
        padding: 2px 8px;
        border-radius: 50%;
        transition: background 0.2s ease-in-out;
    }

    .clear-btn:hover {
        background: var(--primary-color-dark);
    }
</style>
