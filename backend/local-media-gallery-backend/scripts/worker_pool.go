package scripts

import (
	"errors"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm/clause"
	"local-media-gallery/database"
	"local-media-gallery/models"
	"log"
	"os"
	"path/filepath"
	"sync"
)

type Task struct {
	Directory string
}

func (t *Task) Execute() error {
	currentDirFullPath := t.Directory

	log.Printf("Processing Directory %s \n", currentDirFullPath)

	entries, err := os.ReadDir(currentDirFullPath)
	if err != nil {
		return fmt.Errorf("error reading content of Directory %s: %w", currentDirFullPath, err)
	}

	// Process the current Directory
	currentDirectory := &models.Directory{
		Name: filepath.Base(currentDirFullPath),
		Path: currentDirFullPath,
		Url:  RemovePathPrefix(currentDirFullPath),
	}

	err = database.DB.Where("path = ?", currentDirFullPath).FirstOrCreate(currentDirectory).Error
	if err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) {
			if mysqlErr.Number == 1062 {
				log.Printf("Warning: Duplicate entry for directory path '%s'. It likely already exists. Error: %v\n", currentDirFullPath, err)
				return nil
			}
		}
		return fmt.Errorf("error FirstOrCreate current directory %s: %w", currentDirFullPath, err)
	}

	// Categorize Child Entries
	var subdirectoriesDirEntry []os.DirEntry
	var videosDirEntry []os.DirEntry
	var imagesDirEntry []os.DirEntry

	// Process each subdirectory, image, and video in the Directory
	for _, entry := range entries {
		// Skip over hidden files
		if entry.Name()[0] == '.' {
			continue
		}

		fullPath := filepath.Join(currentDirFullPath, entry.Name())

		if entry.IsDir() {
			subdirectoriesDirEntry = append(subdirectoriesDirEntry, entry)
		} else if !entry.Type().IsRegular() {
			fmt.Printf("Skipping non-regular file %s \n", fullPath)
		} else {
			fileExtInLowercase := GetFileExtInLowercase(entry)
			switch fileExtInLowercase {
			case ".jpg", ".jpeg", ".png", ".gif", ".webp":
				imagesDirEntry = append(imagesDirEntry, entry)
			case ".mp4", ".m4v", ".mov", ".webm":
				videosDirEntry = append(videosDirEntry, entry)
			default:
				log.Printf("Unknown file extension: %s \n", fullPath)
			}
		}
	}

	// Prepare Maps for Direct Children and Videos
	directoriesMap := make(map[string]*models.Directory)
	videosMap := make(map[string]*models.Video)

	// Populate directoriesMap with child directories, setting their ParentDirectoryId
	for _, dirEntry := range subdirectoriesDirEntry {
		name := dirEntry.Name()
		fullPath := GetDirEntryFullPath(currentDirFullPath, &dirEntry)

		directoriesMap[name] = &models.Directory{
			Name:              name,
			Path:              fullPath,
			Url:               RemovePathPrefix(fullPath),
			ParentDirectoryId: &currentDirectory.Id,
		}
	}

	// Populate videosMap with child videos
	for _, videoEntry := range videosDirEntry {
		name := videoEntry.Name()
		baseName := GetFileBasename(&videoEntry)
		fullPath := GetDirEntryFullPath(currentDirFullPath, &videoEntry)

		videosMap[baseName] = &models.Video{
			Name:        name,
			Path:        fullPath,
			Url:         GetStaticResourceUrl(fullPath),
			DirectoryId: currentDirectory.Id,
		}
	}

	// Process images (potential thumbnails or standalone images)
	var images []*models.Image

	for _, imageEntry := range imagesDirEntry {
		name := imageEntry.Name()
		basename := GetFileBasename(&imageEntry)
		fullPath := GetDirEntryFullPath(currentDirFullPath, &imageEntry)

		width, height, err := GetImageDimensions(fullPath)
		if err != nil {
			return fmt.Errorf("error getting image dimensions: %s \n %v \n", fullPath, err)
		}

		i := models.Image{
			Name:        name,
			Path:        fullPath,
			Url:         GetStaticResourceUrl(fullPath),
			Width:       width,
			Height:      height,
			DirectoryId: currentDirectory.Id,
		}

		if directory, ok := directoriesMap[basename]; ok {
			directory.Thumbnail = &models.Thumbnail{
				Image: i,
			}
			continue
		} else if video, ok := videosMap[basename]; ok {
			video.Thumbnail = &models.Thumbnail{
				Image: i,
			}
		} else {
			images = append(images, &i)
		}
	}

	if len(directoriesMap) != 0 {
		for _, directory := range directoriesMap {
			err = database.DB.Clauses(clause.OnConflict{
				Columns:   []clause.Column{{Name: "path"}},
				DoUpdates: clause.AssignmentColumns([]string{"parent_directory_id", "thumbnail_id"}),
			}).Create(&directory).Error
			if err != nil {
				return fmt.Errorf("error creating/updating child directories for %s: %w", currentDirFullPath, err)
			}
		}
	}

	if len(videosMap) != 0 {
		var videos []*models.Video
		for _, video := range videosMap {
			videos = append(videos, video)
		}
		err = database.DB.Create(&videos).Error
		if err != nil {
			return fmt.Errorf("error creating/updating videos for %s: %w", currentDirFullPath, err)
		}
	}

	if len(images) != 0 {
		err = database.DB.Create(&images).Error
		if err != nil {
			return fmt.Errorf("error creating/updating images for %s: %w", currentDirFullPath, err)
		}
	}

	return nil
}

type WorkerPool struct {
	Tasks             []*Task
	ConcurrencyDegree int
	tasksChan         chan *Task
	waitGroup         sync.WaitGroup
}

func (workerPool *WorkerPool) worker() {
	for task := range workerPool.tasksChan {
		err := task.Execute()
		if err != nil {
			fmt.Printf("Error processing Directory: %s \n %v \n", task.Directory, err.Error())
		}
		workerPool.waitGroup.Done()
	}
}

func (workerPool *WorkerPool) Run() {
	workerPool.tasksChan = make(chan *Task, len(workerPool.Tasks))

	for i := 0; i < workerPool.ConcurrencyDegree; i++ {
		go workerPool.worker()
	}

	workerPool.waitGroup.Add(len(workerPool.Tasks))
	for _, task := range workerPool.Tasks {
		workerPool.tasksChan <- task
	}
	close(workerPool.tasksChan)

	workerPool.waitGroup.Wait()
}
