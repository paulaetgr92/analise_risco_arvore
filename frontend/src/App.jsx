import { useEffect, useState } from "react";
import NavBar from "./components/NavBar.jsx";
import TreeForm from "./components/TreeForm.jsx";
import TreeList from "./components/TreeList.jsx";
import { getSeattleTrees } from "./service/seattletrees"; // ✅ corrigido: lowercase igual ao nome real do arquivo

function App() {
  const [trees, setTrees] = useState([]);
  const [page, setPage] = useState("home");

  // 🔥 busca árvores da API
  const fetchTrees = async (newTree = null) => {
    // 👉 se criou nova árvore
    if (newTree) {
      setTrees((prev) => [newTree, ...prev]);
      return;
    }

    const data = await getSeattleTrees();
    setTrees(data);
  };

  useEffect(() => {
    fetchTrees();
  }, []);

  return (
    <>
      <NavBar setPage={setPage} currentPage={page} />

      {page === "home" && (
        <div className="container">
          <h1>🌳 Tree Risk Monitor</h1>
          <TreeForm onCreated={fetchTrees} />
        </div>
      )}

      {page === "arvores" && <TreeList trees={trees} setPage={setPage} />}

      {/* ✅ página mapa: evita tela em branco */}
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
