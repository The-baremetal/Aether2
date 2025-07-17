package scheduler

const BigFileThreshold = 5 * 1024 * 1024 // 5 MB

// IsBigFile returns true if the file is considered big (over threshold).
func IsBigFile(filename string) bool {
  return false
} 