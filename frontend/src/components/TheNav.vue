<script setup lang="ts">
import { availableLocales, loadLanguageAsync } from '~/modules/i18n'

const { t, locale } = useI18n()

async function toggleLocales() {
  // change to some real logic
  const locales = availableLocales
  const newLocale = locales[(locales.indexOf(locale.value) + 1) % locales.length]
  await loadLanguageAsync(newLocale)
  locale.value = newLocale
}
</script>

<template>
  <nav flex="~ gap-12" justify-center py-10 text-xl>
    <RouterLink icon-btn to="/" :title="t('button.home')">
      <div i-carbon:earth-southeast-asia-filled />
    </RouterLink>

    <button icon-btn :title="t('button.toggle-dark')" @click="toggleDark()">
      <div i="carbon-sun dark:carbon-moon" />
    </button>

    <a icon-btn :title="t('button.toggle-langs')" @click="toggleLocales()">
      <div i-carbon-language />
    </a>

    <RouterLink icon-btn to="/transfer-index" :title="t('button.transfer-index')">
      <div i-bx:transfer />
    </RouterLink>

    <RouterLink icon-btn to="/balance" :title="t('button.balance')">
      <div i-cryptocurrency-color:eth />
    </RouterLink>

    <RouterLink icon-btn to="/account" :title="t('button.account')">
      <div i-logos:metamask-icon />
    </RouterLink>
  </nav>
</template>
