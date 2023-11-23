import { createRouter, createWebHistory } from 'vue-router'

const routes = [
  {
    path: '/login',
    component: () => import('../views/LoginView.vue'),
  },
  {
    path: '/register',
    component: () => import('../views/RegisterView.vue'),
  },
  {
    path: '/',
    component: () => import('../views/GroupListView.vue'),
  },
  {
    path: '/group/:id(\\d+)',
    component: () => import('../views/GroupView.vue'),
  },
  {
    path: '/key_user/:id(\\d+)',
    component: () => import('../views/KeyUserView.vue'),
  },
  {
    path: '/key_users',
    component: () => import('../views/KeyUserListView.vue'),
  },
  {
    path: '/keywords',
    component: () => import('../views/KeywordListView.vue'),
  },
  {
    path: '/hyper_group',
    component: () => import('../views/HyperGroupView.vue'),
  },
  {
    path: '/fraud_detect',
    component: () => import('../views/FraudDetectView.vue')
  }
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

export default router
