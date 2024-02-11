<template>
  <fwb-navbar>
    <template #logo>
      <a href="/" class="flex items-center space-x-3 rtl:space-x-reverse">
        <img src="@/assets/images/logo.png" class="h-8" alt="Search Scraper Logo" />
        <span class="self-center text-2xl font-semibold whitespace-nowrap dark:text-white">Search Scraper</span>
      </a>
    </template>
    <template #default="{isShowMenu}">
      <fwb-navbar-collapse :is-show-menu="isShowMenu">
        <fwb-navbar-link link="/sign-out">
          Logout
        </fwb-navbar-link>
      </fwb-navbar-collapse>
    </template>
  </fwb-navbar>
</template>

<script setup>
import {
  FwbNavbar,
  FwbNavbarCollapse,
  FwbNavbarLink,
  FwbNavbarLogo,
} from 'flowbite-vue'
</script>