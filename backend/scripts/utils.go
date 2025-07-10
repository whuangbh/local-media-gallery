package scripts

import (
	"errors"
	"fmt"
	"image"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

func GetImageDimensions(imagePath string) (int, int, error) {
	file, err := os.Open(imagePath)
	if err != nil {
		return 0, 0, fmt.Errorf("could not open image file: %w", err)
	}
	defer file.Close()

	// DecodeConfig reads only the image header to get dimensions
	config, _, err := image.DecodeConfig(file)
	if err != nil {
		// Check for specific error if it's not a recognized format
		if errors.Is(err, image.ErrFormat) {
			return 0, 0, fmt.Errorf("unrecognized image format for %s: %w", imagePath, err)
		}
		return 0, 0, fmt.Errorf("could not decode image config for %s: %w", imagePath, err)
	}

	return config.Width, config.Height, nil
}

func GetFileBasename(dirEntry *os.DirEntry) string {
	name := (*dirEntry).Name()
	extension := filepath.Ext(name)

	return strings.TrimSuffix(name, extension)
}

func GetDirEntryFullPath(parentDirPath string, dirEntry *os.DirEntry) string {
	return filepath.Join(parentDirPath, (*dirEntry).Name())
}

func GetFileExtInLowercase(entry os.DirEntry) string {
	return strings.ToLower(filepath.Ext(entry.Name()))
}

func GetStaticResourceUrl(fullPath string) string {
	// TODO
	return "/static" + RemovePathPrefix(fullPath)
}

func RemovePathPrefix(fullPath string) string {
	prefix := "/static_data"

	processedStr := strings.Replace(fullPath, prefix, "", 1)
	if len(processedStr) == 0 {
		return "/"
	} else {
		return processedStr
	}
}

func GetAllSubdirectories(roots []string) ([]string, error) {
	var result []string

	for _, root := range roots {
		err := filepath.WalkDir(root, func(path string, file fs.DirEntry, err error) error {
			if err != nil {
				return err
			}

			if file.IsDir() {
				result = append(result, path)
			}

			return nil
		})

		if err != nil {
			return nil, err
		}
	}

	return result, nil
}
