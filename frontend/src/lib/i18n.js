import { register, init, getLocaleFromNavigator } from 'svelte-i18n';

register('de', () => import('./locales/de/course.json'));
register('de', () => import('./locales/de/common.json'));
register('de', () => import('./locales/de/archive.json'));
register('de', () => import('./locales/de/statistics.json'));
register('de', () => import('./locales/de/participants.json'))


init({
  initialLocale: 'de',
  fallbackLocale: 'de'
});
