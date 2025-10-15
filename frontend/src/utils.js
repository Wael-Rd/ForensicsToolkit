// Utility functions for the frontend

/**
 * Copy text to clipboard
 * @param {string} text - Text to copy
 * @returns {Promise<boolean>} - Success status
 */
export async function copyToClipboard(text) {
  try {
    await navigator.clipboard.writeText(text)
    return true
  } catch (error) {
    console.error('Failed to copy to clipboard:', error)
    return false
  }
}

/**
 * Format file size in human readable format
 * @param {number} bytes - File size in bytes
 * @returns {string} - Formatted file size
 */
export function formatFileSize(bytes) {
  if (bytes === 0) return '0 Bytes'
  
  const k = 1024
  const sizes = ['Bytes', 'KB', 'MB', 'GB', 'TB']
  const i = Math.floor(Math.log(bytes) / Math.log(k))
  
  return parseFloat((bytes / Math.pow(k, i)).toFixed(2)) + ' ' + sizes[i]
}

/**
 * Validate file extension
 * @param {string} filename - Name of the file
 * @param {string[]} allowedExtensions - Array of allowed extensions
 * @returns {boolean} - Whether file extension is valid
 */
export function validateFileExtension(filename, allowedExtensions) {
  const extension = filename.split('.').pop().toLowerCase()
  return allowedExtensions.includes(extension)
}

/**
 * Format timestamp to readable date
 * @param {number} timestamp - Unix timestamp
 * @returns {string} - Formatted date string
 */
export function formatTimestamp(timestamp) {
  return new Date(timestamp * 1000).toLocaleString()
}

/**
 * Debounce function calls
 * @param {Function} func - Function to debounce
 * @param {number} wait - Wait time in milliseconds
 * @returns {Function} - Debounced function
 */
export function debounce(func, wait) {
  let timeout
  return function executedFunction(...args) {
    const later = () => {
      clearTimeout(timeout)
      func(...args)
    }
    clearTimeout(timeout)
    timeout = setTimeout(later, wait)
  }
}

/**
 * Generate random ID
 * @returns {string} - Random ID string
 */
export function generateId() {
  return Math.random().toString(36).substr(2, 9)
}

/**
 * Validate hex string
 * @param {string} hex - Hex string to validate
 * @returns {boolean} - Whether string is valid hex
 */
export function isValidHex(hex) {
  return /^[0-9A-Fa-f]+$/.test(hex)
}

/**
 * Convert bytes to hex string
 * @param {Uint8Array} bytes - Byte array
 * @returns {string} - Hex string
 */
export function bytesToHex(bytes) {
  return Array.from(bytes, byte => byte.toString(16).padStart(2, '0')).join('')
}

/**
 * Convert hex string to bytes
 * @param {string} hex - Hex string
 * @returns {Uint8Array} - Byte array
 */
export function hexToBytes(hex) {
  const bytes = new Uint8Array(hex.length / 2)
  for (let i = 0; i < hex.length; i += 2) {
    bytes[i / 2] = parseInt(hex.substr(i, 2), 16)
  }
  return bytes
}
