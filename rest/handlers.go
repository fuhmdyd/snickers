package rest

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/flavioribeiro/snickers/db"
	"github.com/flavioribeiro/snickers/db/memory"
)

func Index(w http.ResponseWriter, r *http.Request) {
	// TODO a great page for API root location
	fmt.Fprint(w, "Snickers")
}

func CreatePreset(w http.ResponseWriter, r *http.Request) {
	var preset db.Preset
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&preset)

	dbInstance, _ := memory.GetDatabase()
	dbInstance.StorePreset(preset)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func UpdatePreset(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "update preset")
}

func ListPresets(w http.ResponseWriter, r *http.Request) {
	dbInstance, err := memory.GetDatabase()
	if err != nil {
		fmt.Fprint(w, "error while creating database")
	}

	var result []string
	for _, preset := range dbInstance.GetPresets() {
		presetJson, _ := json.Marshal(preset)
		result = append(result, string(presetJson))
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	fmt.Fprintf(w, "%s", result)
}

func GetPresetDetails(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "get preset details")
}

func CreateJob(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "create job")
}

func ListJobs(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "list jobs")
}

func GetJobDetails(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "get job details")
}
