package camera

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Framesize uint8

const (
	FRAMESIZE_96X96   Framesize = iota // 96x96
	FRAMESIZE_QQVGA                    //160x120
	FRAMESIZE_QCIF                     //176x144
	FRAMESIZE_HQVGA                    //240x176
	FRAMESIZE_240X240                  //240x240
	FRAMESIZE_QVGA                     //320x240
	FRAMESIZE_CIF                      //400x296
	FRAMESIZE_HVGA                     //480x320
	FRAMESIZE_VGA                      //640x480
	FRAMESIZE_SVGA                     //800x600
	FRAMESIZE_XGA                      //1024x768
	FRAMESIZE_HD                       //1280x720
	FRAMESIZE_SXGA                     //1280x1024
	FRAMESIZE_UXGA                     //1600x1200
	// 3MP Sensors
	FRAMESIZE_FHD   //1920x1080
	FRAMESIZE_P_HD  // 720x1280
	FRAMESIZE_P_3MP // 864x1536
	FRAMESIZE_QXGA  //2048x1536
	// 5MP Sensors
	FRAMESIZE_QHD   //2560x1440
	FRAMESIZE_WQXGA //2560x1600
	FRAMESIZE_P_FHD //1080x1920
	FRAMESIZE_QSXGA //2560x1920
	FRAMESIZE_INVALID
)

var framesizeText = []string{
	"96x96",
	"QQVGA 160x120",
	"QCIF 176x144",
	"HQVGA 240x176",
	"240x240",
	"QVGA 320x240",
	"CIF 400x296",
	"HVGA 480x320",
	"VGA 640x480",
	"SVGA 800x600",
	"XGA 1024x768",
	"HD 1280x720",
	"SXGA 1280x1024",
	"UXGA 1600x1200",
	"FHD 1920x1080",
	"P_HD 720x1280",
	"P_3MP 864x1536",
	"QXGA 2048x1536",
	"QHD 2560x1440",
	"WQXGA 2560x1600",
	"P_FHD 1080x1920",
	"QSXGA 2560x1920",
	"Invalid framesize",
}

func (fr Framesize) String() string {
	if fr >= FRAMESIZE_INVALID {
		return framesizeText[FRAMESIZE_INVALID]
	}
	return framesizeText[fr]
}

type CameraStatus struct {
	Framesize     Framesize `json:"framesize"`  //0 - 10
	Quality       uint8     `json:"quality"`    //0 - 63
	Brightness    int8      `json:"brightness"` //-2 - 2
	Contrast      int8      `json:"contrast"`   //-2 - 2
	Saturation    int8      `json:"saturation"` //-2 - 2
	Sharpness     int8      `json:"sharpness"`  //-2 - 2
	Denoise       uint8     `json:"denoise"`
	SpecialEffect uint8     `json:"special_effect"` //0 - 6
	WbMode        uint8     `json:"wb_mode"`        //0 - 4
	Awb           uint8     `json:"awb"`
	AwbGain       uint8     `json:"awb_gain"`
	Aec           uint8     `json:"aec"`
	Aec2          uint8     `json:"aec2"`
	AeLevel       int8      `json:"ae_level"`  //-2 - 2
	AecValue      int16     `json:"aec_value"` //0 - 1200
	Agc           uint8     `json:"agc"`
	AgcGain       uint8     `json:"agc_gain"`    //0 - 30
	GainCeiling   uint8     `json:"gainceiling"` //0 - 6
	Bpc           uint8     `json:"bpc"`
	Wpc           uint8     `json:"wpc"`
	RawGma        uint8     `json:"raw_gma"`
	Lenc          uint8     `json:"lenc"`
	Hmirror       uint8     `json:"hmirror"`
	Vflip         uint8     `json:"vflip"`
	Dcw           uint8     `json:"dcw"`
	Colorbar      uint8     `json:"colorbar"`
}

func (cam *Camera) GetStatus() (err error) {
	var data []byte
	data, err = request(cam.statusUrl)
	if err != nil {
		return
	}
	err = json.Unmarshal(data, &cam.Settings)
	if err != nil {
		return
	}
	fmt.Println(cam.Settings)
	return
}

func (cam *Camera) SetFrameSize(val int) error {
	return cam.set(val, "framesize", 0, 10, func(val int) {
		cam.Settings.Framesize = Framesize(val)
	})
}

func (cam *Camera) SetQuality(val int) error {
	return cam.set(val, "quality", 0, 63, func(val int) {
		cam.Settings.Quality = uint8(val)
	})
}

func (cam *Camera) SetBrightness(val int) error {
	return cam.set(val, "brightness", -2, 2, func(val int) {
		cam.Settings.Brightness = int8(val)
	})

}

func (cam *Camera) SetContrast(val int) error {
	return cam.set(val, "contrast", -2, 2, func(val int) {
		cam.Settings.Contrast = int8(val)
	})
}

func (cam *Camera) SetSaturation(val int) error {
	return cam.set(val, "saturation", -2, 2, func(val int) {
		cam.Settings.Saturation = int8(val)
	})
}

