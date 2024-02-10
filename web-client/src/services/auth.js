import axios from 'axios';
const API_URI = window.apiBaseUrl;

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

/**
 * Send a GET request to retrieve user information from the backend.
 * @param {String} token User's token
 * @returns {Promise<Object>} The response from the server, containing user information.
 */
export async function getUserInfo(token) {
    try {
        const response = await axios.get(`${API_URI}/api/auth/me`, {
            headers: {
                'Authorization': `Bearer ${token}`,
            }
        });
        return response.data;
    } catch (error) {
        console.error("Get user info error:", error.response ? `${error.response.status} ${error.response.statusText}` : error);
        throw error; // Rethrow to allow caller to handle
    }
}

/**
 * Send a signup request to the backend.
 * @param {String} username User's username
 * @param {String} password User's password
 * @param {String} name User's name
 * @returns {Promise<Object>} The response from the server, including user data if successful.
 */
export async function signUp(username, password, name) {
    try {
        const response = await axios.post(`${API_URI}/api/auth/signup`, {
            username,
            password,
            name
        }, {
            headers: {
                'Content-Type': 'application/json',
            }
        });
        return response.data;
    } catch (error) {
        console.error("Signup error:", error.response ? `${error.response.status} ${error.response.statusText}` : error);
        throw error; // Rethrow to allow caller to handle
    }
}