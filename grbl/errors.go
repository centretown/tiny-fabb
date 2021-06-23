package grbl

const (
	Ok                          uint = 0
	ExpectedCommandLetter       uint = 1
	BadNumberFormat             uint = 2
	InvalidStatement            uint = 3
	NegativeValue               uint = 4
	SettingDisabled             uint = 5
	SettingStepPulseMin         uint = 6
	SettingReadFail             uint = 7
	IdleError                   uint = 8
	SystemGcLock                uint = 9
	SoftLimitError              uint = 10
	Overflow                    uint = 11
	MaxStepRateExceeded         uint = 12
	CheckDoor                   uint = 13
	LineLengthExceeded          uint = 14
	TravelExceeded              uint = 15
	InvalidJogCommand           uint = 16
	SettingDisabledLaser        uint = 17
	HomingNoCycles              uint = 18
	GcodeUnsupportedCommand     uint = 20
	GcodeModalGroupViolation    uint = 21
	GcodeUndefinedFeedRate      uint = 22
	GcodeCommandValueNotInteger uint = 23
	GcodeAxisCommandConflict    uint = 24
	GcodeWordRepeated           uint = 25
	GcodeNoAxisWords            uint = 26
	GcodeInvalidLineNumber      uint = 27
	GcodeValueWordMissing       uint = 28
	GcodeUnsupportedCoordSys    uint = 29
	GcodeG53InvalidMotionMode   uint = 30
	GcodeAxisWordsExist         uint = 31
	GcodeNoAxisWordsInPlane     uint = 32
	GcodeInvalidTarget          uint = 33
	GcodeArcRadiusError         uint = 34
	GcodeNoOffsetsInPlane       uint = 35
	GcodeUnusedWords            uint = 36
	GcodeG43DynamicAxisError    uint = 37
	GcodeMaxValueExceeded       uint = 38
	PParamMaxExceeded           uint = 39
	FsFailedMount               uint = 60 // SD Failed to mount
	FsFailedRead                uint = 61 // SD Failed to read file
	FsFailedOpenDir             uint = 62 // SD card failed to open directory
	FsDirNotFound               uint = 63 // SD Card directory not found
	FsFileEmpty                 uint = 64 // SD Card file empty
	FsFileNotFound              uint = 65 // SD Card file not found
	FsFailedOpenFile            uint = 66 // SD card failed to open file
	FsFailedBusy                uint = 67 // SD card is busy
	FsFailedDelDir              uint = 68
	FsFailedDelFile             uint = 69
	BtFailBegin                 uint = 70 // Bluetooth failed to start
	WifiFailBegin               uint = 71 // WiFi failed to start
	NumberRange                 uint = 80 // Setting number range problem
	InvalidValue                uint = 81 // Setting string problem
	MessageFailed               uint = 90
	NvsSetFailed                uint = 100
	NvsGetStatsFailed           uint = 101
	AuthenticationFailed        uint = 110
	Eol                         uint = 111
	AnotherInterfaceBusy        uint = 120
	JogCancelled                uint = 130
)

