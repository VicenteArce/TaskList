<template>
    <div class="max-w-3xl mx-auto space-y-6">
        <!-- Filtro -->
        <div class="flex justify-between items-center bg-white shadow-md rounded-lg p-4 border border-gray-200">
            <h2 class="text-xl font-semibold text-gray-700">Mis Tareas</h2>
            <div class="flex space-x-4">
                <BaseButton variant="secondary" @click="filterTasks('all')">Todas</BaseButton>
                <BaseButton variant="secondary" @click="filterTasks('completed')">Completadas</BaseButton>
                <BaseButton variant="secondary" @click="filterTasks('pending')">Pendientes</BaseButton>
            </div>
        </div>
        <!-- Mensaje de éxito -->
        <div v-if="successMessage" class="bg-white shadow border-green-800 rounded-lg p-4 flex flex-col gap-2 border-l-4">
            <p class="text font-semibold text-green">{{ successMessage }}</p>
        </div> 
        <!-- Lista de tareas -->
        <div class="space-y-4">
            <TaskItem v-for="task in tasks" :key="task.id" :task="task" @markComplete="markTaskAsComplete"
                @editTask="openEditModal" @deleteTask="deleteTask"/>
            <!-- Optional: Message if no tasks -->
            <p v-if="tasks.length === 0" class="text-center text-gray-500 italic">No hay tareas pendientes.</p>
        </div>
    </div>
    <EditTaskModal v-if="showModal" :task="selectedTask" @close="showModal = false" @save="updateTask" />
</template>

<script setup>
import { ref } from 'vue'
import BaseButton from '@/components/BaseButton.vue'
import TaskItem from '@/components/TaskCard.vue'
import taskService from '@/services/taskService.ts'
import EditTaskModal from '@/components/EditTaskModal.vue'

const successMessage = ref('') // Mensaje de éxito para mostrar al usuario
const tasks = ref([])           // Array para almacenar las tareas
const completedTasks = ref([])  // Array para almacenar las tareas completadas
const pendingTasks = ref([])    // Array para almacenar las tareas pendientes
const filter = ref('all')       // Filtro de tareas
const showModal = ref(false)
const selectedTask = ref(null)



const filterTasks = (status) => {
    filter.value = status
    console.log('Filtrar tareas por estado:', status)
    if (status === 'completed') {
        tasks.value = completedTasks.value
    } else if (status === 'pending') {
        tasks.value = pendingTasks.value
    } else {
        tasks.value = [...pendingTasks.value, ...completedTasks.value]
    }
}

const fetchTasks = async () => {
    try {
        const response = await taskService.getAll()
        if (response && response.data) {
            tasks.value = response.data // Asignar las tareas a la variable reactiva
            completedTasks.value = response.data.filter(task => task.completed)
            pendingTasks.value = response.data.filter(task => !task.completed)
            tasks.value = [...pendingTasks.value, ...completedTasks.value] // Combinar tareas pendientes y completadas
        } else {
            console.error('Respuesta inesperada del servidor:', response)
        }
    } catch (error) {
        console.error('Error al obtener las tareas:', error)
    }
}

const markTaskAsComplete = async (taskId) => {
    try {
        const response = await taskService.complete(taskId)
        if (response && response.data) {
            // Actualizar la tarea en el array de tareas
            const taskIndex = tasks.value.findIndex(task => task.id === taskId)
            if (taskIndex !== -1) {
                tasks.value[taskIndex].completed = true
            }
            // Filtrar las tareas nuevamente
            filterTasks(filter.value)
        } else {
            console.error('Respuesta inesperada del servidor:', response)
        }
    } catch (error) {
        console.error('Error al marcar la tarea como completada:', error)
    }
}

const openEditModal = (taskId) => {
    selectedTask.value = tasks.value.find(t => t.id === taskId)
    showModal.value = true
}

const updateTask = async (updatedTask) => {
    try {
        const response = await taskService.update(updatedTask.id, updatedTask)
        if (response && response.data) {
            showModal.value = false
            await fetchTasks() // Recargar todas las tareas desde el backend
            filterTasks(filter.value)

            // Mostrar mensaje de éxito
            successMessage.value = '¡Tarea actualizada exitosamente!'
            // Ocultar el mensaje después de 3 segundos
            setTimeout(() => {
                successMessage.value = ''
            }, 2500)
        }
    } catch (error) {
        console.error('Error al actualizar la tarea:', error)
    }
}

const deleteTask = async (taskId) => {
    try {
        const response = await taskService.remove(taskId)
        if (response && response.data) {
            // Mostrar mensaje de éxito
            successMessage.value = '¡Tarea eliminada exitosamente!'
            
            // Recargar las tareas desde el backend
            await fetchTasks()

            // Ocultar el mensaje después de 3 segundos
            setTimeout(() => {
                successMessage.value = ''
            }, 2500)
        } else {
            console.error('Respuesta inesperada del servidor:', response)
        }
    } catch (error) {
        console.error('Error al eliminar la tarea:', error)
    }
}




// Llamar a fetchTasks al montar el componente
fetchTasks()
</script>