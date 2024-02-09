import { createRouter, createWebHistory } from 'vue-router';
import { useAuthStore } from '@/stores/auth.js';
import HomeView from '@/views/HomeView.vue';
import SignInView from '@/views/SignInView.vue';

const router = createRouter({
    history: createWebHistory(import.meta.env.BASE_URL),
    routes: [
        {
            path: '/',
            name: 'Home',
            component: HomeView,
            meta: { requiresAuth: true }
        },
        {
            path: '/sign-in',
            name: 'Sign in',
            component: SignInView,
            meta: { requiresUnAuth: true }
        },
        {
            path: '/sign-out',
            name: 'Sign out',
            beforeEnter: (to, from, next) => {
                const authStore = useAuthStore();
                authStore.logout().then(() => {
                    next({ name: 'Sign in' });
                });
            }
        },
        {
            path: '/:pathMatch(.*)*',
            redirect: '/',
        },
    ]
});

router.beforeEach((to, from, next) => {
    const authStore = useAuthStore();
    const isAuthenticated = authStore.isAuthenticated;
    if (to.meta.requiresAuth && !isAuthenticated) {
        next({ name: 'Sign in' });
    } else if (to.meta.requiresUnAuth && isAuthenticated) {
        next({ name: 'Home' });
    } else {
        next();
    }
});

export default router;
