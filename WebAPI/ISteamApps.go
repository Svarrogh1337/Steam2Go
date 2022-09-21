package Steam2Go

type GetAppListResponse struct {
	Applist Applist `json:"applist"`
}

type Applist struct {
	Apps App `json:"apps"`
}

type Apps struct {
	App []App `json:"app"`
}

type App struct {
	Appid int    `json:"appid"`
	Name  string `json:"name"`
}
