/** @type {import('tailwindcss').Config} */
/* eslint-env node */
module.exports = {
    content: [
        'node_modules/flowbite-vue/**/*.{js,jsx,ts,tsx,vue}',
        'node_modules/flowbite/**/*.{js,jsx,ts,tsx}',
        './index.html',
        './src/**/*.{vue,js,ts,jsx,tsx}'],
    theme: {
        extend: {},
    },
    variants: {
        extend: {},
    },
    plugins: [
        require('flowbite/plugin')
    ],
}