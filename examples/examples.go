package main

import (
	"github.com/gorilla/mux"
	log "github.com/mdean75/simple-log"
	"net/http"
	"os"
)

func customLogger() {
	settings := log.NewEnabled(false, true, true)
	log.CustomLogger(*settings, os.Stderr)
}
func main() {

	testData := struct {
		Name    string `json:"name"`
		Age     int    `json:"age"`
		Married bool   `json:"married"`
	}{Name: "Michael", Age: 47, Married: true}

	log.WithCaller().WithStruct(testData).Info("test with caller directly")
	log.WithStruct(testData).WithCaller().Info("test with struct directly the with caller")
	log.SetLongFile().WithCaller().Info("test setting long file then caller without entry")
	log.WithStruct(testData).SetOutStream(os.Stderr).WithCaller().Info("line 25")

	log.Entry().SetLongFile().WithCaller().WithStruct(testData).Info(" chained calls")
	log.Entry().WithStruct(testData).SetLongFile().WithCaller().Info("this should have long file format")
	log.Info("this should be the default without caller")
	log.Entry().WithCaller().Info("this should have caller on the default logger by calling logger")
	log.Entry().Info("test line 27")
	customLogger()
	log.Entry().Info("test line 29")
	log.Info("test line 30")

	log.Entry().WithStruct(testData).Debug("test with struct")
	log.Info("test ... this should have caller info without calling withCaller")
	log.Entry().SetLongFile().WithStruct(testData).WithCaller().Info("re-arranging chained calls")
	log.Entry().SetLongFile().WithStruct(testData).SetOutStream(os.Stderr).SetShortFile().SetLongFile().WithCaller().Info("garbage ... a mess of chaining short and long format calls")
	//log.Info("test")
	//log.Info("test 2")
	//log.Entry().WithCaller().WithStruct(testData).Info("test with caller")
	//log.Entry().SetLongFile().WithCaller().Info("long file new test")
	//log.Entry().WithCaller().Info("test new after setting long file")
	//log.Entry().SetShortFile().WithCaller().Info("with short file")
	//log.Entry().SetLongFile().WithCaller().Info("test default out stream")
	//log.Entry().SetOutStream(os.Stderr).WithCaller().Info("test out stream stderr")
	//log.Info("back to testing default")

	r := mux.NewRouter()
	r.HandleFunc("/health", health())

	err := http.ListenAndServe(":4000", r)
	if err != nil {
		log.Entry().WithCaller().Info("error from listen and serve")
	}
}

func health() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusAccepted)

		log.Entry().SetLongFile().WithCaller().SetOutStream(w).Info("test health check")
	}
}
