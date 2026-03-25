export default function TreeList({ trees }) {
    if (!Array.isArray(trees)) {
        return <p>Erro ao carregar árvores</p>;
    }

    return (
        <div className="card">
            <h2>Árvores cadastradas</h2>

            {trees.length === 0 ? (
                <p>Nenhuma árvore cadastrada</p>
            ) : (
                trees.map((tree, index) => (
                    <div key={tree.id ?? index} className="tree-item">
                        <strong>{tree.species || "Sem espécie"}</strong>

                        <p>Altura: {tree.height} m</p>
                        <p>Latitude: {tree.latitude}</p>
                        <p>Longitude: {tree.longitude}</p>
                    </div>
                ))
            )}
        </div>
    );
}