<template>
  <div>
    <label>Upload csv file to scrape keyword</label>
    <div class="relative">
      <input @change="onFileChange"
             class="block w-full text-lg text-gray-900 border border-gray-300 rounded-lg cursor-pointer bg-gray-50 dark:text-gray-400 focus:outline-none dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400"
             aria-describedby="file_input_help" id="file_input" type="file" accept="text/csv">
      <div class="absolute inset-y-0 right-0 flex items-center pr-2.5">
        <FwbButton color="green" @click="upload">Scrape</FwbButton>
      </div>
    </div>
    <p class="mt-1 text-sm text-gray-500 dark:text-gray-300" id="file_input_help">CSV (MAX 100 Keywords, 2MB).</p>
  </div>
</template>
<script setup>
import {ref} from 'vue';
import {FwbButton} from "flowbite-vue";
import { showToast } from '@/utility/toast.js';
const emit = defineEmits(["upload"])
const file = ref(null);

const onFileChange = (e) => {
  const files = e.target.files || e.dataTransfer.files;
  console.log(files)
  if (files.length > 0) {
    file.value = files[0];
  } else {
    file.value = null;
  }
}

const
    upload = () => {
      if (file.value != null) {
        console.log(file.value)
        emit("upload", file.value)
      } else {
        showToast.warning("Please select some csv file for upload first.");
      }
    }

</script>
<style scoped>
</style>
