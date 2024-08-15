import catppuccin from "@catppuccin/tailwindcss";

/** @type {import('tailwindcss').Config} */
export default {
  darkMode: "selector",
  content: ["./pages/**/*.templ"],
  theme: {
    extend: {}
  },
  safelist: [
    "border-ctp-sapphire",
    "border-ctp-surface2",

    "bg-ctp-mantle",
    "bg-ctp-crust",
    
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
