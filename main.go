package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

var (
	threads      int
	outputFolder string
)

func main() {

	fmt.Fprintln(os.Stderr, `

	██████╗  ██████╗ ██╗    ██╗███╗   ██╗     ██╗███████╗
	██╔══██╗██╔═══██╗██║    ██║████╗  ██║     ██║██╔════╝
	██║  ██║██║   ██║██║ █╗ ██║██╔██╗ ██║     ██║███████╗
	██║  ██║██║   ██║██║███╗██║██║╚██╗██║██   ██║╚════██║
	██████╔╝╚██████╔╝╚███╔███╔╝██║ ╚████║╚█████╔╝███████║
	╚═════╝  ╚═════╝  ╚══╝╚══╝ ╚═╝  ╚═══╝ ╚════╝ ╚══════╝
														 
	
	`)
	fmt.Fprintln(os.Stderr, "[i] by: deeplooklabs.com")
	fmt.Fprintln(os.Stderr, "[i] Waiting for download...")

	flag.StringVar(&outputFolder, "o", "downjs_output", "Set Output folder")
	flag.IntVar(&threads, "t", 20, "max threads to download files")
	flag.Parse()

	urls := make(chan string)
	var wg sync.WaitGroup
	sem := make(chan struct{}, threads)

	for i := 0; i < threads; i++ {
		go func() {
			for url := range urls {
				baseURL, err := extractBaseURL(url)

				if err != nil {
					return

				}

				if checkMap(baseURL) {
					downloadAndSave(baseURL + ".map")
				} else {
					downloadAndSave(baseURL)
				}

				wg.Done()
				<-sem
			}
		}()
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		url := scanner.Text()
		wg.Add(1)
		sem <- struct{}{}
		urls <- url
	}

	close(urls)
	wg.Wait()
}

func extractBaseURL(rawURL string) (string, error) {
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		return "", err
	}

	baseURL := parsedURL.Scheme + "://" + parsedURL.Host + parsedURL.Path

	return baseURL, nil
}

func checkMap(url string) bool {
	baseURL, err := extractBaseURL(url)
	resp, err := http.Get(baseURL + ".map")
	if err != nil {
		return false
	}

	defer resp.Body.Close()
	contentType := resp.Header.Get("Content-Type")

	if resp.StatusCode != http.StatusOK {
		return false
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false
	}
	bodyString := string(body)

	if contentType == "application/json" && strings.Contains(bodyString, "version\":") || contentType == "text/plain" && strings.Contains(bodyString, "version\":") {
		return true
	} else {
		return false
	}
}

func downloadAndSave(url string) {
	baseURL, err := extractBaseURL(url)
	resp, err := http.Get(baseURL)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return
	}

	makedir(outputFolder)

	fileName := strings.Replace(url, "://", "-", -1)
	fileName = strings.Replace(fileName, "/", "_", -1)

	fileName = filepath.Base(fileName)
	filePath := filepath.Join(outputFolder, fileName)

	file, err := os.Create(filePath)
	if err != nil {
		return
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return
	}

	fmt.Printf("Download of %s done and save on %s\n", url, filePath)
}

func makedir(name string) {
	if _, err := os.Stat(name); os.IsNotExist(err) {
		err := os.Mkdir(name, os.ModePerm)
		if err != nil {
			fmt.Println("Error to make output folder:", err)
			return
		}
	} else {
		return
	}
}
