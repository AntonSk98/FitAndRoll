import { register, init, getLocaleFromNavigator } from 'svelte-i18n';

register('de', () => import('./locales/de.json'));
register('de', () => import('./locales/de/common.json'))

init({
  initialLocale: 'de',
  fallbackLocale: 'de'
});
