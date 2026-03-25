// Inicializa mapa
const map = L.map('map').setView([-23.5505, -46.6333], 13);

// Tile
L.tileLayer('https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', {
    attribution: '&copy; OpenStreetMap'
}).addTo(map);

let latSelecionada = null;
let lngSelecionada = null;

// Clique no mapa
map.on('click', function (e) {
    latSelecionada = e.latlng.lat;
    lngSelecionada = e.latlng.lng;

    document.getElementById("form-container").classList.remove("hidden");
});

// Fecha modal
function fecharForm() {
    document.getElementById("form-container").classList.add("hidden");
}

// Salvar árvore
function salvarArvore() {
    const especie = document.getElementById("especie").value;

    const novaArvore = {
        especie: especie,
        latitude: latSelecionada,
        longitude: lngSelecionada
    };

    fetch("http://localhost:8080/arvores", {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify(novaArvore)
    })
        .then(res => res.json())
        .then(data => {
            adicionarMarcador(data);
            fecharForm();
        })
        .catch(err => {
            console.error("Erro:", err);
        });
}

// Adiciona marcador
function adicionarMarcador(arvore) {
    L.marker([arvore.latitude, arvore.longitude])
        .addTo(map)
        .bindPopup(`🌳 ${arvore.especie}`);
}

// Carrega árvores do backend
function carregarArvores() {
    fetch("http://localhost:8080/arvores")
        .then(res => res.json())
        .then(data => {
            data.forEach(arvore => adicionarMarcador(arvore));
        })
        .catch(err => console.error(err));
}

// Inicializa
carregarArvores();