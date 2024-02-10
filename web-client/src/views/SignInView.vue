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
          <div class="flex flex-col gap-2" v-if="isSignUpMode">
            <label for="confirmPassword" class="block text-sm text-gray-800">Confirm Password</label>
            <FwbInput :disabled="loading" v-model="confirmPassword" required type="password" id="confirmPassword" placeholder="Confirm Your Password"/>
            <div v-if="password !== confirmPassword && confirmPassword.length > 0" class="text-sm text-red-500">
              Passwords do not match.
            </div>
          </div>
          <div v-if="passwordPatternError || usernamePatternError" class="p-3 bg-red-100 text-red-700 rounded">
            <div v-if="usernamePatternError">{{ usernamePatternErrorMessage }}</div>
            <div v-if="passwordPatternError">{{ passwordPatternErrorMessage }}</div>
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
import { FwbInput, FwbButton } from 'flowbite-vue'
import { ref } from 'vue'
import { useRouter } from 'vue-router';
import { useAuthStore } from "@/stores/auth.js";
import { showToast } from "@/utility/toast.js";

const router = useRouter();
const authStore = useAuthStore();
const loading = ref(false);
const name = ref("");
const username = ref("");
const password = ref("");
const confirmPassword = ref("");
const isSignUpMode = ref(false);
const errorMessage = ref("");
const passwordPatternError = ref(false);
const usernamePatternError = ref(false);
const passwordPatternErrorMessage = "Password must be at least 8 characters long and include a mix of letters, numbers, and special characters.";
const usernamePatternErrorMessage = "Username must be 3-16 characters long and can contain letters, numbers, and underscores.";

const usernamePattern = /^[a-zA-Z0-9_]{3,16}$/;
const passwordPattern = /^(?=.*[A-Za-z])(?=.*\d)(?=.*[@$!%*#?&])[A-Za-z\d@$!%*#?&]{8,}$/;

async function formSubmit() {
  errorMessage.value = "";
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
  try {
    await authStore.login(username.value, password.value)
    await router.push({name: 'Home'})
  } catch (error) {
    if (error.response && error.response.status === 401) {
      errorMessage.value = "Invalid username or password.";
    } else {
      showToast.error("Unable to login. Something went wrong.");
    }
  }
}

function clearInput() {
  name.value = ""
  username.value = ""
  password.value = ""
}

async function submitSignUpForm() {
  passwordPatternError.value = false; // Reset pattern error
  usernamePatternError.value = false; // Reset pattern error

  if (!usernamePattern.test(username.value)) {
    usernamePatternError.value = true;
    return;
  }

  if (!passwordPattern.test(password.value)) {
    passwordPatternError.value = true;
    return;
  }

  if (isSignUpMode.value && password.value !== confirmPassword.value) {
    errorMessage.value = "Passwords do not match.";
    return;
  }

  try {
    await authStore.signupUser(username.value, password.value, name.value)
    clearInput();
    toggleMode();
  } catch (error) {
    showToast.error("Unable to sign up. Something went wrong.");
  }
}

function toggleMode() {
  isSignUpMode.value = !isSignUpMode.value;
  errorMessage.value = "";
}
</script>

<style scoped>
</style>
