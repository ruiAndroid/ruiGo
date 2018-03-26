package model

type UserBugTackerInfo struct {
	UserId string `json:"user_id"`
	DownloadTime string `json:"download_time"`
	PhoneModel string `json:"phone_model"`
	SdCardMemory string `json:"sd_card_memory"`
	OriginalZipSize string `json:"original_zip_size"`
	WordId []string `json:"word_id"`
	AudioFileCount string `json:"audio_file_count"`
	PicFileCount string `json:"pic_file_count"`
	BugWordId string `json:"bug_word_id"`

}
