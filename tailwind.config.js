import catppuccin from "@catppuccin/tailwindcss";

/** @type {import('tailwindcss').Config} */
export default {
  darkMode: "selector",
  content: ["./pages/**/*.templ", "./admin/**/*.templ"],
  theme: {
    extend: {}
  },
  safelist: [
    "border-ctp-sapphire",
    "border-ctp-surface2",

    "bg-ctp-base",
    "bg-ctp-text",
    "bg-ctp-crust",
    "bg-ctp-surface2",

    "text-ctp-base",
    "text-ctp-crust",
    "text-ctp-sapphire",
    "text-ctp-surface2",
    "text-ctp-peach",
    "text-ctp-text",
    "text-ctp-overlay0",

    "ctp-mocha",
    "ctp-latte",
  ],
  plugins: [
    catppuccin({
      prefix: "ctp",
    })
  ],
}
