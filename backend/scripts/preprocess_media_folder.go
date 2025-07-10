package scripts

import (
	"local-media-gallery/database"
	"log"
)

func PreprocessMediaFolder() {
	err := database.DropAllTables()
	if err != nil {
		log.Fatalf("error dropping all tables: \n %v \n", err)
	}

	err = database.PerformMigration()
	if err != nil {
		log.Fatalf("error performing migration: \n %v \n", err)
	}

	rootDirectories := []string{
		"/static_data",
	}

	// List all the subdirectories
	directoriesToProcess, err := GetAllSubdirectories(rootDirectories)
	if err != nil {
		log.Fatalf("error trying to get all subdirectories: \n %v \n", err)
	}

	// Wrap the directories into tasks
	tasks := make([]*Task, 0, len(directoriesToProcess))
	for _, parentDirFullPath := range directoriesToProcess {
		tasks = append(tasks, &Task{
			Directory: parentDirFullPath,
		})
	}

	wp := WorkerPool{
		Tasks:             tasks,
		ConcurrencyDegree: 4,
	}

	wp.Run()

	log.Println("All directories processed!")
}
