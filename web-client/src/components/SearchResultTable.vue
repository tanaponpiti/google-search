<template>
  <div class="relative">
    <fwb-table v-if="searchData.length > 0">
      <fwb-table-head>
        <fwb-table-head-cell>Keyword</fwb-table-head-cell>
        <fwb-table-head-cell>Links</fwb-table-head-cell>
        <fwb-table-head-cell>AdWords</fwb-table-head-cell>
        <fwb-table-head-cell>Search Result</fwb-table-head-cell>
        <fwb-table-head-cell>Status</fwb-table-head-cell>
        <fwb-table-head-cell>HTML Cache</fwb-table-head-cell>
      </fwb-table-head>
      <fwb-table-body>
        <fwb-table-row v-for="(data, index) in searchData" :key="index">
          <fwb-table-cell>{{ data.KeywordText }}</fwb-table-cell>
          <fwb-table-cell>{{ data.SearchResults[0]?.TotalLinks }}</fwb-table-cell>
          <fwb-table-cell>{{ data.SearchResults[0]?.AdWordsCount }}</fwb-table-cell>
          <fwb-table-cell>{{ data.SearchResults[0]?.TotalResults }}</fwb-table-cell>
          <fwb-table-cell>
            <fwb-badge :type="getBadgeClass(data.SearchResults[0]?.Status)">{{
                data.SearchResults[0]?.Status
              }}
            </fwb-badge>
          </fwb-table-cell>
          <fwb-table-cell>
            <div class="flex justify-center items-center">
              <fwb-a href="#" @click="download(data.SearchResults[0]?.ID)">
                <FontAwesomeIcon :icon="faDownload"/>
              </fwb-a>
            </div>
          </fwb-table-cell>
        </fwb-table-row>
      </fwb-table-body>
    </fwb-table>
    <div v-else class="flex justify-center items-center text-center p-5" style="height: 100%;">
      <div class="flex flex-col gap-5">
        <FontAwesomeIcon :icon="faSearchMinus" class="fa-5x text-gray-300"/>
        <p class="text-gray-300">Keyword not found</p>
      </div>
    </div>
    <div v-if="loading"
         class="absolute top-0 left-0 right-0 bottom-0 bg-black bg-opacity-20 flex justify-center items-center z-50">
      <FwbSpinner size="10"></FwbSpinner>
    </div>
  </div>
</template>

<script setup>
import {
  FwbA, FwbBadge, FwbSpinner,
  FwbTable,
  FwbTableBody,
  FwbTableCell,
  FwbTableHead,
  FwbTableHeadCell,
  FwbTableRow,
} from 'flowbite-vue';
import {faDownload, faSearchMinus} from "@fortawesome/free-solid-svg-icons";
import {FontAwesomeIcon} from '@fortawesome/vue-fontawesome'

const emit = defineEmits(['download'])

const props = defineProps({
  loading: {
    type: Boolean,
    default: false
  },
  searchData: {
    type: Array,
    default: () => []
  }
});

const getBadgeClass = (status) => {
  switch (status) {
    case 'COMPLETED':
      return 'green';
    case 'PENDING':
      return 'yellow';
    case 'FAILED':
      return 'red';
    default:
      return '';
  }
}

const download = (searchResultId) => {
  if (searchResultId != null) {
    emit('download', searchResultId)
  }
}
</script>

<style scoped>
/* Add styles if needed */
</style>
