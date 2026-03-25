export async function getSeattleTrees() {
    const url =
        "https://services.arcgis.com/ZOyb2t4B0UYuYNYH/arcgis/rest/services/SDOT_Trees_(Active)/FeatureServer/0/query" +
        "?where=1=1" +
        "&outFields=*" +
        "&outSR=4326" +
        "&f=json" +
        "&resultRecordCount=500";

    try {
        const res = await fetch(url);
        const data = await res.json();

        console.log("🌳 dados brutos:", data); // 👈 DEBUG

        return data.features.map((item) => ({
            lat: item.geometry.y,
            lng: item.geometry.x,
            species: item.attributes.SCIENTIFIC_NAME || "Desconhecida",
            common: item.attributes.COMMON_NAME || "",
            height: item.attributes.HEIGHT_M || "N/A",
        }));
    } catch (err) {
        console.error("Erro Seattle API:", err);
        return [];
    }
}