{{define "camera-bar"}}
<div id="camera-bar-container" class="pic-content w3-container w3-border w3-theme-d1">
    {{range $index, $camera := .}}
        <div class="pics-items pic-content w3-animate-opacity">
            {{template "camera-servo" $camera}}
            <div class="w3-bar">
                <span class="w3-bar-item">
                    {{$camera.Title}}
                </span>
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

{{define "camera-servo"}}
<div id="{{.ID}}" class="w3-display-container entries">
    <img id="{{.ID}}-stream" 
        class="pic w3-image"
        src="/{{.ID}}/mjpeg">
    <div class="w3-display-topleft w3-container"
        style="margin-top:16px;">
        <div class="w3-cell-row">
            <div class="w3-container w3-btn w3-ripple w3-cell servo-control"
                onclick="panTilt({{.ID}},30,30,50)">
                <i class="bi bi-arrow-up-left"></i>
            </div>
            <div class="w3-container w3-btn w3-ripple w3-cell servo-control"
                onclick="panTilt({{.ID}},0,30,50)">
                <i class="bi bi-arrow-up"></i>
            </div>
            <div class="w3-container w3-btn w3-ripple w3-cell servo-control"
                onclick="panTilt({{.ID}},-30,30,50)">
                <i class="bi bi-arrow-up-right"></i>
            </div>
        </div>
        <div class="w3-cell-row">
            <div class="w3-container w3-btn w3-ripple w3-cell servo-control"
                onclick="panTilt({{.ID}},30,0,50)">
                <i class="bi bi-arrow-left"></i>
            </div>
            <div class="w3-container w3-cell servo-control">
                <i class="bi bi-arrows-move"></i>
            </div>
            <div class="w3-container w3-btn w3-ripple w3-cell servo-control"
                onclick="panTilt({{.ID}},-30,0,50)">
                <i class="bi bi-arrow-right"></i>
            </div>
        </div>
        <div class="w3-cell-row">
            <div class="w3-container w3-btn w3-ripple w3-cell servo-control"
                onclick="panTilt({{.ID}},30,-30,50)">
                <i class="bi bi-arrow-down-left"></i>
            </div>
            <div class="w3-container w3-btn w3-ripple w3-cell servo-control"
                onclick="panTilt({{.ID}},0,-30,50)">
                <i class="bi bi-arrow-down"></i>
            </div>
            <div class="w3-container w3-btn w3-ripple w3-cell servo-control"
                onclick="panTilt({{.ID}},-30,-30,50)">
                <i class="bi bi-arrow-down-right"></i>
            </div>
        </div>
    </div>
</div>
{{end}}