func (cam *Camera) SetSharpness(val int) error {
	return cam.set(val, "sharpness", -2, 2, func(val int) {
		cam.Settings.Sharpness = int8(val)
	})
}

func (cam *Camera) SetDenoise(val int) error {
	return cam.set(val, "denoise", 0, 1, func(val int) {
		cam.Settings.Denoise = uint8(val)
	})
}

func (cam *Camera) SetSpecialEffect(val int) error {
	return cam.set(val, "special_effect", 0, 6, func(val int) {
		cam.Settings.SpecialEffect = uint8(val)
	})
}

func (cam *Camera) SetWbMode(val int) error {
	return cam.set(val, "wb_mode", 0, 4, func(val int) {
		cam.Settings.WbMode = uint8(val)
	})
}

func (cam *Camera) SetAwb(val int) error {
	return cam.set(val, "awb", 0, 1, func(val int) {
		cam.Settings.Awb = uint8(val)
	})
}

func (cam *Camera) SetAwbGain(val int) error {
	return cam.set(val, "awb_gain", 0, 1, func(val int) {
		cam.Settings.AwbGain = uint8(val)
	})
}
func (cam *Camera) SetAec(val int) error {
	return cam.set(val, "aec", 0, 1, func(val int) {
		cam.Settings.Aec = uint8(val)
	})
}

func (cam *Camera) SetAec2(val int) error {
	return cam.set(val, "aec2", 0, 1, func(val int) {
		cam.Settings.Aec2 = uint8(val)
	})
}

func (cam *Camera) SetAeLevel(val int) error {
	return cam.set(val, "ae_level", -2, 2, func(val int) {
		cam.Settings.AeLevel = int8(val)
	})
}

func (cam *Camera) SetAecValue(val int) error {
	return cam.set(val, "aec_value", 0, 1200, func(val int) {
		cam.Settings.AecValue = int16(val)
	})
}

func (cam *Camera) SetAgc(val int) error {
	return cam.set(val, "agc", 0, 1, func(val int) {
		cam.Settings.Agc = uint8(val)
	})
}

func (cam *Camera) SetAgcGain(val int) error {
	return cam.set(val, "agc_gain", 0, 30, func(val int) {
		cam.Settings.AgcGain = uint8(val)
	})
}

func (cam *Camera) SetGainCeiling(val int) error {
	return cam.set(val, "gainceiling", 0, 6, func(val int) {
		cam.Settings.GainCeiling = uint8(val)
	})
}

func (cam *Camera) SetBpc(val int) error {
	return cam.set(val, "bpc", 0, 1, func(val int) {
		cam.Settings.Bpc = uint8(val)
	})
}

func (cam *Camera) SetWpc(val int) error {
	return cam.set(val, "wpc", 0, 1, func(val int) {
		cam.Settings.Wpc = uint8(val)
	})
}

func (cam *Camera) SetRawGma(val int) error {
	return cam.set(val, "raw_gma", 0, 1, func(val int) {
		cam.Settings.RawGma = uint8(val)
	})
}

func (cam *Camera) SetLenc(val int) error {
	return cam.set(val, "lenc", 0, 1, func(val int) {
		cam.Settings.Lenc = uint8(val)
	})
}

func (cam *Camera) SetHmirror(val int) error {
	return cam.set(val, "hmirror", 0, 1, func(val int) {
		cam.Settings.Hmirror = uint8(val)
	})
}

func (cam *Camera) SetVflip(val int) error {
	return cam.set(val, "vflip", 0, 1, func(val int) {
		cam.Settings.Vflip = uint8(val)
	})
}

func (cam *Camera) SetDcw(val int) error {
	return cam.set(val, "dcw", 0, 1, func(val int) {
		cam.Settings.Dcw = uint8(val)
	})
}

func (cam *Camera) SetColorbar(val int) error {
	return cam.set(val, "colorbar", 0, 1, func(val int) {
		cam.Settings.Colorbar = uint8(val)
	})
}

func (cam *Camera) set(val int, code string, min, max int, update func(int)) (err error) {
	if val < min || val > max {
		err = fmt.Errorf("%s: %d out of range (%d-%d)", code, val, min, max)
		return
	}
	_, err = request(fmt.Sprintf("%s?var=%s&val=%d", cam.controlUrl, code, val))
	if err == nil {
		update(val)
	}
	return
}

func (cam *Camera) MoveLeft() (err error) {
	return cam.move("left")
}
func (cam *Camera) MoveRight() (err error) {
	return cam.move("right")
}
func (cam *Camera) MoveUp() (err error) {
	return cam.move("up")
}
func (cam *Camera) MoveDown() (err error) {
	return cam.move("down")
}

func (cam *Camera) move(val string) (err error) {
	s := fmt.Sprintf("%s?go=%s", cam.controlUrl, val)
	fmt.Println(s)
	_, err = request(s)
	return
}

func request(u string) (data []byte, err error) {
	req, err := http.NewRequest("GET", u, nil)
	if err != nil {
		return
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	data, err = io.ReadAll(resp.Body)
	return
}