// GrblErrors is based on https://grblminicnc.blogspot.com/2017/04/grbl-error-list.html
var GrblErrors = map[uint]string{
	ExpectedCommandLetter:       "Expected command letter G-code words consist of a letter and a value. Letter was not found.",
	BadNumberFormat:             "Bad number format Missing the expected G-code word value or numeric value format is not valid.",
	InvalidStatement:            "Invalid statement Grbl '$' system command was not recognized or supported.",
	NegativeValue:               "Value < 0 Negative value received for an expected positive value.",
	SettingDisabled:             "Setting disabled Homing cycle failure. Homing is not enabled via settings.",
	SettingStepPulseMin:         "Value < 3 usec Minimum step pulse time must be greater than 3usec.",
	SettingReadFail:             "EEPROM read fail. Using defaults An EEPROM read failed. Auto-restoring affected EEPROM to default values.",
	IdleError:                   "Not idle Grbl '$' command cannot be used unless Grbl is IDLE. Ensures smooth operation during a job.",
	SystemGcLock:                "G-code lock G-code commands are locked out during alarm or jog state.",
	SoftLimitError:              "Homing not enabled Soft limits cannot be enabled without homing also enabled.",
	Overflow:                    "Line overflow Max characters per line exceeded. Received command line was not executed.",
	MaxStepRateExceeded:         "Step rate > 30kHz Grbl '$' setting value cause the step rate to exceed the maximum supported.",
	CheckDoor:                   "Check Door Safety door detected as opened and door state initiated.",
	LineLengthExceeded:          "Line length exceeded Build info or startup line exceeded EEPROM line length limit. Line not stored.",
	TravelExceeded:              "Travel exceeded Jog target exceeds machine travel. Jog command has been ignored.",
	InvalidJogCommand:           "Invalid jog command Jog command has no '=' or contains prohibited g-code.",
	SettingDisabledLaser:        "Setting disabled Laser mode requires PWM output.",
	HomingNoCycles:              "Unsupported command Unsupported or invalid g-code command found in block.",
	GcodeUnsupportedCommand:     "Modal group violation More than one g-code command from same modal group found in block.",
	GcodeUndefinedFeedRate:      "Undefined feed rate Feed rate has not yet been set or is undefined.",
	GcodeCommandValueNotInteger: "Invalid gcode ID:23 G-code command in block requires an integer value.",
	GcodeAxisCommandConflict:    "Invalid gcode ID:24 More than one g-code command that requires axis words found in block.",
	GcodeWordRepeated:           "Invalid gcode ID:25 Repeated g-code word found in block.",
	GcodeNoAxisWords:            "Invalid gcode ID:26 No axis words found in block for g-code command or current modal state which requires them.",
	GcodeInvalidLineNumber:      "Invalid gcode ID:27 Line number value is invalid.",
	GcodeValueWordMissing:       "Invalid gcode ID:28 G-code command is missing a required value word.",
	GcodeUnsupportedCoordSys:    "Invalid gcode ID:29 G59.x work coordinate systems are not supported.",
	GcodeG53InvalidMotionMode:   "Invalid gcode ID:30 G53 only allowed with G0 and G1 motion modes.",
	GcodeAxisWordsExist:         "Invalid gcode ID:31 Axis words found in block when no command or current modal state uses them.",
	GcodeNoAxisWordsInPlane:     "Invalid gcode ID:32 G2 and G3 arcs require at least one in-plane axis word.",
	GcodeInvalidTarget:          "Invalid gcode ID:33 Motion command target is invalid.",
	GcodeArcRadiusError:         "Invalid gcode ID:34 Arc radius value is invalid.",
	GcodeNoOffsetsInPlane:       "Invalid gcode ID:35 G2 and G3 arcs require at least one in-plane offset word.",
	GcodeUnusedWords:            "Invalid gcode ID:36 Unused value words found in block.",
	GcodeG43DynamicAxisError:    "Invalid gcode ID:37 G43.1 dynamic tool length offset is not assigned to configured tool length axis.",
	GcodeMaxValueExceeded:       "Invalid gcode ID:38 Tool number greater than max supported value.",
	PParamMaxExceeded:           "Maximum parameters exceeded.",
	FsFailedMount:               "SD Failed to mount.",
	FsFailedRead:                "SD Failed to read file.",
	FsFailedOpenDir:             "SD card failed to open directory.",
	FsDirNotFound:               "SD Card directory not found.",
	FsFileEmpty:                 "SD Card file empty.",
	FsFileNotFound:              "SD Card file not found.",
	FsFailedOpenFile:            "SD card failed to open file.",
	FsFailedBusy:                "SD card is busy.",
	FsFailedDelDir:              "SD card failed to delete directory.",
	FsFailedDelFile:             "SD card failed to delete file.",
	BtFailBegin:                 "Bluetooth failed to start.",
	WifiFailBegin:               "WiFi failed to start.",
	NumberRange:                 "Setting out of range.",
	InvalidValue:                "Setting invalid",
	MessageFailed:               "Message failed.",
	NvsSetFailed:                "NVS set failed.",
	NvsGetStatsFailed:           "NVS get stats failed.",
	AuthenticationFailed:        "Authentication failed.",
	Eol:                         "End of line.",
	AnotherInterfaceBusy:        "Another interface busy",
	JogCancelled:                "Jog cancelled.",
}
