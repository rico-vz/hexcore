package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"gopkg.in/yaml.v3"
)

const (
	swaggerURL  = "https://raw.githubusercontent.com/dysolix/hasagi-types/refs/heads/main/swagger.json"
	outputYAML  = "api.yaml"
	tagToRemove = "Plugin riot-messaging-service" // Gives issues in generated code when not removed cus {c}
)

func main() {
	fmt.Println("[1] Downloading latest swagger...")
	resp, err := http.Get(swaggerURL)
	if err != nil {
		fmt.Printf("[Error] Error downloading swagger: %v\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("[Error] Bad status code: %d\n", resp.StatusCode)
		os.Exit(1)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("[Error] Couldn't read response body: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Download complete.")

	var swaggerData map[string]interface{}
	if err := json.Unmarshal(body, &swaggerData); err != nil {
		fmt.Printf("[Error] Failed parsing the JSON: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("[2] Removing paths with tag '%s'...\n", tagToRemove)
	paths, ok := swaggerData["paths"].(map[string]interface{})
	if !ok {
		fmt.Println("[Error] Couldnt find paths in the json")
		os.Exit(1)
	}

	pathsToDelete := make([]string, 0)

	for path, pathDataObject := range paths {
		pathMethods, ok := pathDataObject.(map[string]interface{})
		if !ok {
			continue
		}

		for _, methodDataObject := range pathMethods {
			method, ok := methodDataObject.(map[string]interface{})
			if !ok {
				continue
			}

			if tagsObject, ok := method["tags"].([]interface{}); ok {
				for _, tagObject := range tagsObject {
					if tag, ok := tagObject.(string); ok && tag == tagToRemove {
						pathsToDelete = append(pathsToDelete, path)
						goto nextPath
					}
				}
			}
		}
	nextPath:
	}

	for _, path := range pathsToDelete {
		delete(paths, path)
	}
	fmt.Printf("Removed %d pathd.\n", len(pathsToDelete))

	fmt.Println("[3] Converting YAML...")
	yamlData, err := yaml.Marshal(swaggerData)
	if err != nil {
		fmt.Printf("[Error] Failed to convert to YAML: %v\n", err)
		os.Exit(1)
	}

	err = os.WriteFile(outputYAML, yamlData, 0644)
	if err != nil {
		fmt.Printf("[Error] Failed to write YAML file: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Done. Output saved to %s.\n", outputYAML)
}
