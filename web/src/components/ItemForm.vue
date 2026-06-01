<script setup>
import { ref, watch } from 'vue'

const props = defineProps({
  item: {
    type: Object,
    default: () => ({
      name: '',
      description: '',
    }),
  },

  editing: {
    type: Boolean,
    default: false,
  },
})

const emit = defineEmits([
  'submit',
  'cancel',
])

const form = ref({
  name: '',
  description: '',
})

watch(
  () => props.item,
  (newValue) => {
    form.value = { ...newValue }
  },
  { immediate: true }
)

function submitForm() {
  emit('submit', form.value)
}
</script>

<template>
  <div
    class="card mb-4"
    :class="editing ? 'border-primary' : ''"
  >
    <div class="card-body">
      <div class="d-flex justify-content-between mb-3">
        <h3 class="h5 mb-0">
          {{ editing ? 'Editar Item' : 'Novo Item' }}
        </h3>

        <span
          v-if="editing"
          class="badge text-bg-primary"
        >
          Modo edição
        </span>
      </div>

      <form @submit.prevent="submitForm">
        <div class="mb-3">
          <label class="form-label">
            Nome
          </label>

          <input
            v-model="form.name"
            class="form-control"
          />
        </div>

        <div class="mb-3">
          <label class="form-label">
            Descrição
          </label>

          <textarea
            v-model="form.description"
            class="form-control"
          ></textarea>
        </div>

        <div class="d-flex gap-2">
          <button class="btn btn-primary">
            {{ editing ? 'Atualizar' : 'Salvar' }}
          </button>

          <button
            v-if="editing"
            type="button"
            class="btn btn-outline-secondary"
            @click="$emit('cancel')"
          >
            Cancelar
          </button>
        </div>
      </form>
    </div>
  </div>
</template>