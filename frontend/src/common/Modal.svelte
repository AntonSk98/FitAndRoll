<script>
    import { fade } from "svelte/transition";

    export let onModalConfirmed;
    export let onModalCanceled;
</script>

<div class="modal-overlay" transition:fade={{ duration: 150 }}>
    <div class="modal-container">
        <div class="modal-content">
            <div class="modal-body">
                <!-- Icon -->
                <svg
                    class="modal-icon"
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
                <h3 class="modal-text">
                    Are you sure you want to perform this action?
                </h3>

                <button on:click={onModalConfirmed} class="modal-confirm-btn"
                    >Yes, I'm sure</button
                >

                <button on:click={onModalCanceled} class="modal-cancel-btn"
                    >No, cancel</button
                >
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
    }

    .modal-container {
        position: relative;
        padding: 16px;
        width: 100%;
        max-width: 400px;
        max-height: 100%;
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
        color: #9ca3af;
        width: 48px;
        height: 48px;
    }

    .modal-text {
        margin-bottom: 20px;
        font-size: 1rem;
        font-weight: 400;
        color: #6b7280;
    }

    .modal-confirm-btn {
        background: #10b981;
        color: white;
        padding: 10px 20px;
        font-size: 0.875rem;
        font-weight: 500;
        border-radius: 8px;
        border: none;
        cursor: pointer;
    }

    .modal-confirm-btn:hover {
        background: #059669;
    }

    .modal-cancel-btn {
        margin-left: 8px;
        padding: 10px 20px;
        font-size: 0.875rem;
        font-weight: 500;
        color: #111827;
        background: white;
        border: 1px solid #e5e7eb;
        border-radius: 8px;
        cursor: pointer;
    }

    .modal-cancel-btn:hover {
        background: #f3f4f6;
        color: #065f46;
    }
</style>
