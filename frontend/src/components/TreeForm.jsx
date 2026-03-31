import { useState } from "react";
import { createTree } from "../service/treeService";

export default function TreeForm({ onCreated }) {
    const [latitude, setLatitude]   = useState("");
    const [longitude, setLongitude] = useState("");
    const [especie, setEspecie]     = useState("");
    const [altura, setAltura]       = useState("");
    const [saving, setSaving]       = useState(false);
    const [error, setError]         = useState(null);

    const handleSubmit = async (e) => {
        e.preventDefault();
        setError(null);

        if (!latitude || !longitude) {
            alert("Preencha latitude e longitude!");
            return;
        }

        if (!altura || parseFloat(altura) <= 0) {
            alert("Preencha uma altura válida maior que zero!");
            return;
        }

        const newTree = {
            lat:     parseFloat(latitude),
            lng:     parseFloat(longitude),
            species: especie || "Desconhecida",
            height:  parseFloat(altura), // <- number, não string
        };

        console.log("🌳 Nova árvore:", newTree);

        setSaving(true);
        try {
            await createTree(newTree);
            onCreated();
            setLatitude("");
            setLongitude("");
            setEspecie("");
            setAltura("");
        } catch (err) {
            console.error("Erro ao cadastrar:", err);
            setError("Erro ao cadastrar árvore. Verifique os dados.");
        } finally {
            setSaving(false);
        }
    };

    return (
        <form onSubmit={handleSubmit} className="form">
            <h2>Cadastrar Árvore</h2>

            <input
                type="number"
                step="any"
                placeholder="Latitude"
                value={latitude}
                onChange={(e) => setLatitude(e.target.value)}
            />
            <input
                type="number"
                step="any"
                placeholder="Longitude"
                value={longitude}
                onChange={(e) => setLongitude(e.target.value)}
            />
            <input
                type="text"
                placeholder="Espécie"
                value={especie}
                onChange={(e) => setEspecie(e.target.value)}
            />
            <input
                type="number"
                step="any"
                min="0.1"
                placeholder="Altura (m)"
                value={altura}
                onChange={(e) => setAltura(e.target.value)}
            />

            {error && <p style={{ color: "red", fontSize: "0.85rem" }}>{error}</p>}

            <button type="submit" disabled={saving}>
                {saving ? "Salvando..." : "Cadastrar"}
            </button>
        </form>
    );
}