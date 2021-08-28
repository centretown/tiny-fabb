package camera

import (
	"fmt"
	"os"

	"gocv.io/x/gocv"
)

var (
	classifier gocv.CascadeClassifier
)

func LoadClassifier() (err error) {
	var usr string
	usr, err = os.UserHomeDir()
	if err != nil {
		return
	}
	// load classifier to recognize faces
	fname := usr + "/src/opencv/data/haarcascades/haarcascade_frontalface_default.xml"
	classifier = gocv.NewCascadeClassifier()
	if !classifier.Load(fname) {
		err = fmt.Errorf("error reading cascade file: %s", fname)
		return
	}
	return
}

func CloseClassifier() {
	classifier.Close()
}
