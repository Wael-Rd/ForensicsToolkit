<template>
  <div class="data-extraction">
    <t-card class="main-card">
      <template #header>
        <div class="card-header">
          <t-icon name="file" size="24px" />
          <span class="card-title">Data Extraction</span>
        </div>
      </template>
      
      <t-space direction="vertical" size="large" style="width: 100%">
        <t-card class="tool-card" bordered>
          <template #header>
            <h3>Data extraction tools will be implemented here</h3>
          </template>
          <p>This section will contain tools for extracting forensic data from various sources.</p>
        </t-card>
      </t-space>
    </t-card>
  </div>
</template>

<script setup>
// Data extraction component - to be implemented
</script>

<style scoped>
.data-extraction {
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

.tool-card {
  background: white;
}

.tool-card h3 {
  margin: 0;
  color: #1976d2;
  font-size: 16px;
}
</style>
