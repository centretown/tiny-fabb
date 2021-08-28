package camera

import "github.com/centretown/tiny-fabb/forms"

var settingsEntries = map[forms.WebId]forms.Entries{
	idFramesize: {
		{
			ID:    idFramesize.String(),
			Code:  "framesize",
			Label: "Framesize",
			Type:  "number",
			Min:   0,
			Max:   10,
			Step:  1,
		},
	},
	idQuality: {
		{
			ID:    idQuality.String(),
			Code:  "quality",
			Label: "Quality",
			Type:  "number",
			Min:   0,
			Max:   63,
			Step:  1,
		},
	},
	idBrightness: {
		{
			ID:    idBrightness.String(),
			Code:  "brightness",
			Label: "Brightness",
			Type:  "number",
			Min:   -2,
			Max:   2,
			Step:  1,
		},
	},
	idContrast: {
		{
			ID:    idContrast.String(),
			Code:  "contrast",
			Label: "Contrast",
			Type:  "number",
			Min:   -2,
			Max:   2,
			Step:  1,
		},
	},
	idSaturation: {
		{
			ID:    idSaturation.String(),
			Code:  "saturation",
			Label: "Saturation",
			Type:  "number",
			Min:   -2,
			Max:   2,
			Step:  1,
		},
	},
	idSharpness: {
		{
			ID:    idSharpness.String(),
			Code:  "sharpness",
			Label: "Sharpness",
			Type:  "number",
			Min:   -2,
			Max:   2,
			Step:  1,
		},
	},
	idDenoise: {
		{
			ID:    idDenoise.String(),
			Code:  "denoise",
			Label: "Denoise",
			Type:  "checkbox",
		},
	},
	idSpecialEffect: {
		{
			ID:    idSharpness.String(),
			Code:  "special_effect",
			Label: "Special Effect",
			Type:  "number",
			Min:   0,
			Max:   6,
			Step:  1,
		},
	},
	idWbMode: {
		{
			ID:    idWbMode.String(),
			Code:  "wb_mode",
			Label: "WbMode",
			Type:  "number",
			Min:   0,
			Max:   4,
			Step:  1,
		},
	},
	idAwb: {
		{
			ID:    idAwb.String(),
			Code:  "awb",
			Label: "Awb",
			Type:  "checkbox",
		},
	},
	idAwbGain: {
		{
			ID:    idAwbGain.String(),
			Code:  "awb_gain",
			Label: "AwbGain",
			Type:  "checkbox",
		},
	},
	idAec: {
		{
			ID:    idAec.String(),
			Code:  "aec",
			Label: "Aec",
			Type:  "checkbox",
		},
	},
	idAec2: {
		{
			ID:    idAec2.String(),
			Code:  "aec2",
			Label: "Aec2",
			Type:  "checkbox",
		},
	},
	idAeLevel: {
		{
			ID:    idAeLevel.String(),
			Code:  "ae_level",
			Label: "Ae Level",
			Type:  "number",
			Min:   -2,
			Max:   2,
			Step:  1,
		},
	},
	idAecValue: {
		{
			ID:    idAecValue.String(),
			Code:  "aec_value",
			Label: "Aec Value",
			Type:  "number",
			Min:   0,
			Max:   1200,
			Step:  10,
		},
	},
	idAgc: {
		{
			ID:    idAgc.String(),
			Code:  "agc",
			Label: "Agc",
			Type:  "checkbox",
		},
	},
	idAgcGain: {
		{
			ID:    idAgcGain.String(),
			Code:  "agc_gain",
			Label: "Agc Gain",
			Type:  "number",
			Min:   0,
			Max:   30,
			Step:  1,
		},
	},
	idGainCeiling: {
		{
			ID:    idGainCeiling.String(),
			Code:  "gainceiling",
			Label: "Gain Ceiling",
			Type:  "number",
			Min:   0,
			Max:   6,
			Step:  1,
		},
	},
	idBpc: {
		{
			ID:    idBpc.String(),
			Code:  "bpc",
			Label: "Bpc",
			Type:  "checkbox",
		},
	},
	idWpc: {
		{
			ID:    idWpc.String(),
			Code:  "wpc",
			Label: "Wpc",
			Type:  "checkbox",
		},
	},
	idRawGma: {
		{
			ID:    idRawGma.String(),
			Code:  "raw_gma",
			Label: "Raw GMA",
			Type:  "checkbox",
		},
	},
	idLenc: {
		{
			ID:    idLenc.String(),
			Code:  "lenc",
			Label: "Lenc",
			Type:  "checkbox",
		},
	},
	idHmirror: {
		{
			ID:    idHmirror.String(),
			Code:  "hmirror",
			Label: "Mirror horizontally",
			Type:  "checkbox",
		},
	},
	idVflip: {
		{
			ID:    idVflip.String(),
			Code:  "vflip",
			Label: "Flip vertically",
			Type:  "checkbox",
		},
	},
	idDcw: {
		{
			ID:    idDcw.String(),
			Code:  "dcw",
			Label: "Dcw",
			Type:  "checkbox",
		},
	},
	idColorbar: {
		{
			ID:    idColorbar.String(),
			Code:  "colorbar",
			Label: "Colorbar",
			Type:  "checkbox",
		},
	},
}
