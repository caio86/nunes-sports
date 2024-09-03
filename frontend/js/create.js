const URL = "http://localhost:5000/api/v1/products";

const form = document.body.querySelector("#createForm");
const errorMsg = document.body.querySelector(".error");

/**
 * @param {Event} event
 * */
async function createProduct(event) {
  event.preventDefault();
  const formData = new FormData(form);
  const data = { error: false };

  formData.forEach((value, key) => {
    if (value.length <= 0) {
      errorMsg.classList.remove("hidden");
      errorMsg.innerHTML = `Erro: ${key} está vazio`;
      data["error"] = true;
    }

    data[key] = value;
  });

  if (data.error) {
    return;
  }

  data.price = parseFloat(data.price);

  await fetch(URL, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(data),
  })
    .then(() => {
      form.reset();
    })
    .catch((err) => {
      errorMsg.classList.remove("hidden");
      errorMsg.innerHTML = `Erro: ${err}`;
      console.error("Error: ", err);
    });
}

function hideError() {
  errorMsg.classList.add("hidden");
}

form.addEventListener("submit", createProduct);
