<template>
  <div class="timestamp-parser">
    <t-card class="main-card">
      <template #header>
        <div class="card-header">
          <t-icon name="time" size="24px" />
          <span class="card-title">Timestamp Conversion</span>
        </div>
      </template>
      
      <t-space direction="vertical" size="large" style="width: 100%">
        <!-- Universal Timestamp Converter -->
        <t-card class="tool-card" bordered>
          <template #header>
            <h3>Universal Timestamp Converter</h3>
          </template>
          
          <t-form :data="universalForm" layout="vertical">
            <t-form-item label="Timestamp" name="timestamp">
              <t-input 
                v-model="universalForm.timestamp" 
                placeholder="Enter timestamp (auto-detects format)"
              />
            </t-form-item>
            
            <t-row :gutter="16">
              <t-col :span="12">
                <t-form-item label="Source Timezone" name="sourceTimezone">
                  <t-select v-model="universalForm.sourceTimezone" :options="timezoneOptions" filterable />
                </t-form-item>
              </t-col>
              <t-col :span="12">
                <t-form-item label="Target Timezone" name="targetTimezone">
                  <t-select v-model="universalForm.targetTimezone" :options="timezoneOptions" filterable />
                </t-form-item>
              </t-col>
            </t-row>
            
            <t-form-item>
              <t-button theme="primary" @click="convertUniversalTimestamp">
                Convert Timestamp
              </t-button>
            </t-form-item>
            
            <t-form-item v-if="universalResult" label="Converted Time">
              <t-input :value="universalResult" readonly>
                <template #suffix>
                  <t-button variant="text" @click="copyToClipboard(universalResult)">
                    <t-icon name="copy" />
                  </t-button>
                </template>
              </t-input>
            </t-form-item>
          </t-form>
        </t-card>

        <!-- Specific Format Converters -->
        <t-card class="tool-card" bordered>
          <template #header>
            <h3>Specific Format Converters</h3>
          </template>
          
          <t-tabs v-model="activeTab">
            <!-- iOS Timestamp -->
            <t-tab-panel value="ios" label="iOS Timestamp">
              <t-form layout="vertical">
                <t-form-item label="iOS Timestamp (seconds since 2001-01-01)">
                  <t-input v-model="iosTimestamp" placeholder="e.g., 656208000.123456" />
                </t-form-item>
                <t-form-item>
                  <t-button theme="primary" @click="convertIosTimestamp">
                    Convert iOS Timestamp
                  </t-button>
                </t-form-item>
                <t-form-item v-if="iosResult" label="Converted Time">
                  <t-input :value="iosResult" readonly>
                    <template #suffix>
                      <t-button variant="text" @click="copyToClipboard(iosResult)">
                        <t-icon name="copy" />
                      </t-button>
                    </template>
                  </t-input>
                </t-form-item>
              </t-form>
            </t-tab-panel>

            <!-- Chrome Timestamp -->
            <t-tab-panel value="chrome" label="Chrome Timestamp">
              <t-form layout="vertical">
                <t-form-item label="Chrome Timestamp (microseconds since 1601-01-01)">
                  <t-input v-model="chromeTimestamp" placeholder="e.g., 13287974400000000" />
                </t-form-item>
                <t-form-item>
                  <t-button theme="primary" @click="convertChromeTimestamp">
                    Convert Chrome Timestamp
                  </t-button>
                </t-form-item>
                <t-form-item v-if="chromeResult" label="Converted Time">
                  <t-input :value="chromeResult" readonly>
                    <template #suffix>
                      <t-button variant="text" @click="copyToClipboard(chromeResult)">
                        <t-icon name="copy" />
                      </t-button>
                    </template>
                  </t-input>
                </t-form-item>
              </t-form>
            </t-tab-panel>

            <!-- Windows File Time -->
            <t-tab-panel value="windows" label="Windows File Time">
              <t-form layout="vertical">
                <t-form-item label="Windows File Time (100-nanosecond intervals since 1601-01-01)">
                  <t-input v-model="windowsTimestamp" placeholder="e.g., 132879744000000000" />
                </t-form-item>
                <t-form-item>
                  <t-button theme="primary" @click="convertWindowsTimestamp">
                    Convert Windows File Time
                  </t-button>
                </t-form-item>
                <t-form-item v-if="windowsResult" label="Converted Time">
                  <t-input :value="windowsResult" readonly>
                    <template #suffix>
                      <t-button variant="text" @click="copyToClipboard(windowsResult)">
                        <t-icon name="copy" />
                      </t-button>
                    </template>
                  </t-input>
                </t-form-item>
              </t-form>
            </t-tab-panel>

            <!-- Apple Timestamp -->
            <t-tab-panel value="apple" label="Apple Timestamp">
              <t-form layout="vertical">
                <t-form-item label="Apple Timestamp (seconds since 2001-01-01, 9-digit format)">
                  <t-input v-model="appleTimestamp" placeholder="e.g., 656208000.123456789" />
                </t-form-item>
                <t-form-item>
                  <t-button theme="primary" @click="convertAppleTimestamp">
                    Convert Apple Timestamp
                  </t-button>
                </t-form-item>
                <t-form-item v-if="appleResult" label="Converted Time">
                  <t-input :value="appleResult" readonly>
                    <template #suffix>
                      <t-button variant="text" @click="copyToClipboard(appleResult)">
                        <t-icon name="copy" />
                      </t-button>
                    </template>
                  </t-input>
                </t-form-item>
              </t-form>
            </t-tab-panel>
          </t-tabs>
        </t-card>

        <!-- Current Time -->
        <t-card class="tool-card" bordered>
          <template #header>
            <h3>Current Time Reference</h3>
          </template>
          
          <t-descriptions :data="currentTimeData" layout="horizontal" />
          
          <t-button theme="default" @click="updateCurrentTime" style="margin-top: 16px;">
            <t-icon name="refresh" />
            Refresh Current Time
          </t-button>
        </t-card>

        <!-- Help Information -->
        <t-card class="tool-card" bordered>
          <template #header>
            <h3>Timestamp Format Information</h3>
          </template>
          
          <t-list size="small" :split="true">
            <t-list-item>
              <strong>Unix Timestamp:</strong> Seconds since 1970-01-01 00:00:00 UTC (10 digits)
            </t-list-item>
            <t-list-item>
              <strong>Unix Milliseconds:</strong> Milliseconds since 1970-01-01 00:00:00 UTC (13 digits)
            </t-list-item>
            <t-list-item>
              <strong>iOS Timestamp:</strong> Seconds since 2001-01-01 00:00:00 UTC
            </t-list-item>
            <t-list-item>
              <strong>Chrome Timestamp:</strong> Microseconds since 1601-01-01 00:00:00 UTC
            </t-list-item>
            <t-list-item>
              <strong>Windows File Time:</strong> 100-nanosecond intervals since 1601-01-01 00:00:00 UTC
            </t-list-item>
            <t-list-item>
              <strong>Apple Timestamp:</strong> Special 9-digit format used in some Apple applications
            </t-list-item>
          </t-list>
        </t-card>
      </t-space>
    </t-card>
  </div>
