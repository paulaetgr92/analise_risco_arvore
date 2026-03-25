import { useState } from "react";
import { createTree } from "../service/api";

export default function TreeForm({ onCreated }) {
    const [form, setForm] = useState({
        latitude: "",
        longitude: "",
        species: "",
        height: "",
    });

    const [error, setError] = useState("");

    const handleChange = (e) => {
        setForm({
            ...form,
            [e.target.name]: e.target.value,
        });
    };

    const handleSubmit = async (e) => {
        e.preventDefault();
        setError("");

        const latitude = parseFloat(form.latitude);
        const longitude = parseFloat(form.longitude);
        const height = parseFloat(form.height);

        if (
            isNaN(latitude) ||
            isNaN(longitude) ||
            isNaN(height)
        ) {
            setError("Latitude, longitude e altura devem ser números válidos.");
            return;
        }

        const newTree = {
            latitude,
            longitude,
            species: form.species,
            height,
        };

        try {
            const createdTree = await createTree(newTree);

            onCreated(createdTree);

            setForm({
                latitude: "",
                longitude: "",
                species: "",
                height: "",
            });

        } catch (err) {
            setError(err.message || "Erro ao cadastrar árvore.");
        }
    };

    return (
        <form onSubmit={handleSubmit} className="card">
            <h2>Cadastrar Árvore</h2>

            {error && <p style={{ color: "red" }}>{error}</p>}

            <input
                name="latitude"
                type="number"
                step="any"
                placeholder="Latitude"
                value={form.latitude}
                onChange={handleChange}
            />

            <input
                name="longitude"
                type="number"
                step="any"
                placeholder="Longitude"
                value={form.longitude}
                onChange={handleChange}
            />

            <input
                name="species"
                placeholder="Espécie"
                value={form.species}
                onChange={handleChange}
            />

            <input
                name="height"
                type="number"
                step="any"
                placeholder="Altura"
                value={form.height}
                onChange={handleChange}
            />

            <button type="submit">Cadastrar</button>
        </form>
    );
}