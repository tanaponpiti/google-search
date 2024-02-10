import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import Vue3Toastify from 'vue3-toastify';
import 'vue3-toastify/dist/index.css';
import { createPinia } from 'pinia'
import { useAuthStore } from '@/stores/auth';
import './index.css'

const app = createApp(App)
const pinia = createPinia()

app.use(pinia)
app.use(router)
app.use(Vue3Toastify, {
    position: "top-right",
    timeout: 3000,
});

const authStore = useAuthStore();
authStore.initialize();

app.mount('#app')