</template>

<script setup>
import { ref, reactive, onMounted } from 'vue'
import { MessagePlugin } from 'tdesign-vue-next'

// Form data
const universalForm = reactive({
  timestamp: '',
  sourceTimezone: 'UTC',
  targetTimezone: 'Asia/Shanghai'
})

// Individual timestamp inputs
const iosTimestamp = ref('')
const chromeTimestamp = ref('')
const windowsTimestamp = ref('')
const appleTimestamp = ref('')

// Results
const universalResult = ref('')
const iosResult = ref('')
const chromeResult = ref('')
const windowsResult = ref('')
const appleResult = ref('')

// UI state
const activeTab = ref('ios')
const currentTimeData = ref([])

// Timezone options
const timezoneOptions = [
  { label: 'UTC', value: 'UTC' },
  { label: 'Asia/Shanghai (CST)', value: 'Asia/Shanghai' },
  { label: 'America/New_York (EST)', value: 'America/New_York' },
  { label: 'Europe/London (GMT)', value: 'Europe/London' },
  { label: 'Asia/Tokyo (JST)', value: 'Asia/Tokyo' },
  { label: 'America/Los_Angeles (PST)', value: 'America/Los_Angeles' }
]

// Backend functions (would be imported from Wails generated code)
const { 
  DefaultTimestampToDatetime,
  IosTimestampToDatetime,
  ChromeTimestampToDatetime,
  WindowsFileTimeToDatetime,
  AppleTimestampToDatetime
} = window.go?.timestamp || {}

// Convert universal timestamp
const convertUniversalTimestamp = async () => {
  if (!universalForm.timestamp) {
    MessagePlugin.warning('Please enter a timestamp')
    return
  }
  
  try {
    const timestamp = parseInt(universalForm.timestamp)
    const result = DefaultTimestampToDatetime ? 
      await DefaultTimestampToDatetime(timestamp, universalForm.sourceTimezone, universalForm.targetTimezone) :
      new Date(timestamp * 1000).toLocaleString()
    
    universalResult.value = result
    MessagePlugin.success('Timestamp converted successfully')
  } catch (error) {
    MessagePlugin.error('Conversion failed: ' + error.message)
  }
}

