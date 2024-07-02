package main

type YouTubeResponse struct {
	Items []struct {
		ID struct {
			VideoID string `json:"videoId"`
		} `json:"id"`
		Snippet struct {
			Title       string `json:"title"`
			Description string `json:"description"`
			Thumbnails  struct {
				Default struct {
					URL string `json:"url"`
				} `json:"default"`
			} `json:"thumbnails"`
		} `json:"snippet"`
	} `json:"items"`
}

type RelevantVideoInfo struct {
	VideoID      string
	Title        string
	ThumbnailURL string
}

type VideoInfoPair struct {
	ContentVideo RelevantVideoInfo
	MagnetVideo  RelevantVideoInfo
}

func NewVideoInfoPair(contentVideo, magnetVideo RelevantVideoInfo) VideoInfoPair {
	return VideoInfoPair{
		ContentVideo: contentVideo,
		MagnetVideo:  magnetVideo,
	}
}
