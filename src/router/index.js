import { createRouter, createWebHistory } from 'vue-router';
import TableView from '../components/TableView.vue';
import LoginView from '../components/LoginView.vue';
import RegisterCompany from '../components/RegisterCompany.vue';
import TableAllCompany from '../components/TableAllCompany.vue';
import TableSectionCompany from '../components/sectionCompany/TableSectionCompany.vue';

const routes = [
  { path: '/', component: LoginView },
  { path: '/api/admins', component: TableView },
  { path: '/api/register', component: RegisterCompany },
  { path: '/api/AllCompany', component: TableAllCompany },
  { path: '/api/company', component: TableSectionCompany},
  
  { path: '/:catchAll(.*)', redirect: '/' } // Agregar una ruta catch-all con una expresión regular personalizada
];

const router = createRouter({
  history: createWebHistory(),
  routes
});

export default router;



