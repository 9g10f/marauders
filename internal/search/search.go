package search

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func Search(logLevel string, sPath string) func (w http.ResponseWriter, r *http.Request) {
	return func (w http.ResponseWriter, r *http.Request) {
		url := r.URL.Path
		logUrl := r.Method + " " + url
		urlQuery := r.URL.Query()

		w.Header().Set("Content-Type", "application/json")

		if r.Method != "GET" {
			json.NewEncoder(w).Encode(map[string]string{"error": "Invalid method"})
			slog.Warn("API misuse", "url", logUrl)

			return 
		}

		txt := urlQuery.Get("txt")
		if txt == "" {
			json.NewEncoder(w).Encode(map[string]string{"error": "No txt"})
			slog.Warn("API misuse", "url", logUrl)

			return 
		}

		scripts, err := os.ReadDir(sPath)
		if err != nil {
			json.NewEncoder(w).Encode(map[string]string{"error": "Internal server error"})
			slog.Error("Error while reading sPath directory", "url", logUrl, "error", err)

			return 
		}

		var results []string

		for _, script := range scripts {
			if script.IsDir() {
				continue
			}

			scriptBytes, err := os.ReadFile(filepath.Join(sPath, script.Name()))
			if err != nil {
				json.NewEncoder(w).Encode(map[string]string{"error": "Internal server error"})
				slog.Error("Error while reading file", "file", filepath.Join(sPath, script.Name()), "url", logUrl, "error", err)

				return 
			}

			if strings.Contains(strings.ToLower(string(scriptBytes)), strings.ToLower(txt)) {
				results = append(results, script.Name())
			}
		}

		json.NewEncoder(w).Encode(map[string]any{"error": nil, "results": results})
		slog.Info("Successfull serve", "url", logUrl)
	}
}