// Convert iOS timestamp
const convertIosTimestamp = async () => {
  if (!iosTimestamp.value) {
    MessagePlugin.warning('Please enter an iOS timestamp')
    return
  }
  
  try {
    const timestamp = parseFloat(iosTimestamp.value)
    const result = IosTimestampToDatetime ? 
      await IosTimestampToDatetime(timestamp, 'UTC', 'Asia/Shanghai') :
      new Date((timestamp + 978307200) * 1000).toLocaleString()
    
    iosResult.value = result
    MessagePlugin.success('iOS timestamp converted successfully')
  } catch (error) {
    MessagePlugin.error('Conversion failed: ' + error.message)
  }
}

// Convert Chrome timestamp
const convertChromeTimestamp = async () => {
  if (!chromeTimestamp.value) {
    MessagePlugin.warning('Please enter a Chrome timestamp')
    return
  }
  
  try {
    const timestamp = parseInt(chromeTimestamp.value)
    const result = ChromeTimestampToDatetime ? 
      await ChromeTimestampToDatetime(timestamp, 'UTC', 'Asia/Shanghai') :
      new Date((timestamp / 1000) - 11644473600000).toLocaleString()
    
    chromeResult.value = result
    MessagePlugin.success('Chrome timestamp converted successfully')
  } catch (error) {
    MessagePlugin.error('Conversion failed: ' + error.message)
  }
}

// Convert Windows timestamp
const convertWindowsTimestamp = async () => {
  if (!windowsTimestamp.value) {
    MessagePlugin.warning('Please enter a Windows file time')
    return
  }
  
  try {
    const timestamp = parseInt(windowsTimestamp.value)
    const result = WindowsFileTimeToDatetime ? 
      await WindowsFileTimeToDatetime(timestamp, 'UTC', 'Asia/Shanghai') :
      new Date((timestamp / 10000) - 11644473600000).toLocaleString()
    
    windowsResult.value = result
    MessagePlugin.success('Windows file time converted successfully')
  } catch (error) {
    MessagePlugin.error('Conversion failed: ' + error.message)
  }
}

// Convert Apple timestamp
const convertAppleTimestamp = async () => {
  if (!appleTimestamp.value) {
    MessagePlugin.warning('Please enter an Apple timestamp')
    return
  }
  
  try {
    const timestamp = parseFloat(appleTimestamp.value)
    const result = AppleTimestampToDatetime ? 
      await AppleTimestampToDatetime(timestamp, 'UTC', 'Asia/Shanghai') :
      new Date((timestamp + 978307200) * 1000).toLocaleString()
    
    appleResult.value = result
    MessagePlugin.success('Apple timestamp converted successfully')
  } catch (error) {
    MessagePlugin.error('Conversion failed: ' + error.message)
  }
}

// Update current time information
const updateCurrentTime = () => {
  const now = new Date()
  const unixSeconds = Math.floor(now.getTime() / 1000)
  const unixMilliseconds = now.getTime()
  const iosTime = unixSeconds - 978307200 // Seconds since 2001-01-01
  const chromeTime = (unixMilliseconds + 11644473600000) * 1000 // Microseconds since 1601-01-01
  const windowsTime = (unixMilliseconds + 11644473600000) * 10000 // 100-nanosecond intervals since 1601-01-01
  
  currentTimeData.value = [
    { label: 'Current Date/Time', value: now.toString() },
    { label: 'Unix Timestamp (seconds)', value: unixSeconds.toString() },
    { label: 'Unix Timestamp (milliseconds)', value: unixMilliseconds.toString() },
    { label: 'iOS Timestamp', value: iosTime.toString() },
    { label: 'Chrome Timestamp', value: chromeTime.toString() },
    { label: 'Windows File Time', value: windowsTime.toString() }
  ]
}

// Copy to clipboard
const copyToClipboard = async (text) => {
  try {
    await navigator.clipboard.writeText(text)
    MessagePlugin.success('Copied to clipboard')
  } catch (error) {
    MessagePlugin.error('Failed to copy to clipboard')
  }
}

// Initialize current time on mount
onMounted(() => {
  updateCurrentTime()
})
</script>

<style scoped>
.timestamp-parser {
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
