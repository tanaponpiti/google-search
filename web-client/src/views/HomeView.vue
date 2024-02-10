<template>
  <div>
    <div class="sticky top-0 z-50 w-full bg-white shadow-md">
      <AppNavbar></AppNavbar>
    </div>
    <KeywordSearchResult></KeywordSearchResult>
  </div>
</template>

<script setup>
import KeywordSearchResult from "@/components/KeywordSearchResult.vue";
import AppNavbar from "@/components/AppNavbar.vue";
</script>
<style scoped>
</style>
