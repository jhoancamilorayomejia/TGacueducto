import { createRouter, createWebHistory } from 'vue-router';
import TableView from '../components/TableView.vue';
import LoginView from '../components/LoginView.vue';
import RegisterCompany from '../components/RegisterCompany.vue';
import RegisterAdmin from '../components/RegisterAdmin.vue';
import TableAllCompany from '../components/TableAllCompany.vue';
import TableSectionCompany from '../components/sectionCompany/TableSectionCompany.vue';
import RegisterCustomer from '../components/sectionCompany/RegisterCustomer.vue';
import EditAdmin from '../components/EditAdmin.vue'; // Para modificar Admin
import EditCompany from '../components/EditCompany.vue';
import EditCustomer from '../components/sectionCompany/EditCustomer.vue'; // Para modificar Customer
import InfoCompany from '../components/CompanyInfo.vue';
import WelcomeCustomer from '../components/sectionCustomer/WelcomeCustomer.vue';
import FactureInfoCustomer from '@/components/sectionCompany/FactureInfoCustomer.vue';
import FactureNew from '@/components/sectionCompany/FactureNew.vue';
import PayView from '../components/sectionCustomer/payView.vue';
//import SuccessMessage from '../components/sectionCustomer/SuccessMessage.vue';
//import FailureMessage from '../components/sectionCustomer/FailureMessage.vue';
//import PendingStatus from '../components/sectionCustomer/PendingStatus.vue';

const routes = [
  { path: '/', component: LoginView },
  { path: '/api/admins/:idadmin', component: TableView },
  { path: '/api/register', component: RegisterCompany },
  { path: '/api/registerAdmin', component: RegisterAdmin },
  { path: '/api/registerCustomer', component: RegisterCustomer },
  { path: '/api/AllCompany', component: TableAllCompany },
  { path: '/api/company/:idcompany', component: TableSectionCompany },
  { path: '/api/welcome/:idcustomer', component: WelcomeCustomer },
  { path: '/api/admin/edit/:idadmin', name: 'EditAdmin', component: EditAdmin }, // Ruta de edición
  { path: '/api/company/edit/:idcompany', name: 'EditCompany', component: EditCompany },
  { path: '/api/customer/edit/:idcustomer', name: 'EditCustomer', component: EditCustomer }, // Ruta de edición para customer
  { path: '/api/company/info/:idcompany', name: 'InfoCompany', component: InfoCompany },
  { path: '/api/customer/info-facture/:idcustomer', name: 'FactureInfoCustomer', component: FactureInfoCustomer },
  { path: '/api/customer/new-facture/:idcustomer', name: 'FactureNew', component: FactureNew },
  { path: '/api/payment/:totalpay', component: PayView },
  //{ path: '/success', component: SuccessMessage },
  //{ path: '/failure', component: FailureMessage },
  //{ path: '/pending', component: PendingStatus },

  //{ path: '/success', component: SuccessMessage },
  //{ path: '/failure', component: FailureMessage },
  //{ path: '/pending', component: PendingStatus },




  { path: '/:catchAll(.*)', redirect: '/' } // Ruta catch-all
];

const router = createRouter({
  history: createWebHistory(),
  routes
});





export default router;


