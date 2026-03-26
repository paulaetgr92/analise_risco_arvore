import { useEffect, useRef, useState } from "react";
import L from "leaflet";
import "leaflet/dist/leaflet.css";

export default function TreeList({ trees, setPage }) {
  const mapRef = useRef(null);
  const [filteredTrees, setFilteredTrees] = useState([]);

  useEffect(() => {
    setFilteredTrees(trees);
  }, [trees]);

  useEffect(() => {
    if (mapRef.current) return;

    const map = L.map("map-filter").setView([-23.5505, -46.6333], 13);
    mapRef.current = map;

    L.tileLayer("https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png", {
      attribution: "© OpenStreetMap",
    }).addTo(map);

    map.on("click", (e) => {
      const { lat, lng } = e.latlng;

      const filtered = trees.filter((tree) => {
        if (!tree.lat || !tree.lng) return false;

        const dist = getDistance(lat, lng, tree.lat, tree.lng);
        return dist < 1; // 1km
      });

      setFilteredTrees(filtered);
    });
  }, [trees]);

  useEffect(() => {
    if (!mapRef.current) return;

    const map = mapRef.current;

    map.eachLayer((layer) => {
      if (layer instanceof L.Marker) {
        map.removeLayer(layer);
      }
    });

    filteredTrees.forEach((tree) => {
      if (tree.lat && tree.lng) {
        L.marker([tree.lat, tree.lng])
          .addTo(map)
          .bindPopup(`🌳 ${tree.species}`);
      }
    });
  }, [filteredTrees]);

  return (
    <div className="container">
      {/* ✅ botão voltar */}
      <button className="btn-back" onClick={() => setPage("home")}>
        ← Voltar
      </button>

      <h2 style={{ marginTop: "15px" }}>🌳 Árvores</h2>

      <p>Clique no mapa para filtrar por localização</p>

      <div
        id="map-filter"
        style={{
          height: "400px",
          width: "100%",
          marginBottom: "20px",
        }}
      ></div>

      <ul>
        {filteredTrees.length === 0 && <p>Nenhuma árvore encontrada</p>}

        {filteredTrees.map((tree, i) => (
          <li key={i}>
            🌳 {tree.species} - {tree.height}m
          </li>
        ))}
      </ul>
    </div>
  );
}

// 📏 Distância real (Haversine)
function getDistance(lat1, lng1, lat2, lng2) {
  const R = 6371;
  const dLat = (lat2 - lat1) * (Math.PI / 180);
  const dLng = (lng2 - lng1) * (Math.PI / 180);

  const a =
    Math.sin(dLat / 2) * Math.sin(dLat / 2) +
    Math.cos(lat1 * (Math.PI / 180)) *
      Math.cos(lat2 * (Math.PI / 180)) *
      Math.sin(dLng / 2) *
      Math.sin(dLng / 2);

  return R * 2 * Math.atan2(Math.sqrt(a), Math.sqrt(1 - a));
}
