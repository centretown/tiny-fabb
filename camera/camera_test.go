package camera

import (
	"html/template"
	"testing"
	"time"

	"github.com/gorilla/mux"
)

var (
	addr     = ":8080"
	interval = 200 * time.Millisecond
	webpage  = `
	<html>
    <head>
        <style>
            .camera {
                padding: 0px;
                margin: 0;
                margin-block-start: 0;
                margin-block-end: 0;
                margin-inline-start: 0;
                margin-inline-end: 0;
            }

            .camera img {
                display: block;
                border-radius: 4px;
                margin-top: 8px;
            }
        </style>
    </head>
    <body>
    <figure class="camera">
        <div id="stream-camera0" class="image-container">
            <img id="camera0-stream" src="/camera0/mjpeg">
        </div>
    </figure>
    <figure class="camera">
        <div id="stream-camera1" class="image-container">
            <img id="camera1-stream" src="/camera1/mjpeg">
        </div>
    </figure>
    </body>
</html>`
)

func getLayout(assetsPath string) (layout *template.Template) {
	layout = template.Must(
		layout.ParseFiles(
			assetsPath+"/layout.html",
			assetsPath+"/layout.go.tpl",
			assetsPath+"/entry.go.tpl",
			assetsPath+"/camera.go.tpl",
			assetsPath+"/servo.go.tpl"))
	return
}

func TestCamera(t *testing.T) {
	router := mux.NewRouter()
	dataSource := "../assets/data/"
	layout := getLayout("../assets")
	servoController := "http://192.168.0.99:80/action"

	camConn := NewConnector(dataSource, layout, 200)
	camConn.Start(router, servoController,
		"/dev/video0", "http://192.168.0.99")
}

const json_test = `{ "framesize":7,
"quality":10,
"brightness:0,
"contrast:0,
"saturation:0,
"sharpness:0,
"special_effect":0,
"wb_mode":0,
"awb":1,
"awb_gain":1,
"aec":1,
"aec2":0,
"ae_level:0,
"aec_value":168,
"agc":1,
"agc_gain":0,
"gainceiling":0,
"bpc":0,
"wpc":1,
"raw_gma":1,
"lenc":1,
"vflip":0,
"hmirror":0,
"dcw":1,
"colorbar":0 }`

const json_test_flat = `{"framesize":7,"quality":10,"brightness":0,"contrast":0,"saturation":0,"sharpness":0,"special_effect":0,"wb_mode":0,"awb":1,"awb_gain":1,"aec":1,"aec2":0,"ae_level":0,"aec_value":168,"agc":1,"agc_gain":0,"gainceiling":0,"bpc":0,"wpc":1,"raw_gma":1,"lenc":1,"vflip":0,"hmirror":0,"dcw":1,"colorbar":0}`

// func testJson(t *testing.T) {
// 	var cam Camera
// 	err := json.Unmarshal([]byte(json_test), &cam.Settings)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	err = json.Unmarshal([]byte(json_test_flat), &cam.Settings)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Println(cam.Settings)
// }
