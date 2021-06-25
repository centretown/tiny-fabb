package grbl

import (
	"fmt"
	"html/template"
	"io"
	"strings"

	"github.com/centretown/tiny-fabb/monitor"
	"github.com/centretown/tiny-fabb/serialio"
	"github.com/centretown/tiny-fabb/web"
	"github.com/golang/glog"
)

// Controller are defined as $nnn=value
type Controller struct {
	Title    string       `json:"title"`
	Active   bool         `json:"active"`
	Port     string       `json:"port"`
	Version  string       `json:"version"`
	Build    string       `json:"build"`
	Settings GrblSettings `json:"settings"`
	Commands GrblCommands `json:"commands"`

	views  map[string]web.Forms
	bus    *monitor.Bus
	layout *template.Template
}

// NewController returns a newly minted GRBL controller
func NewController(layout *template.Template) (gctl *Controller) {
	gctl = &Controller{
		layout: layout,
	}

	gctl.views = make(map[string]web.Forms)
	gctl.views["settings"] = gctl.bindSettings()
	gctl.views["commands"] = gctl.bindCommands()

	return
}

func (gctl *Controller) Describe() *web.Descriptor {
	d := &web.Descriptor{
		Title:       gctl.Title,
		Description: "GRBL",
		Port:        gctl.Port,
		Active:      gctl.Active,
		Version:     gctl.Version + " " + gctl.Build,
	}
	return d
}

func (gctl *Controller) ListViews() []string {
	return []string{"Settings", "Commands", "Status"}
}

func (gctl *Controller) Upload(w io.Writer, files []string) (err error) {
	return
}

func (gctl *Controller) List(w io.Writer, viewName string) (err error) {
	tmpl, err := gctl.getTemplate("list")
	if err != nil {
		return
	}

	view, err := gctl.getView(viewName)
	if err != nil {
		return
	}
	err = tmpl.Execute(w, view)
	return
}

func (gctl *Controller) Edit(w io.Writer, viewName, key string) (err error) {
	tmpl, err := gctl.getTemplate("edit")
	if err != nil {
		return
	}

	form, err := gctl.getForm(viewName, key)
	if err != nil {
		return
	}
	err = tmpl.Execute(w, form)
	return
}

func (gctl *Controller) Apply(viewName, key string, vals map[string][]string) (updated []*web.Updated, err error) {
	form, err := gctl.getForm(viewName, key)
	if err != nil {
		return
	}

	if len(form.Entries) < 1 {
		return
	}

	err = form.Update(vals)
	if err != nil {
		return
	}

	if gctl.Active {
		err = gctl.Update(form)
		if err != nil {
			return
		}
	}

	updated = form.GetUpdated()
	return
}

const (
	termOK  = "ok"
	termErr = "error"
)

var terminators = []string{termOK, termErr}

func isTerminated(result string) (terminated bool, err error) {
	if strings.HasPrefix(result, termOK) {
		terminated = true
	}
	if strings.HasPrefix(result, termErr) {
		terminated = true
		err = fmt.Errorf("%s", result)
		var errcode uint
		count, _ := fmt.Sscanf(result, "error:%d", &errcode)
		fmt.Println(count, errcode)
		if count == 1 {
			s, ok := GrblErrors[errcode]
			if ok {
				err = fmt.Errorf("error: %s", s)
			}
		}
	}
	return
}

func (gctl *Controller) Query(view string, key string) (err error) {
	var (
		id      web.WebId
		val     string
		form    *web.Form
		ent     *web.Entry
		results []string
		done    bool
	)

	form, err = gctl.getForm(view, key)
	if err != nil {
		return err
	}

	ent = form.Entries[0]
	results, err = gctl.bus.Capture(ent.Code, terminators...)
	if err != nil {
		return err
	}
	// fmt.Println(results)
	for _, result := range results {
		done, err = isTerminated(result)
		if err != nil {
			glog.Warningln(err)
		}
		if done {
			break
		}

		switch view {
		case "settings":
			var count int
			count, err = fmt.Sscanf(result, "$%d=%s", &id, &val)
			if err == nil && count == 2 {
				form, err = gctl.getForm(view, id.String())
				if err == nil {
					ent = form.Entries[0]
					err = ent.ScanInput(val, form.Value)
				}
			}

		case "commands":
			val += result
		}

		if err != nil {
			glog.Warningln(err.Error(), view, ent.Code)
			return
		}
	}

	if view == "commands" {
		// fmt.Println(val)
		form.Value = &val
	}
	return
}

func (gctl *Controller) Update(form *web.Form) (err error) {
	var (
		ent  *web.Entry
		val  interface{}
		code string
	)

	ent = form.Entries[0]
	val = ent.Value(form.Value)
	switch t := val.(type) {
	case bool:
		var val int = 0
		if t {
			val = 1
		}
		code = fmt.Sprintf("%s=%v", ent.Code, val)
	default:
		code = fmt.Sprintf("%s=%v", ent.Code, val)
	}

	if !gctl.Active {
		err = fmt.Errorf("error updating '%s:' %v is inactive",
			code, gctl.Title)
		return
	}

	var (
		result  string
		results []string
		done    bool
	)
	results, err = gctl.bus.Capture(code, terminators...)
	for _, result = range results {
		done, err = isTerminated(result)
		if done {
			break
		}
	}

	return
}

