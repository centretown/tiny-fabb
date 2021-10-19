{{define "servo-settings"}}
{{$camera:=.}}
{{$svos:=.Servos}}
<div class="w3-modal-content w3-theme-d5">
    <header class="w3-container w3-theme-l1 w3-padding-16">
        <h5 style="display:inline-block">Servos</h5>
        <span class="w3-right">
            <button class="w3-btn w3-round w3-left w3-theme-d5"
                onclick="closeDialog()">x</button>
        </span>
    </header>
    <div class="content w3-theme-l3">
        {{template "camera-servo" $camera}}
        <div class="content" style="display:block;width:35%;">
            {{template "servo-edit" $svos}}
        </div>
    </div>
    <footer class="w3-container w3-theme-l1">
        <button
            type="button"
            onclick="closeDialog()"
            class="w3-btn w3-round w3-left w3-margin w3-theme-d5">
            Close
        </button>
        <button
            type="button"
            class="w3-btn w3-round w3-right w3-margin w3-theme-d5"
            onclick="applyServo()">
            Apply
        </button>
    </footer>
</div> 
{{end}}

{{define "servo-edit"}}
{{$svos:=.}}
<form class="w3-container entries">
    <label for="servo-select">Servo</label>
    <select id="servo-select" name="servo-select"
        class="w3-theme-l1 select-controller w3-padding">
        {{range $index, $svo := $svos}}
        <option value="{{$svo.Index}}">{{$svo.Title}}</option>
        {{end}}
    </select>
    <label for="servo-command">Command</label>
    <select id="servo-command" name="servo-command"
        class="w3-theme-l1 select-controller w3-padding"
        value="3">
        <option value="home">Home</option>
        <option value="move">Move</option>
        <option value="ease" selected>Ease</option>
        <option value="test" selected>Test</option>
        <option value="stop" selected>Stop</option>
    </select>

    <label for="servo-angle">Angle</label>
    <input id="servo-angle" name="servo-angle" 
        type="number" min="0" max="180"
        value="90" style="width:100%"/>

    <label for="servo-speed">Speed</label>
    <input id="servo-speed" name="servo-speed" 
        type="number"  min="0" max="255"
        value="50" style="width:100%"/>

    <label for="servo-ease-type">Easing Method</label>
    <select id="servo-ease-type" name="servo-ease-type"
        class="w3-theme-l1 select-controller w3-padding"
        value="1">
        <option value="0">Linear</option>
        <option value="1" selected>Quadradic</option>
        <option value="2">Cubic</option>
        <option value="3">Quartic</option>
    </select>

    <label for="servo-angle2">Angle 2</label>
    <input id="servo-angle2" name="servo-angle2" 
        type="number"  min="0" max="255"
        value="50" style="width:100%"/>

</form>
{{end}}

