package services

import (
	"github.com/ricochet2200/go-disk-usage/du"

	"github.com/go-mock-api/internal/utils/constants"
	"github.com/go-mock-api/internal/core/model"

)

var KB = uint64(1024)

type HealthLog struct {
	Pattern    string `json:"pattern"`
	Duration   int64  `json:"duration"`
	StatusCode int    `json:"statusCode"`
}

func CheckHealth() model.ManagerHealth {
	hDB, upDB := healthDB()
	hDisk, upDisk := healthDiskUsage()
	status := constants.UP

	if !upDB || !upDisk {
		status = constants.DOWN
	}

	return model.ManagerHealth{
		Status:    status,
		DB:        hDB,
		DiskSpace: hDisk,
	}
}

func healthDB() (model.ManagerHealthDB, bool) {
	statusDB := constants.NO_DEPLOY

	m := model.ManagerHealthDB{
		Status: statusDB,
	}

	return m, true
}

func healthDiskUsage() (model.ManagerHealthDiskSpace, bool) {
	usage := du.NewDiskUsage("/")
	statusDisk := constants.UP

	if (usage.Usage() * 100) > 90 {
		statusDisk = constants.DOWN
	}

	m := model.ManagerHealthDiskSpace{
		Free:   usage.Free() / (KB * KB),
		Total:  usage.Size() / (KB * KB),
		Status: statusDisk,
	}

	return m, statusDisk == constants.UP
}
