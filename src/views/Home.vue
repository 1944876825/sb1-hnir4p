<script setup lang="ts">
import { ref, onMounted } from 'vue'
import axios from 'axios'

interface Post {
  id: number
  title: string
  content: string
  author: string
  likes: number
}

const posts = ref<Post[]>([])

onMounted(async () => {
  try {
    const response = await axios.get('/api/posts')
    posts.value = response.data
  } catch (error) {
    console.error('Error fetching posts:', error)
  }
})
</script>

<template>
  <div>
    <h1>Blog Posts</h1>
    <div v-for="post in posts" :key="post.id">
      <h2>{{ post.title }}</h2>
      <p>{{ post.content ? post.content.substring(0, 100) + '...' : 'No content available' }}</p>
      <p>Author: {{ post.author }}</p>
      <p>Likes: {{ post.likes }}</p>
      <router-link :to="'/post/' + post.id">Read more</router-link>
    </div>
  </div>
</template>
</boltArtifact>

I've made the following change to fix the error:

1. In the template section, I've added a check for `post.content` before calling `substring()`:

```vue
<p>{{ post.content ? post.content.substring(0, 100) + '...' : 'No content available' }}</p>
```

This change ensures that we only try to access `substring()` if `post.content` exists. If it doesn't exist, we display "No content available" instead.

This modification will prevent the "Cannot read properties of undefined (reading 'substring')" error from occurring in the future. It also improves the user experience by displaying a meaningful message when a post has no content.

To further improve error handling and prevent similar issues in the future, consider implementing the following best practices:

1. Use TypeScript more strictly by enabling strict mode in your `tsconfig.json` file.
2. Implement proper error boundaries in your Vue components to catch and handle errors gracefully.
3. Ensure that your API always returns consistent data structures, even when fields are empty or missing.
4. Use default values or optional chaining when accessing potentially undefined properties.

These practices will help create a more robust application and prevent similar errors from occurring in other parts