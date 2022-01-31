package model

type ManagerInfo struct {
	App *ManagerInfoApp `json:"app"`
}

type ManagerInfoApp struct {
	Name 				string `json:"name"`
	Description 		string `json:"description"`
	Version 			string `json:"version"`
}

type ManagerHealthDiskSpace struct {
	Status    string `json:"status"`
	Total     uint64 `json:"total"`
	Free      uint64 `json:"free"`
	Threshold uint64 `json:"threshold"`
}

type ManagerHealth struct {
	Status    string                 `json:"status"`
	DB        ManagerHealthDB        `json:"db"`
	DiskSpace ManagerHealthDiskSpace `json:"diskSpace"`
}

type ManagerHealthDB struct {
	Status string `json:"status"`
}
