export default function NavBar({ setPage, currentPage }) {
  return (
    <nav className="navbar">
      <h2 className="logo">🌳 TreeRisk</h2>

      <div className="links">
        <button
          onClick={() => setPage("home")}
          className={currentPage === "home" ? "nav-active" : ""}
        >
          Home
        </button>
        <button
          onClick={() => setPage("arvores")}
          className={currentPage === "arvores" ? "nav-active" : ""}
        >
          Árvores
        </button>
        <button
          onClick={() => setPage("mapa")}
          className={currentPage === "mapa" ? "nav-active" : ""}
        >
          Mapa
        </button>
      </div>
    </nav>
  );
}
