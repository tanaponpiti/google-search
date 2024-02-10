<template>
  <div>
    <div class="sticky top-0 z-50 w-full bg-white shadow-md">
      <AppNavbar></AppNavbar>
    </div>
    <div class="max-w-screen-xl mx-auto p-4">
      <KeywordSearchResult></KeywordSearchResult>
    </div>
  </div>
</template>

<script setup>
import KeywordSearchResult from "@/components/KeywordSearchResult.vue";
import AppNavbar from "@/components/AppNavbar.vue";
</script>
<style scoped>
</style>
