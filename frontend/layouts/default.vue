<template>
    <div class="min-h-screen bg-main text-main flex flex-col">
        <!-- Navbar -->
        <nav class="bg-navbar text-navbar px-6 py-4 flex items-center justify-center relative shadow-md">
            <div class="absolute left-6 flex items-center">
                <!-- Envolver la imagen en un NuxtLink -->
                <NuxtLink to="/">
                    <img src="../assets/images/logo.png" alt="Logo" class="h-12" />
                </NuxtLink>
            </div>
            <div class="flex items-center">
                <NuxtLink to="/createTask" class="hover:bg-navbar-hover text-lg px-3 py-2 rounded transition">
                    Crear Tarea
                </NuxtLink>
                <NuxtLink to="/" class="hover:bg-navbar-hover text-lg px-3 py-2 rounded transition">
                    Home
                </NuxtLink>
                <NuxtLink to="/myTasks" class="hover:bg-navbar-hover text-lg px-3 py-2 rounded transition">
                    Mis Tareas
                </NuxtLink>
            </div>
        </nav>

        <!-- Página -->
        <main class="p-6 flex-grow">
            <slot />
        </main>

        <footer class="bg-navbar text-navbar px-6 py-4 flex items-center justify-center">
            <div class="flex items-center">
                <p>Capacitación</p>
                <span class="mx-2">|</span>
                <p>Desarrollado por VicenteArce</p>
            </div>
        </footer>
    </div>
</template>
