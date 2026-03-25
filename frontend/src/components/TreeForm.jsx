import { useState } from "react";

export default function TreeForm({ onCreated }) {
    const [latitude, setLatitude] = useState("");
    const [longitude, setLongitude] = useState("");
    const [especie, setEspecie] = useState("");
    const [altura, setAltura] = useState("");

    const handleSubmit = (e) => {
        e.preventDefault();

        if (!latitude || !longitude) {
            alert("Preencha latitude e longitude!");
            return;
        }

        const newTree = {
            lat: parseFloat(latitude),
            lng: parseFloat(longitude),
            species: especie || "Desconhecida",
            height: altura || "N/A",
        };

        console.log("🌳 Nova árvore:", newTree);

        // 🔥 adiciona na lista
        onCreated(newTree);

        // limpa campos
        setLatitude("");
        setLongitude("");
        setEspecie("");
        setAltura("");
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
                placeholder="Altura (m)"
                value={altura}
                onChange={(e) => setAltura(e.target.value)}
            />

            <button type="submit">Cadastrar</button>
        </form>
    );
}