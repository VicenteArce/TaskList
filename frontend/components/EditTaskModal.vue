<template>
    <div class="fixed inset-0 bg-white/30 backdrop-blur-sm flex items-center justify-center z-50">
        <div class="bg-white p-6 rounded-lg w-full max-w-md space-y-4 shadow-xl">
            <h2 class="text-xl font-bold text-gray-700">Editar Tarea</h2>
            <div class="space-y-2">
                <label class="block text-sm">Título</label>
                <input v-model="editedTask.title" class="w-full border p-2 rounded" />

                <label class="block text-sm">Descripción</label>
                <textarea v-model="editedTask.description" class="w-full border p-2 rounded"></textarea>
            </div>
            <div class="flex justify-end gap-2">
                <BaseButton variant="secondary" @click="$emit('close')">Cancelar</BaseButton>
                <BaseButton variant="primary" @click="saveChanges">Guardar</BaseButton>
            </div>
        </div>
    </div>
</template>

<script setup>
import { ref, watch } from 'vue'
import BaseButton from '@/components/BaseButton.vue'

const props = defineProps({
    task: Object,
})

const emits = defineEmits(['close', 'save'])

const editedTask = ref({ ...props.task })

watch(() => props.task, (newTask) => {
    editedTask.value = { ...newTask }
})

const saveChanges = () => {
    emits('save', editedTask.value)
}
</script>