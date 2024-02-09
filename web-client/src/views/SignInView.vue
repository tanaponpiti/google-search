<template>
  <div class="flex items-center justify-center min-h-screen bg-gray-100">
    <div class="w-full max-w-sm p-6 m-auto bg-white rounded-md shadow-md transition-all duration-500">
      <h1 class="text-3xl font-semibold text-center text-gray-700" v-if="isSignUpMode">Sign Up</h1>
      <h1 class="text-3xl font-semibold text-center text-gray-700" v-else>Login</h1>
      <form class="mt-6" @submit.prevent="formSubmit">
        <div class="flex flex-col gap-4">
          <div v-if="errorMessage" class="p-3 bg-red-100 text-red-700 rounded">
            {{ errorMessage }}
          </div>
          <div v-if="isSignUpMode" class="flex flex-col gap-2">
            <label for="name" class="block text-sm text-gray-800">Name</label>
            <FwbInput :disabled="loading" v-model="name" required type="text" id="name" placeholder="Your Name"/>
          </div>
          <div class="flex flex-col gap-2">
            <label for="username" class="block text-sm text-gray-800">Username</label>
            <FwbInput :disabled="loading" v-model="username" required type="text" id="username"
                      placeholder="Your Username"/>
          </div>
          <div class="flex flex-col gap-2">
            <label for="password" class="block text-sm text-gray-800">Password</label>
            <FwbInput :disabled="loading" v-model="password" required type="password" id="password"
                      placeholder="Your Password"/>
          </div>
          <div>
            <FwbButton loading-position="suffix" :loading="loading" :disabled="loading" type="submit"
                       class="w-full justify-center px-4 py-2 tracking-wide">{{
                isSignUpMode ? 'Sign Up' : 'Sign In'
              }}
              <template #suffix>
              </template>
            </FwbButton>
          </div>
        </div>
        <div class="mt-4 text-center">
          <button :disabled="loading" type="button" @click="toggleMode" class="text-blue-600 hover:underline">
            {{ isSignUpMode ? 'Already have an account? Sign In' : 'Don\'t have an account? Sign Up' }}
          </button>
        </div>
      </form>
    </div>
  </div>
</template>

<script setup>
import {FwbInput, FwbButton} from 'flowbite-vue'
import {ref} from 'vue'
import {useRouter} from 'vue-router';
import {useAuthStore} from "@/stores/auth.js";

const router = useRouter();
const authStore = useAuthStore();
const loading = ref(false)
const name = ref("")
const username = ref("")
const password = ref("")
const isSignUpMode = ref(false)
const errorMessage = ref("") // Define a ref for storing error messages

async function formSubmit() {
  errorMessage.value = ""; // Reset error message on new form submission
  try {
    loading.value = true;
    if (isSignUpMode.value) {
      await submitSignUpForm();
    } else {
      await submitLoginForm();
    }
  } finally {
    loading.value = false;
  }
}

async function submitLoginForm() {
  console.log("Login with:", username.value, password.value);
  try {
    await authStore.login(username.value, password.value)
    await router.push({name: 'Home'})
  } catch (error) {
    errorMessage.value = "Invalid username or password.";
  }
}

function submitSignUpForm() {
  console.log("Sign up with:", name.value, username.value, password.value);
}

function toggleMode() {
  isSignUpMode.value = !isSignUpMode.value;
  errorMessage.value = "";
}
</script>

<style scoped>
</style>
