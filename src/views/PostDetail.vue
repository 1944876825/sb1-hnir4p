<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import axios from 'axios'

const route = useRoute()
const post = ref(null)
const comment = ref('')

onMounted(async () => {
  try {
    const response = await axios.get(`/api/posts/${route.params.id}`)
    post.value = response.data
  } catch (error) {
    console.error('Error fetching post:', error)
  }
})

const addComment = async () => {
  try {
    await axios.post(`/api/posts/${route.params.id}/comments`, { content: comment.value })
    comment.value = ''
    // Refresh comments
    const response = await axios.get(`/api/posts/${route.params.id}`)
    post.value = response.data
  } catch (error) {
    console.error('Error adding comment:', error)
  }
}

const likePost = async () => {
  try {
    await axios.post(`/api/posts/${route.params.id}/like`)
    post.value.likes++
  } catch (error) {
    console.error('Error liking post:', error)
  }
}
</script>

<template>
  <div v-if="post">
    <h1>{{ post.title }}</h1>
    <p>{{ post.content }}</p>
    <p>Author: {{ post.author }}</p>
    <p>Likes: {{ post.likes }} <button @click="likePost">Like</button></p>
    
    <h3>Comments</h3>
    <ul>
      <li v-for="comment in post.comments" :key="comment.id">
        {{ comment.content }} - {{ comment.author }}
      </li>
    </ul>
    
    <form @submit.prevent="addComment">
      <textarea v-model="comment" placeholder="Add a comment" required></textarea>
      <button type="submit">Add Comment</button>
    </form>
  </div>
</template>