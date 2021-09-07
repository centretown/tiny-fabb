package camera

import (
	"fmt"
	"time"

	"github.com/gorilla/mux"
)

type Cameras map[string]*Camera

func (cams Cameras) Start(router *mux.Router, interval time.Duration, urls ...string) {
	// LoadClassifier()
	const (
		format = "camera%d"
	)
	var title string

	for i, url := range urls {
		title = fmt.Sprintf(format, i)
		cam := &Camera{
			Title:      title,
			controlUrl: url + ":80/action",
			captureUrl: url + ":80/capture",
			statusUrl:  url + ":80/status",
			streamUrl:  url + ":81/stream",
			interval:   interval,
		}
		cams[title] = cam

		cam.SetFrameSize(10)
		cam.SetQuality(10)
		cam.GetStatus()

		// err := testMove(cam.MoveUp)
		// if err != nil {
		// 	glog.Info(err)
		// }
		// time.Sleep(500)
		// err = testMove(cam.MoveLeft)
		// if err != nil {
		// 	glog.Info(err)
		// }
		// time.Sleep(500)
		// err = testMove(cam.MoveDown)
		// if err != nil {
		// 	glog.Info(err)
		// }
		// time.Sleep(500)
		// err = testMove(cam.MoveRight)
		// if err != nil {
		// 	glog.Info(err)
		// }
		// time.Sleep(500)
		cam.Start(router)
	}
}

// func testMove(m func() error) (err error) {
// 	for i := 0; i < 10; i++ {
// 		err = m()
// 		if err != nil {
// 			return
// 		}
// 	}
// 	return
// }

func (cams Cameras) Stop() {
	for _, cam := range cams {
		cam.Stop()
	}
	// CloseClassifier()
}
