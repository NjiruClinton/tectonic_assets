package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

// Asset configuration
type Assets struct {
	Path string
}

var assets Assets

var rootCmd = &cobra.Command{
	Use:   "flutter-asset",
	Short: "Manage Flutter assets",
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&assets.Path, "path", "p", "assets", "Asset folder path")
}


var browseCmd = &cobra.Command{
	Use:   "browse",
	Short: "Browse assets folder",
	Run: func(cmd *cobra.Command, args []string) {
		assets.Browse()
	},
}

var optimizeCmd = &cobra.Command{
	Use:   "optimize",
	Short: "Optimize image assets",
	Run: func(cmd *cobra.Command, args []string) {
		assets.OptimizeImages()
	},
}

var manifestCmd = &cobra.Command{
	Use:   "manifest",
	Short: "Generate asset manifest",
	Run: func(cmd *cobra.Command, args []string) {
		assets.GenerateManifest()
	},
}


func (a *Assets) Browse() {
	files, err := ioutil.ReadDir(a.Path)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Println(file.Name())
	}
}

func (a *Assets) OptimizeImages() {

	files, _ := ioutil.ReadDir(a.Path)

	for _, file := range files {

		if strings.HasSuffix(file.Name(), ".png") {

			input := filepath.Join(a.Path, file.Name())
			output := filepath.Join(a.Path, "optimized", file.Name())

			err := exec.Command("tinypng", input, output).Run()
			if err != nil {
				log.Fatal(err)
			}

		}

	}

}

func GetAssetHash(name string) string {
	// hash asset contents soon...
	return "hash"
}

func (a *Assets) GenerateManifest() {

	manifest := map[string]string{}

	var manifestFile *os.File
	manifestFile, err := os.Create("manifest.json")
	if err != nil {
		log.Fatal(err)
	}

	files, _ := ioutil.ReadDir(a.Path)

	for _, file := range files {
		name := file.Name()
		hash := GetAssetHash(name)
		manifest[name] = hash
	}

	json.NewEncoder(manifestFile).Encode(manifest)

	manifestFile.Close()

}

// Hooks
//TODO: dont forget about hooks

func PostBuildHook() {
	assets := Assets{Path: "build/assets"}
	assets.GenerateManifest()
}

func main() {

	rootCmd.AddCommand(browseCmd)
	rootCmd.AddCommand(optimizeCmd)
	rootCmd.AddCommand(manifestCmd)

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}

}
