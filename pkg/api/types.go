package api

type DataPoint struct {
	ID           int
	FriendlyName string
	Unit         string
}

type DxsEntry struct {
	ID           int     `json:"dxsId"`
	Value        float32 `json:"value"`
	FriendlyName string
	Unit         string
}

type DxsSession struct {
	SessionID int `json:"sessionId"`
	RoleID    int `json:"roleId"`
}

type DxsStatus struct {
	Code int `json:"code"`
}

type DxsRespone struct {
	Entries []DxsEntry `json:"dxsEntries"`
	Session DxsSession `json:"session"`
	Status  DxsStatus  `json:"status"`
}
