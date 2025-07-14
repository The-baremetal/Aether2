package analysis

import (
	"os"
	"path/filepath"
)

// ResolveStdlibImport tries to resolve a stdlib package from AETHERROOT.
// Returns the resolved path and true if found, or "" and false if not found.
func ResolveStdlibImport(pkgName string) (string, bool) {
	rootPkgs := GetAetherRootPackagesDir()
	if rootPkgs == "" {
		return "", false
	}
	// Assume stdlib packages are in $AETHERROOT/packages/pkgName/
	pkgPath := filepath.Join(rootPkgs, pkgName)
	if stat, err := os.Stat(pkgPath); err == nil && stat.IsDir() {
		return pkgPath, true
	}
	// Also check for a single-file package: $AETHERROOT/packages/pkgName.ae or .aeth
	for _, ext := range []string{".ae", ".aeth"} {
		filePath := filepath.Join(rootPkgs, pkgName+ext)
		if stat, err := os.Stat(filePath); err == nil && !stat.IsDir() {
			return filePath, true
		}
	}
	return "", false
}

// (Optional) ListStdlibPackages returns a list of available stdlib packages in AETHERROOT.
func ListStdlibPackages() ([]string, error) {
	rootPkgs := GetAetherRootPackagesDir()
	if rootPkgs == "" {
		return nil, nil
	}
	dirs, err := os.ReadDir(rootPkgs)
	if err != nil {
		return nil, err
	}
	var pkgs []string
	for _, entry := range dirs {
		if entry.IsDir() {
			pkgs = append(pkgs, entry.Name())
		} else if name := entry.Name(); len(name) > 3 && (name[len(name)-3:] == ".ae" || name[len(name)-5:] == ".aeth") {
			pkgs = append(pkgs, name[:len(name)-3])
		}
	}
	return pkgs, nil
} 