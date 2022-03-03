package responseModel

type DownloadFilesRes struct {
	RequestId   string
	Content     []byte
	ContentType string
	Data        string
}
