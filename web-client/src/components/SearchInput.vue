<template>
  <fwb-input
      label="Search"
      :disabled="loading"
      :modelValue="query"
      @update:modelValue="query => updateQuery(query)"
      placeholder="Search existing keyword"
      size="lg"
      @keyup.enter="handleSearch"
  >
    <template #prefix>
      <svg aria-hidden="true" class="w-5 h-5 text-gray-500 dark:text-gray-400" fill="none" stroke="currentColor"
           viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
        <path d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" stroke-linecap="round" stroke-linejoin="round"
              stroke-width="2"/>
      </svg>
    </template>
    <template #suffix>
      <fwb-button @click="handleSearch" :disabled="loading">Search</fwb-button>
    </template>
  </fwb-input>
</template>

<script lang="ts" setup>
import {ref, watch} from 'vue'
import {FwbButton, FwbInput} from 'flowbite-vue'

const emit = defineEmits(['update:modelValue', 'search'])
const props = defineProps({
  modelValue: String,
  loading: {
    type: Boolean,
    default: false
  }
});
const query = ref(props.modelValue)

// Watch for changes on the query and emit an update to the parent
watch(query, (newValue) => {
  emit('update:modelValue', newValue)
});

// Update query based on internal changes or props update
const updateQuery = (value) => {
  query.value = value;
};

const handleSearch = () => {
  emit('search', query.value)
}

// Watch for external changes to modelValue prop
watch(() => props.modelValue, (newVal) => {
  query.value = newVal;
});
</script>
