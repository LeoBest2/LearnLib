import { createRouter, createWebHashHistory } from 'vue-router'
import Home from '../views/Home.vue'
import Page1 from "../views/Page1.vue";
import Page2 from "../views/Page2.vue";

const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home,
    meta: { displayName: '首页' }
  },
  {
    path: '/page1',
    name: 'Page1',
    component: Page1,
    meta: { displayName: '页面1' }
  },
  {
    path: '/page2',
    name: 'Page2',
    component: Page2,
    meta: { displayName: '页面2' }
  },
  {
    path: '/about',
    name: 'About',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () => import(/* webpackChunkName: "about" */ '../views/About.vue'),
    meta: { displayName: '关于页面' }
  },
  {
    path: '/login',
    name: 'Login',
    component: () => import('../views/Login.vue'),
    meta: { displayName: '登陆页面' }
  }
]

const router = createRouter({
  history: createWebHashHistory(),
  routes
})


export default router
