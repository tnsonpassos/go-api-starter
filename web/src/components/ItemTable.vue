<script setup>
defineProps({
  items: {
    type: Array,
    default: () => [],
  },
})

defineEmits([
  'edit',
  'delete',
])
</script>

<template>
  <div class="card">
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
          <tr
            v-for="item in items"
            :key="item.id"
          >
            <td>{{ item.id }}</td>
            <td>{{ item.name }}</td>
            <td>{{ item.description }}</td>
            <td>
              <button
                class="btn btn-sm btn-outline-primary me-2"
                @click="$emit('edit', item)"
              >
                Editar
              </button>

              <button
                class="btn btn-sm btn-danger"
                @click="$emit('delete', item.id)"
              >
                Excluir
              </button>
            </td>
          </tr>

          <tr v-if="items.length === 0">
            <td colspan="4" class="text-center text-muted py-4">
              Nenhum item cadastrado.
            </td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>