package main

import (
	log "github.com/mdean75/simple-log"
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

	log.Entry().WithStruct(testData).WithCaller().SetLongFile().Info("this should have long file format")
	log.Info("this should be the default without caller")
	log.Entry().WithCaller().Info("this should have caller on the default logger by calling logger")
	customLogger()

	log.Entry().WithStruct(testData).Debug("test with struct")
	log.Info("test ... this should have caller info without calling withCaller")
	//log.Info("test")
	//log.Info("test 2")
	//log.Entry().WithCaller().WithStruct(testData).Info("test with caller")
	//log.Entry().SetLongFile().WithCaller().Info("long file new test")
	//log.Entry().WithCaller().Info("test new after setting long file")
	//log.Entry().SetShortFile().WithCaller().Info("with short file")
	//log.Entry().SetLongFile().WithCaller().Info("test default out stream")
	//log.Entry().SetOutStream(os.Stderr).WithCaller().Info("test out stream stderr")
	//log.Info("back to testing default")

}
