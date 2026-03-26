import L from "leaflet";
import "leaflet/dist/leaflet.css";

let map;
let markersLayer;

export function initMap() {

    // evita recriar mapa (React re-render)
    if (map) return;

    // 🌎 Seattle
    map = L.map("map").setView([47.6062, -122.3321], 13);

    L.tileLayer("https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png", {
        attribution: "&copy; OpenStreetMap"
    }).addTo(map);

    markersLayer = L.layerGroup().addTo(map);

    let selectedLat = null;
    let selectedLng = null;

    // 📍 clique no mapa
    map.on("click", async function (e) {

        selectedLat = e.latlng.lat;
        selectedLng = e.latlng.lng;

        console.log("📍 Clique:", selectedLat, selectedLng);

        markersLayer.clearLayers();

        L.marker([selectedLat, selectedLng])
            .addTo(markersLayer)
            .bindPopup("📍 Local selecionado")
            .openPopup();

        await loadTreesByLocation(selectedLat, selectedLng);
    });

    // 🌳 botão cadastrar
    const btn = document.getElementById("btn-map");

    if (btn) {
        btn.onclick = () => {

            const species = document.getElementById("species").value;
            const height = document.getElementById("height").value;

            if (!species || !height || !selectedLat) {
                alert("Preencha tudo e clique no mapa");
                return;
            }

            const tree = { species, height, lat: selectedLat, lng: selectedLng };

            const trees = JSON.parse(localStorage.getItem("trees")) || [];
            trees.push(tree);
            localStorage.setItem("trees", JSON.stringify(trees));

            renderTrees([tree]);
        };
    }
}

// 🔍 buscar árvores
async function loadTreesByLocation(lat, lng) {
    try {
        const res = await fetch(`http://localhost:8080/trees/near?lat=${lat}&lng=${lng}`);
        const apiTrees = await res.json();

        const localTrees = JSON.parse(localStorage.getItem("trees")) || [];

        const filteredLocal = localTrees.filter(t =>
            distanceInKm(t.lat, t.lng, lat, lng) < 1
        );

        const allTrees = [...apiTrees, ...filteredLocal];

        renderTrees(allTrees);

    } catch (err) {
        console.error(err);
    }
}

// 🌳 renderizar markers
function renderTrees(trees) {

    markersLayer.clearLayers();

    trees.forEach(t => {
        L.marker([t.lat, t.lng])
            .addTo(markersLayer)
            .bindPopup(`🌳 ${t.species || "Desconhecido"}<br>Altura: ${t.height || "N/A"}`);
    });
}

// 📏 distância real
function distanceInKm(lat1, lng1, lat2, lng2) {
    const R = 6371;

    const dLat = (lat2 - lat1) * Math.PI / 180;
    const dLng = (lng2 - lng1) * Math.PI / 180;

    const a =
        Math.sin(dLat / 2) ** 2 +
        Math.cos(lat1 * Math.PI / 180) *
        Math.cos(lat2 * Math.PI / 180) *
        Math.sin(dLng / 2) ** 2;

    return R * (2 * Math.atan2(Math.sqrt(a), Math.sqrt(1 - a)));
}