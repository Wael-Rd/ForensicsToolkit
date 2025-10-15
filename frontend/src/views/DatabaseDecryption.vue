<template>
  <div class="database-decryption">
    <t-card class="main-card">
      <template #header>
        <div class="card-header">
          <t-icon name="data-base" size="24px" />
          <span class="card-title">Database Decryption</span>
        </div>
      </template>
      
      <t-space direction="vertical" size="large" style="width: 100%">
        <!-- Database Type Selection -->
        <t-card class="tool-card" bordered>
          <template #header>
            <h3>Select Database Type</h3>
          </template>
          
          <t-radio-group v-model="selectedDbType" variant="default-filled">
            <t-radio value="wechat">WeChat (EnMicroMsg)</t-radio>
            <t-radio value="sqlcipher3">SQLCipher 3</t-radio>
            <t-radio value="sqlcipher4">SQLCipher 4</t-radio>
            <t-radio value="ftsindex">FTS Index</t-radio>
            <t-radio value="wcdb">WCDB</t-radio>
          </t-radio-group>
        </t-card>

        <!-- File Selection and Password Input -->
        <t-card class="tool-card" bordered>
          <template #header>
            <h3>Database File and Credentials</h3>
          </template>
          
          <t-form :data="decryptForm" layout="vertical">
            <t-form-item label="Database File Path" name="dbPath">
              <t-input 
                v-model="decryptForm.dbPath" 
                placeholder="Select or enter database file path"
                readonly
              >
                <template #suffix>
                  <t-button variant="text" @click="selectDatabaseFile">
                    <t-icon name="folder-open" />
                    Browse
                  </t-button>
                </template>
              </t-input>
            </t-form-item>
            
            <t-form-item label="Decryption Password" name="password">
              <t-input 
                v-model="decryptForm.password" 
                placeholder="Enter decryption password"
                type="password"
                show-password
              />
            </t-form-item>

            <!-- Output Directory -->
            <t-form-item label="Output Directory (Optional)" name="outputDir">
              <t-input 
                v-model="decryptForm.outputDir" 
                placeholder="Leave empty to save alongside original file"
              >
                <template #suffix>
                  <t-button variant="text" @click="selectOutputDirectory">
                    <t-icon name="folder-open" />
                    Browse
                  </t-button>
                </template>
              </t-input>
            </t-form-item>

            <t-form-item>
              <t-space>
                <t-button 
                  theme="primary" 
                  @click="decryptDatabase"
                  :loading="isDecrypting"
                  :disabled="!decryptForm.dbPath || !decryptForm.password"
                >
                  {{ isDecrypting ? 'Decrypting...' : 'Decrypt Database' }}
                </t-button>
                <t-button theme="default" @click="clearForm">
                  Clear
                </t-button>
              </t-space>
            </t-form-item>
          </t-form>
        </t-card>

        <!-- Results -->
        <t-card v-if="decryptResult" class="tool-card" bordered>
          <template #header>
            <h3>Decryption Results</h3>
          </template>
          
          <t-form layout="vertical">
            <t-form-item v-if="decryptResult.err" label="Error">
              <t-textarea 
                :value="decryptResult.err" 
                readonly
                theme="error"
                :autosize="{ minRows: 2, maxRows: 6 }"
              />
            </t-form-item>
            
            <t-form-item v-if="decryptResult.save_path" label="Decrypted Database Path">
              <t-input :value="decryptResult.save_path" readonly>
                <template #suffix>
                  <t-space>
                    <t-button variant="text" @click="copyToClipboard(decryptResult.save_path)">
                      <t-icon name="copy" />
                    </t-button>
                    <t-button variant="text" @click="openFileLocation(decryptResult.save_path)">
                      <t-icon name="folder-open" />
                    </t-button>
                  </t-space>
                </template>
              </t-input>
            </t-form-item>
            
            <t-form-item v-if="decryptResult.wxid" label="WeChat ID (WXID)">
              <t-input :value="decryptResult.wxid" readonly>
                <template #suffix>
                  <t-button variant="text" @click="copyToClipboard(decryptResult.wxid)">
                    <t-icon name="copy" />
                  </t-button>
                </template>
              </t-input>
            </t-form-item>

            <t-alert v-if="decryptResult.save_path && !decryptResult.err" theme="success" message="Database decrypted successfully!" />
          </t-form>
        </t-card>

        <!-- Batch Decryption -->
        <t-card class="tool-card" bordered>
          <template #header>
            <h3>Batch Decryption</h3>
          </template>
          
          <t-form layout="vertical">
            <t-form-item label="Database Directory">
              <t-input 
                v-model="batchForm.directory" 
                placeholder="Select directory containing database files"
                readonly
              >
                <template #suffix>
                  <t-button variant="text" @click="selectBatchDirectory">
                    <t-icon name="folder-open" />
                    Browse
                  </t-button>
                </template>
              </t-input>
            </t-form-item>
            
            <t-form-item label="File Pattern (Optional)">
              <t-input 
                v-model="batchForm.pattern" 
                placeholder="e.g., *.db, EnMicroMsg*.db (leave empty for all .db files)"
              />
            </t-form-item>
            
            <t-form-item label="Default Password">
              <t-input 
                v-model="batchForm.password" 
                placeholder="Password to try for all databases"
                type="password"
                show-password
              />
            </t-form-item>

            <t-form-item>
              <t-button 
                theme="primary" 
                @click="batchDecrypt"
                :loading="isBatchDecrypting"
                :disabled="!batchForm.directory || !batchForm.password"
              >
                {{ isBatchDecrypting ? 'Processing...' : 'Start Batch Decryption' }}
              </t-button>
            </t-form-item>
          </t-form>

          <!-- Batch Results -->
          <t-table 
            v-if="batchResults.length > 0"
            :data="batchResults" 
            :columns="batchColumns"
            stripe
            style="margin-top: 16px;"
          />
        </t-card>

        <!-- Help Section -->
        <t-card class="tool-card" bordered>
          <template #header>
            <h3>Database Type Information</h3>
          </template>
          
          <t-list size="small" :split="true">
            <t-list-item>
              <strong>WeChat (EnMicroMsg):</strong> Android WeChat encrypted message database
            </t-list-item>
            <t-list-item>
              <strong>SQLCipher 3:</strong> Databases encrypted with SQLCipher version 3.x
            </t-list-item>
            <t-list-item>
              <strong>SQLCipher 4:</strong> Databases encrypted with SQLCipher version 4.x (default for newer apps)
            </t-list-item>
            <t-list-item>
              <strong>FTS Index:</strong> Full-text search index databases with specific encryption parameters
            </t-list-item>
            <t-list-item>
              <strong>WCDB:</strong> WeChat Database format used by some applications
            </t-list-item>
          </t-list>
        </t-card>
      </t-space>
    </t-card>
  </div>
