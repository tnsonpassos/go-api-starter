<script setup>
import { onMounted, ref } from 'vue'
import { getItems, createItem, deleteItem, updateItem } from '../services/api'
import ItemForm from '../components/ItemForm.vue'
import ItemTable from '../components/ItemTable.vue'

const items = ref([])
const loading = ref(true)
const error = ref('')
const success = ref('')

const form = ref({
  name: '',
  description: '',
})

const editingItemId = ref(null)

function clearMessages() {
  error.value = ''
  success.value = ''
}

async function submitForm(payload) {
  try {
    clearMessages()

    if (editingItemId.value) {
      await updateItem(editingItemId.value, payload)
      success.value = 'Item atualizado com sucesso'
    } else {
      await createItem(payload)
      success.value = 'Item criado com sucesso'
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
    clearMessages()
    const confirmDelete = confirm('Deseja realmente excluir este item?')

    if (!confirmDelete) {
        return
    }

    try {
        await deleteItem(id)
        success.value = 'Item excluído com sucesso'
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

    <ItemForm
    :item="form"
    :editing="!!editingItemId"
    @submit="submitForm"
    @cancel="cancelEdit"
    />

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

    <ItemTable
      :items="items"
      @edit="editItem"
      @delete="removeItem"
    />

    <div v-if="success" class="alert alert-success">
        {{ success }}
    </div>
    
  </div>
</template>