package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var (
	updateMirror    string
	downloadNightly bool
)

var UpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Check for Aether updates and manage versions/forks",
	Run: func(cmd *cobra.Command, args []string) {
		doUpdate()
	},
}

func init() {
	UpdateCmd.Flags().StringVar(&updateMirror, "mirror", "@https://github.com/The-baremetal/Aether2", "Set update mirror URL")
	UpdateCmd.Flags().BoolVar(&downloadNightly, "nightly", false, "Download latest nightly instead of stable")
}

func doUpdate() {
	currentVersion := getCurrentVersion()
	mirror := strings.TrimPrefix(updateMirror, "@")
	tag, assetURL := getLatestRelease(mirror, downloadNightly)
	if tag == "" || assetURL == "" {
		fmt.Println("Could not fetch latest release or asset from mirror:", mirror)
		return
	}
	fmt.Println("Current version:", currentVersion)
	fmt.Println("Latest release:", tag)
	if currentVersion != tag {
		fmt.Printf("Update available! Downloading %s with aria2...\n", tag)
		if downloadWithAria2(assetURL, "aether2-latest") {
			fmt.Println("Downloaded new version to aether2-latest")
			// Optionally, handle version switching here
		} else {
			fmt.Println("Failed to download new version with aria2.")
		}
	} else {
		fmt.Println("Aether is up to date!")
	}
}

func getCurrentVersion() string {
	verFile := filepath.Join(".aether_version")
	data, err := ioutil.ReadFile(verFile)
	if err != nil {
		return "unknown"
	}
	return strings.TrimSpace(string(data))
}

func getLatestRelease(mirror string, wantNightly bool) (string, string) {
	api := mirror + "/releases"
	resp, err := http.Get(api)
	if err != nil {
		return "", ""
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return "", ""
	}
	var releases []map[string]interface{}
	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(&releases); err != nil {
		return "", ""
	}
	var chosenTag, chosenAsset string
	for _, rel := range releases {
		tag, _ := rel["tag_name"].(string)
		assets, _ := rel["assets"].([]interface{})
		if len(assets) == 0 {
			continue
		}
		asset, _ := assets[0].(map[string]interface{})
		assetURL, _ := asset["browser_download_url"].(string)
		if wantNightly && strings.HasSuffix(tag, "-nightly") {
			chosenTag, chosenAsset = tag, assetURL
			break
		}
		if !wantNightly && strings.HasSuffix(tag, "-stable") {
			chosenTag, chosenAsset = tag, assetURL
			break
		}
	}
	return chosenTag, chosenAsset
}

func downloadWithAria2(url, out string) bool {
	cmd := exec.Command("aria2c", "-o", out, url)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run() == nil
}
