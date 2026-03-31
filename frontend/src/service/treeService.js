// src/service/treeService.js

const BASE = "http://localhost:8080";

export async function loadAllTrees() {
    try {
        const res = await fetch(`${BASE}/trees`);
        if (!res.ok) throw new Error(`Erro HTTP: ${res.status}`);
        return await res.json();
    } catch (err) {
        console.error("Erro ao carregar árvores:", err);
        return [];
    }
}

export async function loadTreesByBbox(lat1, lat2, lng1, lng2) {
    try {
        const url = `${BASE}/trees/bbox?lat1=${lat1}&lat2=${lat2}&lng1=${lng1}&lng2=${lng2}`;
        const res = await fetch(url);
        if (!res.ok) throw new Error(`Erro HTTP: ${res.status}`);
        return await res.json();
    } catch (err) {
        console.error("Erro ao carregar árvores por área:", err);
        return [];
    }
}

export async function createTree(tree) {
    const res = await fetch(`${BASE}/trees`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({
            latitude:  tree.lat,
            longitude: tree.lng,
            species:   tree.species,
            height:    parseFloat(tree.height) || 0,
        }),
    });
    if (!res.ok) throw new Error(`Erro HTTP: ${res.status}`);
    return await res.json();
}