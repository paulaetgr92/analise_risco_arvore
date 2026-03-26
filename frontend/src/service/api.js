import { getSeattleTrees } from "./seattletrees"; // ✅ corrigido: lowercase

export async function getTrees() {
  return await getSeattleTrees();
}

// opcional (mantém se quiser usar depois)
export async function createTree(tree) {
  console.log("Simulação de criação:", tree);
  return tree;
}
