/** @type {import('tailwindcss').Config} */
export default {
  content: [
    "./index.html",
    "./src/**/*.{js,ts,jsx,tsx}",
  ],
  theme: {
    extend: {
      colors: {
        darkBg: "#0B0C10",
        darkCard: "#1F2833",
        neonCyan: "#66FCF1",
        neonTeal: "#45A29E",
        neonPurple: "#B026FF",
        textColor: "#C5C6C7",
      },
      boxShadow: {
        'neon': '0 0 10px #45A29E, 0 0 20px #45A29E',
        'neon-hover': '0 0 15px #66FCF1, 0 0 30px #66FCF1',
      }
    },
  },
  plugins: [],
}
