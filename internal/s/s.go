package s

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func Script(logLevel string, sPath string) func (w http.ResponseWriter, r *http.Request) {
	return func (w http.ResponseWriter, r *http.Request) {
		url := r.URL.Path
		logUrl := r.Method + " " + url

		if r.Method != "GET" {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{"error": "Invalid method"})

			slog.Warn("API misuse", "url", logUrl)

			return 
		}

		if len(strings.SplitAfter(url, "/")) != 2 || url == "/" {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{"error": "Invalid path"})

			slog.Warn("API misuse", "url", logUrl)

			return 
		}

		filePath := filepath.Join(sPath, url)

		_, err := os.Stat(filePath)
		if err == nil {
			scriptBytes, err := os.ReadFile(filePath)
			if err != nil {
				slog.Error("Error while reading file", "file", filePath, "url", logUrl, "error", err)

				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(map[string]any{"error": "Internal server error"})

				return 
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]any{"error": nil, "script": string(scriptBytes)})

			slog.Info("Successfull serve", "url", logUrl)
		} else if errors.Is(err, os.ErrNotExist) {
			slog.Info("Game script not found", "file", filePath, "url", logUrl)

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]any{"error": "Game script not found"})

			return 
		} else {
			slog.Error("Error while opening file", "file", filePath, "url", logUrl, "error", err)

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]any{"error": "Internal server error"})
		}
	}
}