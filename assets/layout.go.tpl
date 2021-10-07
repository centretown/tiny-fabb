
{{define "controller-bar"}}
<div id="controller-bar" class="w3-bar w3-border w3-theme-d3">
{{range $i, $ctl := .}}
    <button 
        id="ctl{{$i}}" 
        class="w3-button w3-bar-item tablink"
        onclick="selectController(event,{{$i}},{{$ctl.Descriptor}})">
        {{.Descriptor}}
    </button>
{{end}}
</div>
{{end}}

{{define "controller"}}
    {{$ctlID := .Descriptor}}
    <div id="{{$ctlID}}" class="controller" style="display:none;">
    {{range $i, $view := .Views}}
       {{$ID := printf "%s-%s" $ctlID $view.ID}}
        <button 
            class="w3-btn w3-block w3-black w3-left-align"
            title="{{$view.Title}}"
            onclick="selectSection({{$ID}},{{$view.ID}})">
            <i class="bi {{$view.Icon}}"></i>
            <label>{{$view.Title}}</label>
        </button>
        <div id="{{$ID}}" class="w3-container w3-hide"></div>
    {{end}}
    </div>
{{end}}

{{define "controllers"}}
    {{range $i, $ctl := .}}
        {{template "controller" $ctl}}
    {{end}}
{{end}}

{{define "app-bar"}}
<div id="app-bar" class="w3-bar w3-border w3-theme-d3">
    <div class="w3-dropdown-hover w3-left w3-bar-item w3-theme-d3">
        <button class="w3-button">
            <i class="bi bi-menu-button"></i>
        </button>
        <div class="w3-dropdown-content w3-bar-block w3-border w3-theme-d1"
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
    <div class="w3-dropdown-hover w3-right w3-bar-item w3-theme-d3">
        <button class="w3-button app-icon">
            <i class="bi bi-person"></i>
        </button>
        <div class="w3-dropdown-content w3-bar-block w3-border w3-theme-d1" 
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
    <div class="w3-theme-d3">
        {{template "controller-bar" .Controllers}}
    </div>
</div>
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
