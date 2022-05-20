import { createRouter, createWebHistory } from 'vue-router'

const routes = [
  {
    path: '/',
    name: 'Index',
    component: () => import('../views/IndexView.vue'),
  },
  {
    path: '/login',
    name: 'Login',
    component: () => import('../views/LoginView.vue'),
  },
  {
    path: '/group/:id(\\d+)',
    name: 'Group',
    component: () => import('../views/GroupView.vue'),
  },
  {
    path: '/user/:id(\\d+)',
    name: 'User',
    component: () => import('../views/UserView.vue'),
  },
  {
    path: '/keywords',
    name: 'Keywords',
    component: () => import('../views/KeywordView.vue'),
  }
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

export default router
