package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var CleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "Clean up build/ and bin/ directories",
	Run: func(cmd *cobra.Command, args []string) {
		doClean()
	},
}

func doClean() {
	fmt.Println("Cleaning build/ and bin/ directories...")
	os.RemoveAll("build")
	os.RemoveAll("bin")
	fmt.Println("Clean complete!")
}
