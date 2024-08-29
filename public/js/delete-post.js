const deleteButtons = document.getElementsByClassName("delete-buttons");

Array.from(deleteButtons).forEach((b) => {
  b.addEventListener("click", (e) => {
    e.preventDefault();

    if (b.innerHTML === "Delete for Real?") {
      const id = b.id.split("-")[0];

      fetch(`/admin/posts/${id}`, {
        method: "DELETE",
        headers: {
          "Content-Type": "application/x-www-form-urlencoded",
        }
      }).then((res) => {
        if (!res.ok) {
          alert("It was not possible to delete the post.");
          return;
        }

        b.parentElement.parentElement.remove();
      })
    } else {
      b.innerHTML = "Delete for Real?";
    }
  });
});
