const API_URL = 'http://localhost:8080'

export async function getItems() {
    const response = await fetch(`${API_URL}/items`)
    const result = await response.json()

    if (!response.ok) {
        throw new Error(result.message || 'Erro ao buscar items')
    }

    return result.data
}