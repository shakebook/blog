import { AppConfig } from 'vue'
import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import store, { ROOT_KEY } from './store'
import 'reset-css'
import './base.css'

createApp(App).use(store, ROOT_KEY).use(router).mount('#app')
