package commands

// import (
// 	"encoding/json"
// 	"fmt"
// 	"io/ioutil"
// 	"log"
// 	"os"
// 	"path/filepath"

// )

// // "github.com/NjiruClinton/tectonic_assets/assets"
// func BrowseAssets() {

// 	// Parse pubspec.yaml
// 	pubspec, err := ParsePubspec()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// Get asset directories from pubspec
// 	assetDirs := pubspec.Flutter.Assets

// 	// Walk asset directories and collect files
// 	files := make([]string, 0)
// 	for _, dir := range assetDirs {
// 		err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
// 			files = append(files, path)
// 			return nil
// 		})
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 	}

// 	// Display files
// 	for _, file := range files {
// 		fmt.Println(file)
// 	}

// }

// // func OptimizeAssets() {

// // 	// Parse pubspec.yaml
// // 	pubspec, err := ParsePubspec()
// // 	if err != nil {
// // 		log.Fatal(err)
// // 	}

// 	// Get asset directories from pubspec
// 	// assetDirs := pubspec.Flutter.Assets

// 	// Walk asset directories and collect files
// 	// files := make([]string, 0)
// 	// for _, dir := range assetDirs {
// 	// 	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
// 	// 		files = append(files, path)
// 	// 		return nil
// 	// 	})
// 	// 	if err != nil {
// 	// 		log.Fatal(err)
// 	// 	}
// 	// }

// 	// Optimize files
// 	// for _, file := range files {
// 	// 	fmt.Println(file)
// 	// }

// // }

// // ParsePubspec parses the pubspec.yaml
// // func ParsePubspec() (assets.Pubspec, error) {
// // 	data, err := ioutil.ReadFile("pubspec.yaml")
// // 	if err != nil {
// // 		return assets.Pubspec{}, err
// // 	}

// // 	var pubspec assets.Pubspec
// // 	if err := json.Unmarshal(data, &pubspec); err != nil {
// // 		return assets.Pubspec{}, err
// // 	}

// // 	return pubspec, nil
// // }
