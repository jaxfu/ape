async function getComponents() {
  const req = await fetch("http://localhost:5000/api/components");
  const res = await req.json();

  return res;
}

const Card = (comp) => {
  const card = document.createElement("div");
  card.className = "card";
  card.innerHTML = `
  <h2>${comp.name}</h2>
  <hr>
  <h3>id: ${comp.component_id}</h3>
`;

  return card;
};

async function main() {
  const root = document.querySelector("#root");

  const components = await getComponents();
  const container = document.createElement("section");
  container.id = "components_container";

  components.forEach((comp) => {
    container.appendChild(Card(comp));
  });

  root.appendChild(container);
}

main();
