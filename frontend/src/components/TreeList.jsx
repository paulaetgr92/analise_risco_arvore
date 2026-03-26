import { useEffect } from "react";
import { initMap } from "./map.jsx";

function TreeList({ trees, setPage }) {

  useEffect(() => {
    // espera renderizar o DOM antes de iniciar mapa
    setTimeout(() => {
      initMap();
    }, 100);
  }, []);

  return (
      <div className="container">

        <button className="btn-back" onClick={() => setPage("home")}>
          ← Voltar
        </button>

        <h2>🌳 Árvores</h2>
        <p>Clique no mapa para filtrar por localização</p>

        {/* 📍 MAPA */}
        <div
            id="map"
            style={{
              height: "400px",
              width: "600px",
              margin: "0 auto",
              borderRadius: "10px"
            }}
        ></div>

        {/* 🌳 LISTA */}
        <ul style={{ marginTop: "20px" }}>
          {trees.map((t, i) => (
              <li key={i}>
                🌳 {t.species || "Desconhecido"} - {t.height || "N/A"}
              </li>
          ))}
        </ul>

      </div>
  );
}

export default TreeList;