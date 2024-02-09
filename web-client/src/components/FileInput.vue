<template>
  <div>
    <label>Upload csv file to scrape keyword</label>
    <div class="relative">
      <FwbFileInput v-model="file">
      </FwbFileInput>
      <div class="absolute inset-y-0 right-0 flex items-center pr-2.5">
        <FwbButton color="green" @click="upload">Scrape</FwbButton>
      </div>
    </div>
  </div>
</template>
<script setup>
import {computed, ref} from 'vue';
import {FwbButton, FwbFileInput} from "flowbite-vue";

const emit = defineEmits(["upload"])
const file = ref(null);

const haveFile = computed(() => {
  return file.value != null
})

const upload = () => {
  if (file.value != null) {
    emit("upload", file.value)
  }else{

  }
}

</script>
<style scoped>
</style>
