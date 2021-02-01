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

// TODO: provide method for status mapping
type OperatingStatus string

const (
	// value 0
	Off OperatingStatus = "off"
	// value 1
	Standby OperatingStatus = "standby"
	// value 2
	Starting OperatingStatus = "starting"
	// value 3
	FeedInMPP OperatingStatus = "Feed in (MPP)"
	// value 4
	FeedInLimited OperatingStatus = "Feed in limited"
	// value 5
	FeedIn OperatingStatus = "Feed In"
	// value 6
	DcPowerLow OperatingStatus = "DC voltage low"
	// value 7
	InsulationMeasurement OperatingStatus = "Insulation measurement"
	// value 8
	WaitingTimeActive OperatingStatus = "Waiting time active!"
)
