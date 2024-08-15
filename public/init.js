let currentTheme = localStorage.getItem("theme");

if (currentTheme === "dark") {
  document.body.classList.add("dark");
} else {
  localStorage.setItem("theme", "light");
}
