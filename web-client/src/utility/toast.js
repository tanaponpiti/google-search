import { toast } from 'vue3-toastify';

const defaultOptions = {
    position: "top-right",
    timeout: 5000,
};

export const showToast = {
    success(message) {
        toast.success(message, {
            ...defaultOptions,
        });
    },
    warning(message) {
        toast.warning(message, {
            ...defaultOptions,
        });
    },
    error(message) {
        toast.error(message, {
            ...defaultOptions,
        });
    },
};
