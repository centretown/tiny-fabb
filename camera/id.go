package camera

import "github.com/centretown/tiny-fabb/forms"

const (
	idBase forms.WebId = iota + 5000
	idFramesize
	idQuality
	idBrightness
	idContrast
	idSaturation
	idSharpness
	idDenoise
	idSpecialEffect
	idWbMode
	idAwb
	idAwbGain
	idAec
	idAec2
	idAeLevel
	idAecValue
	idAgc
	idAgcGain
	idGainCeiling
	idBpc
	idWpc
	idRawGma
	idLenc
	idHmirror
	idVflip
	idDcw
	idColorbar
)

const (
	idFrameWidth forms.WebId = iota + 5100
	idFrameHeight
	idFrameRate
	idFormat
	idMode
	// idBrightness
	// idContrast
	// idSaturation
	idHue
	idGain
	idMonochrome
	// idSharpness
	idAutoExposure
	idGamma
	idTemperature
)
