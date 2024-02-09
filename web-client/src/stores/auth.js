import { defineStore } from 'pinia';
import { login as apiLogin,logout as apiLogout } from '@/services/auth.js';

export const useAuthStore = defineStore('auth', {
    state: () => ({
        token: null,
    }),
    getters: {
        isAuthenticated(state) {
            return !!state.token;
        }
    },
    actions: {
        saveToken(newToken) {
            this.token = newToken;
            localStorage.setItem('userToken', newToken);
        },
        loadToken() {
            const storedToken = localStorage.getItem('userToken');
            if (storedToken) {
                this.token = storedToken;
            }
        },
        async login(username, password) {
            const data = await apiLogin(username, password);
            this.saveToken(data.token); // Save token
            console.log("Login successful, token stored.");
        },
        async logout() {
            try {
                await apiLogout(this.token);
                localStorage.removeItem('userToken');
                this.token = null;
            } catch (error) {
                console.error('Logout failed:', error);
            }
        },
        initialize() {
            this.loadToken();
        }
    }
});
