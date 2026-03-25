import { getTrees } from "../service/api.js";

export async function loadAllTrees() {
    const container = document.getElementById("tree-list");

    if (!container) return;

    container.innerHTML = "Carregando...";

    try {
        const trees = await getTrees();

        if (!trees || trees.length === 0) {
            container.innerHTML = "<p>Nenhuma árvore cadastrada 🌱</p>";
            return;
        }

        container.innerHTML = "";

        trees.forEach(tree => {
            const div = document.createElement("div");
            div.className = "tree-card";

            div.innerHTML = `
                <h3>🌳 ${tree.species}</h3>
                <p>Altura: ${tree.height}m</p>
                <p>📍 (${tree.latitude?.toFixed(4)}, ${tree.longitude?.toFixed(4)})</p>
            `;

            container.appendChild(div);
        });

    } catch (error) {
        console.error(error);
        container.innerHTML = "<p>Erro ao carregar árvores ❌</p>";
    }
}