import { useEffect } from "react";
import L from "leaflet";
import "leaflet/dist/leaflet.css";

function App() {

    useEffect(() => {

        // Criar mapa
        const map = L.map("map").setView([-23.55, -46.63], 15);

        // Camada base
        L.tileLayer("https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png", {
            attribution: "&copy; OpenStreetMap",
        }).addTo(map);

        // Camada para árvores OSM
        const osmLayer = L.layerGroup().addTo(map);

        async function carregarArvores() {
            osmLayer.clearLayers();

            const center = map.getCenter();

            const query = `
        [out:json];
        node["natural"="tree"](around:500,${center.lat},${center.lng});
        out;
      `;

            try {
                const response = await fetch(
                    "https://overpass-api.de/api/interpreter",
                    {
                        method: "POST",
                        body: query,
                    }
                );

                const data = await response.json();

                data.elements.forEach((tree) => {
                    if (tree.lat && tree.lon) {
                        L.circleMarker([tree.lat, tree.lon], {
                            radius: 5,
                            color: "green",
                        })
                            .addTo(osmLayer)
                            .bindPopup("🌳 Árvore (OSM)");
                    }
                });

            } catch (err) {
                console.error("Erro Overpass:", err);
            }
        }

        carregarArvores();
        map.on("moveend", carregarArvores);

        // 🔥 CORREÇÃO IMPORTANTE (React evita duplicar mapa)
        return () => {
            map.remove();
        };

    }, []);

    return (
        <div style={{ padding: "20px", textAlign: "center" }}>
            <h1>🌳 Tree Risk</h1>

            <div
                id="map"
                style={{
                    height: "500px",
                    width: "100%",
                    borderRadius: "12px",
                }}
            ></div>
        </div>
    );
}

export default App;