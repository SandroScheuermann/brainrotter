package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

const searchURL = "https://www.googleapis.com/youtube/v3/search?"
const apiKey = ""

func GetVideoPairInfo(ch chan<- []VideoInfoPair) {

	contentParams := url.Values{}

	contentParams.Add("part", "snippet")                  // Para forçar a resposta com o "snippet" incluído.
	contentParams.Add("videoDuration", "short")           // Obter somente vídeos curtos
	contentParams.Add("maxResults", "50")                 // Obter no máximo 50 vídeos
	contentParams.Add("type", "video")                    // Obter somente vídeos (sem info de playlist, etc)
	contentParams.Add("location", "-14.235004,-51.92528") // Coordenadas centrais do Brasil
	contentParams.Add("locationRadius", "500km")          // Raio de 500 km
	contentParams.Add("order", "rating")                  // Relevance é default, testar "rating" depois
	contentParams.Add("q", "fatos de jogos|fatos de filmes|fatos de ciencia|#shorts")
	contentParams.Add("key", apiKey) // Ajustar API key para env dps

	ytContentResp, err := search(contentParams)

	if err != nil {
		fmt.Println("Erro ao buscar informação bruta de vídeos content")
		ch <- []VideoInfoPair{}
	}

	magnetParams := url.Values{}

	magnetParams.Add("part", "snippet")  // Para forçar a resposta com o "snippet" incluído.
	magnetParams.Add("maxResults", "10") // Obter no máximo 10 vídeos
	magnetParams.Add("type", "video")    // Obter somente vídeos (sem info de playlist, etc)
	magnetParams.Add("order", "relevance")
	magnetParams.Add("q", "gta v parkour gameplay no copyright")
	magnetParams.Add("key", apiKey) // Ajustar API key para env dps

	ytMagnetResp, err := search(magnetParams)

	if err != nil {
		fmt.Println("Erro ao buscar informação bruta de vídeos content")
		ch <- []VideoInfoPair{}
	}

	videoPairSlice := []VideoInfoPair{}

	j := 0

	for i, contentVid := range ytContentResp.Items {

		if i%5 == 0 && i != 0 {
			j++
		}

		contentVideoInfo := RelevantVideoInfo{
			VideoID:      contentVid.ID.VideoID,
			Title:        contentVid.Snippet.Title,
			ThumbnailURL: contentVid.Snippet.Thumbnails.Default.URL,
		}

		magnetVideoInfo := RelevantVideoInfo{
			VideoID:      ytMagnetResp.Items[j].ID.VideoID,
			Title:        ytMagnetResp.Items[j].Snippet.Title,
			ThumbnailURL: ytMagnetResp.Items[j].Snippet.Thumbnails.Default.URL,
		}

		videoPair := NewVideoInfoPair(contentVideoInfo, magnetVideoInfo)

		videoPairSlice = append(videoPairSlice, videoPair)
	}

	ch <- videoPairSlice
}

func search(params url.Values) (YouTubeResponse, error) {

	resp, err := http.Get(searchURL + params.Encode())

	if err != nil {
		fmt.Println("Erro realizando a request", err)
		return YouTubeResponse{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Retorno de erro da API", resp.Status, searchURL+params.Encode())
		return YouTubeResponse{}, err
	}

	var ytResp YouTubeResponse

	if err := json.NewDecoder(resp.Body).Decode(&ytResp); err != nil {
		fmt.Println("Erro realizando o decode da resposta da API", err)
	}

	return ytResp, nil
}
