import L from "leaflet";
import "leaflet/dist/leaflet.css";

let map;
let markersLayer;

export function initMap(onMapClick) {
    const container = document.getElementById("map-filter");
    if (!container) return;

    if (map) {
        map.remove();
        map = null;
    }

    map = L.map("map-filter").setView([47.6062, -122.3321], 13);

    L.tileLayer("https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png", {
        attribution: "© OpenStreetMap"
    }).addTo(map);

    markersLayer = L.layerGroup().addTo(map);

    map.on("click", function (e) {
        const lat = e.latlng.lat;
        const lng = e.latlng.lng;

        markersLayer.clearLayers();

        L.marker([lat, lng])
            .addTo(markersLayer)
            .bindPopup("📍 Local selecionado")
            .openPopup();

        // Chama o callback do React com as coordenadas
        if (typeof onMapClick === "function") {
            onMapClick(lat, lng);
        }
    });
}

window.focusTreeOnMap = function (tree) {
    if (!map || !tree) return;
    const { lat, lng, species, scientific, height } = tree;
    if (!lat || !lng) return;

    map.setView([lat, lng], 17);
    markersLayer.clearLayers();

    L.marker([lat, lng])
        .addTo(markersLayer)
        .bindPopup(`
            🌳 <b>${species}</b><br/>
            ${scientific ? `🔬 ${scientific}<br/>` : ""}
            Altura: ${height}
        `)
        .openPopup();
};