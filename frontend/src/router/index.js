import { createRouter, createWebHistory } from 'vue-router'
import PasswordCalculation from '../views/PasswordCalculation.vue'
import DatabaseDecryption from '../views/DatabaseDecryption.vue'
import DataExtraction from '../views/DataExtraction.vue'
import RegistryAnalysis from '../views/RegistryAnalysis.vue'
import BruteForce from '../views/BruteForce.vue'
import TimestampParser from '../views/TimestampParser.vue'
import About from '../views/About.vue'

const routes = [
  {
    path: '/',
    redirect: '/password-calculation'
  },
  {
    path: '/password-calculation',
    name: 'PasswordCalculation',
    component: PasswordCalculation
  },
  {
    path: '/database-decryption',
    name: 'DatabaseDecryption',
    component: DatabaseDecryption
  },
  {
    path: '/data-extraction',
    name: 'DataExtraction',
    component: DataExtraction
  },
  {
    path: '/registry-analysis',
    name: 'RegistryAnalysis',
    component: RegistryAnalysis
  },
  {
    path: '/brute-force',
    name: 'BruteForce',
    component: BruteForce
  },
  {
    path: '/timestamp-parser',
    name: 'TimestampParser',
    component: TimestampParser
  },
  {
    path: '/about',
    name: 'About',
    component: About
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

export default router
