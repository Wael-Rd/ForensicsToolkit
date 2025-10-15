<template>
  <div class="password-calculation">
    <t-card class="main-card">
      <template #header>
        <div class="card-header">
          <t-icon name="key" size="24px" />
          <span class="card-title">Password Calculation</span>
        </div>
      </template>
      
      <t-space direction="vertical" size="large" style="width: 100%">
        <!-- WeChat Password Calculation -->
        <t-card class="tool-card" bordered>
          <template #header>
            <h3>WeChat Database Password</h3>
          </template>
          
          <t-form :data="wechatForm" layout="vertical">
            <t-form-item label="User ID (UIN)" name="uin">
              <t-input v-model="wechatForm.uin" placeholder="Enter WeChat User ID" />
            </t-form-item>
            <t-form-item label="IMEI (Optional)" name="imei">
              <t-input v-model="wechatForm.imei" placeholder="Leave empty for default: 1234567890ABCDEF" />
            </t-form-item>
            <t-form-item>
              <t-button theme="primary" @click="calculateWechatPassword">
                Calculate WeChat Password
              </t-button>
            </t-form-item>
            <t-form-item v-if="wechatResult" label="Result">
              <t-input v-model="wechatResult" readonly>
                <template #suffix>
                  <t-button variant="text" @click="copyToClipboard(wechatResult)">
                    <t-icon name="copy" />
                  </t-button>
                </template>
              </t-input>
            </t-form-item>
          </t-form>
        </t-card>

        <!-- WeChat Index Database Password -->
        <t-card class="tool-card" bordered>
          <template #header>
            <h3>WeChat Index Database Password</h3>
          </template>
          
          <t-form :data="wechatIndexForm" layout="vertical">
            <t-form-item label="User ID (UIN)" name="uin">
              <t-input v-model="wechatIndexForm.uin" placeholder="Enter WeChat User ID" />
            </t-form-item>
            <t-form-item label="WeChat ID (WXID)" name="wxid">
              <t-input v-model="wechatIndexForm.wxid" placeholder="Enter WeChat internal ID" />
            </t-form-item>
            <t-form-item label="IMEI (Optional)" name="imei">
              <t-input v-model="wechatIndexForm.imei" placeholder="Leave empty for default: 1234567890ABCDEF" />
            </t-form-item>
            <t-form-item>
              <t-button theme="primary" @click="calculateWechatIndexPassword">
                Calculate WeChat Index Password
              </t-button>
            </t-form-item>
            <t-form-item v-if="wechatIndexResult" label="Result">
              <t-input v-model="wechatIndexResult" readonly>
                <template #suffix>
                  <t-button variant="text" @click="copyToClipboard(wechatIndexResult)">
                    <t-icon name="copy" />
                  </t-button>
                </template>
              </t-input>
            </t-form-item>
          </t-form>
        </t-card>

        <!-- WildFire IM Password -->
        <t-card class="tool-card" bordered>
          <template #header>
            <h3>WildFire IM Password</h3>
          </template>
          
          <t-form :data="wildfireForm" layout="vertical">
            <t-form-item label="Token" name="token">
              <t-textarea 
                v-model="wildfireForm.token" 
                placeholder="Enter token from shared_prefs/config.xml"
                :autosize="{ minRows: 3, maxRows: 6 }"
              />
            </t-form-item>
            <t-form-item>
              <t-button theme="primary" @click="calculateWildfirePassword">
                Calculate WildFire Password
              </t-button>
            </t-form-item>
            <t-form-item v-if="wildfireResult.password" label="Password">
              <t-input v-model="wildfireResult.password" readonly>
                <template #suffix>
                  <t-button variant="text" @click="copyToClipboard(wildfireResult.password)">
                    <t-icon name="copy" />
                  </t-button>
                </template>
              </t-input>
            </t-form-item>
            <t-form-item v-if="wildfireResult.instruction" label="Instructions">
              <t-input v-model="wildfireResult.instruction" readonly />
            </t-form-item>
          </t-form>
        </t-card>

        <!-- Quick Access Tools -->
        <t-card class="tool-card" bordered>
          <template #header>
            <h3>Quick Password Tools</h3>
          </template>
          
          <t-space direction="horizontal" size="medium" style="flex-wrap: wrap;">
            <t-button theme="default" @click="getXiaoHongShuPassword">
              XiaoHongShu Password
            </t-button>
            <t-button theme="default" @click="showMostoneForm = !showMostoneForm">
              Mostone Password
            </t-button>
            <t-button theme="default" @click="showTiktokForm = !showTiktokForm">
              TikTok Password
            </t-button>
          </t-space>

          <!-- Mostone Form -->
          <t-collapse-panel v-if="showMostoneForm" style="margin-top: 16px;">
            <t-form layout="inline">
              <t-form-item label="UID">
                <t-input v-model="mostoneUid" placeholder="Enter Mostone UID" style="width: 200px;" />
              </t-form-item>
              <t-form-item>
                <t-button theme="primary" @click="calculateMostonePassword">Calculate</t-button>
              </t-form-item>
            </t-form>
          </t-collapse-panel>

          <!-- TikTok Form -->
          <t-collapse-panel v-if="showTiktokForm" style="margin-top: 16px;">
            <t-form layout="inline">
              <t-form-item label="UID">
                <t-input v-model="tiktokUid" placeholder="Enter TikTok UID" style="width: 200px;" />
              </t-form-item>
              <t-form-item>
                <t-button theme="primary" @click="calculateTiktokPassword">Calculate</t-button>
              </t-form-item>
            </t-form>
          </t-collapse-panel>

          <!-- Quick Results -->
          <t-form v-if="quickResult.password" :data="quickResult" layout="vertical" style="margin-top: 16px;">
            <t-form-item label="Password">
              <t-input v-model="quickResult.password" readonly>
                <template #suffix>
                  <t-button variant="text" @click="copyToClipboard(quickResult.password)">
                    <t-icon name="copy" />
                  </t-button>
                </template>
              </t-input>
            </t-form-item>
            <t-form-item v-if="quickResult.instruction" label="Instructions">
              <t-input v-model="quickResult.instruction" readonly />
            </t-form-item>
          </t-form>
        </t-card>
      </t-space>
    </t-card>
  </div>
