const API_URL = 'http://localhost:8080'

export async function getItems() {
    const response = await fetch(`${API_URL}/items`)
    const result = await response.json()

    if (!response.ok) {
        throw new Error(result.message)
    }

    return result.data
}

export async function createItem(item) {
    const response = await fetch(`${API_URL}/items`, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(item),
    })

    const result = await response.json()

    if (!response.ok) {
        throw result
    }

    return result.data
}

export async function deleteItem(id) {
    const response = await fetch(`${API_URL}/items/${id}`, {
        method: 'DELETE',
    })

    const result = await response.json()

    if (!response.ok) {
        throw new Error(result.message)
    }

    return result.data
}

export async function updateItem(id, item) {
    const response = await fetch(`${API_URL}/items/${id}`, {
        method: 'PUT',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(item),
    })

    const result = await response.json()

    if (!response.ok) {
        throw result
    }

    return result.data
}