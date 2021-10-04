{{define "camera-bar"}}
<div id="camera-bar-container" class="pic-content w3-container w3-border w3-theme-d1">
    {{range $index, $camera := .}}
        <div class="pics-items pic-content w3-animate-opacity">
            <div class="w3-margin w3-padding">{{$camera.Title}}</div>
            <img id={{$index}}
                class="pic w3-image w3-margin" 
                src="/{{$index}}/mjpeg" 
                onclick="openCameraFull({{$index}})">
            <div class="w3-bar">
                <span class="w3-button w3-bar-item"
                    title="Servos"
                    style="padding-right:4px"
                    onclick="openServoSettings({{$index}})">
                    <i class="bi bi-gear-wide-connected"></i>
                </span>
                <span class="w3-button w3-bar-item" 
                    title="Configure"
                    style="padding-right:4px"
                    onclick="openCameraSettings({{$index}})">
                    <i class="bi bi-grid-1x2"></i>
                </span>
                <span class="w3-button w3-bar-item"
                    title="Full View"
                    style="padding-right:4px"
                    onclick="openCameraFull({{$index}})">
                    <i class="bi bi-fullscreen"></i>
                </span>
            </div>
        </div>
    {{end}}
</div>
{{end}}

{{define "camera-full"}}
<div class="full w3-animate-opacity">
    <div class="w3-bar w3-border w3-theme-d3">
        <span class="w3-bar-item w3-left">{{.ID}} Full</span>
        <span class="w3-bar-item w3-right">
            <button class="w3-btn w3-round w3-left w3-theme-l2"
                onclick="closeDialog()">x</button>
        </span>
    </div>
    <div class="fullview">
        <img id="{{.ID}}-stream" 
            class="fullImage" 
            onclick="closeDialog()"
            src="/{{.ID}}/mjpeg">
    </div>
</div>
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
        <div class="entries">
            <img id="{{.ID}}-stream" 
                class="pic w3-image"
                src="/{{.ID}}/mjpeg">
        </div>
        <div class="content">
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
{{range $id, $form := .Forms}}
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

