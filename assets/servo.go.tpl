{{define "servo-settings"}}
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
        <div class="entries">
            <img id="{{.ID}}-stream" 
                class="pic w3-image"
                src="/{{.ID}}/mjpeg">
        </div>
        <div class="content" style="display:block;width:30%;">
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
    </footer>
</div> 
{{end}}

{{define "servo-edit"}}
{{$svos:=.}}
<div class="w3-container entries">
{{range $index, $svo := $svos}}
    {{range $webid, $form := .Forms}}
        {{$value := $form.Value}}
        {{$first := index $form.Entries 0}}
        {{range $i, $ent := $form.Entries}}
        <div class="entries dlgtab">
            {{$f:=$ent.FormatInput $value $first}}
            {{$id:=printf "%s-%d" $f.ID $index}}
            <label for="{{$id}}">{{$ent.Label}}</label>
            {{if $f.ReadOnly}}
               <span>{{$value}}</span>
            {{else}}
                <input
                    id="{{$id}}"
                    name="{{$id}}"
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
            {{end}}
        </div>
        {{end}}
    {{end}}
    <div>
        <div>
            <button class="w3-button w3-center"
                style="width:100%"
                onclick="">
                Apply
            </button>
        </div>
    </div>
{{end}}
</div>
{{end}}

