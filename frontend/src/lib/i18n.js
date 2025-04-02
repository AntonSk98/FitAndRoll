import { register, init, getLocaleFromNavigator } from 'svelte-i18n';

register('de', () => import('./locales/de/course.json'));
register('de', () => import('./locales/de/common.json'));
register('de', () => import('./locales/de/archive.json'));
register('de', () => import('./locales/de/statistics.json'));
register('de', () => import('./locales/de/participants.json'));

register('en', () => import('./locales/en/course.json'));
register('en', () => import('./locales/en/common.json'));
register('en', () => import('./locales/en/archive.json'));
register('en', () => import('./locales/en/statistics.json'));
register('en', () => import('./locales/en/participants.json'));


init({
  initialLocale: 'de',
  fallbackLocale: 'de'
});
