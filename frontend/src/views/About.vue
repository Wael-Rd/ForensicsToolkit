<template>
  <div class="about">
    <t-card class="main-card">
      <template #header>
        <div class="card-header">
          <t-icon name="help-circle" size="24px" />
          <span class="card-title">About Digital Forensics Tool</span>
        </div>
      </template>
      
      <t-space direction="vertical" size="large" style="width: 100%">
        <!-- Application Info -->
        <t-card class="info-card" bordered>
          <template #header>
            <h3>Application Information</h3>
          </template>
          
          <t-descriptions :data="appInfo" layout="vertical" />
        </t-card>

        <!-- Features -->
        <t-card class="info-card" bordered>
          <template #header>
            <h3>Key Features</h3>
          </template>
          
          <t-list size="medium" :split="true">
            <t-list-item>
              <t-list-item-meta 
                title="Password Calculation"
                description="Generate decryption passwords for various messaging applications including WeChat, TikTok, and more."
              />
            </t-list-item>
            <t-list-item>
              <t-list-item-meta 
                title="Database Decryption"
                description="Decrypt SQLCipher databases with support for multiple encryption versions and formats."
              />
            </t-list-item>
            <t-list-item>
              <t-list-item-meta 
                title="Data Extraction"
                description="Extract forensic artifacts from mobile application databases and file systems."
              />
            </t-list-item>
            <t-list-item>
              <t-list-item-meta 
                title="Registry Analysis"
                description="Analyze Windows registry entries for forensic evidence and system information."
              />
            </t-list-item>
            <t-list-item>
              <t-list-item-meta 
                title="Brute Force Tools"
                description="Password cracking utilities for forensic investigations."
              />
            </t-list-item>
            <t-list-item>
              <t-list-item-meta 
                title="Timestamp Conversion"
                description="Convert various timestamp formats used in digital forensics."
              />
            </t-list-item>
          </t-list>
        </t-card>

        <!-- Supported Applications -->
        <t-card class="info-card" bordered>
          <template #header>
            <h3>Supported Applications</h3>
          </template>
          
          <t-row :gutter="16">
            <t-col :span="6" v-for="app in supportedApps" :key="app.name">
              <t-card size="small" :bordered="false" class="app-card">
                <div class="app-info">
                  <strong>{{ app.name }}</strong>
                  <p>{{ app.description }}</p>
                </div>
              </t-card>
            </t-col>
          </t-row>
        </t-card>

        <!-- Legal Disclaimer -->
        <t-card class="info-card" bordered>
          <template #header>
            <h3>Legal Disclaimer</h3>
          </template>
          
          <t-alert theme="warning" message="Important Legal Notice">
            <template #default>
              <p>This tool is intended for educational and research purposes in digital forensics only. </p>
              <p><strong>Do not use for illegal activities.</strong> Users are responsible for ensuring compliance with applicable laws and regulations.</p>
              <p>The developers are not responsible for any misuse of this software.</p>
            </template>
          </t-alert>
        </t-card>

        <!-- Technology Stack -->
        <t-card class="info-card" bordered>
          <template #header>
            <h3>Technology Stack</h3>
          </template>
          
          <t-row :gutter="16">
            <t-col :span="12">
              <h4>Backend</h4>
              <t-tag v-for="tech in backendTech" :key="tech" style="margin: 4px;">{{ tech }}</t-tag>
            </t-col>
            <t-col :span="12">
              <h4>Frontend</h4>
              <t-tag v-for="tech in frontendTech" :key="tech" style="margin: 4px;">{{ tech }}</t-tag>
            </t-col>
          </t-row>
        </t-card>

        <!-- Credits -->
        <t-card class="info-card" bordered>
          <template #header>
            <h3>Acknowledgments</h3>
          </template>
          
          <p>This project uses modified versions of the following libraries:</p>
          <t-list size="small">
            <t-list-item>
              <t-link href="https://github.com/WXjzcccc/registry" target="_blank">WXjzcccc/registry</t-link> - Windows Registry reading library
            </t-list-item>
            <t-list-item>
              <t-link href="https://github.com/WXjzcccc/go-sqlcipher" target="_blank">WXjzcccc/go-sqlcipher</t-link> - SQLCipher decryption library
            </t-list-item>
          </t-list>
          
          <p style="margin-top: 16px;">Special thanks to <t-link href="https://www.jetbrains.com" target="_blank">JetBrains</t-link> for providing open source licenses.</p>
        </t-card>
      </t-space>
    </t-card>
  </div>
</template>

<script setup>
import { reactive } from 'vue'

// Application information
const appInfo = reactive([
  { label: 'Application Name', value: 'Digital Forensics Tool' },
  { label: 'Version', value: '2.0.0 (English Edition)' },
  { label: 'Framework', value: 'Wails v2' },
  { label: 'Author', value: 'MiniMax Agent' },
  { label: 'Original Project', value: 'WXjzcccc/ForensicsTool' },
  { label: 'License', value: 'Open Source' }
])

// Supported applications
const supportedApps = reactive([
  { name: 'WeChat', description: 'Android WeChat databases and message history' },
  { name: 'XiaoHongShu', description: 'Little Red Book social media app data' },
  { name: 'TikTok', description: 'TikTok application databases' },
  { name: 'MosGram/Bubble', description: 'Bubble messaging application' },
  { name: 'WuKong IM', description: 'WuKong instant messaging series' },
  { name: 'WildFire IM', description: 'WildFire messaging applications' },
  { name: 'Mostone', description: 'Mostone messaging platform' },
  { name: 'Custom Apps', description: 'SQLCipher encrypted databases' }
])

// Technology stacks
const backendTech = ['Go', 'Wails v2', 'SQLCipher', 'Crypto Libraries']
const frontendTech = ['Vue 3', 'TDesign', 'Vite', 'JavaScript']
</script>

<style scoped>
.about {
  padding: 20px;
}

.main-card {
  max-width: 1000px;
  margin: 0 auto;
}

.card-header {
  display: flex;
  align-items: center;
  gap: 8px;
}

.card-title {
  font-size: 20px;
  font-weight: 600;
}

.info-card {
  background: white;
}

.info-card h3 {
  margin: 0;
  color: #1976d2;
  font-size: 16px;
}

.info-card h4 {
  margin: 0 0 8px 0;
  color: #666;
  font-size: 14px;
}

.app-card {
  height: 100%;
  background: #f8f9fa;
}

.app-info strong {
  color: #1976d2;
}

.app-info p {
  margin: 4px 0 0 0;
  font-size: 12px;
  color: #666;
}
</style>
