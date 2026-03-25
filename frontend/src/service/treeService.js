// src/service/treeService.js

export async function loadAllTrees() {
    try {
        const res = await fetch("http://localhost:8080/trees");
        const apiTrees = await res.json();

        const localTrees =
            JSON.parse(localStorage.getItem("trees")) || [];

        return [...apiTrees, ...localTrees];
    } catch (err) {
        console.error("Erro ao carregar árvores:", err);
        return [];
    }
}