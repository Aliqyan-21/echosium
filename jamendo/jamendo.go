package jamendo

type Track struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
	TrackUrl      string `json:"song"`
	Artist        string `json:"artist"`
	Album         string `json:"album"`
	Image         string `json:"image"`
	AudioFormat   string `json:"audioformat"`
	AudioDownload bool   `json:"audiodownload"`
}

func getTracks() {

}
