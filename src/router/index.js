import TestView from '../views/TestView.vue'
import { createRouter, createWebHistory } from 'vue-router'
import DashboardViewVue from '@/views/DashboardView.vue'
import DataView from '../views/DataView.vue'
import MetricViewVue from '@/views/MetricView.vue'
import InfoViewVue from '@/views/InfoView.vue'
//import InfoViewVue from '@/views/InfoView.vue'

const routes = [
  {
    path: '/',
    name: 'home',
    component: DashboardViewVue
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
  },
  {
    path: "/metric",
    name: "metric",
    component: MetricViewVue
  },
  {
    path: "/info",
    name: "info",
    component: InfoViewVue
  } 
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

export default router
