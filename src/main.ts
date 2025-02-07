/**
 * main.ts
 *
 * Bootstraps Vuetify and other plugins then mounts the App`
 */



// Components
import App from './App.vue'
import * as components from "vuetify/components";
import * as directives from "vuetify/directives";
import '@mdi/font/css/materialdesignicons.css'
import 'vuetify/styles'

// Composables
import { createApp } from 'vue'
import router from './router'
import { createVuetify } from 'vuetify'
import {createPinia} from "pinia";


const vuetify = createVuetify({
  theme: {
    defaultTheme: 'light',
  },
  components,
  directives,
})

const app = createApp(App)

app
.use(vuetify)
.use(router)
.use(createPinia())


app.mount('#app')
