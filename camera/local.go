package camera

import (
	"github.com/centretown/tiny-fabb/forms"
	"github.com/mattn/go-mjpeg"
	"gocv.io/x/gocv"
)

type LocalSettings struct {
	FrameWidth   float64 `json:"frameWidth"`
	FrameHeight  float64 `json:"frameHeight"`
	FrameRate    float64 `json:"framerate"`
	Format       float64 `json:"format"`
	Mode         float64 `json:"mode"`
	Brightness   float64 `json:"brightness"`
	Contrast     float64 `json:"contrast"`
	Saturation   float64 `json:"saturation"`
	Hue          float64 `json:"hue"`
	Gain         float64 `json:"gain"`
	Monochrome   float64 `json:"monochrome"`
	Sharpness    float64 `json:"sharpness"`
	AutoExposure float64 `json:"autoExposure"`
	Gamma        float64 `json:"gamma"`
	Temperature  float64 `json:"temperature"`
}

type LocalCam struct {
	Settings  LocalSettings `json:"settings"`
	StreamUrl string        `json:"streamUrl"`
	Stream    *mjpeg.Stream `json:"-"`

	vc  *gocv.VideoCapture
	mat gocv.Mat
}

func (lcam *LocalCam) Open() (err error) {
	lcam.vc, err = gocv.OpenVideoCapture(lcam.StreamUrl)
	lcam.mat = gocv.NewMat()
	lcam.vc.Set(gocv.VideoCaptureFrameWidth, 1920)
	lcam.vc.Set(gocv.VideoCaptureFrameHeight, 1080)
	return
}

func (lcam *LocalCam) Read() (buf []byte, err error) {
	if ok := lcam.vc.Read(&lcam.mat); !ok {
		return
	}
	nbuf, err := gocv.IMEncode(".jpg", lcam.mat)
	if err != nil {
		return
	}
	buf = nbuf.GetBytes()
	return
}

func (lcam *LocalCam) Close() (err error) {
	lcam.vc.Close()
	return
}

func (lcam *LocalCam) SetProperty(ent *forms.Entry, val string) (err error) {
	return
}

func (lcam *LocalCam) UpdateProperties() (err error) {
	vc := lcam.vc
	lcam.Settings.FrameHeight = vc.Get(gocv.VideoCaptureFrameHeight)
	lcam.Settings.FrameWidth = vc.Get(gocv.VideoCaptureFrameWidth)
	lcam.Settings.FrameRate = vc.Get(gocv.VideoCaptureFPS)
	lcam.Settings.Format = vc.Get(gocv.VideoCaptureFormat)
	lcam.Settings.Mode = vc.Get(gocv.VideoCaptureMode)
	lcam.Settings.Brightness = vc.Get(gocv.VideoCaptureBrightness)
	lcam.Settings.Contrast = vc.Get(gocv.VideoCaptureContrast)
	lcam.Settings.Saturation = vc.Get(gocv.VideoCaptureSaturation)
	lcam.Settings.Hue = vc.Get(gocv.VideoCaptureHue)
	lcam.Settings.Gain = vc.Get(gocv.VideoCaptureGain)
	lcam.Settings.Monochrome = vc.Get(gocv.VideoCaptureMonochrome)
	lcam.Settings.Sharpness = vc.Get(gocv.VideoCaptureSharpness)
	lcam.Settings.AutoExposure = vc.Get(gocv.VideoCaptureAutoExposure)
	lcam.Settings.Gamma = vc.Get(gocv.VideoCaptureGain)
	lcam.Settings.Temperature = vc.Get(gocv.VideoCaptureTemperature)
	return
}

func (lcam *LocalCam) BindProperties() (f forms.Forms) {
	f = forms.Forms{
		idFrameWidth: {
			ID:      idFrameWidth,
			Value:   &lcam.Settings.FrameWidth,
			Entries: settingsEntries[idFrameWidth],
		},
		idFrameHeight: {
			ID:      idFrameHeight,
			Value:   &lcam.Settings.FrameHeight,
			Entries: settingsEntries[idFrameHeight],
		},
		idFrameRate: {
			ID:      idFrameRate,
			Value:   &lcam.Settings.FrameRate,
			Entries: settingsEntries[idFrameRate],
		},
		idFormat: {
			ID:      idFormat,
			Value:   &lcam.Settings.Format,
			Entries: settingsEntries[idFormat],
		},
		idMode: {
			ID:      idMode,
			Value:   &lcam.Settings.Mode,
			Entries: settingsEntries[idMode],
		},
		idBrightness: {
			ID:      idBrightness,
			Value:   &lcam.Settings.Brightness,
			Entries: settingsEntries[idBrightness],
		},
		idContrast: {
			ID:      idContrast,
			Value:   &lcam.Settings.Contrast,
			Entries: settingsEntries[idContrast],
		},
		idSaturation: {
			ID:      idSaturation,
			Value:   &lcam.Settings.Saturation,
			Entries: settingsEntries[idSaturation],
		},
		idHue: {
			ID:      idHue,
			Value:   &lcam.Settings.Hue,
			Entries: settingsEntries[idHue],
		},
		idGain: {
			ID:      idGain,
			Value:   &lcam.Settings.Gain,
			Entries: settingsEntries[idGain],
		},
		idMonochrome: {
			ID:      idMonochrome,
			Value:   &lcam.Settings.Monochrome,
			Entries: settingsEntries[idMonochrome],
		},
		idSharpness: {
			ID:      idSharpness,
			Value:   &lcam.Settings.Sharpness,
			Entries: settingsEntries[idSharpness],
		},
		idAutoExposure: {
			ID:      idAutoExposure,
			Value:   &lcam.Settings.AutoExposure,
			Entries: settingsEntries[idAutoExposure],
		},
		idGamma: {
			ID:      idGamma,
			Value:   &lcam.Settings.Gamma,
			Entries: settingsEntries[idGamma],
		},
		idTemperature: {
			ID:      idTemperature,
			Value:   &lcam.Settings.Temperature,
			Entries: settingsEntries[idTemperature],
		},
	}
	return
}
