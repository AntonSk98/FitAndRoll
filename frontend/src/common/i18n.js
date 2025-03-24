import { t, locale, unwrapFunctionStore } from "svelte-i18n";

// Unwrap the `t` function to use it globally
const i18n = unwrapFunctionStore(t);

// Function to change the app language
function setLocale(newLocale) {
    locale.set(newLocale);
}

// Export the translation function
export { i18n, setLocale };