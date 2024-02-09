<template>
  <div class="flex gap-5 flex-col items-center p-5">
    <div class="w-[100%]">
      <FileInput></FileInput>
    </div>
    <div class="w-[100%]">
      <SearchInput v-model="searchQuery" @search="onSearch" :loading="loading" class="max-w-none"></SearchInput>
    </div>
    <FwbCard class="max-w-none w-[100%]">
      <SearchResultTable @download="onDownload" :loading="loading" :search-data="searchData"/>
    </FwbCard>
    <FwbPagination v-model="currentPage" @page-changed="onChangePage" :total-pages="totalPage"
                   show-icons></FwbPagination>
  </div>
</template>

<script setup>
import {ref} from 'vue'
import SearchResultTable from "@/components/SearchResultTable.vue";
import SearchInput from "@/components/SearchInput.vue";
import {FwbCard, FwbFileInput, FwbPagination} from "flowbite-vue";
import {getKeywordPage} from "@/services/keyword.js";
import FileInput from "@/components/FileInput.vue";

const file = ref()
const searchQuery = ref("");
const loading = ref(false);
const totalPage = ref(10)
const pageSize = ref(10)
const currentPage = ref(1)
const searchData = ref([
  {
    "ID": 44,
    "CreatedAt": "2024-02-08T22:14:53.039476+07:00",
    "UpdatedAt": "2024-02-09T00:51:33.591324+07:00",
    "KeywordText": "fault",
    "SearchResults": [
      {
        "ID": 209,
        "CreatedAt": "2024-02-09T00:51:33.608023+07:00",
        "UpdatedAt": "2024-02-09T00:51:44.596094+07:00",
        "KeywordID": 44,
        "AdWordsCount": 0,
        "TotalLinks": 115,
        "TotalResults": "About 1,430,000,000 results (0.43 seconds) ",
        "SearchDate": "2024-02-09T00:51:44.556365+07:00",
        "PageData": null,
        "Status": "COMPLETED"
      },
      {
        "ID": 108,
        "CreatedAt": "2024-02-08T22:27:50.347159+07:00",
        "UpdatedAt": "2024-02-08T22:28:02.525749+07:00",
        "KeywordID": 44,
        "AdWordsCount": 0,
        "TotalLinks": 124,
        "TotalResults": "About 1,370,000,000 results (0.30 seconds) ",
        "SearchDate": "2024-02-08T22:28:02.50669+07:00",
        "PageData": null,
        "Status": "COMPLETED"
      },
      {
        "ID": 9,
        "CreatedAt": "2024-02-08T22:14:53.056427+07:00",
        "UpdatedAt": "2024-02-08T22:15:22.263472+07:00",
        "KeywordID": 44,
        "AdWordsCount": 0,
        "TotalLinks": 124,
        "TotalResults": "About 1,370,000,000 results (0.37 seconds) ",
        "SearchDate": "2024-02-08T22:15:22.225739+07:00",
        "PageData": null,
        "Status": "COMPLETED"
      }
    ]
  },
  {
    "ID": 43,
    "CreatedAt": "2024-02-08T22:14:53.039476+07:00",
    "UpdatedAt": "2024-02-09T00:51:33.591323+07:00",
    "KeywordText": "student",
    "SearchResults": [
      {
        "ID": 208,
        "CreatedAt": "2024-02-09T00:51:33.605711+07:00",
        "UpdatedAt": "2024-02-09T00:51:41.760111+07:00",
        "KeywordID": 43,
        "AdWordsCount": 0,
        "TotalLinks": 100,
        "TotalResults": "About 10,370,000,000 results (0.26 seconds) ",
        "SearchDate": "2024-02-09T00:51:41.714815+07:00",
        "PageData": null,
        "Status": "COMPLETED"
      },
      {
        "ID": 107,
        "CreatedAt": "2024-02-08T22:27:50.344167+07:00",
        "UpdatedAt": "2024-02-08T22:27:59.699582+07:00",
        "KeywordID": 43,
        "AdWordsCount": 0,
        "TotalLinks": 99,
        "TotalResults": "About 12,750,000,000 results (0.31 seconds) ",
        "SearchDate": "2024-02-08T22:27:59.664506+07:00",
        "PageData": null,
        "Status": "COMPLETED"
      },
      {
        "ID": 8,
        "CreatedAt": "2024-02-08T22:14:53.055277+07:00",
        "UpdatedAt": "2024-02-08T22:15:17.445139+07:00",
        "KeywordID": 43,
        "AdWordsCount": 0,
        "TotalLinks": 99,
        "TotalResults": "About 12,810,000,000 results (0.39 seconds) ",
        "SearchDate": "2024-02-08T22:15:17.415001+07:00",
        "PageData": null,
        "Status": "COMPLETED"
      }
    ]
  },
  {
    "ID": 42,
    "CreatedAt": "2024-02-08T22:14:53.039476+07:00",
    "UpdatedAt": "2024-02-09T00:51:33.591322+07:00",
    "KeywordText": "enlarge",
    "SearchResults": [
      {
        "ID": 207,
        "CreatedAt": "2024-02-09T00:51:33.601214+07:00",
        "UpdatedAt": "2024-02-09T00:51:41.708954+07:00",
        "KeywordID": 42,
        "AdWordsCount": 0,
        "TotalLinks": 106,
        "TotalResults": "About 609,000,000 results (0.25 seconds) ",
        "SearchDate": "2024-02-09T00:51:41.68632+07:00",
        "PageData": null,
        "Status": "COMPLETED"
      },
      {
        "ID": 106,
        "CreatedAt": "2024-02-08T22:27:50.342116+07:00",
        "UpdatedAt": "2024-02-08T22:27:58.798999+07:00",
        "KeywordID": 42,
        "AdWordsCount": 0,
        "TotalLinks": 87,
        "TotalResults": "About 431,000,000 results (0.29 seconds) ",
        "SearchDate": "2024-02-08T22:27:58.77166+07:00",
        "PageData": null,
        "Status": "COMPLETED"
      },
      {
        "ID": 7,
        "CreatedAt": "2024-02-08T22:14:53.054084+07:00",
        "UpdatedAt": "2024-02-08T22:15:20.683935+07:00",
        "KeywordID": 42,
        "AdWordsCount": 0,
        "TotalLinks": 166,
        "TotalResults": "About 431,000,000 results (0.33 seconds) ",
        "SearchDate": "2024-02-08T22:15:20.637762+07:00",
        "PageData": null,
        "Status": "COMPLETED"
      }
    ]
  },
  {
    "ID": 41,
    "CreatedAt": "2024-02-08T22:14:53.039476+07:00",
    "UpdatedAt": "2024-02-09T00:51:33.591321+07:00",
    "KeywordText": "loop",
    "SearchResults": [
      {
        "ID": 206,
        "CreatedAt": "2024-02-09T00:51:33.59763+07:00",
        "UpdatedAt": "2024-02-09T00:51:42.341441+07:00",
        "KeywordID": 41,
        "AdWordsCount": 0,
        "TotalLinks": 173,
        "TotalResults": "About 2,160,000,000 results (0.49 seconds) ",
        "SearchDate": "2024-02-09T00:51:42.286393+07:00",
        "PageData": null,
        "Status": "COMPLETED"
      },
      {
        "ID": 105,
        "CreatedAt": "2024-02-08T22:27:50.337392+07:00",
        "UpdatedAt": "2024-02-08T22:27:56.029729+07:00",
        "KeywordID": 41,
        "AdWordsCount": 0,
        "TotalLinks": 112,
        "TotalResults": "About 2,160,000,000 results (0.32 seconds) ",
        "SearchDate": "2024-02-08T22:27:55.996226+07:00",
        "PageData": null,
        "Status": "COMPLETED"
      },
      {
        "ID": 6,
        "CreatedAt": "2024-02-08T22:14:53.052126+07:00",
        "UpdatedAt": "2024-02-08T22:15:20.756673+07:00",
        "KeywordID": 41,
        "AdWordsCount": 0,
        "TotalLinks": 182,
        "TotalResults": "About 2,160,000,000 results (0.32 seconds) ",
        "SearchDate": "2024-02-08T22:15:20.721239+07:00",
        "PageData": null,
        "Status": "COMPLETED"
      }
    ]
  },
  {
    "ID": 90,
    "CreatedAt": "2024-02-08T22:14:53.039476+07:00",
    "UpdatedAt": "2024-02-09T00:51:27.112375+07:00",
    "KeywordText": "bible",
    "SearchResults": [
      {
        "ID": 205,
        "CreatedAt": "2024-02-09T00:51:27.125593+07:00",
        "UpdatedAt": "2024-02-09T00:51:33.207471+07:00",
        "KeywordID": 90,
        "AdWordsCount": 0,
        "TotalLinks": 122,
        "TotalResults": "About 1,620,000,000 results (0.38 seconds) ",
        "SearchDate": "2024-02-09T00:51:33.173005+07:00",
        "PageData": null,
        "Status": "COMPLETED"
      },
      {
        "ID": 204,
        "CreatedAt": "2024-02-08T22:31:39.259486+07:00",
        "UpdatedAt": "2024-02-08T22:31:42.755101+07:00",
        "KeywordID": 90,
        "AdWordsCount": 0,
        "TotalLinks": 115,
        "TotalResults": "About 1,600,000,000 results (0.42 seconds) ",
        "SearchDate": "2024-02-08T22:31:42.701006+07:00",
        "PageData": null,
        "Status": "COMPLETED"
      },
      {
        "ID": 154,
        "CreatedAt": "2024-02-08T22:27:50.411181+07:00",
        "UpdatedAt": "2024-02-08T22:29:46.811552+07:00",
        "KeywordID": 90,
        "AdWordsCount": null,
        "TotalLinks": null,
        "TotalResults": null,
        "SearchDate": "2024-02-08T22:29:46.811006+07:00",
        "PageData": null,
        "Status": "FAILED"
      },
      {
        "ID": 55,
        "CreatedAt": "2024-02-08T22:14:53.107481+07:00",
        "UpdatedAt": "2024-02-08T22:17:26.827972+07:00",
        "KeywordID": 90,
        "AdWordsCount": null,
        "TotalLinks": null,
        "TotalResults": null,
        "SearchDate": "2024-02-08T22:17:26.827215+07:00",
        "PageData": null,
        "Status": "FAILED"
      }
    ]
  },
  {
    "ID": 138,
    "CreatedAt": "2024-02-08T22:14:53.039476+07:00",
    "UpdatedAt": "2024-02-08T22:27:50.312239+07:00",
    "KeywordText": "kinship",
    "SearchResults": [
      {
        "ID": 202,
        "CreatedAt": "2024-02-08T22:27:50.514154+07:00",
        "UpdatedAt": "2024-02-08T22:29:46.615772+07:00",
        "KeywordID": 138,
        "AdWordsCount": 0,
        "TotalLinks": 104,
        "TotalResults": "About 92,100,000 results (0.31 seconds) ",
        "SearchDate": "2024-02-08T22:29:46.582446+07:00",
        "PageData": null,
        "Status": "COMPLETED"
      },
      {
        "ID": 103,
        "CreatedAt": "2024-02-08T22:14:53.136591+07:00",
        "UpdatedAt": "2024-02-08T22:16:58.263832+07:00",
        "KeywordID": 138,
        "AdWordsCount": 0,
        "TotalLinks": 104,
        "TotalResults": "About 92,100,000 results (0.36 seconds) ",
        "SearchDate": "2024-02-08T22:16:58.226852+07:00",
        "PageData": null,
        "Status": "COMPLETED"
      }
    ]
  },
  {
    "ID": 139,
    "CreatedAt": "2024-02-08T22:14:53.039476+07:00",
    "UpdatedAt": "2024-02-08T22:27:50.312239+07:00",
    "KeywordText": "devote",
    "SearchResults": [
      {
        "ID": 203,
        "CreatedAt": "2024-02-08T22:27:50.514884+07:00",
        "UpdatedAt": "2024-02-08T22:29:46.772796+07:00",
        "KeywordID": 139,
        "AdWordsCount": 0,
        "TotalLinks": 79,
        "TotalResults": "About 169,000,000 results (0.31 seconds) ",
        "SearchDate": "2024-02-08T22:29:46.732291+07:00",
        "PageData": null,
        "Status": "COMPLETED"
      },
      {
        "ID": 104,
        "CreatedAt": "2024-02-08T22:14:53.138314+07:00",
        "UpdatedAt": "2024-02-08T22:17:03.052939+07:00",
        "KeywordID": 139,
        "AdWordsCount": 0,
        "TotalLinks": 79,
        "TotalResults": "About 169,000,000 results (0.35 seconds) ",
        "SearchDate": "2024-02-08T22:17:03.01282+07:00",
        "PageData": null,
        "Status": "COMPLETED"
      }
    ]
  },
  {
    "ID": 135,
    "CreatedAt": "2024-02-08T22:14:53.039476+07:00",
    "UpdatedAt": "2024-02-08T22:27:50.312238+07:00",
    "KeywordText": "dressing",
    "SearchResults": [
      {
        "ID": 199,
        "CreatedAt": "2024-02-08T22:27:50.511912+07:00",
        "UpdatedAt": "2024-02-08T22:29:44.068335+07:00",
        "KeywordID": 135,
        "AdWordsCount": 0,
        "TotalLinks": 141,
        "TotalResults": "About 1,280,000,000 results (0.36 seconds) ",
        "SearchDate": "2024-02-08T22:29:44.005515+07:00",
        "PageData": null,
        "Status": "COMPLETED"
      },
      {
        "ID": 100,
        "CreatedAt": "2024-02-08T22:14:53.134956+07:00",
        "UpdatedAt": "2024-02-08T22:16:53.073889+07:00",
        "KeywordID": 135,
        "AdWordsCount": 0,
        "TotalLinks": 141,
        "TotalResults": "About 1,290,000,000 results (0.42 seconds) ",
        "SearchDate": "2024-02-08T22:16:53.026681+07:00",
        "PageData": null,
        "Status": "COMPLETED"
      }
    ]
  },
  {
    "ID": 137,
    "CreatedAt": "2024-02-08T22:14:53.039476+07:00",
    "UpdatedAt": "2024-02-08T22:27:50.312238+07:00",
    "KeywordText": "edge",
    "SearchResults": [
      {
        "ID": 201,
        "CreatedAt": "2024-02-08T22:27:50.513454+07:00",
        "UpdatedAt": "2024-02-08T22:29:45.232728+07:00",
        "KeywordID": 137,
        "AdWordsCount": 0,
        "TotalLinks": 126,
        "TotalResults": "About 6,610,000,000 results (0.36 seconds) ",
        "SearchDate": "2024-02-08T22:29:45.192357+07:00",
        "PageData": null,
        "Status": "COMPLETED"
      },
      {
        "ID": 102,
        "CreatedAt": "2024-02-08T22:14:53.136019+07:00",
        "UpdatedAt": "2024-02-08T22:16:55.22285+07:00",
        "KeywordID": 137,
        "AdWordsCount": 0,
        "TotalLinks": 134,
        "TotalResults": "About 6,610,000,000 results (0.33 seconds) ",
        "SearchDate": "2024-02-08T22:16:55.182519+07:00",
        "PageData": null,
        "Status": "COMPLETED"
      }
    ]
  },
  {
    "ID": 134,
    "CreatedAt": "2024-02-08T22:14:53.039476+07:00",
    "UpdatedAt": "2024-02-08T22:27:50.312238+07:00",
    "KeywordText": "paralyzed",
    "SearchResults": [
      {
        "ID": 198,
        "CreatedAt": "2024-02-08T22:27:50.511132+07:00",
        "UpdatedAt": "2024-02-08T22:29:43.04277+07:00",
        "KeywordID": 134,
        "AdWordsCount": 0,
        "TotalLinks": 119,
        "TotalResults": "About 374,000,000 results (0.33 seconds) ",
        "SearchDate": "2024-02-08T22:29:42.998119+07:00",
        "PageData": null,
        "Status": "COMPLETED"
      },
      {
        "ID": 99,
        "CreatedAt": "2024-02-08T22:14:53.134415+07:00",
        "UpdatedAt": "2024-02-08T22:16:50.549448+07:00",
        "KeywordID": 134,
        "AdWordsCount": 0,
        "TotalLinks": 131,
        "TotalResults": "About 374,000,000 results (0.45 seconds) ",
        "SearchDate": "2024-02-08T22:16:50.510014+07:00",
        "PageData": null,
        "Status": "COMPLETED"
      }
    ]
  }
])

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


</script>
<style scoped>
</style>
