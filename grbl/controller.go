package grbl

import (
	"fmt"
	"html/template"
	"io"
	"strings"

	"github.com/centretown/tiny-fabb/forms"
	"github.com/centretown/tiny-fabb/monitor"
	"github.com/golang/glog"
)

type Controller struct {
	ID       string       `json:"id"`
	Title    string       `json:"title"`
	Profile  *Profile     `json:"profile"`
	Active   bool         `json:"active"`
	Port     string       `json:"port"`
	Settings GrblSettings `json:"settings"`
	Commands GrblCommands `json:"-"`

	views    map[string]forms.Forms
	viewList []string
	bus      *monitor.Bus
	layout   *template.Template
}

func NewController(bus *monitor.Bus, layout *template.Template) (gctl *Controller) {
	gctl = &Controller{}
	gctl.initialize(bus, layout)
	return
}

func (gctl *Controller) initialize(bus *monitor.Bus, layout *template.Template) {
	gctl.layout = layout
	gctl.bus = bus
	gctl.views = make(map[string]forms.Forms)
	gctl.views["settings"] = gctl.bindSettings()
	gctl.views["commands"] = gctl.bindCommands()
	gctl.viewList = []string{"settings", "commands", "status"}
}

func (gctl *Controller) Views() []string {
	return gctl.viewList
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

func (gctl *Controller) Apply(viewName, key string, vals map[string][]string) (updated []*forms.Updated, err error) {
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
		id      forms.WebId
		val     string
		form    *forms.Form
		ent     *forms.Entry
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

	if view == "settings" {
		for _, result := range results {
			done, err = isTerminated(result)
			if err != nil {
				glog.Warningln(err)
			}
			if done {
				break
			}

			var count int
			count, err = fmt.Sscanf(result, "$%d=%s", &id, &val)
			if err == nil && count == 2 {
				form, err = gctl.getForm(view, id.String())
				if err == nil {
					ent = form.Entries[0]
					err = ent.ScanInput(val, form.Value)
				}
			}

			if err != nil {
				glog.Warningln(err.Error(), view, ent.Code)
				return
			}
		}
	} else if view == "commands" {
		form.Value = results
	}
	return
}

func (gctl *Controller) Update(form *forms.Form) (err error) {
	var (
		ent  *forms.Entry
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

func (gctl *Controller) startup() (err error) {
	err = gctl.Query("settings", idSettings.String())
	if err != nil {
		return
	}

	cmds := []string{
		idParameters.String(),
		idParserState.String(),
		idBuildInfo.String(),
		idStartupBlocks.String(),
	}
	for _, cmd := range cmds {
		err = gctl.Query("commands", cmd)
		if err != nil {
			return
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

func (gctl *Controller) getView(viewName string) (view forms.Forms, err error) {
	view = gctl.views[viewName]
	if view == nil {
		err = fmt.Errorf("view '%s' not found", viewName)
	}
	return
}

func (gctl *Controller) getForm(viewName, key string) (form *forms.Form, err error) {
	view, err := gctl.getView(viewName)
	if err != nil {
		return
	}

	id := forms.ToWebId(key)
	form, ok := view[id]
	if !ok {
		err = fmt.Errorf("form '%s:%s' not found", viewName, key)
		return
	}

	return
}
