import { createRouter, createWebHistory } from 'vue-router';
import TableView from '../components/TableView.vue';
import LoginView from '../components/LoginView.vue';
import RegisterCompany from '../components/RegisterCompany.vue';

const routes = [
  { path: '/', component: LoginView },
  { path: '/api/admins', component: TableView },
  { path: '/api/register', component: RegisterCompany },
  // Agregar una ruta catch-all con una expresión regular personalizada
  { path: '/:catchAll(.*)', redirect: '/' }
];

const router = createRouter({
  history: createWebHistory(),
  routes
});

export default router;



