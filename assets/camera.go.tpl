{{define "camera-bar"}}
<div class="pic-group w3-theme-d1">
    {{range $index, $camera := .}}
        <div class="pic-item w3-animate-opacity">
            {{template "camera-servo" $camera}}
            <div class="w3-bar">
                <label class="w3-bar-item">
                    {{$camera.Title}}
                </label>
                <span class="w3-button w3-bar-item"
                    title="Servos"
                    onclick="openServoSettings({{$index}})">
                    <i class="bi bi-gear-wide-connected"></i>
                </span>
                <span class="w3-button w3-bar-item" 
                    title="Configure"
                    onclick="openCameraSettings({{$index}})">
                    <i class="bi bi-grid-1x2"></i>
                </span>
                <span class="w3-button w3-bar-item"
                    title="Full View"
                    onclick="openCameraFull({{$index}})">
                    <i class="bi bi-fullscreen"></i>
                </span>
            </div>
        </div>
    {{end}}
</div>
{{end}}

{{define "camera-servo"}}
<div class="w3-display-container">
    <img id="{{.ID}}" 
        class="pic w3-image"
        src="/{{.ID}}/mjpeg">
    {{template "pan-tilt" .}}
</div>
{{end}}

{{define "pan-tilt"}}
<div class="pantilt-container">
    {{template "servo-controls" .}}
</div>
{{end}}

{{define "servo-controls"}}
{{$ctls := .ServoControls 20}}
    {{range $index, $ctl := $ctls}}
        <div class="w3-ripple servo-control"
            onmousedown="panTiltDown({{$ctl.ID}},{{$ctl.PanStep}},{{$ctl.TiltStep}},{{$ctl.Speed}})"
            onmouseup="panTiltUp()"
            ontouchstart="panTiltDown({{$ctl.ID}},{{$ctl.PanStep}},{{$ctl.TiltStep}},{{$ctl.Speed}})"
             ontouchend="panTiltUp()">
            <i class="bi {{$ctl.Icon}} iservo"></i>
        </div>
    {{end}}
    {{$len := len .Servos}}
    {{if ge $len 1}}
        {{$svo := index .Servos 0}}
        <input id="{{.ID}}-{{$svo.Index}}" class="hslide" type="range" min="0" max="180" value="90">
    {{end}}
    {{if ge $len 2}}
        {{$svo := index .Servos 1}}
        <input id="{{.ID}}-{{$svo.Index}}" class="vslide" type="range" min="0" max="180" value="90" orient="vertical">
    {{end}}
{{end}}


{{define "camera-settings"}}
<div class="w3-modal-content w3-theme-d5 w3-container">
    <header class="w3-container w3-theme-l1 w3-padding-16">
        <h5 style="display:inline-block">{{.Title}}</h5>
        <span class="w3-bar-item w3-right">
            <button class="w3-btn w3-round w3-left w3-theme-l2"
                onclick="closeDialog()">x</button>
        </span>
    </header>
    <div class="content w3-theme-l3">
        {{template "camera-servo" .}}
        <div class="content" style="display:block;width:35%;">
            {{template "camera-edit" .}}
        </div>
    </div>
    <footer class="w3-container w3-theme-l1">
        <button
            type="button"
            onclick="closeDialog()"
            class="w3-btn w3-round w3-left w3-margin w3-theme-d5">
            Close
        </button>
    </footer>
</div>
{{end}}

{{define "camera-edit"}}
{{$cam:=.}}
<table class="w3-table entries dlgtab">
{{range $id, $form := .CommandForms}}
    {{$value := $form.Value}}
    {{$first := index $form.Entries 0}}
    {{range $i, $ent := $form.Entries}}
    <tr class="">
        {{$f:=$ent.FormatInput $value $first}}
        <td>{{$ent.Label}}</td>
        <td>
            <input
                id="{{$f.ID}}"
                name="{{$f.ID}}"
                class="{{$f.Class}}"
                type="{{$f.Type}}"

                {{if $f.HasChecked}}
                    checked
                {{end}}

                {{if $f.HasRange}}
                    min="{{$ent.Min}}"
                    max="{{$ent.Max}}"
                {{end}}
                
                {{if $f.HasStep}}
                    step="{{$ent.Step}}"
                {{end}}
                value="{{$value}}"/>
        </td>
        <td>
            {{if eq $ent.Type "checkbox"}}
            <button class="w3-button"
                onclick="applyCameraCheckbox({{$cam.ID}},{{$f.ID}})">
                Apply
            </button>
            {{else}}
            <button class="w3-button"
                onclick="applyCameraSetting({{$cam.ID}},{{$f.ID}})">
                Apply
            </button>
            {{end}}
        </td>
    </tr>
    {{end}}
{{end}}
</table>
{{end}}
