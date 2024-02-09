import axios from 'axios';
import {API_URI} from "@/config/config.js";

/**
 * Send a login request to the backend.
 * @param {String} username User's username
 * @param {String} password User's password
 * @returns {Promise<Object>} The response from the server, including a token if successful.
 */
export async function login(username, password) {
    try {
        const response = await axios.post(`${API_URI}/api/auth/login`, {
            username,
            password
        }, {
            headers: {
                'Content-Type': 'application/json',
            }
        });
        return response.data;
    } catch (error) {
        console.error("Login error:", error.response ? `${error.response.status} ${error.response.statusText}` : error);
        throw error; // Rethrow to allow caller to handle
    }
}

/**
 * Send a logout request to the backend.
 * @param {String} token User's token
 */
export async function logout(token) {
    try {
        await axios.post(`${API_URI}/api/auth/logout`, {}, {
            headers: {
                'Authorization': `Bearer ${token}`
            }
        });
    } catch (error) {
        console.error("Logout error:", error);
    }
}

