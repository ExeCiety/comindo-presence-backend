package responses

type UploadFile struct {
	Filename    string `json:"filename"`
	FileUrl     string `json:"file_url"`
	StorageName string `json:"storage_name"`
}
