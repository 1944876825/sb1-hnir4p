<script setup lang="ts">
import { ref } from 'vue'
import { useUserStore } from '../stores/user'
import { useRouter } from 'vue-router'

const username = ref('')
const password = ref('')
const error = ref('')
const userStore = useUserStore()
const router = useRouter()

const login = async () => {
  try {
    error.value = ''
    await userStore.login(username.value, password.value)
    router.push('/')
  } catch (err) {
    console.error('Login failed:', err)
    error.value = 'Invalid username or password. Please try again.'
  }
}
</script>

<template>
  <div class="login-container">
    <h2>Login</h2>
    <form @submit.prevent="login" class="login-form">
      <div class="form-group">
        <label for="username">Username</label>
        <input 
          v-model="username" 
          type="text" 
          id="username"
          placeholder="Enter your username" 
          required
        >
      </div>
      <div class="form-group">
        <label for="password">Password</label>
        <input 
          v-model="password" 
          type="password" 
          id="password"
          placeholder="Enter your password" 
          required
        >
      </div>
      <button type="submit" class="login-button">Login</button>
    </form>
    <p v-if="error" class="error-message">{{ error }}</p>
    <p class="register-link">
      Don't have an account? 
      <router-link to="/register">Register here</router-link>
    </p>
  </div>
</template>

<style scoped>
.login-container {
  max-width: 400px;
  margin: 0 auto;
  padding: 20px;
  border: 1px solid #ddd;
  border-radius: 5px;
  background-color: #f9f9f9;
}

.login-form {
  display: flex;
  flex-direction: column;
}

.form-group {
  margin-bottom: 15px;
}

label {
  display: block;
  margin-bottom: 5px;
  font-weight: bold;
}

input {
  width: 100%;
  padding: 8px;
  border: 1px solid #ddd;
  border-radius: 4px;
}

.login-button {
  background-color: #4CAF50;
  color: white;
  padding: 10px;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  font-size: 16px;
}

.login-button:hover {
  background-color: #45a049;
}

.error-message {
  color: red;
  margin-top: 10px;
}

.register-link {
  margin-top: 15px;
  text-align: center;
}
</style>