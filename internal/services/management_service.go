package services

import (
	"time"

	"github.com/ricochet2200/go-disk-usage/du"

	"github.com/go-mock-api/internal/utils/constants"
	"github.com/go-mock-api/internal/core/model"
	"github.com/go-mock-api/internal/utils/loggers"

)

var KB = uint64(1024)

type HealthLog struct {
	Pattern    string `json:"pattern"`
	Duration   int64  `json:"duration"`
	StatusCode int    `json:"statusCode"`
}

func CheckHealth() model.ManagerHealth {
	loggers.GetLogger().Named(constants.Service).Info("CheckHealth")
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
	loggers.GetLogger().Named(constants.Service).Info("healthDB")
	statusDB := constants.NO_DEPLOY

	m := model.ManagerHealthDB{
		Status: statusDB,
	}

	return m, true
}

func healthDiskUsage() (model.ManagerHealthDiskSpace, bool) {
	loggers.GetLogger().Named(constants.Service).Info("healthDiskUsage")
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

func CpuStress(count int) string {
	loggers.GetLogger().Named(constants.Service).Info("CpuStress") 
	start := time.Now()

	for n := 0; n <= count; n++ {
		f := make([]int, count+1, count+2)
		if count < 2 {
			f = f[0:2]
		}
		f[0] = 0
		f[1] = 1
		for i := 2; i <= count; i++ {
			f[i] = f[i-1] + f[i-2]
		}
    }

	t := time.Now()
	elapsed := t.Sub(start)

	return "Done in " + elapsed.String()
}