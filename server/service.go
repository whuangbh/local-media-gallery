package server

import (
	"local-media-gallery/database"
	"local-media-gallery/models"
)

func QueryFolderInfo(parentDirectoryUrl string) (
	*models.Directory,
	[]*models.Directory,
	[]*models.Video,
	[]*models.Image,
	error,
) {
	parentDirectory := models.GetDirectoryByUrl(database.DB, parentDirectoryUrl)

	directories, err := models.GetDirectoriesByParentDirId(database.DB, &parentDirectory.Id)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	videos, err := models.GetVideosByParentDirId(database.DB, &parentDirectory.Id)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	images, err := models.GetImagesByParentDirId(database.DB, &parentDirectory.Id)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	return parentDirectory, directories, videos, images, nil
}

func SearchMediaByKeyword(query string) (
	[]*models.Video,
	[]*models.Image,
	error,
) {

	videos, err := models.SearchVideosByName(database.DB, query)
	if err != nil {
		return nil, nil, err
	}

	images, err := models.SearchImagesByName(database.DB, query)
	if err != nil {
		return nil, nil, err
	}

	return videos, images, nil
}