</template>

<script setup>
import { ref, reactive } from 'vue'
import { MessagePlugin } from 'tdesign-vue-next'

// WeChat forms
const wechatForm = reactive({
  uin: '',
  imei: ''
})

const wechatIndexForm = reactive({
  uin: '',
  wxid: '',
  imei: ''
})

const wildfireForm = reactive({
  token: ''
})

// Results
const wechatResult = ref('')
const wechatIndexResult = ref('')
const wildfireResult = reactive({
  password: '',
  instruction: ''
})

const quickResult = reactive({
  password: '',
  instruction: ''
})

// UI state
const showMostoneForm = ref(false)
const showTiktokForm = ref(false)
const mostoneUid = ref('')
const tiktokUid = ref('')

// Import backend functions (these would be generated by Wails)
// For demo purposes, we'll simulate the backend calls
const { CalWechat, CalWechatIndex, CalWildFire, CalMostone, CalTiktok, CalXiaoHongShu } = window.go?.passwdCalc || {}

// Calculate WeChat password
const calculateWechatPassword = async () => {
  if (!wechatForm.uin) {
    MessagePlugin.warning('Please enter User ID (UIN)')\n    return\n  }\n  \n  try {\n    // This would call the backend function\n    const result = CalWechat ? await CalWechat(wechatForm.uin, wechatForm.imei) : 'demo_password'\n    wechatResult.value = result\n    MessagePlugin.success('WeChat password calculated successfully')\n  } catch (error) {\n    MessagePlugin.error('Failed to calculate password: ' + error.message)\n  }\n}\n\n// Calculate WeChat Index password\nconst calculateWechatIndexPassword = async () => {\n  if (!wechatIndexForm.uin || !wechatIndexForm.wxid) {\n    MessagePlugin.warning('Please enter both User ID (UIN) and WeChat ID (WXID)')\n    return\n  }\n  \n  try {\n    const result = CalWechatIndex ? \n      await CalWechatIndex(wechatIndexForm.uin, wechatIndexForm.wxid, wechatIndexForm.imei) : \n      'demo_index_password'\n    wechatIndexResult.value = result\n    MessagePlugin.success('WeChat index password calculated successfully')\n  } catch (error) {\n    MessagePlugin.error('Failed to calculate password: ' + error.message)\n  }\n}\n\n// Calculate WildFire password\nconst calculateWildfirePassword = async () => {\n  if (!wildfireForm.token) {\n    MessagePlugin.warning('Please enter the token')\n    return\n  }\n  \n  try {\n    const result = CalWildFire ? await CalWildFire(wildfireForm.token) : ['demo_password', 'Use SQLCipher4 for decryption']\n    wildfireResult.password = result[0]\n    wildfireResult.instruction = result[1]\n    MessagePlugin.success('WildFire password calculated successfully')\n  } catch (error) {\n    MessagePlugin.error('Failed to calculate password: ' + error.message)\n  }\n}\n\n// Get XiaoHongShu password (static)\nconst getXiaoHongShuPassword = () => {\n  quickResult.password = 'xhsdev'\n  quickResult.instruction = 'Use SQLCipher3 for decryption'\n  MessagePlugin.success('XiaoHongShu password retrieved')\n}\n\n// Calculate Mostone password\nconst calculateMostonePassword = async () => {\n  if (!mostoneUid.value) {\n    MessagePlugin.warning('Please enter Mostone UID')\n    return\n  }\n  \n  try {\n    const result = CalMostone ? await CalMostone(mostoneUid.value) : ['DEMO12', 'Use SQLCipher3 for decryption']\n    quickResult.password = result[0]\n    quickResult.instruction = result[1]\n    MessagePlugin.success('Mostone password calculated successfully')\n  } catch (error) {\n    MessagePlugin.error('Failed to calculate password: ' + error.message)\n  }\n}\n\n// Calculate TikTok password\nconst calculateTiktokPassword = async () => {\n  if (!tiktokUid.value) {\n    MessagePlugin.warning('Please enter TikTok UID')\n    return\n  }\n  \n  try {\n    const result = CalTiktok ? await CalTiktok(tiktokUid.value) : [`byte${tiktokUid.value}imwcdb${tiktokUid.value}dance`, 'Use WCDB for decryption']\n    quickResult.password = result[0]\n    quickResult.instruction = result[1]\n    MessagePlugin.success('TikTok password calculated successfully')\n  } catch (error) {\n    MessagePlugin.error('Failed to calculate password: ' + error.message)\n  }\n}\n\n// Copy to clipboard\nconst copyToClipboard = async (text) => {\n  try {\n    await navigator.clipboard.writeText(text)\n    MessagePlugin.success('Password copied to clipboard')\n  } catch (error) {\n    MessagePlugin.error('Failed to copy to clipboard')\n  }\n}\n</script>\n\n<style scoped>\n.password-calculation {\n  padding: 20px;\n}\n\n.main-card {\n  max-width: 1000px;\n  margin: 0 auto;\n}\n\n.card-header {\n  display: flex;\n  align-items: center;\n  gap: 8px;\n}\n\n.card-title {\n  font-size: 20px;\n  font-weight: 600;\n}\n\n.tool-card {\n  background: white;\n}\n\n.tool-card h3 {\n  margin: 0;\n  color: #1976d2;\n  font-size: 16px;\n}\n</style>
