/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "/app/**/*.{js,ts,jsx,tsx,mdx}",
    "./views/**/*.{js,ts,jsx,tsx,mdx}",
    "./components/**/*.{js,ts,jsx,tsx,mdx}",
 
    // Or if using `src` directory:
    "./src/**/*.{js,ts,jsx,tsx,mdx}",
  ],
  theme: {
    colors:{
      'custom-orange': '#FFA13A',
      'custom-purple': '#2E329B'
    },
    extend: {
    },
  },
  plugins: [],
}
