const URL = "http://localhost:5000/api/v1/products";

const form = document.body.querySelector("#createForm");

/**
 * @param {Event} event
 * */
function createProduct(event) {
  event.preventDefault();
  const formData = new FormData(form);
  const data = {};

  formData.forEach((value, key) => {
    data[key] = value;
  });

  data.price = parseFloat(data.price);

  fetch(URL, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(data),
  })
    .then((res) => {
      console.log(res);
    })
    .catch((err) => {
      console.error("Error: ", err);
    });
}

form.addEventListener("submit", createProduct);
