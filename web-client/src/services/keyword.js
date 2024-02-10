import axios from 'axios';
import {API_URI} from "@/config/config.js";
import {useAuthStore} from "@/stores/auth.js";
import {showToast} from "@/utility/toast.js";

/**
 * Send a login request to the backend.
 * @param {Integer} page page number of keyword
 * @param {Integer} pageSize Number of keyword per page
 * @param {String} keywordSearch keyword search filter
 * @returns {Promise<Object>} The response from the server, including a token if successful.
 */
export async function getKeywordPage(page, pageSize = 10, keywordSearch = null) {
    const authStore = useAuthStore()
    try {
        //TODO might want to separate useAuthStore out as dependency injection for easier testing in the future
        const response = await axios.post(`${API_URI}/api/keyword/search`, {
            "page": page,
            "pageSize": pageSize,
            "filter": {
                "keywordSearch": keywordSearch
            }
        }, {
            headers: {
                'Authorization': `Bearer ${authStore.getToken}`
            }
        });
        return response.data;
    } catch (error) {
        if (error.response && error.response.status === 401) {
            showToast.warning("Session expire");
            await authStore.forceLogout();
        } else {
            showToast.error("Unable to get keyword results");
            console.error("Get keyword page error:", error.response ? `${error.response.status} ${error.response.statusText}` : error);
            throw error; // Rethrow to allow caller to handle
        }
    }
}


/**
 * Send keywords to the backend for processing.
 * @param {Array<String>} keywords Array of keywords to be processed.
 * @returns {Promise<Object>} The response from the server.
 */
export async function addKeywords(keywords) {
    const authStore = useAuthStore();
    try {
        const response = await axios.post(`${API_URI}/api/keyword`, {
            keywords: keywords
        }, {
            headers: {
                'Authorization': `Bearer ${authStore.getToken}`,
                'Content-Type': 'application/json',
            }
        });
        return response.data;
    } catch (error) {
        if (error.response && error.response.status === 401) {
            await authStore.forceLogout();
        } else {
            showToast.error("Unable to add keyword");
            console.error("Add keywords error:", error.response ? `${error.response.status} ${error.response.statusText}` : error);
            throw error; // Rethrow to allow caller to handle
        }
    }
}

/**
 * Uploads a CSV file to the backend for processing.
 * @param {File} file CSV file to be uploaded.
 * @returns {Promise<Object>} The response from the server.
 */
export async function uploadCsv(file) {
    const authStore = useAuthStore();
    const formData = new FormData();
    formData.append('file', file);
    try {
        const response = await axios.post(`${API_URI}/api/keyword/csv-upload`, formData, {
            headers: {
                'Authorization': `Bearer ${authStore.getToken}`,
                'Content-Type': 'multipart/form-data', // Axios sets the correct multipart boundary automatically
            }
        });
        return response.data;
    } catch (error) {
        if (error.response && error.response.status === 401) {
            await authStore.forceLogout();
        } else {
            showToast.error("Unable to upload csv");
            console.error("CSV upload error:", error.response ? `${error.response.status} ${error.response.statusText}` : error);
            throw error; // Rethrow to allow caller to handle
        }
    }
}
