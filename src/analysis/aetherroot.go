package analysis

import (
	"os"
	"path/filepath"
)

// GetAetherRoot returns the path to the AETHERROOT directory, or an empty string if not set.
func GetAetherRoot() string {
	return os.Getenv("AETHERROOT")
}

// ExistsAetherRoot returns true if AETHERROOT is set and exists on disk.
func ExistsAetherRoot() bool {
	root := GetAetherRoot()
	if root == "" {
		return false
	}
	stat, err := os.Stat(root)
	return err == nil && stat.IsDir()
}

// GetAetherRootPackagesDir returns the path to the stdlib packages directory in AETHERROOT.
func GetAetherRootPackagesDir() string {
	root := GetAetherRoot()
	if root == "" {
		return ""
	}
	return filepath.Join(root, "packages")
} 