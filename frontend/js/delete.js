/**
 * @param {String} id
 * @returns {string}
 * */
const getProductUrl = (id) => `http://localhost:5000/api/v1/products/${id}`;

/**
 * @param {string} id
 * */
async function getProduct(id) {
  const data = await fetch(getProductUrl(id))
    .then((res) => res.json())
    .catch((err) => {
      return { error: true, errMsg: err };
    });

  return data;
}

function hideError() {
  errElement.classList.add("hidden");
}

/**
 * @param {String} errMsg
 * */
function showError(errMsg) {
  errElement.classList.remove("hidden");
  errElement.innerHTML = `Error: ${errMsg}`;
  console.error("Error: ", errMsg);
}

/**
 * @param {Object} data
 * @param {String} data.Id
 * @param {string} data.Name
 * @param {string} data.Description
 * @param {Number} data.Price
 * @param {Boolean} data.error
 * @param {String} data.errMsg
 * */
function fillFormWithProduct(data) {
  if (data.error) {
    showError(data.errMsg);

    form["name"].value = "";
    form["description"].value = "";
    form["price"].value = "";

    return;
  }

  form["name"].value = data.Name;
  form["description"].value = data.Description;
  form["price"].value = data.Price;
}

async function onUpdateForm() {
  hideError();

  let id = form["id"].value;

  const data = await getProduct(id);

  fillFormWithProduct(data);
}

const form = document.forms["deleteForm"];
const errElement = document.body.querySelector(".error");

/**
 * @param {Event} event
 * */
async function deleteProduct(event) {
  event.preventDefault();

  const id = form["id"].value;

  await fetch(getProductUrl(id), {
    method: "DELETE",
  })
    .then(() => {
      form.reset();
    })
    .catch((err) => {
      showError(err);
    });
}

form.addEventListener("submit", deleteProduct);
