import { getTrees, createTree } from "./service/api.js";

// 🌎 SEATTLE
const map = L.map('map').setView([47.6062, -122.3321], 13);

// 🗺️ mapa base
L.tileLayer('https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', {
    attribution: '&copy; OpenStreetMap'
}).addTo(map);

// 📌 estado
let clickEnabled = false;
let markers = [];

// 🔘 botão cadastrar
window.enableMapClick = () => {
    clickEnabled = true;
    setStatus("Clique no mapa para cadastrar...");
};

// 🧠 status
function setStatus(msg) {
    document.getElementById("status").innerText = msg;
}

// 🧹 limpar inputs
function clearInputs() {
    document.getElementById("species").value = "";
    document.getElementById("height").value = "";
}

// 📍 adicionar marcador
function addMarker(tree) {
    const marker = L.marker([tree.latitude, tree.longitude]).addTo(map);

    marker.bindPopup(`
        <b>${tree.species}</b><br>
        Altura: ${tree.height}m
    `);

    markers.push(marker);
}

// 🧹 limpar markers
function clearMarkers() {
    markers.forEach(m => map.removeLayer(m));
    markers = [];
}

// 🌳 carregar árvores por localização
async function loadTreesByLocation() {
    try {
        const center = map.getCenter();

        const trees = await getTrees(center.lat, center.lng);

        clearMarkers();

        trees.forEach(tree => addMarker(tree));

    } catch (err) {
        console.error(err);
        setStatus("Erro ao carregar árvores");
    }
}

// 🔄 quando mover mapa
map.on("moveend", () => {
    loadTreesByLocation();
});

// 🖱️ clique para cadastrar
map.on('click', async function (e) {

    if (!clickEnabled) return;

    const species = document.getElementById("species").value;
    const height = document.getElementById("height").value;

    if (!species || !height) {
        alert("Preencha os campos!");
        return;
    }

    const tree = {
        species,
        height: Number(height),
        latitude: e.latlng.lat,
        longitude: e.latlng.lng
    };

    try {
        const saved = await createTree(tree);

        addMarker(saved);

        setStatus("Árvore cadastrada 🌳");

        clearInputs();
        clickEnabled = false;

    } catch (err) {
        alert(err.message);
    }
});

// 🔍 expandir mapa
window.toggleMap = () => {
    const wrapper = document.querySelector(".map-wrapper");
    wrapper.classList.toggle("expanded");

    setTimeout(() => map.invalidateSize(), 300);
};

// 🚀 inicial
loadTreesByLocation();