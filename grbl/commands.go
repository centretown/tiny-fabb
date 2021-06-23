package grbl

// var scanSetting = func(g *GrblController, s string) (err error) {
// 	err = g.scanSetting(s)
// 	return
// }

// var scanInfo = func(g *GrblController, s string) (err error) {
// 	err = g.scanInfo(s)
// 	return
// }

type BuildInfo struct {
	Version  string
	Options  string
	Messages []string
}

type GCodeParameters struct {
	G54 [3]float32
	G55 [3]float32
	G56 [3]float32
	G57 [3]float32
	G58 [3]float32
	G59 [3]float32
	G28 [3]float32
	G30 [3]float32
	TLB float32
}

type GrblCommands struct {
	Settings        string `json:"settings"`
	Parameters      string `json:"parameters"`
	ParserState     string `json:"parserState"`
	BuildInfo       string `json:"buildInfo"`
	StartupBlocks   string `json:"startupBlocks"`
	CodeMode        string `json:"codeMode"`
	KillAlarm       string `json:"killAlarm"`
	RunHomingCycle  string `json:"runHomingCycle"`
	RunJoggingCycle string `json:"runJoggingCycle"`
	EraseRestore    string `json:"eraseRestore"`
	EraseZero       string `json:"eraseZero"`
	ClearRestore    string `json:"clearRestore"`
}
