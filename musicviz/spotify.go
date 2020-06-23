package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type AudioFeatures struct {
	Duration_ms      int64   `json:"duration_ms"`
	Key              int64   `json:"key"`
	Mode             int64   `json:"mode"`
	Time_Signature   int64   `json:"time_signature"`
	Acousticness     float64 `json:"acousticness"`
	Danceability     float64 `json:"danceability"`
	Energy           float64 `json:"energy"`
	Instrumentalness float64 `json:"instrumentalness"`
	Liveness         float64 `json:"liveness"`
	Loudness         float64 `json:"loudness"`
	Speechiness      float64 `json:"speechiness"`
	Valence          float64 `json:"valence"`
	Tempo            float64 `json:"tempo"`
	ID               string  `json:"id"`
	Uri              string  `json:"uri"`
	Track_href       string  `json:"track_href"`
	analysis_url     string  `json:"analysis_url"`
	Type             string  `json:"type"`
}
type AuthToken struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
	Scope       string `json:"scope"`
}
type AudioAnalysisResponse struct {
	Bars     []TimeInterval `json:"bars"`
	Beats    []TimeInterval `json:"beats"`
	Sections []Section      `json:"sections"`
	Segments []Segments     `json:"segments"`
	Tatums   []TimeInterval `json:"tatums"`
}
type Section struct {
	Start                    float64 `json:"start"`
	Duration                 float64 `json:"duration"`
	Loudness                 float64 `json:"loudness"`
	Tempo                    float64 `json:"tempo"`
	Tempo_Confidence         float64 `json:"tempo_confidence"`
	Key                      int     `json:"key"`
	Key_Confidence           float64 `json:"key_confidence"`
	Mode                     int64   `json:"mode"`
	Mode_Confidence          float64 `json:"mode_confidence"`
	Time_Signature           int64   `json:"time_signature"`
	Time_Signature_Confdence float64 `json:"time_signature_confidence"`
}
type TimeInterval struct {
	Start      float64 `json:"start"`
	Duration   float64 `json:"duration"`
	Confidence float64 `json:"confidence"`
}
type Segments struct {
	Start             float64   `json:"start"`
	Duration          float64   `json:"duration"`
	Confidence        float64   `json:"confidence"`
	Loudness_Start    float64   `json:"loudness_start"`
	Loudness_Max      float64   `json:"loudness_max"`
	Loudness_Max_Time float64   `json:"loudness_max_time"`
	Loudness_End      float64   `json:"loudness_end"`
	Pitches           []float64 `json:"pitches"`
	Timbre            []float64 `json:"timbre`
}

const client_id string = "b4da79f085be478fabb6f49bd562d9b4"

func authenticate() AuthToken {
	client := &http.Client{}
	data := url.Values{}
	data.Set("grant_type", "client_credentials")
	base64Id := base64.StdEncoding.EncodeToString([]byte(client_id + ":" + client_secret))
	req, err := http.NewRequest("POST", "https://accounts.spotify.com/api/token", strings.NewReader(data.Encode()))
	req.Header.Add(http.CanonicalHeaderKey("Authorization"), "Basic "+base64Id)
	req.Header.Add(http.CanonicalHeaderKey("Content-Type"), "application/x-www-form-urlencoded")
	resp, err := client.Do(req)
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	var token AuthToken
	jsonErr := json.Unmarshal(body, &token)
	if jsonErr != nil {
		fmt.Println("cannot convert to json")
	}
	if err != nil {
		fmt.Println(err)
	}
	return token

}

//Just searches for a song, TODO
func searchSong(token *AuthToken, artist string, song string) {
	client := &http.Client{}

	noSpacesArtistName := strings.Replace(artist, " ", "%20", -1)
	noSpacesSongName := strings.Replace(artist, " ", "%20", -1)

	req, err := http.NewRequest("GET", "https://api.spotify.com/v1/search?q="+noSpacesArtistName+"%20"+noSpacesSongName+"&type=track", strings.NewReader(""))
	if err != nil {
		fmt.Println("Something failed in the http request")
	}
	req.Header.Add(http.CanonicalHeaderKey("Authorization"), "Bearer "+token.AccessToken)
	req.Header.Add(http.CanonicalHeaderKey("Accept"), "application/json")
	req.Header.Add(http.CanonicalHeaderKey("Content-Type"), "application/json")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Couldnt read all")
	}
	//TODO WRONG UNMARSHAL
	var AudioAnalysis AudioAnalysisResponse
	unmarshalErr := json.Unmarshal(body, &AudioAnalysis)
	fmt.Println(AudioAnalysis)
	if unmarshalErr != nil {
		fmt.Println("Could not unmarshal into JSON")
	}
}
func getAudioAnalysis(token *AuthToken, inputURI string) AudioAnalysisResponse {

	spotifyURI := getURI(inputURI)
	if spotifyURI == "" {
		return AudioAnalysisResponse{}
	}
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api.spotify.com/v1/audio-analysis/"+spotifyURI, strings.NewReader(""))

	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add(http.CanonicalHeaderKey("Authorization"), "Bearer "+token.AccessToken)
	req.Header.Add(http.CanonicalHeaderKey("Accept"), "application/json")
	req.Header.Add(http.CanonicalHeaderKey("Content-Type"), "application/json")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Couldnt read all")
	}
	var AudioAnalysis AudioAnalysisResponse
	unmarshalErr := json.Unmarshal(body, &AudioAnalysis)
	if unmarshalErr != nil {
		fmt.Println("Could not unmarshal into JSON")
	}
	fmt.Println(AudioAnalysis)
	return AudioAnalysis
}
func getAudioFeatures(token *AuthToken, inputURI string) AudioFeatures {
	spotifyURI := getURI(inputURI)
	if spotifyURI == "" {
		return AudioFeatures{}
	}
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api.spotify.com/v1/audio-features/"+spotifyURI, strings.NewReader(""))

	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add(http.CanonicalHeaderKey("Authorization"), "Bearer "+token.AccessToken)
	req.Header.Add(http.CanonicalHeaderKey("Accept"), "application/json")
	req.Header.Add(http.CanonicalHeaderKey("Content-Type"), "application/json")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Couldnt read all")
	}
	var AudioFeatures AudioFeatures
	unmarshalErr := json.Unmarshal(body, &AudioFeatures)
	if unmarshalErr != nil {
		fmt.Println("Could not unmarshal into JSON")
	}
	fmt.Println(AudioFeatures)
	return AudioFeatures
}
func getURI(inputURI string) string {
	idxOfSlashOrColon := len(inputURI) - 1
	if idxOfSlashOrColon < 0 {
		return ""
	}
	spotifyURI := ""
	for {
		//TODO this looks hacky
		if rune(inputURI[idxOfSlashOrColon]) == rune(":"[0]) || rune(inputURI[idxOfSlashOrColon]) == rune("/"[0]) {
			spotifyURI = inputURI[idxOfSlashOrColon+1 : len(inputURI)]
			break
		}
		idxOfSlashOrColon--
	}
	return spotifyURI
}