func (gctl *Controller) getTemplate(tmplName string) (tmpl *template.Template, err error) {
	tmpl = gctl.layout.Lookup(tmplName)
	if tmpl == nil {
		err = fmt.Errorf("template '%s' not found", tmplName)
	}
	return
}

func (gctl *Controller) getView(viewName string) (view web.Forms, err error) {
	view = gctl.views[viewName]
	if view == nil {
		err = fmt.Errorf("view '%s' not found", viewName)
	}
	return
}

func (gctl *Controller) getForm(viewName, key string) (form *web.Form, err error) {
	view, err := gctl.getView(viewName)
	if err != nil {
		return
	}

	id := web.ToWebId(key)
	form, ok := view[id]
	if !ok {
		err = fmt.Errorf("form '%s:%s' not found", viewName, key)
		return
	}

	return
}

func (gctl *Controller) ActivateController() (err error) {
	if gctl.Active {
		err = fmt.Errorf("%s already active", gctl.Title)
		return
	}

	var sio *serialio.SerialIO
	sio, err = serialio.GetSerialIO(gctl.Port)
	if err != nil {
		return
	}

	gctl.bus = monitor.NewBus()
	go monitor.Monitor(sio, gctl.bus)
	gctl.Active = true
	glog.Infof("Monitoring %v: %v...\n", gctl.Title, gctl.Port)

	err = gctl.Query("settings", idSettings.String())
	return
}

// func DeactivateController(g *grbl.GrblController) (err error) {
// 	if !g.Active {
// 		err = fmt.Errorf("%s not active", g.Title)
// 		return
// 	}

// 	g.Bus.Done <- true
// 	time.Sleep(time.Millisecond)
// 	g.Active = false

// 	return
// }

// func (gctl *Controller) scanInfo(s string) (err error) {

// 	const (
// 		ver = "[VER:"
// 		opt = "[OPT:"
// 		sfx = "]"
// 	)

// 	fmt.Println("scanInfo", s)
// 	if strings.HasPrefix(s, ver) {
// 		s = strings.TrimPrefix(s, ver)
// 		gctl.BuildVer = strings.TrimSuffix(s, sfx)
// 	} else if strings.HasPrefix(s, opt) {
// 		s = strings.TrimPrefix(s, opt)
// 		gctl.BuildOpt = strings.TrimSuffix(s, sfx)
// 	}
// 	return
// }

// func (gctl *Controller) UpdateSetting(src map[string][]string) (err error) {
// 	fmt.Println(src)
// 	var key GrblId
// 	for k, values := range src {
// 		for _, value := range values {
// 			fmt.Println(k, value)
// 			fmt.Sscanf(k, "%d", &key)
// 			setting, ok := gctl.Settings[key]
// 			if !ok {
// 				err = fmt.Errorf("key not found: %d", key)
// 				return
// 			}
// 			err = settingctl.Update(value)
// 			if err != nil {
// 				return
// 			}
// 			topic := fmt.Sprintf("$%d=%s", key, value)
// 			fmt.Println(topic)
// 			// gctl.Capture(topic)
// 		}
// 	}

// 	return
// }

// func (gctl *Controller) scanSetting(s string) (err error) {
// 	var (
// 		key   GrblId
// 		count int
// 	)
// 	count, err = fmt.Sscanf(s, "$%d=", &key)
// 	if err != nil {
// 		return
// 	}

// 	if count != 1 {
// 		err = fmt.Errorf("scanSetting invalid count: %d", count)
// 		return
// 	}

// 	setting, ok := gctl.Settings[key]
// 	if ok {
// 		v := gctl.Values[key]
// 		_, err = fmt.Sscanf(s, "$%d=%v", &key, v)
// 	}
// 	return
// }

// // SettingValue -
// func (gctl *Controller) SettingValue(key int) (s string) {
// 	v := gctl.Settings[key].Value
// 	switch t := v.(type) {
// 	case *uint:
// 		s = fmt.Sprintf("%d", *t)
// 	case *float32:
// 		s = fmt.Sprintf("%f", *t)
// 	case *bool:
// 		s = fmt.Sprintf("%v", *t)
// 	}
// 	return
// }

// // SettingValues -
// // func (gctl *Controller) SettingValues() (s []string) {
// // 	s = make([]string, len(ordered))
// // 	for j, i := range ordered {
// // 		s[j] = fmt.Sprintf("($%d) %s = %s",
// // 			i, settingEntries[j].Label, gctl.SettingValue(i))
// // 	}
// // 	return
// // }

// // Capture -
// func (gctl *Controller) Capture(topic GrblId) {
// 	c, ok := commandEntries[topic]
// 	if ok {
// 		gctl.Bus.Out <- monitor.SerReq{Topic: c.Code, Immediate: c.Type == "immediate"}
// 	} else {
// 		gctl.Bus.Out <- monitor.SerReq{Topic: c.Code}
// 	}

// 	for {
// 		select {
// 		case s := <-gctl.Bus.In:
// 			s = strings.TrimSpace(s)
// 			if strings.HasPrefix(s, "ok") {
// 				return
// 			}

// 			if strings.HasPrefix(s, "error") {
// 				glogctl.Warningln(s)
// 				return
// 			}
// 			c.Scan(g, s)
// 		default:
// 			time.Sleep(time.Millisecond)
// 		}
// 	}
// }
