const BASE_URL =
    "https://services.arcgis.com/ZOyb2t4B0UYuYNYH/arcgis/rest/services/SDOT_Trees_CDL/FeatureServer/0/query";

export async function getSeattleTrees() {
    const url =
        BASE_URL +
        "?where=CURRENT_STATUS%3D%27INSVC%27" +
        "&outFields=OBJECTID,COMMON_NAME,SCIENTIFIC_NAME,CONDITION" +
        "&outSR=4326" +
        "&f=json" +
        "&resultRecordCount=500";

    try {
        const res = await fetch(url);
        if (!res.ok) throw new Error(`Erro HTTP: ${res.status}`);
        const data = await res.json();
        if (data.error) throw new Error(data.error.message);
        if (!data.features) return [];

        return data.features.map((item) => ({
            lat: item.geometry.y,
            lng: item.geometry.x,
            species: item.attributes.SCIENTIFIC_NAME || "Desconhecida",
            common: item.attributes.COMMON_NAME || "",
            height: "N/A",
            condition: item.attributes.CONDITION || "",
        }));
    } catch (err) {
        console.error("Erro Seattle API:", err);
        return [];
    }
}