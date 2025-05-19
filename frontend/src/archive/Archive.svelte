<script>
    import {i18n} from "../common/i18n"
    import { ExportData } from "../../wailsjs/go/archive/ExportDataHandler";
    import { toastError, toastSuccess } from "../toast/toastStore.js";
    import ArchivedDataComponent from "./ArchivedDataComponent.svelte";
    import {
        FindArchivedCourses,
        FindArchivedParticipants,
        UnarchiveCourse,
        UnarchiveParticipant,
        PurgeParticipant
    } from "../../wailsjs/go/archive/ArchivedDataHandler";
    import {LogError} from "../../wailsjs/runtime/runtime"

    let displayParticipantArchive = false;
    let displayCourseArchive = false;

    function resetView() {
        displayParticipantArchive = false;
        displayCourseArchive = false;
    }

    function exportExcel() {
        ExportData()
            .then(() => toastSuccess())
            .catch((err) => {
                LogError(`Export failed: ${err}`);
                toastError();
            });
    }

    function archivedParticipantsSupplier(filter, pagination) {
        return FindArchivedParticipants(filter, pagination);
    }

    function unarchiveParticipantSupplier(participantId) {
        return UnarchiveParticipant(participantId);
    }

    function archivedCoursesSupplier(filter, pagination) {
        return FindArchivedCourses(filter, pagination);
    }

    function unarchiveCourseSupplier(courseId) {
        return UnarchiveCourse(courseId);
    }

    function purgeParticipantSupplier(participantId) {
        return PurgeParticipant(participantId);
    }
</script>

{#if !displayCourseArchive && !displayParticipantArchive}
    <div class="outer-archive-container">
        <div class="archive-container">
            <h1 class="text-4xl text-[var(--text-color)] font-bold mb-3">
                {i18n("archive.header")}
            </h1>
            <div class="flex flex-col items-start gap-2">
                <button
                    class="archive-button"
                    on:click={() => (displayParticipantArchive = true)}
                    >{i18n("archive.navigation.participants")}</button
                >
                <button
                    class="archive-button"
                    on:click={() => (displayCourseArchive = true)}
                    >{i18n("archive.navigation.courses")}</button
                >
                <div class="flex flex-col items-start">
                    <button class="archive-button" on:click={exportExcel}
                        >{i18n("archive.export")}</button
                    >
                    <span class="mt-1 text-[var(--text-color)]">
                        <svg
                            xmlns="http://www.w3.org/2000/svg"
                            fill="none"
                            viewBox="0 0 24 24"
                            stroke-width="1.5"
                            stroke="currentColor"
                            class="w-6 h-6 text-[var(--primary-color-dark)] inline"
                        >
                            <path
                                stroke-linecap="round"
                                stroke-linejoin="round"
                                d="m11.25 11.25.041-.02a.75.75 0 0 1 1.063.852l-.708 2.836a.75.75 0 0 0 1.063.853l.041-.021M21 12a9 9 0 1 1-18 0 9 9 0 0 1 18 0Zm-9-3.75h.008v.008H12V8.25Z"
                            />
                        </svg>
                        {i18n("archive.exportInfo")}
                    </span>
                </div>
            </div>
        </div>
    </div>
{/if}

{#if displayParticipantArchive}
    <ArchivedDataComponent
        header={i18n("archive.archivedParticipants.header")}
        loadArchivedDataPromise={archivedParticipantsSupplier}
        unarchiveEntryPromise={unarchiveParticipantSupplier}
        purgeEntryPromise={purgeParticipantSupplier}
        onComponentDestroyed={() => resetView()}
    />
{/if}

{#if displayCourseArchive}
    <ArchivedDataComponent
        header="{i18n("archive.archivedCourses.header")}"
        loadArchivedDataPromise={archivedCoursesSupplier}
        unarchiveEntryPromise={unarchiveCourseSupplier}
        onComponentDestroyed={() => resetView()}
    />
{/if}

<style>
    .outer-archive-container {
        max-width: var(--container-max-width);
        margin: 2rem auto;
        cursor: default;
        border-left: 0.3rem solid var(--primary-color);
    }
    .archive-container {
        padding: 1rem 0 5rem 1.5rem;
        max-width: 500px;
    }

    .archive-button {
        cursor: pointer;
        padding: 0.6rem 1rem;
        border-radius: 0.5rem;
        background-color: var(--primary-color);
        color: white;
        border: none;
        transition:
            background-color 0.5s ease-in-out,
            box-shadow 0.2s ease-in-out,
            transform 0.5s ease-in-out;
        box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
    }

    .archive-button:hover {
        background-color: var(--primary-color-dark);
        transform: scale(1.08);
    }
</style>
