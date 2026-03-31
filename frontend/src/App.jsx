import { useState } from "react";
import NavBar from "./components/NavBar.jsx";
import TreeForm from "./components/TreeForm.jsx";
import TreeList from "./components/TreeList.jsx";

function App() {
  const [page, setPage] = useState("home");

  return (
      <>
        <NavBar setPage={setPage} currentPage={page} />

        {page === "home" && (
            <div className="container">
              <h1>🌳 Tree Risk Monitor</h1>
              <TreeForm onCreated={() => {}} />
            </div>
        )}

        {page === "arvores" && <TreeList setPage={setPage} />}

        {page === "mapa" && (
            <div className="container">
              <button className="btn-back" onClick={() => setPage("home")}>
                ← Voltar
              </button>
              <h2 style={{ marginTop: "15px" }}>🗺️ Mapa</h2>
              <p style={{ color: "#ccc", marginTop: "10px" }}>
                Use a aba <strong>Árvores</strong> para ver o mapa interativo com
                filtro por localização.
              </p>
            </div>
        )}
      </>
  );
}

export default App;