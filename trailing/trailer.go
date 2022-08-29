package trailing

import (
	"bufio"
	"io"
	"log"
	"os"
	"time"
)

type TrailerManager struct {
	trailingFileName string
	lineOut          chan string
	stopReading      chan any
	fileSize         int64
}

func Initalize(fileName string, line chan string) TrailerManager {
	t := TrailerManager{trailingFileName: fileName, lineOut: line, stopReading: make(chan any)}
	go t.trail()
	return t
}

func (t TrailerManager) ChangeFile() {
	close(t.lineOut)
	t.stopReading = make(chan any)
	go t.trail()
}

func getFileSize(f *os.File) (int64, error) {
	stat, err := f.Stat()
	if err != nil {
		log.Println(err)
		return 0, err
	}
	return stat.Size(), nil
}
func (t TrailerManager) trail() {
	file, err := os.Open(t.trailingFileName)
	if err != nil {
		log.Println(err)
		return
	}
	size, err := getFileSize(file)
	if err != nil {
		log.Println(err)
		return
	}

	t.fileSize = size

	reader := bufio.NewReader(file)
	for {
		select {
		case <-t.stopReading:
			file.Close()
			return
		default:
		}
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				s, err := getFileSize(file)
				if err != nil {
					log.Println(err)
					return
				}
				if s < t.fileSize {
					log.Println("truncated")
					_, err = file.Seek(0, io.SeekStart)
					t.fileSize = s
				} else {
					log.Println("EOF")
				}
				time.Sleep(3 * time.Second)
				continue
			}
			log.Println(err)
			break
		}
		t.lineOut <- line

	}
}
