import {defineStore} from 'pinia';
import router from '@/router';
import {login as apiLogin, logout as apiLogout, getUserInfo as apiGetUserInfo} from '@/services/auth.js';
export const useAuthStore = defineStore('auth', {
    state: () => ({
        token: null,
    }),
    getters: {
        isAuthenticated(state) {
            return !!state.token;
        },
        getToken(state) {
            return state.token;
        }
    },
    actions: {
        saveToken(newToken) {
            this.token = newToken;
            localStorage.setItem('userToken', newToken);
        },
        async loadToken() {
            const storedToken = localStorage.getItem('userToken');
            if (storedToken) {
                this.token = storedToken;
            }
        },
        async checkTokeValidity() {
            if (this.token) {
                try {
                    await apiGetUserInfo(this.token)
                } catch (e) {
                    localStorage.removeItem('userToken');
                    this.token = null;
                }
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
        async forceLogout() {
            try {
                await this.logout()
            } catch (error) {
                console.error('Logout failed:', error);
            } finally {
                await router.push({name: 'Sign in'})
            }
        },
        async initialize() {
            await this.loadToken();
        }
    }
});
