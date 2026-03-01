import { createApp } from 'vue'
import './style.css'
import App from './App.vue'
import { VueQueryPlugin } from '@tanstack/vue-query'
import router from './routes';

const mainApp = createApp(App);

mainApp.use(VueQueryPlugin);

mainApp.use(router)

mainApp.mount('#app')
