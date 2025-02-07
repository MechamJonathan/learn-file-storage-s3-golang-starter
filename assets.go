package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func (cfg apiConfig) ensureAssetsDir() error {
	//create main directory
	if _, err := os.Stat(cfg.assetsRoot); os.IsNotExist(err) {
		return os.Mkdir(cfg.assetsRoot, 0755)
	}

	// Create thumbnails subdirectory
	thumbnailsPath := filepath.Join(cfg.assetsRoot, "thumbnails")
	if _, err := os.Stat(thumbnailsPath); os.IsNotExist(err) {
		if err := os.Mkdir(thumbnailsPath, 0755); err != nil {
			return err
		}
	}
	return nil
}

func getAssetPath(filename string) string {
	return "thumbnails/" + filename
}

func (cfg apiConfig) getAssetDiskPath(assetPath string) string {
	return filepath.Join(cfg.assetsRoot, assetPath)
}

func (cfg apiConfig) getAssetURL(assetPath string) string {
	return fmt.Sprintf("http://localhost:%s/assets/%s", cfg.port, assetPath)
}

func mediaTypeToExt(mediaType string) string {
	parts := strings.Split(mediaType, "/")
	if len(parts) != 2 {
		return ".bin"
	}
	return "." + parts[1]
}
