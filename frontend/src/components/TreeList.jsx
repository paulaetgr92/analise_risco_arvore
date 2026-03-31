import { useEffect, useState, useRef } from "react";
import { initMap } from "./map";
import { getSeattleTrees } from "../service/seattletrees";
import { loadAllTrees, loadTreesByBbox } from "../service/treeService";

function buildEnvelope(lat, lng, radiusMeters) {
    const deltaLat = radiusMeters / 111320;
    const deltaLng = radiusMeters / (111320 * Math.cos(lat * Math.PI / 180));
    return {
        xmin: lng - deltaLng,
        ymin: lat - deltaLat,
        xmax: lng + deltaLng,
        ymax: lat + deltaLat,
    };
}

async function fetchSeattleTreesNear(lat, lng, radiusMeters = 1000) {
    const env = buildEnvelope(lat, lng, radiusMeters);
    const geometryParam = encodeURIComponent(JSON.stringify({
        ...env,
        spatialReference: { wkid: 4326 },
    }));

    const url =
        "https://services.arcgis.com/ZOyb2t4B0UYuYNYH/arcgis/rest/services/SDOT_Trees_CDL/FeatureServer/0/query" +
        `?geometry=${geometryParam}` +
        "&geometryType=esriGeometryEnvelope" +
        "&inSR=4326" +
        "&spatialRel=esriSpatialRelIntersects" +
        "&where=1%3D1" +
        "&outFields=OBJECTID,COMMON_NAME,SCIENTIFIC_NAME,CONDITION" +
        "&outSR=4326" +
        "&returnGeometry=true" +
        "&f=json" +
        "&resultRecordCount=100";

    const res = await fetch(url);
    const data = await res.json();

    if (data.error) throw new Error(data.error.message);
    if (!data.features || data.features.length === 0) return [];

    return data.features.map((item) => ({
        id:         item.attributes.OBJECTID,
        species:    item.attributes.COMMON_NAME || "Espécie desconhecida",
        scientific: item.attributes.SCIENTIFIC_NAME || "",
        height:     "N/A",
        condition:  item.attributes.CONDITION || "",
        lat:        item.geometry?.y ?? lat,
        lng:        item.geometry?.x ?? lng,
        source:     "seattle-api",
    }));
}

function mapDbTree(t) {
    return {
        id:         `db-${t.id}`,
        species:    t.species || "Desconhecida",
        scientific: "",
        height:     t.height ? `${t.height}m` : "N/A",
        condition:  t.health || "",
        lat:        t.latitude,
        lng:        t.longitude,
        source:     "local-db",
    };
}

export default function TreeList() {
    const [trees, setTrees]                     = useState([]);
    const [loading, setLoading]                 = useState(false);
    const [loadingAll, setLoadingAll]           = useState(false);
    const [error, setError]                     = useState(null);
    const [clickedLocation, setClickedLocation] = useState(null);
    const handleMapClickRef                     = useRef(null);

    handleMapClickRef.current = async (lat, lng) => {
        setClickedLocation({ lat, lng });
        setLoading(true);
        setError(null);
        setTrees([]);

        const env = buildEnvelope(lat, lng, 1000);
        console.log("📦 Envelope:", env);
        console.log("🏗️ BD bbox params — lat1:", env.ymin, "lat2:", env.ymax, "lng1:", env.xmin, "lng2:", env.xmax);

        try {
            const [apiTrees, dbTrees] = await Promise.all([
                fetchSeattleTreesNear(lat, lng, 1000),
                loadTreesByBbox(env.ymin, env.ymax, env.xmin, env.xmax),
            ]);

            console.log("🌳 Seattle API:", apiTrees.length, "árvores");
            console.log("💾 BD:", dbTrees?.length, "árvores", dbTrees);

            const fromDB = (Array.isArray(dbTrees) ? dbTrees : []).map(mapDbTree);

            setTrees([...fromDB, ...apiTrees]);
        } catch (err) {
            console.error("Erro ao buscar as árvores:", err);
            setError("Não foi possível carregar as árvores desta área.");
        } finally {
            setLoading(false);
        }
    };

    useEffect(() => {
        initMap((lat, lng) => handleMapClickRef.current(lat, lng));
    }, []);

    const handleListAll = async () => {
        setLoadingAll(true);
        setError(null);
        setClickedLocation(null);

        try {
            const [dbTrees, apiData] = await Promise.all([
                loadAllTrees(),
                getSeattleTrees(),
            ]);

            console.log("💾 BD listar todas:", dbTrees?.length, "árvores");
            console.log("🌳 Seattle API listar todas:", apiData?.length, "árvores");

            const fromDB = (Array.isArray(dbTrees) ? dbTrees : []).map(mapDbTree);

            const fromAPI = (Array.isArray(apiData) ? apiData : []).map((t) => ({
                id:         `api-${t.lat}-${t.lng}`,
                species:    t.common || t.species || "Espécie desconhecida",
                scientific: t.species || "",
                height:     "N/A",
                condition:  t.condition || "",
                lat:        t.lat,
                lng:        t.lng,
                source:     "seattle-api",
            }));

            setTrees([...fromDB, ...fromAPI]);
        } catch (err) {
            console.error("Erro ao listar todas:", err);
            setError("Não foi possível carregar todas as árvores.");
        } finally {
            setLoadingAll(false);
        }
    };

    return (
        <div>
            <div id="map-filter" style={{ height: "400px", width: "100%" }} />

            <button id="btn-load" onClick={handleListAll} disabled={loadingAll}>
                {loadingAll ? "Carregando..." : "Listar todas"}
            </button>

            {loading    && <p>🔍 Buscando árvores na área...</p>}
            {loadingAll && <p>🌳 Carregando inventário completo...</p>}
            {error      && <p style={{ color: "red" }}>{error}</p>}

            {!loading && !loadingAll && !error && clickedLocation && trees.length === 0 && (
                <p>Nenhuma árvore encontrada nesta área.</p>
            )}

            {!loading && !loadingAll && trees.length > 0 && (
                <>
                    <p style={{ fontSize: "0.85rem", color: "#666" }}>
                        {trees.length} árvore(s)
                        {clickedLocation ? " em um raio de 1km" : " no inventário"}
                    </p>
                    <ul>
                        {trees.map((tree) => (
                            <li
                                key={tree.id}
                                onClick={() => window.focusTreeOnMap?.(tree)}
                                style={{ cursor: "pointer" }}
                            >
                                <strong>{tree.species}</strong>
                                {tree.scientific && <em> — {tree.scientific}</em>}
                                <span> | Altura: {tree.height}</span>
                                {tree.condition && (
                                    <span style={{ marginLeft: 8, fontSize: "0.8rem" }}>
                                        | Condição: {tree.condition}
                                    </span>
                                )}
                                {tree.source === "seattle-api" && (
                                    <span style={{ marginLeft: 8, fontSize: "0.75rem", color: "#888" }}>
                                        📡 SDOT
                                    </span>
                                )}
                                {tree.source === "local-db" && (
                                    <span style={{ marginLeft: 8, fontSize: "0.75rem", color: "#4a9e5c" }}>
                                        💾 BD
                                    </span>
                                )}
                            </li>
                        ))}
                    </ul>
                </>
            )}
        </div>
    );
}