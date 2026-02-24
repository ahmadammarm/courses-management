import { createApp } from 'vue'
import './style.css'
import App from './App.vue'
import { VueQueryPlugin } from '@tanstack/vue-query'

const mainApp = createApp(App);

mainApp.use(VueQueryPlugin);

mainApp.mount('#app')
