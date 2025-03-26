import { i18n } from '../common/i18n'
import { writable } from 'svelte/store'

export const toasts = writable([])

export const dismissToast = (id) => {
    toasts.update((all) => all.filter((t) => t.id !== id))
}

export const toastSuccess = (message) => {
    addToast({ type: "success", message: message ?? `${i18n('toast.success')}!`});
}

export const toastError = (message) => {
    addToast({type: 'error', message: message ?? `${i18n('toast.error')}!`});
}

const addToast = (toast) => {
    const id = Math.floor(Math.random() * 10000)

    const defaults = {
        id
    }


    const toastWithDefaults = { ...defaults, ...toast }

    toasts.update((all) => [toastWithDefaults, ...all])

    setTimeout(() => dismissToast(id), 5000)
}
