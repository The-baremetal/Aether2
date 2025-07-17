package buildcache

import (
  "crypto/sha256"
  "encoding/json"
  "io"
  "io/ioutil"
  "os"
  "fmt"
)

type BuildCacheEntry struct {
  Hash         string            `json:"hash"`
  Output       string            `json:"output"`
  Deps         []string          `json:"deps"`
  DepHashes    map[string]string `json:"dep_hashes"`
  LastBuild    int64             `json:"last_build"`
}

type BuildCache struct {
  Files map[string]BuildCacheEntry `json:"files"`
}

func LoadCache(path string) (*BuildCache, error) {
  data, err := ioutil.ReadFile(path)
  if err != nil {
    return &BuildCache{Files: make(map[string]BuildCacheEntry)}, nil
  }
  var cache BuildCache
  if err := json.Unmarshal(data, &cache); err != nil {
    return &BuildCache{Files: make(map[string]BuildCacheEntry)}, nil
  }
  return &cache, nil
}

func SaveCache(path string, cache *BuildCache) error {
  data, err := json.MarshalIndent(cache, "", "  ")
  if err != nil {
    return err
  }
  return ioutil.WriteFile(path, data, 0644)
}

func FileHash(path string) (string, error) {
  f, err := os.Open(path)
  if err != nil {
    return "", err
  }
  h := sha256.New()
  if _, err := io.Copy(h, f); err != nil {
    f.Close()
    return "", err
  }
  f.Close()
  return fmt.Sprintf("%x", h.Sum(nil)), nil
} 