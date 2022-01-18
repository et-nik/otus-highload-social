import { createRouter, createWebHistory } from 'vue-router'

import store from '@/store'

import Home from '../views/Home.vue'

import SignIn from "@/components/SignIn";
import SignUp from "@/components/SignUp";
import SignOut from "@/components/SignOut";

const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home,
    meta: {
      guest: true
    }
  },
  {
    path: '/sign-in',
    name: 'Sign In',
    component: SignIn,
    meta: {
      guest: true
    }
  },
  {
    path: '/sign-out',
    name: 'Sign Out',
    component: SignOut,
    meta: {
      guest: false
    }
  },
  {
    path: '/sign-up',
    name: 'Sign Up',
    component: SignUp,
    meta: {
      guest: true
    }
  },
  {
    path: '/profile',
    name: 'Profile',
    component: () => import(/* webpackChunkName: "profile" */ '../views/Profile.vue'),
    meta: {
      guest: false,
    },
  },
  {
    path: '/friends',
    name: 'Friends',
    component: () => import(/* webpackChunkName: "friends" */ '../views/Friends.vue'),
    meta: {
      guest: false,
    }
  },
  {
    path: '/users',
    name: 'Users',
    component: () => import(/* webpackChunkName: "users" */ '../views/Users.vue'),
    meta: {
      guest: false,
    }
  },
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

router.beforeEach((to, from, next) => {
  if (to.matched.some(record => !record.meta.guest)) {
    if (store.getters.isLoggedIn) {
      next()
      return
    }
    next('/sign-in')
  } else {
    next()
  }
})

export default router
