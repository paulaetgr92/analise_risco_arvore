const API_URL = "http://localhost:8080";

export const getTrees = async () => {
    const res = await fetch(`${API_URL}/trees`);
    const data = await res.json();
    return data;
};

export const createTree = async (tree) => {
    const res = await fetch(`${API_URL}/trees`, {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify(tree),
    });

    const data = await res.json();

    if (!res.ok) {
        throw new Error(data?.error || "Erro ao cadastrar árvore");
    }

    return data;
};