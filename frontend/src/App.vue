<template>
  <t-layout>
    <t-layout class="full-height-layout">
      <SideBar />
      <MainContent class="content-layout"/>
      <t-sticky-tool
          :style="{ position: '', overflow: 'hidden' }"
          placement="left-bottom"
          shape="round"
          :offset="[-30,30]"
          @click="handleFloatButtonClick"
      >
        <t-sticky-item label="Tips">
          <template #icon><t-icon name="sticky-note" /></template>
        </t-sticky-item>
      </t-sticky-tool>
      <t-dialog v-model:visible="visible" header="Forensics Tips" class="dialog"
                :cancelBtn="null"
                :confirmBtn="null"
                :preventScrollThrough="false"
      >
        <t-list size="small" :split="true" class="tip-container">
          <t-list-item>1. XiaoHongShu: Password is "xhsdev", use SQLCipher3 for direct decryption</t-list-item>
          <t-list-item>2. WeChat IMEI: Can now be obtained through files/KeyInfo.bin, requires file decryption using RC4 algorithm with key "_wEcHAT_"</t-list-item>
          <t-list-item>3. MosGram (Bubble): Password is MD5 hash of cust_id found in sp directory's account_config.xml file, use SQLCipher4 parameters</t-list-item>
          <t-list-item>4. WuKong IM series chat databases: Database name is wk_userID.db, decryption password is the userID, use SQLCipher4 parameters</t-list-item>
        </t-list>
      </t-dialog>
    </t-layout>
  </t-layout>
</template>
<script setup>
import MainContent from "@/components/MainContent.vue";
import SideBar from "@/components/SideBar.vue";
import {ref} from "vue";
const visible = ref(false)
const handleFloatButtonClick = ()=>{
  visible.value = true
}
</script>
<style>
#logo {
  display: block;
  width: 50%;
  height: 50%;
  margin: auto;
  padding: 10% 0 0;
  background-position: center;
  background-repeat: no-repeat;
  background-size: 100% 100%;
  background-origin: content-box;
}
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  color: #2c3e50;
  height: 100vh;
}

.full-height-layout {
  height: 100vh;
}

.content-layout {
  flex: 1;
  min-height: 0; /* Key property to fix flex container scroll issues */
}

/* Responsive sidebar collapse state */
.t-layout--collapsed .floating-button-container {
  left: 64px; /* Sidebar width when collapsed */
}
.dialog{
  text-align: left;
}

.tip-container{
  overflow-y: scroll;
  height: 20vh;
}
</style>
