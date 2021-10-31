package camera

import (
	"bytes"
	"encoding/json"
	"fmt"
	"image"
	"image/jpeg"

	"github.com/centretown/tiny-fabb/forms"
	"github.com/mattn/go-mjpeg"
)

type EspProperties struct {
	Framesize     uint8 `json:"framesize"`  //0 - 10
	Quality       uint8 `json:"quality"`    //0 - 63
	Brightness    int8  `json:"brightness"` //-2 - 2
	Contrast      int8  `json:"contrast"`   //-2 - 2
	Saturation    int8  `json:"saturation"` //-2 - 2
	Sharpness     int8  `json:"sharpness"`  //-2 - 2
	Denoise       uint8 `json:"denoise"`
	SpecialEffect uint8 `json:"special_effect"` //0 - 6
	WbMode        uint8 `json:"wb_mode"`        //0 - 4
	Awb           uint8 `json:"awb"`
	AwbGain       uint8 `json:"awb_gain"`
	Aec           uint8 `json:"aec"`
	Aec2          uint8 `json:"aec2"`
	AeLevel       int8  `json:"ae_level"`  //-2 - 2
	AecValue      int16 `json:"aec_value"` //0 - 1200
	Agc           uint8 `json:"agc"`
	AgcGain       uint8 `json:"agc_gain"`    //0 - 30
	GainCeiling   uint8 `json:"gainceiling"` //0 - 6
	Bpc           uint8 `json:"bpc"`
	Wpc           uint8 `json:"wpc"`
	RawGma        uint8 `json:"raw_gma"`
	Lenc          uint8 `json:"lenc"`
	Hmirror       uint8 `json:"hmirror"`
	Vflip         uint8 `json:"vflip"`
	Dcw           uint8 `json:"dcw"`
	Colorbar      uint8 `json:"colorbar"`
}

type EspCam struct {
	Settings   EspProperties `json:"settings"`
	ControlUrl string        `json:"controlUrl"`
	StreamUrl  string        `json:"streamUrl"`
	StatusUrl  string        `json:"statusUrl"`
	CaptureUrl string        `json:"captureUrl"`
	dec        *mjpeg.Decoder
}

func (ncam *EspCam) Open() (err error) {
	ncam.dec, err = mjpeg.NewDecoderFromURL(ncam.StreamUrl)
	return
}

func (ncam *EspCam) Read(buf *bytes.Buffer) (err error) {
	var img image.Image
	img, err = ncam.dec.Decode()
	if err == nil {
		err = jpeg.Encode(buf, img, nil)
	}
	return
}

func (ncam *EspCam) Close() (err error) {
	return
}

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

func (ncam *EspCam) UpdateProperties() (err error) {
	var data []byte
	data, err = forms.Request(ncam.StatusUrl)
	if err != nil {
		return
	}
	err = json.Unmarshal(data, &ncam.Settings)
	if err != nil {
		return
	}
	return
}

func (ncam *EspCam) SetProperty(ent *forms.Entry, val string) (err error) {
	_, err = forms.Request(
		fmt.Sprintf("%s?var=%s&val=%s", ncam.ControlUrl, ent.Code, val))
	return
}

func (ncam *EspCam) BindProperties() (f forms.Forms) {
	f = forms.Forms{
		idFramesize: {
			ID:      idFramesize,
			Value:   &ncam.Settings.Framesize,
			Entries: settingsEntries[idFramesize],
		},
		idQuality: {
			ID:      idQuality,
			Value:   &ncam.Settings.Quality,
			Entries: settingsEntries[idQuality],
		},
		idBrightness: {
			ID:      idBrightness,
			Value:   &ncam.Settings.Brightness,
			Entries: settingsEntries[idBrightness],
		},
		idContrast: {
			ID:      idContrast,
			Value:   &ncam.Settings.Contrast,
			Entries: settingsEntries[idContrast],
		},
		idSaturation: {
			ID:      idSaturation,
			Value:   &ncam.Settings.Saturation,
			Entries: settingsEntries[idSaturation],
		},
		idSharpness: {
			ID:      idSharpness,
			Value:   &ncam.Settings.Sharpness,
			Entries: settingsEntries[idSharpness],
		},
		idDenoise: {
			ID:      idDenoise,
			Value:   &ncam.Settings.Denoise,
			Entries: settingsEntries[idDenoise],
		},
		idSpecialEffect: {
			ID:      idSpecialEffect,
			Value:   &ncam.Settings.SpecialEffect,
			Entries: settingsEntries[idSpecialEffect],
		},
		idWbMode: {
			ID:      idWbMode,
			Value:   &ncam.Settings.WbMode,
			Entries: settingsEntries[idWbMode],
		},
		idAwb: {
			ID:      idAwb,
			Value:   &ncam.Settings.Awb,
			Entries: settingsEntries[idAwb],
		},
		idAwbGain: {
			ID:      idAwbGain,
			Value:   &ncam.Settings.AwbGain,
			Entries: settingsEntries[idAwbGain],
		},
		idAec: {
			ID:      idAec,
			Value:   &ncam.Settings.Aec,
			Entries: settingsEntries[idAec],
		},
		idAec2: {
			ID:      idAec2,
			Value:   &ncam.Settings.Aec2,
			Entries: settingsEntries[idAec2],
		},
		idAeLevel: {
			ID:      idAeLevel,
			Value:   &ncam.Settings.AeLevel,
			Entries: settingsEntries[idAeLevel],
		},
		idAecValue: {
			ID:      idAecValue,
			Value:   &ncam.Settings.AecValue,
			Entries: settingsEntries[idAecValue],
		},
		idAgc: {
			ID:      idAgc,
			Value:   &ncam.Settings.Agc,
			Entries: settingsEntries[idAgc],
		},
		idAgcGain: {
			ID:      idAgcGain,
			Value:   &ncam.Settings.AgcGain,
			Entries: settingsEntries[idAgcGain],
		},
		idGainCeiling: {
			ID:      idGainCeiling,
			Value:   &ncam.Settings.GainCeiling,
			Entries: settingsEntries[idGainCeiling],
		},
		idBpc: {
			ID:      idBpc,
			Value:   &ncam.Settings.Bpc,
			Entries: settingsEntries[idBpc],
		},
		idWpc: {
			ID:      idWpc,
			Value:   &ncam.Settings.Wpc,
			Entries: settingsEntries[idWpc],
		},
		idRawGma: {
			ID:      idRawGma,
			Value:   &ncam.Settings.RawGma,
			Entries: settingsEntries[idRawGma],
		},
		idLenc: {
			ID:      idLenc,
			Value:   &ncam.Settings.Lenc,
			Entries: settingsEntries[idLenc],
		},
		idHmirror: {
			ID:      idHmirror,
			Value:   &ncam.Settings.Hmirror,
			Entries: settingsEntries[idHmirror],
		},
		idVflip: {
			ID:      idVflip,
			Value:   &ncam.Settings.Vflip,
			Entries: settingsEntries[idVflip],
		},
		idDcw: {
			ID:      idDcw,
			Value:   &ncam.Settings.Dcw,
			Entries: settingsEntries[idDcw],
		},
		idColorbar: {
			ID:      idColorbar,
			Value:   &ncam.Settings.Colorbar,
			Entries: settingsEntries[idColorbar],
		},
	}
	return
}
