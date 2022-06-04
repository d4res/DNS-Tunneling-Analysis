import TestView from '../views/TestView.vue'
import { createRouter, createWebHistory } from 'vue-router'
import DashboardViewVue from '@/views/DashboardView.vue'
import DataView from '../views/DataView.vue'

const routes = [
  {
    path: '/',
    name: 'home',
    component: DashboardViewVue
  },
  {
    path: '/about',
    name: 'about',
    component: () => import(/* webpackChunkName: "about" */ '../views/AboutView.vue')
  },
  {
    path: '/test',
    name: 'test',
    component: TestView
  },
  {
    path: "/data",
    name: "data", 
    component: DataView
  }
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

export default router
