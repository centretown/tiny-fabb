{{define "select-bar"}}
<div class="w3-bar-item w3-theme-d3 w3-border">
    <select
        id="selected-controller"
        name="selected-controller"
        class="w3-select w3-theme-d3 w3-bar-item"
        onchange="selectController();">

        {{range $index, $element := .Controllers}}
            <option value="{{$index}}">
            {{.Title}} {{.Port}}
            </option>
        {{end}}
    </select>
</div>
{{end}}

{{define "button-bar"}}
<div id="button-bar" class="w3-bar w3-border w3-theme-d3">
    <div class="w3-bar-item w3-left">
        <button
            href="#" 
            class="w3-button"
            title="Upload">
            <i class="bi bi-upload"></i>
        </button>

        <button 
            id="check-camera" 
            class="w3-button"
            name="check-camera" 
            title="Cameras"
            onclick="selectCamera()">
            <i class="bi bi-camera"></i>
        </button>
    </div>

    <div class="radio-toolbar w3-bar-item w3-right">
       <input 
            type="radio" 
            id="radio-commands" 
            name="radio-view" 
            value="commands"
            title="Commands"
           onchange="selectView(value)">
        <label for="radio-commands">
            <i class="bi bi-command"></i>
        </label>

        <input 
            type="radio" 
            id="radio-settings" 
            name="radio-view" 
            value="settings"
            onchange="selectView(value)"
            title="Settings"
            checked>
        <label for="radio-settings">
            <i class="bi bi-grid"></i>
        </label>
    </div>
</div>
{{end}}

{{define "app-bar"}}
<div id="app-bar" class="w3-bar w3-border w3-theme-d3">
    <div class="w3-theme-d3">
        <select
            id="selected-controller"
            name="selected-controller"
            class="w3-theme-d3 select-controller w3-padding"
            onchange="selectController();">

            {{range $index, $element := .Controllers}}
                <option value="{{$index}}">
                {{.Title}} {{.Port}}
                </option>
            {{end}}
        </select>
    </div>
    <div class="w3-dropdown-hover w3-left w3-theme-d3 w3-bar-item">
        <button class="w3-button app-icon">
            <i class="bi bi-menu-button"></i>
        </button>
        <div id="menu-left" 
            class="w3-dropdown-content w3-bar-block w3-border w3-theme-d1"
            style="left:0">

            <button class="w3-bar-item w3-button" onclick="openOptions()">
                <i class="bi bi-sliders"></i>
                <label>Options</label>
            </button>
            <a href="#" class="w3-bar-item w3-button">
                <i class="bi bi-upload"></i>
                <label>Upload</label>
            </a>
            <a href="#" class="w3-bar-item w3-button">
                <i class="bi bi-alarm"></i>
                <label>Alarms</label>
            </a>
        </div>
    </div>
    <div class="w3-right w3-dropdown-hover w3-bar-item w3-theme-d3">
        <button class="w3-button app-icon">
            <i class="bi bi-person"></i>
        </button>
        <div id="menu-right" 
            class="w3-dropdown-content w3-bar-block w3-border w3-theme-d1" 
            style="right:0">

            <a href="#" class="w3-bar-item w3-button">
                <i class="bi bi-box-arrow-in-left"></i>
                <label>Sign In</label>
            </a>
            <a href="#" class="w3-bar-item w3-button">
                <i class="bi bi-shield-lock"></i>
                <label>Lock</label>
            </a>
        </div>
    </div>
</div>
{{end}}

{{define "camera-bar"}}
<div id="camera-bar-container" class="pic-content w3-container w3-border w3-theme-d1">
    {{range $index, $camera := .}}
        <div class="pics-items pic-content w3-animate-opacity">
            <div class="w3-margin w3-padding">{{$camera.Title}}</div>
            <img id={{$index}}
                class="pic w3-image w3-margin w3-padding" 
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
        <span class="w3-bar-item w3-left">{{.Name}} Full</span>
        <span class="w3-bar-item w3-right">
            <button class="w3-btn w3-round w3-left w3-theme-l2"
                onclick="closeDialog()">x</button>
        </span>
    </div>
    <div class="fullview">
        <img id="{{.Name}}-stream" 
            class="fullImage" 
            onclick="closeDialog()"
            src="/{{.Name}}/mjpeg">
    </div>
</div>
{{end}}

{{define "servo-settings"}}
<div class="full w3-animate-opacity">
    <div class="w3-bar w3-border w3-theme-d3">
        <span class="w3-bar-item w3-left">{{.Name}} Settings</span>
        <span class="w3-bar-item w3-right">
            <button class="w3-btn w3-round w3-left w3-theme-l2"
                onclick="closeDialog()">x</button>
        </span>
    </div>
    <div class="fullview">
        <img id="{{.Name}}-stream" 
            class="w3-image w3-center" 
            src="/{{.Name}}/mjpeg">
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
                <img id="{{.Name}}-stream" 
                    class="w3-image w3-padding"
                    src="/{{.Name}}/mjpeg">
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
                    name="{{$f.Name}}"
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
                    onclick="applyCameraCheckbox({{$cam.Name}},{{$f.ID}})">
                    Apply
                </button>
                {{else}}
                <button class="w3-button"
                    onclick="applyCameraSetting({{$cam.Name}},{{$f.ID}})">
                    Apply
                </button>
                {{end}}
            </td>
        </tr>
        {{end}}
    {{end}}
    </table>
{{end}}

{{define "options"}}
{{$themeColor := .Color}}
<div class="w3-modal-content">
    <header class="w3-container w3-theme-d1 w3-padding-16">
        <span>Options</span>
        <span class="w3-right">
            <button class="w3-btn w3-round w3-left w3-theme-l2"
                onclick="closeOptions()">x</button>
        </span>
    </header>
<br />
    <form class="w3-container">
        <label class="w3-margin-top w3-text" for="selected-theme">
            Theme
        </label>
        <select
            id="selected-theme"
            name="selected-theme"
            onchange=""
            value="{{$themeColor}}"
            class="w3-select w3-{{$themeColor}}">
        {{range $color, $element := .Themes}}
            <option
                class="w3-{{$color}}"
                value="{{$color}}"
                {{if eq $color $themeColor}}
                    selected
                {{end}}>
                {{$color}}
            </option>
        {{end}}
        </select>
    </form>

    <footer class="w3-container w3-theme-d3">
        <button
            onclick="closeOptions()"
            class="w3-btn w3-round w3-left w3-theme-d2 w3-margin">
            Cancel
        </button>

        <button
            onclick="applyOptions()"
            class="w3-btn w3-round w3-right w3-theme-d2 w3-margin">
            Apply
        </button>
    </footer>
</div>
{{end}}
