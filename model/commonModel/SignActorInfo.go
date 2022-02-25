package commonModel

type SignActorInfo struct {
	SignActor *Actor `json:"signActor,omitempty"`
	BlockHere bool  `json:"blockHere"`
	RequestVerifyFree bool `json:"requestVerifyFree"`
	VerifyMethods []string  `json:"verifyMethods,omitempty"`
}
