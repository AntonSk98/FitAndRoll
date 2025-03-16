<script>
    export let onDateRangePicked;
    import { DatePicker } from "@svelte-plugins/datepicker";
    import { format } from "date-fns";

    let isOpenFrom = false;
    let isOpenTo = false;

    let from = new Date();
    let to = new Date();

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
        selectedDate.setHours(
            dateType === "from" ? 0 : 23,
            dateType === "from" ? 0 : 59,
            dateType === "from" ? 0 : 59,
        );

        if (dateType === "from") {
            from = selectedDate;
            formattedFrom = format(from, "yyyy/MM/dd");
        } else {
            to = selectedDate;
            formattedTo = format(to, "yyyy/MM/dd");
        }
    };

    const clearDate = (dateType) => {
        if (dateType === "from") {
            from = null;
            formattedFrom = "";
        } else if (dateType === "to") {
            to = null;
            formattedTo = "";
        }
    };
</script>

<div class="flex flex-col gap-y-1 w-fit">
    <DatePicker
        bind:isOpen={isOpenFrom}
        bind:startDate={from}
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
                placeholder="From"
                bind:value={formattedFrom}
                on:click={() => toggleDatePicker("from")}
                readonly
            />
        </div>
    </DatePicker>

    <DatePicker
        bind:isOpen={isOpenTo}
        bind:startDate={to}
        enableFutureDates={true}
        onDateChange={(date) => handleDateChange(date, "to")}
    >
        <div class="input-container">
            {#if formattedTo}
                <button class="clear-btn" on:click={() => clearDate("to")}
                    >&times;</button
                >
            {/if}
            <input
                class="max-w-24"
                type="text"
                placeholder="To"
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
        padding: 0.375rem;
        font-size: 0.875rem;
        border: 1px solid var(--secondary-color-darker);
        border-radius: 0.5rem;
        background-color: var(--secondary-color-light);
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
        max-width: 5rem;
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