</template>

<script setup>
import { ref, reactive } from 'vue'
import { MessagePlugin } from 'tdesign-vue-next'

// Form data
const selectedDbType = ref('wechat')
const decryptForm = reactive({
  dbPath: '',
  password: '',
  outputDir: ''
})

const batchForm = reactive({
  directory: '',
  pattern: '*.db',
  password: ''
})

// State
const isDecrypting = ref(false)
const isBatchDecrypting = ref(false)
const decryptResult = ref(null)
const batchResults = ref([])

// Batch results table columns
const batchColumns = [
  { colKey: 'filename', title: 'Database File', width: '200' },
  { colKey: 'status', title: 'Status', width: '100' },
  { colKey: 'result', title: 'Result/Error', ellipsis: true },
  { colKey: 'outputPath', title: 'Output Path', ellipsis: true }
]

// Backend functions (would be imported from Wails generated code)
const { 
  DecryptEnMicroMsg, 
  DecryptSQLCipher3, 
  DecryptSQLCipher4, 
  DecryptFTSIndex, 
  DecryptWCdb 
} = window.go?.database || {}

// File selection functions (would use Wails file dialogs)
const selectDatabaseFile = async () => {
  // In real implementation, this would use Wails file dialog
  const input = document.createElement('input')
  input.type = 'file'\n  input.accept = '.db,.sqlite,.sqlite3'\n  input.onchange = (e) => {\n    const file = e.target.files[0]\n    if (file) {\n      decryptForm.dbPath = file.path || file.name\n    }\n  }\n  input.click()\n}\n\nconst selectOutputDirectory = async () => {\n  // In real implementation, this would use Wails directory dialog\n  MessagePlugin.info('Output directory selection would use native file dialog')\n}\n\nconst selectBatchDirectory = async () => {\n  // In real implementation, this would use Wails directory dialog\n  MessagePlugin.info('Batch directory selection would use native file dialog')\n}\n\n// Main decryption function\nconst decryptDatabase = async () => {\n  isDecrypting.value = true\n  decryptResult.value = null\n  \n  try {\n    let result\n    \n    switch (selectedDbType.value) {\n      case 'wechat':\n        result = DecryptEnMicroMsg ? \n          await DecryptEnMicroMsg(decryptForm.dbPath, decryptForm.password) :\n          { save_path: '/demo/decrypted.db', wxid: 'demo_wxid', err: '' }\n        break\n      case 'sqlcipher3':\n        result = DecryptSQLCipher3 ? \n          await DecryptSQLCipher3(decryptForm.dbPath, decryptForm.password) :\n          { save_path: '/demo/decrypted.db', wxid: '', err: '' }\n        break\n      case 'sqlcipher4':\n        result = DecryptSQLCipher4 ? \n          await DecryptSQLCipher4(decryptForm.dbPath, decryptForm.password) :\n          { save_path: '/demo/decrypted.db', wxid: '', err: '' }\n        break\n      case 'ftsindex':\n        result = DecryptFTSIndex ? \n          await DecryptFTSIndex(decryptForm.dbPath, decryptForm.password) :\n          { save_path: '/demo/decrypted.db', wxid: '', err: '' }\n        break\n      case 'wcdb':\n        result = DecryptWCdb ? \n          await DecryptWCdb(decryptForm.dbPath, decryptForm.password) :\n          { save_path: '/demo/decrypted.db', wxid: '', err: '' }\n        break\n      default:\n        throw new Error('Unknown database type')\n    }\n    \n    decryptResult.value = result\n    \n    if (result.err) {\n      MessagePlugin.error('Decryption failed')\n    } else {\n      MessagePlugin.success('Database decrypted successfully!')\n    }\n  } catch (error) {\n    MessagePlugin.error('Decryption error: ' + error.message)\n    decryptResult.value = { err: error.message }\n  } finally {\n    isDecrypting.value = false\n  }\n}\n\n// Batch decryption\nconst batchDecrypt = async () => {\n  isBatchDecrypting.value = true\n  batchResults.value = []\n  \n  try {\n    // In real implementation, this would scan the directory and process each file\n    MessagePlugin.info('Batch decryption would process all matching files in the directory')\n    \n    // Demo results\n    batchResults.value = [\n      { filename: 'EnMicroMsg.db', status: 'Success', result: 'Decrypted successfully', outputPath: '/path/to/EnMicroMsg_dec.db' },\n      { filename: 'index.db', status: 'Failed', result: 'Invalid password', outputPath: '' },\n      { filename: 'message.db', status: 'Success', result: 'Decrypted successfully', outputPath: '/path/to/message_dec.db' }\n    ]\n  } catch (error) {\n    MessagePlugin.error('Batch decryption error: ' + error.message)\n  } finally {\n    isBatchDecrypting.value = false\n  }\n}\n\n// Clear form\nconst clearForm = () => {\n  decryptForm.dbPath = ''\n  decryptForm.password = ''\n  decryptForm.outputDir = ''\n  decryptResult.value = null\n}\n\n// Utility functions\nconst copyToClipboard = async (text) => {\n  try {\n    await navigator.clipboard.writeText(text)\n    MessagePlugin.success('Copied to clipboard')\n  } catch (error) {\n    MessagePlugin.error('Failed to copy to clipboard')\n  }\n}\n\nconst openFileLocation = (path) => {\n  // In real implementation, this would use Wails to open file explorer\n  MessagePlugin.info(`Would open file location: ${path}`)\n}\n</script>\n\n<style scoped>\n.database-decryption {\n  padding: 20px;\n}\n\n.main-card {\n  max-width: 1000px;\n  margin: 0 auto;\n}\n\n.card-header {\n  display: flex;\n  align-items: center;\n  gap: 8px;\n}\n\n.card-title {\n  font-size: 20px;\n  font-weight: 600;\n}\n\n.tool-card {\n  background: white;\n}\n\n.tool-card h3 {\n  margin: 0;\n  color: #1976d2;\n  font-size: 16px;\n}\n</style>
