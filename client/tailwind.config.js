/** @type {import('tailwindcss').Config} */
export default {
    content: ["./index.html","./src/**/*.{html,js}", "./src/main.tsx", "./src/App.tsx", "./**/*.{js,ts,jsx,tsx}", '!./**/node_modules/**'],
    theme: {
      extend: {
        colors: {
          richBlack: '#00171F',
          prussian: '#003459',
          cerulean: '#007EA7'
          
        },
        fontFamily: {
          main: ["Jost"],
        },
      },
    },
    plugins: [],
  }