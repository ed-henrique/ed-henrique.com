const form = document.getElementsByTagName("form");

form[0].addEventListener("submit", (e) => {
  e.preventDefault();

  fetch(`/admin/posts/`, {
    method: "POST",
    headers: {
      "Content-Type": "application/x-www-form-urlencoded",
    },
    body: new URLSearchParams(new FormData(form[0])),
  }).then((res) => {
    if (!res.ok) {
      alert("It was not possible to create the post.");
      return;
    }

    window.location.href = "/admin";
  });
});
