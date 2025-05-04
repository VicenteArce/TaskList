<template>
    <button
      :class="[
        'base-button cursor-pointer px-4 py-2 rounded text-white font-semibold',
        variant === 'primary' ? 'btn-primary' :
        variant === 'secondary' ? 'btn-secondary' :
        variant === 'success' ? 'btn-success' :
        variant === 'danger' ? 'btn-danger' : ''
      ]"
      v-bind="$attrs"
      
    >
      <slot />
    </button>
  </template>
  
  <script setup>
  defineProps({
    variant: {
      type: String,
      default: 'primary', // primary | secondary | success | danger
    }
  })
  </script>
  