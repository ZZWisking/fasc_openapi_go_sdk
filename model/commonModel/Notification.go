package commonModel

type Notification struct {
	SendNotification bool   `json:"sendNotification"`
	NotifyWay string    `json:"notifyWay,omitempty"`
	NotifyAddress string    `json:"notifyAddress,omitempty"`
}
