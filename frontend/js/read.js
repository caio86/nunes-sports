const URL = "http://localhost:5000/api/v1/products";

const tabela = document.body.querySelector("#readTable > tbody");

async function getProducts() {
  const data = await fetch(URL).then((res) => res.json());

  data.map(
    /**
     * @param {Object} produto
     * @param {Number} produto.ID
     * @param {string} produto.Name
     * @param {string} produto.Description
     * @param {Number} produto.Price
     * */
    (produto) => {
      let item = document.createElement("tr");
      let itemID = document.createElement("td");
      let itemName = document.createElement("td");
      let itemDesc = document.createElement("td");
      let itemPrice = document.createElement("td");

      itemID.innerHTML = produto.ID;
      itemName.innerHTML = produto.Name;
      itemDesc.innerHTML = produto.Description;
      itemPrice.innerHTML = produto.Price.toFixed(2);

      item.appendChild(itemID);
      item.appendChild(itemName);
      item.appendChild(itemDesc);
      item.appendChild(itemPrice);

      tabela.appendChild(item);
    },
  );
}

getProducts();
