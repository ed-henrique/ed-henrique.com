const form = document.getElementsByTagName("form");

form[0].addEventListener("submit", (e) => {
  e.preventDefault();

  const id = window.location.href.split("/").at(-1);
  fetch(`/admin/posts/${id}`, {
    method: "PATCH",
    headers: {
      "Content-Type": "application/x-www-form-urlencoded",
    },
    body: new URLSearchParams(new FormData(form[0])),
  }).then((res) => {
    if (!res.ok) {
      alert("It was not possible to update the post.");
      return;
    }

    window.location.href = "/admin";
  });
});
