import { register, init, getLocaleFromNavigator } from 'svelte-i18n';

register('de', () => import('./locales/de.json'));

init({
  initialLocale: 'de',
  fallbackLocale: 'de'
});
