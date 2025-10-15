<template>
  <t-aside class="sidebar-container">
    <t-menu
        class="menu"
        theme="light"
        defaultValue="PasswordCalculation"
        :collapsed="false"
        @change="handleMenuChange"
        width="15vw"
    >
      <t-menu-item value="PasswordCalculation">
        <template #icon>
          <t-icon name="key" />
        </template>
        Password Calculation
      </t-menu-item>
      <t-menu-item value="DatabaseDecryption">
        <template #icon>
          <t-icon name="data-base" />
        </template>
        Database Decryption
      </t-menu-item>
      <t-menu-item value="DataExtraction">
        <template #icon>
          <t-icon name="file" />
        </template>
        Data Extraction
      </t-menu-item>
      <t-menu-item value="RegistryAnalysis">
        <template #icon>
          <t-icon name="system-regulation" />
        </template>
        Registry Analysis
      </t-menu-item>
      <t-menu-item value="BruteForce">
        <template #icon>
          <t-icon name="lock-on" />
        </template>
        Brute Force
      </t-menu-item>
      <t-menu-item value="TimestampParser">
        <template #icon>
          <t-icon name="time" />
        </template>
        Timestamp Conversion
      </t-menu-item>
      <t-menu-item value="About">
        <template #icon>
          <t-icon name="help-circle" />
        </template>
        About
      </t-menu-item>
    </t-menu>
  </t-aside>
</template>

<script setup>
import { useRouter } from 'vue-router'
import {MessagePlugin} from "tdesign-vue-next";
const router = useRouter()

const handleMenuChange = (value) => {
  router.push({ name: value })
}
</script>

<style scoped>
.sidebar-container {
  width: 15vw;
  height: 92vh;
  overflow-y: clip;
  overflow-x: visible;
  /*border-right: 1px solid var(--td-component-stroke);*/
  flex-shrink: 0;
  padding-top: 3vh;
  padding-bottom: 5vh;
}
.menu {
  border: 3px solid blue;
  border-radius: 10px;
  margin-left: 0.5vh;
}
</style>
