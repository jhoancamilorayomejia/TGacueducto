import { createRouter, createWebHistory } from 'vue-router';
import TableView from '../components/TableView.vue';
import LoginView from '../components/LoginView.vue';
import RegisterCompany from '../components/RegisterCompany.vue';
import TableAllCompany from '../components/TableAllCompany.vue';
import TableSectionCompany from '../components/sectionCompany/TableSectionCompany.vue';
import RegisterCustomer from '../components/RegisterCustomer.vue';
import EditAdmin from '../components/EditAdmin.vue'; // Para modificar Admin
import EditCompany from '../components/EditCompany.vue';
import EditCustomer from '../components/sectionCompany/EditCustomer.vue'; // Para modificar Customer

const routes = [
  { path: '/', component: LoginView },
  { path: '/api/admins', component: TableView },
  { path: '/api/register', component: RegisterCompany },
  { path: '/api/registerCustomer', component: RegisterCustomer },
  { path: '/api/AllCompany', component: TableAllCompany },
  { path: '/api/company', component: TableSectionCompany },
  { path: '/api/customer', component: RegisterCustomer },
  { path: '/api/admin/edit/:idadmin', name: 'EditAdmin', component: EditAdmin }, // Ruta de edición
  { path: '/api/company/edit/:idcompany', name: 'EditCompany', component: EditCompany },
  { path: '/api/customer/edit/:idcustomer', name: 'EditCustomer', component: EditCustomer }, // Ruta de edición para customer
  { path: '/:catchAll(.*)', redirect: '/' } // Ruta catch-all
];

const router = createRouter({
  history: createWebHistory(),
  routes
});

export default router;


