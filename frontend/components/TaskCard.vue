<template>
    <div class="bg-white shadow rounded-lg p-4 flex flex-col gap-2 border-l-4"
        :class="task.completed ? 'border-green-800' : 'border-yellow-500'">
        <div class="flex justify-between items-center">
            <h3 class="text-lg font-bold">{{ task.title }}</h3>
            <span :class="task.completed ? 'text-green-600' : 'text-yellow-600'" class="text-sm font-semibold">
                {{ task.completed ? 'Completada' : 'Pendiente' }}
            </span>
        </div>
        <div class="flex justify-between items-center">
            <info>
                <p class="text-sm text-gray-600">{{ task.description }}</p>
                <p class="text-xs text-gray-400">Creada el: {{ formatDate(task.created_at) }}</p>
            </info>
            <BaseButton v-if="task.completed" @click="$emit('deleteTask', task.id)" variant="danger" class="mt-2 self-start">
                Eliminar
            </BaseButton>
        </div>

        
        <div class="flex justify-between items-center mt-2">
            <!-- Emit markComplete when the button is clicked -->
            <BaseButton v-if="!task.completed" @click="$emit('markComplete', task.id)" variant="success"
                class="mt-2 self-start">
                Marcar como completada
            </BaseButton>
            <!-- Emit editTask when button is clicked-->
            <BaseButton v-if="!task.completed" @click="$emit('editTask', task.id)" variant="danger"
                class="mt-2 self-start">
                Editar
            </BaseButton>

        </div>
    </div>
</template>

<script setup>
import BaseButton from '@/components/BaseButton.vue'

const formatDate = (dateString) => {
    const date = new Date(dateString)
    return new Intl.DateTimeFormat('es-ES', {
        year: 'numeric',
        month: 'long',
        day: 'numeric',
        hour: '2-digit',
        minute: '2-digit',
    }).format(date)
}

// Define the props the component expects
defineProps({
    task: {
        type: Object,
        required: true,
    }
})

// Define the events the component can emit
defineEmits(['markComplete'])
</script>