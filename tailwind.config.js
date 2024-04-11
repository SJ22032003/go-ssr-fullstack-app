/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./views/**/*.templ"],
  daisyui: {
    themes: ["light", "dark", "cupcake", "sunset", "garden"],
  },
  theme: {
    extend: {},
  },
  plugins: [],
};
