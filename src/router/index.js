import Vue from 'vue';
import VueRouter from 'vue-router';
import TableView from '../components/TableView.vue';

Vue.use(VueRouter);

const routes = [
  { path: '/api/admins', component: TableView }
];

const router = new VueRouter({
  routes
});

export default router;


/*import { createRouter, createWebHistory } from 'vue-router'
import LoginView from '@/components/LoginView.vue'
import HomeView from '@/components/TableView.vue'

const routes = [
  {
    path: '/',
    name: 'Login',
    component: LoginView
  },
  {
    path: '/home',
    name: 'Home',
    component: HomeView
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

export default router
*/
