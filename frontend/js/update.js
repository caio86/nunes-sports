const url = (id) => `http://localhost:5000/api/v1/products/${id}`;

const form = document.forms["updateForm"];
const errorMsg = document.body.querySelector(".error");

/**
 * @param {Event} event
 * */
async function updateProduct(event) {
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

  await fetch(url(data.id), {
    method: "PUT",
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

async function fillFormWithProduct() {
  let id = form["id"].value;

  const data = await getProduct(id);

  if (data.error) {
    errorMsg.classList.remove("hidden");
    errorMsg.innerHTML = `Error: ${data.errMsg}`;
    console.error("Error: ", data.errMsg);
  }

  form["name"].value = data.Name;
  form["description"].value = data.Description;
  form["price"].value = data.Price;
}

async function getProduct(id) {
  const data = await fetch(url(id))
    .then((res) => res.json())
    .catch((err) => {
      return { error: true, errMsg: err };
    });

  return data;
}

function hideError() {
  errorMsg.classList.add("hidden");
}

form.addEventListener("submit", updateProduct);
