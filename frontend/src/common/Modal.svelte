<script>
    import { onMount } from "svelte";
    import { fade } from "svelte/transition";
    import { t } from "svelte-i18n";


    export let modalSize = "modal-small";
    export let onModalConfirmed = null;
    export let onModalCanceled = null;

    onMount(() => {
        document.documentElement.style.overflow = "hidden";

        return () => {
            document.documentElement.style.overflow = ""; // Restore scrolling when unmounted
        };
    });
</script>

<div class="modal-overlay" transition:fade={{ duration: 150 }}>
    <div class="modal-container {modalSize}">
        <div class="modal-content">
            <div class="modal-body">
                <div class="modal-icon">
                    <slot name="icon">
                        <svg
                            aria-hidden="true"
                            xmlns="http://www.w3.org/2000/svg"
                            fill="none"
                            viewBox="0 0 20 20"
                        >
                            <path
                                stroke="currentColor"
                                stroke-linecap="round"
                                stroke-linejoin="round"
                                stroke-width="2"
                                d="M10 11V6m0 8h.01M19 10a9 9 0 1 1-18 0 9 9 0 0 1 18 0Z"
                            />
                        </svg>
                    </slot>
                </div>
                <h3 class="modal-text">
                    <slot name="body">
                        {$t('modal.body')}
                    </slot>
                </h3>

                <div class="flex gap-2 flex-wrap justify-center">
                    {#if onModalConfirmed}
                        <button
                            on:click={onModalConfirmed}
                            class="modal-confirm-btn"
                        >
                            <slot name="confirm">{$t('modal.actions.confirm')}</slot>
                        </button>
                    {/if}

                    {#if onModalCanceled}
                        <button
                            on:click={onModalCanceled}
                            class="modal-cancel-btn"
                        >
                            <slot name="cancel">{$t('modal.actions.cancel')}</slot>
                        </button>
                    {/if}
                </div>
            </div>
        </div>
    </div>
</div>

<style>
    .modal-overlay {
        position: fixed;
        inset: 0;
        z-index: 50;
        display: flex;
        justify-content: center;
        align-items: center;
        width: 100%;
        height: 100%;
        background: rgba(0, 0, 0, 0.5);
        overflow: hidden;
    }

    .modal-container {
        position: relative;
        padding: 16px;
        width: 100%;
        max-width: 40%;
        max-height: 85vh;
        overflow: auto;
        cursor: default;
    }

    .modal-small {
        max-width: 30vw;
    }

    @media (max-width: 1400px) {
        .modal-small {
            max-width: 40vw;
        }
    }

    @media (max-width: 1200px) {
        .modal-small {
            max-width: 50vw;
        }
    }

    @media (max-width: 1000px) {
        .modal-small {
            max-width: 60vw;
        }
    }

    @media (max-width: 800px) {
        .modal-small {
            max-width: 70vw;
        }
    }

    @media (max-width: 600px) {
        .modal-small {
            max-width: 80vw;
        }
    }

    .modal-medium {
        max-width: 50vw;
    }

    .modal-content {
        background: white;
        border-radius: 8px;
        box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
    }

    .modal-body {
        padding: 16px;
        text-align: center;
    }

    .modal-icon {
        margin: 0 auto 16px;
        color: var(--secondary-color-dark);
        width: 48px;
        height: 48px;
    }

    .modal-text {
        margin-bottom: 20px;
        font-size: 1rem;
        font-weight: 400;
        color: var(--text-color);
    }

    .modal-confirm-btn {
        background: var(--primary-color);
        color: white;
        padding: 10px 20px;
        font-size: 0.875rem;
        font-weight: 500;
        border-radius: 8px;
        border: none;
        cursor: pointer;
        transition: all 0.4s;
    }

    .modal-confirm-btn:hover {
        background: var(--primary-color-dark);
    }

    .modal-cancel-btn {
        padding: 10px 20px;
        font-size: 0.875rem;
        font-weight: 500;
        color: var(--text-color);
        background: white;
        border: 1px solid var(--secondary-color-darker);
        border-radius: 8px;
        cursor: pointer;
        transition: all 0.4s;
    }

    .modal-cancel-btn:hover {
        background-color: var(--secondary-color-darker);
    }
</style>
