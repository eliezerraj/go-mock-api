package model

type ManagerInfo struct {
	App 		*ManagerInfoApp `json:"app"`
	Server     	Server     		`json:"servers"`
	Setup		Setup			`json:"setup_behaviour"`
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

type Server struct {
	Port 			int `json:"port"`
	ReadTimeout		int `json:"readTimeout"`
	WriteTimeout	int `json:"writeTimeout"`
	IdleTimeout		int `json:"idleTimeout"`
	CtxTimeout		int `json:"ctxTimeout"`
}

type Setup struct {
    ResponseTime 		int `json:"response_time"`
    ResponseStatusCode  int `json:"response_status_code"`
	IsRandomTime		bool `json:"is_random_time"`
	Count				int `json:"count"`
}