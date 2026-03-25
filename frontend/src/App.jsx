import { useEffect, useState } from "react";
import NavBar from "./components/NavBar.jsx";
import TreeForm from "./components/TreeForm.jsx";
import TreeList from "./components/TreeList.jsx";
import { getSeattleTrees } from "./service/seattleTrees";

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
            <NavBar setPage={setPage} />

            {page === "home" && (
                <div className="container">
                    <h1>🌳 Tree Risk Monitor</h1>
                    <TreeForm onCreated={fetchTrees} />
                </div>
            )}

            {page === "arvores" && <TreeList trees={trees} />}
        </>
    );
}

export default App;