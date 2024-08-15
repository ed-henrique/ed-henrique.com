document.getElementById("toggle-theme").addEventListener("click", (e) => {
  e.preventDefault();

  let currentTheme = localStorage.getItem("theme");

  if (currentTheme === null) {
    currentTheme = "dark";
    localStorage.setItem("theme", "dark");
  }

  if (currentTheme === "dark") {
    localStorage.setItem("theme", "light");
    document.body.classList.remove("dark");
  } else {
    localStorage.setItem("theme", "dark");
    document.body.classList.add("dark");
  }
});

document.getElementById("toggle-theme-mobile").addEventListener("click", (e) => {
  e.preventDefault();

  let currentTheme = localStorage.getItem("theme");

  if (currentTheme === null) {
    currentTheme = "dark";
    localStorage.setItem("theme", "dark");
  }

  if (currentTheme === "dark") {
    localStorage.setItem("theme", "light");
    document.body.classList.remove("dark");
  } else {
    localStorage.setItem("theme", "dark");
    document.body.classList.add("dark");
  }
});
