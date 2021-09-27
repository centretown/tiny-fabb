package camera

import "github.com/centretown/tiny-fabb/forms"

func (cam *Camera) bindSettings() {
	cam.Forms = forms.Forms{
		idFramesize: {
			ID:      idFramesize,
			Value:   &cam.Settings.Framesize,
			Entries: settingsEntries[idFramesize],
		},
		idQuality: {
			ID:      idQuality,
			Value:   &cam.Settings.Quality,
			Entries: settingsEntries[idQuality],
		},
		idBrightness: {
			ID:      idBrightness,
			Value:   &cam.Settings.Brightness,
			Entries: settingsEntries[idBrightness],
		},
		idContrast: {
			ID:      idContrast,
			Value:   &cam.Settings.Contrast,
			Entries: settingsEntries[idContrast],
		},
		idSaturation: {
			ID:      idSaturation,
			Value:   &cam.Settings.Saturation,
			Entries: settingsEntries[idSaturation],
		},
		idSharpness: {
			ID:      idSharpness,
			Value:   &cam.Settings.Sharpness,
			Entries: settingsEntries[idSharpness],
		},
		idDenoise: {
			ID:      idDenoise,
			Value:   &cam.Settings.Denoise,
			Entries: settingsEntries[idDenoise],
		},
		idSpecialEffect: {
			ID:      idSpecialEffect,
			Value:   &cam.Settings.SpecialEffect,
			Entries: settingsEntries[idSpecialEffect],
		},
		idWbMode: {
			ID:      idWbMode,
			Value:   &cam.Settings.WbMode,
			Entries: settingsEntries[idWbMode],
		},
		idAwb: {
			ID:      idAwb,
			Value:   &cam.Settings.Awb,
			Entries: settingsEntries[idAwb],
		},
		idAwbGain: {
			ID:      idAwbGain,
			Value:   &cam.Settings.AwbGain,
			Entries: settingsEntries[idAwbGain],
		},
		idAec: {
			ID:      idAec,
			Value:   &cam.Settings.Aec,
			Entries: settingsEntries[idAec],
		},
		idAec2: {
			ID:      idAec2,
			Value:   &cam.Settings.Aec2,
			Entries: settingsEntries[idAec2],
		},
		idAeLevel: {
			ID:      idAeLevel,
			Value:   &cam.Settings.AeLevel,
			Entries: settingsEntries[idAeLevel],
		},
		idAecValue: {
			ID:      idAecValue,
			Value:   &cam.Settings.AecValue,
			Entries: settingsEntries[idAecValue],
		},
		idAgc: {
			ID:      idAgc,
			Value:   &cam.Settings.Agc,
			Entries: settingsEntries[idAgc],
		},
		idAgcGain: {
			ID:      idAgcGain,
			Value:   &cam.Settings.AgcGain,
			Entries: settingsEntries[idAgcGain],
		},
		idGainCeiling: {
			ID:      idGainCeiling,
			Value:   &cam.Settings.GainCeiling,
			Entries: settingsEntries[idGainCeiling],
		},
		idBpc: {
			ID:      idBpc,
			Value:   &cam.Settings.Bpc,
			Entries: settingsEntries[idBpc],
		},
		idWpc: {
			ID:      idWpc,
			Value:   &cam.Settings.Wpc,
			Entries: settingsEntries[idWpc],
		},
		idRawGma: {
			ID:      idRawGma,
			Value:   &cam.Settings.RawGma,
			Entries: settingsEntries[idRawGma],
		},
		idLenc: {
			ID:      idLenc,
			Value:   &cam.Settings.Lenc,
			Entries: settingsEntries[idLenc],
		},
		idHmirror: {
			ID:      idHmirror,
			Value:   &cam.Settings.Hmirror,
			Entries: settingsEntries[idHmirror],
		},
		idVflip: {
			ID:      idVflip,
			Value:   &cam.Settings.Vflip,
			Entries: settingsEntries[idVflip],
		},
		idDcw: {
			ID:      idDcw,
			Value:   &cam.Settings.Dcw,
			Entries: settingsEntries[idDcw],
		},
		idColorbar: {
			ID:      idColorbar,
			Value:   &cam.Settings.Colorbar,
			Entries: settingsEntries[idColorbar],
		},
	}
}
