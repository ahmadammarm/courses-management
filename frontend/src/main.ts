import { createApp } from 'vue'
import './style.css'
import App from './App.vue'
import { VueQueryPlugin } from '@tanstack/vue-query'
import router from './routes';
import { setupSession } from './lib/session';

const mainApp = createApp(App);

mainApp.use(VueQueryPlugin);

mainApp.use(router)

// Periksa session/token saat aplikasi dimulai dan atur auto-logout
setupSession();

mainApp.mount('#app')
