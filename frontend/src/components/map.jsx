import L from "leaflet";

export function initMap() {

    const map = L.map("map").setView([-23.5505, -46.6333], 13);

    L.tileLayer("https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png", {
        attribution: "&copy; OpenStreetMap"
    }).addTo(map);

    let selectedLat = null;
    let selectedLng = null;

    loadMarkers(map);

    // 📍 clique no mapa
    map.on("click", async function (e) {

        selectedLat = e.latlng.lat;
        selectedLng = e.latlng.lng;

        console.log("📍 Clique:", selectedLat, selectedLng);

        L.marker([selectedLat, selectedLng]).addTo(map);

        if (window.loadTreesByLocation) {
            await window.loadTreesByLocation(selectedLat, selectedLng);
        }
    });

    // 🌳 cadastrar árvore
    const btn = document.getElementById("btn-map");

    if (btn) {
        btn.addEventListener("click", () => {

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

            L.marker([tree.lat, tree.lng]).addTo(map);

            document.getElementById("status").innerText = "✅ Árvore cadastrada!";
        });
    }

    // 🔍 filtro global
    window.loadTreesByLocation = async function (lat, lng) {

        const res = await fetch(`http://localhost:8080/trees/near?lat=${lat}&lng=${lng}`);
        const apiTrees = await res.json();

        const localTrees = JSON.parse(localStorage.getItem("trees")) || [];

        const filteredLocal = localTrees.filter(t => {
            const dist = Math.sqrt((t.lat - lat)**2 + (t.lng - lng)**2);
            return dist < 0.01;
        });

        const all = [...apiTrees, ...filteredLocal];

        if (window.renderTrees) {
            window.renderTrees(all);
        }
    };
}

// 📍 markers existentes
function loadMarkers(map) {
    const trees = JSON.parse(localStorage.getItem("trees")) || [];

    trees.forEach(t => {
        L.marker([t.lat, t.lng]).addTo(map);
    });
}