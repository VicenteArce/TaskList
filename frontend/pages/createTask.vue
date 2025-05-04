<template>
    <div class="max-w-3xl mx-auto space-y-6">
        <form @submit.prevent="addTask" class="bg-white shadow-md rounded-lg p-6 space-y-4 border border-gray-200">
            <h2 class="text-xl font-semibold text-gray-700">Crear nueva tarea</h2>
            <input type="text" placeholder="Título de la tarea" v-model="newTask.title"
                class="w-full border border-gray-300 rounded px-3 py-2 focus:outline-none focus:ring-2 focus:ring-indigo-300" />
            <textarea placeholder="Descripción" rows="3" v-model="newTask.description"
                class="w-full border border-gray-300 rounded px-3 py-2 focus:outline-none focus:ring-2 focus:ring-indigo-300"></textarea>
            <BaseButton variant="primary" type="submit">Agregar Tarea</BaseButton>
        </form>
    </div>

</template>

<script setup>
import { ref } from 'vue'
import BaseButton from '@/components/BaseButton.vue'
import taskService from '@/services/taskService.ts' 
import { useRouter } from 'vue-router'

const router = useRouter()

const newTask = ref({
    title: '',
    description: '',
});


const addTask = async () => {
    try {
        const response = await taskService.create(newTask.value)
        if (response && response.data) {
            router.push({ name: 'myTasks' }) // Redirigir a la página de tareas después de crear
        } else {
            console.error('Respuesta inesperada del servidor:', response)
        }
    } catch (error) {
        console.error('Error al crear la tarea:', error)
    }
}
// Limpiar el formulario después de agregar la tarea
newTask.value = {
    title: '',
    description: '',
    dueDate: ''
}

</script>