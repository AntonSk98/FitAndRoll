import { writable } from 'svelte/store'

export const toasts = writable([])

export const dismissToast = (id) => {
    toasts.update((all) => all.filter((t) => t.id !== id))
}

export const addToast = (toast) => {
    const id = Math.floor(Math.random() * 10000)

    const defaults = {
        id,
        type: 'info',
        message: 'Unexpected error occurred...'
    }

    const toastWithDefaults = { ...defaults, ...toast }
    toasts.update((all) => [toastWithDefaults, ...all])

    setTimeout(() => dismissToast(id), 50000)
}
