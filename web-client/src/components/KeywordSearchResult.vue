<template>
  <div class="flex gap-5 flex-col items-center p-5">
    <div class="w-[100%]">
      <FileInput @upload="onUpload"></FileInput>
    </div>
    <div class="w-[100%]">
      <SearchInput v-model="searchQuery" @search="onSearch" :loading="loading" class="max-w-none"></SearchInput>
    </div>
    <FwbCard class="result-table-card">
      <SearchResultTable @download="onDownload" :loading="loading" :search-data="searchData"/>
    </FwbCard>
    <FwbPagination v-model="currentPage" @page-changed="onChangePage" :total-pages="totalPage"
                   show-icons></FwbPagination>
  </div>
</template>

<script setup>
import {onMounted, ref} from 'vue'
import SearchResultTable from "@/components/SearchResultTable.vue";
import SearchInput from "@/components/SearchInput.vue";
import {FwbCard, FwbPagination} from "flowbite-vue";
import {getKeywordPage, uploadCsv} from "@/services/keyword.js";
import FileInput from "@/components/FileInput.vue";
import {showToast} from "@/utility/toast.js";

const searchQuery = ref("");
const loading = ref(false);
const totalPage = ref(10)
const pageSize = ref(10)
const currentPage = ref(1)
const searchData = ref([])

const onSearch = () => {
  if (!loading.value) {
    loadData()
  }
}
const onChangePage = () => {
  if (!loading.value) {
    loadData()
  }
}

const onUpload = async (file) => {
  try {
    loading.value = true;
    await uploadCsv(file);
    await loadData()
    showToast.success("Keyword successfully upload. Please wait for scraping to complete.");
  } finally {
    loading.value = false
  }
}

const onDownload = (searchResultId) => {
  console.log(searchResultId)
}

const loadData = async () => {
  try {
    loading.value = true;
    const response = await getKeywordPage(currentPage.value, pageSize.value, searchQuery.value)
    searchData.value = response.data;
    searchQuery.value = response.pagination?.filter?.KeywordSearch || "";
    pageSize.value = response.pagination?.pageSize || 10;
    currentPage.value = response.pagination?.page || 1;
  } finally {
    loading.value = false
  }
}

onMounted(() => {
  loadData()
})


</script>
<style scoped>
.result-table-card {
  width: 100%;
  max-width: none;
}
</style>
