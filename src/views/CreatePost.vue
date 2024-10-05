<script setup lang="ts">
import { ref } from 'vue'
import axios from 'axios'
import { useRouter } from 'vue-router'

const title = ref('')
const content = ref('')
const router = useRouter()

const createPost = async () => {
  try {
    await axios.post('/api/posts', { title: title.value, content: content.value })
    router.push('/')
  } catch (error) {
    console.error('Failed to create post:', error)
  }
}
</script>

<template>
  <div>
    <h2>Create New Post</h2>
    <form @submit.prevent="createPost">
      <input v-model="title" type="text" placeholder="Title" required>
      <textarea v-model="content" placeholder="Content" required></textarea>
      <button type="submit">Create Post</button>
    </form>
  </div>
</template>