package main

var test01 = `$0=10	Step pulse, microseconds
$1=25	Step idle delay, milliseconds
$2=0	Step port invert, mask
$3=0	Direction port invert, mask
$4=0	Step enable invert, boolean
$5=0	Limit pins invert, boolean
$6=0	Probe pin invert, boolean
$10=1	Status report, masktest01
$25=500.000	Homing seek, mm/min
$26=250	Homing debounce, milliseconds
$27=1.000	Homing pull-off, mm
$30=1000.	Max spindle speed, RPM
$31=0.	Min spindle speed, RPM
$32=1	Laser mode, boolean
$100=250.000	X steps/mm
$101=250.000	Y steps/mm
$102=250.000	Z steps/mm
$110=500.000	X Max rate, mm/min
$111=500.000	Y Max rate, mm/min
$112=500.000	Z Max rate, mm/min
$120=10.000	X Acceleration, mm/sec^2
$121=10.000	Y Acceleration, mm/sec^2
$122=10.000	Z Acceleration, mm/sec^2
$130=200.000	X Max travel, mm
$131=200.000	Y Max travel, mm
$132=200.000	Z Max travel, mm
`

// func testScan(t *testing.T) {
// 	g := grbl.NewController(nil)
// 	scanSettings(g, test01)
// 	fmt.Println(g.grblSettings)
// }

// func scanSettings(g *grbl.GrblController, in string) {
// 	var (
// 		key, idx, count int
// 		err             error
// 	)
// 	for s := in; len(s) > 0; {
// 		// scan the key
// 		count, err = fmt.Sscanf(s, "$%d=", &key)
// 		if err == nil && count == 1 {
// 			// find the value and scan
// 			setting, ok := g.Settings[key]
// 			if ok {
// 				v := setting.Value
// 				fmt.Sscanf(s, "$%d=%v", &key, v)
// 			}
// 		}

// 		// next line
// 		idx = strings.Index(s, "\n")
// 		if idx == -1 {
// 			break
// 		}
// 		s = s[idx+1:]
// 	}
// }
