package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	apiKey := "sk_95d61733e9554e2d434d8b65da70495fc4b03265" 
	htmlFilePath := "/home/chandra/Music/resume-pdf-converter/DevOpsResume.html"
	outputPDFPath := "resume_cv_single_page54.pdf"

	htmlContent, err := os.ReadFile(htmlFilePath)
	if err != nil {
		log.Fatalf(" Failed to read HTML file: %v", err)
	}

	
	payload := map[string]interface{}{
		"source":    string(htmlContent),
		"format":    "A5",       
		"landscape": false,
		"use_print": true,
		"zoom":      0.589,      
		"margin": map[string]string{
			"top":    "0in",
			"bottom": "0in",
			"left":   "0in",
			"right":  "0in",
		},
	}

	// Prepare the request
	jsonData, err := json.Marshal(payload)
	if err != nil {
		log.Fatalf("Failed to marshal payload: %v", err)
	}

	req, err := http.NewRequest("POST", "https://api.pdfshift.io/v3/convert/pdf", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatalf("Failed to create request: %v", err)
	}
	req.Header.Set("X-API-Key", apiKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf(" Request error: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalf(" Failed to read response: %v", err)
		}
		if err := ioutil.WriteFile(outputPDFPath, body, 0644); err != nil {
			log.Fatalf(" Failed to write output: %v", err)
		}
		log.Printf(" PDF saved as: %s", outputPDFPath)
	} else {
		var errResp map[string]interface{}
		json.NewDecoder(resp.Body).Decode(&errResp)
		log.Printf(" API Error Response (Status %d): %v", resp.StatusCode, errResp)
	}
}
