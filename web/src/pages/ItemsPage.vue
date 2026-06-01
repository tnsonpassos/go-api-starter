<script setup>
import { onMounted, ref } from 'vue'
import { getItems } from '../services/api'

const items = ref([])
const loading = ref(true)
const error = ref('')

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
            </tr>
          </thead>

          <tbody>
            <tr v-for="item in items" :key="item.id">
              <td>{{ item.id }}</td>
              <td>{{ item.name }}</td>
              <td>{{ item.description }}</td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  </div>
</template>