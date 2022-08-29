package main

import (
	_ "embed"
	"encoding/json"
	events "logreader/events"
	filewatcher "logreader/file-watcher"
	parser "logreader/parser"
	trailer "logreader/trailing"
)

//go:embed config.json
var config []byte

type Config struct {
	LogFolder   string
	LogFileName string
}

func main() {
	var c Config
	json.Unmarshal(config, &c)

	readLine := make(chan string)
	fileRenamed := make(chan any)
	eventHistory := events.Initialize()

	trailerManager := trailer.Initalize(c.LogFolder+"/"+c.LogFileName, readLine)
	go filewatcher.WatchDir(c.LogFolder, c.LogFileName, fileRenamed)

	for {
		select {
		case x := <-readLine:
			event, err := events.ReadEvent(parser.ParseLine(x))
			if err == nil {
				eventHistory.RecordEvent(event)
			}
		case <-fileRenamed:
			trailerManager.ChangeFile()
		default:
		}
	}
}
