import { register, init, getLocaleFromNavigator } from 'svelte-i18n';

register('de', () => import('./locales/de/course.json'));
register('de', () => import('./locales/de/common.json'));
register('de', () => import('./locales/de/archive.json'))


init({
  initialLocale: 'de',
  fallbackLocale: 'de'
});
