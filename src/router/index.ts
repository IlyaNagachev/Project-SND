/**
 * router/index.ts
 *
 * Automatic routes for `./src/pages/*.vue`
 */

// Composables

import { createRouter, createWebHistory } from 'vue-router'
import MainPage from '../components/general/MainPage.vue'
import ProductPage from "../components/product/ProductPage.vue";
import CatalogPage from "../components/product/GlobalCatalog.vue";

const router = createRouter({
  history: createWebHistory(),
  routes: [{path: "/main", component: MainPage},
	  {path: "/product", component: ProductPage},
	  {path: "/catalog", component: CatalogPage}],
})

// Workaround for https://github.com/vitejs/vite/issues/11804
// router.onError((err, to) => {
//   if (err?.message?.includes?.('Failed to fetch dynamically imported module')) {
//     if (!localStorage.getItem('vuetify:dynamic-reload')) {
//       console.log('Reloading page to fix dynamic import error')
//       localStorage.setItem('vuetify:dynamic-reload', 'true')
//       location.assign(to.fullPath)
//     } else {
//       console.error('Dynamic import error, reloading page did not fix it', err)
//     }
//   } else {
//     console.error(err)
//   }
// })

router.isReady().then(() => {
  localStorage.removeItem('vuetify:dynamic-reload')
})

export default router
