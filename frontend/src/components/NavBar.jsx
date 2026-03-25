export default function NavBar({ setPage }) {
    return (
        <nav className="navbar">
            <h2 className="logo">🌳 TreeRisk</h2>

            <div className="links">
                <button onClick={() => setPage("home")}>Home</button>
                <button onClick={() => setPage("arvores")}>Árvores</button>
                <button onClick={() => setPage("mapa")}>Mapa</button>
            </div>
        </nav>
    );
}