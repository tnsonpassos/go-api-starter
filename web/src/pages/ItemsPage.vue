<script setup>
import { onMounted, ref } from 'vue'
import { getItems, createItem, deleteItem, updateItem } from '../services/api'

const items = ref([])
const loading = ref(true)
const error = ref('')

const form = ref({
  name: '',
  description: '',
})

const editingItemId = ref(null)

async function submitForm() {
  try {
    if (editingItemId.value) {
      await updateItem(editingItemId.value, form.value)
    } else {
      await createItem(form.value)
    }

    form.value = {
      name: '',
      description: '',
    }

    editingItemId.value = null

    await loadItems()
  } catch (err) {
    console.error(err)
  }
}

function editItem(item) {
  editingItemId.value = item.id

  form.value = {
    name: item.name,
    description: item.description,
  }
}

function cancelEdit() {
  editingItemId.value = null

  form.value = {
    name: '',
    description: '',
  }
}

async function removeItem(id) {
  const confirmDelete = confirm('Deseja realmente excluir este item?')

  if (!confirmDelete) {
    return
  }

  try {
    await deleteItem(id)
    await loadItems()
  } catch (err) {
    error.value = err.message
  }
}

async function loadItems() {
  try {
    loading.value = true
    items.value = await getItems()
  } catch (err) {
    error.value = err.message
  } finally {
    loading.value = false
  }
}

onMounted(() => {
  loadItems()
})
</script>

<template>
  <div>

    <div class="card mb-4">
    <div class="card-body">
        <h3 class="h5 mb-3">
            {{ editingItemId ? 'Editar Item' : 'Novo Item' }}
        </h3>

        <form @submit.prevent="submitForm">
        <div class="mb-3">
            <label class="form-label">Nome</label>

            <input
            v-model="form.name"
            class="form-control"
            required
            />
        </div>

        <div class="mb-3">
            <label class="form-label">Descrição</label>

            <textarea
            v-model="form.description"
            class="form-control"
            rows="3"
            ></textarea>
        </div>

        <div class="d-flex gap-2">
        <button class="btn btn-primary">
            {{ editingItemId ? 'Atualizar' : 'Salvar' }}
        </button>

        <button
            v-if="editingItemId"
            type="button"
            class="btn btn-outline-secondary"
            @click="cancelEdit"
        >
            Cancelar
        </button>
        </div>

        </form>
    </div>
    </div>

    <div class="mb-4">
      <h2 class="fw-bold">Items</h2>
      <p class="text-muted">Lista de itens vindos da API Go.</p>
    </div>

    <div v-if="loading" class="alert alert-info">
      Carregando items...
    </div>

    <div v-if="error" class="alert alert-danger">
      {{ error }}
    </div>

    <div v-if="!loading && !error" class="card">
      <div class="card-body">
        <table class="table table-striped mb-0">
          <thead>
            <tr>
              <th>ID</th>
              <th>Nome</th>
              <th>Descrição</th>
              <th>Ações</th>
            </tr>
          </thead>

          <tbody>
            <tr v-for="item in items" :key="item.id">
              <td>{{ item.id }}</td>
              <td>{{ item.name }}</td>
              <td>{{ item.description }}</td>
              <td>
                <button
                    class="btn btn-sm btn-outline-primary me-2"
                    @click="editItem(item)"
                >
                    Editar
                </button>   

                <button
                  class="btn btn-sm btn-danger"
                  @click="removeItem(item.id)"
                >
                  Excluir
                </button>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>

  </div>
</template>