package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomHandler(w http.ResponseWriter, r *http.Request) {
	firstPart := []string{"DSC", "MOV", "IMG", "100", "MVI"}
	separator := []string{" ", "_", ""}

	randomFirstPart := firstPart[rand.Intn(len(firstPart))]
	randomSeparator := separator[rand.Intn(len(separator))]
	numberBase := rand.Intn(9999)

	padToFour := func(number int) string {
		return fmt.Sprintf("%04d", number)
	}

	numbers := padToFour(numberBase)

	randomQuery := randomFirstPart + randomSeparator + numbers

	videoID, err := searchRandomVideo(randomQuery)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"videoId": videoID})
}

func searchRandomVideo(query string) (string, error) {
	apiKey := os.Getenv("YT_API_KEY")
	baseURL := "https://www.googleapis.com/youtube/v3/search?part=snippet&maxResults=4&q=%s&key=%s"

	url := fmt.Sprintf(baseURL, url.QueryEscape(query), apiKey)

	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("error making request: %v", err)
	}
	defer resp.Body.Close()

	var data struct {
		Items []struct {
			ID struct {
				VideoID string `json:"videoId"`
			} `json:"id"`
			Snippet struct {
				Title string `json:"title"`
			} `json:"snippet"`
		} `json:"items"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return "", fmt.Errorf("error decoding response: %v", err)
	}

	if len(data.Items) == 0 {
		return "", fmt.Errorf("no videos found")
	}

	for _, item := range data.Items {
		if item.Snippet.Title == query {
			return item.ID.VideoID, nil
		}
	}

	return "", fmt.Errorf("no video found with exact title match")
}

func main() {
	fmt.Println("Hello go...")

	fmt.Println("YT_API_KEY:", os.Getenv("YT_API_KEY"))

	http.HandleFunc("/random", randomHandler)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Serving index.html")
		http.ServeFile(w, r, "index.html")
	})
	
	log.Fatal(http.ListenAndServe(":3000", nil))
}
