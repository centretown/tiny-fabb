
{{define "command-list"}}
<div class="w3-container">
    {{range $id, $form := .}}
        {{$ent := index $form.Entries 0}}
        {{$val := $form.Value}}
        <div class="w3-cell-row">
            <div class="w3-cell" style="width:12rem;">
                <button id="{{$ent.ID}}-update" class="w3-btn w3-ripple w3-theme-l1">
                    Update
                </button>
            </div>
            <div class="w3-cell w3-left">
                <label for="{{$ent.ID}}">{{$ent.Label}}</label>
                <div id="{{$ent.ID}}">
                    {{$val}}
                </div>
            </div>
        </div>
    {{end}}
</div>
{{end}}

{{define "list"}}
    <div class="w3-responsive">
        <table class="w3-table" style="margin-top:.5rem;margin-bottom:.5rem">
        {{range $id, $form := .}}
            {{range $i, $ent := $form.Entries}}
            <tr>
                <td>
                    <button id="{{$ent.ID}}-edit" class="w3-btn w3-ripple w3-theme-l1"
                        onclick="openEdit({{$id}})">
                        Edit
                    </button>
                </td>
                <td>
                    <label>{{$ent.Label}}</label>
                </td>
                <td id="{{$ent.ID}}">
                    {{$ent.Value $form.Value}}
                </td>
            </tr>
            {{end}}
        {{end}}
        </table>
    </div>
{{end}}


{{define "pop-detail"}}
    <h4>{{.Link}}: {{.Title}}</h4>
    <ul>
    {{range $idx, $text := .Text}}
        <li>{{$text}}</li>
    {{end}}
    </ul>
    {{range $idx, $doc := .Subs}}
        {{template "pop-detail" $doc}}
    {{end}}
{{end}}

{{define "pop-doc"}}
    {{template "pop-detail" .}}
{{end}}

{{define "edit"}} 
{{$value := .Value}}
{{$first := index .Entries 0}}
{{$doc := .FindDoc $first.Code}}
<div class="w3-modal-content w3-theme-d5">
    <form
        class="w3-container"
        id="settings-form"
        name="{{$first.ID}}"
        method="post"
        action="">

        <header class="w3-container w3-theme-l1 w3-padding-16">
            <h5 style="display:inline-block">{{$first.Code}}: {{$first.Label}}</h5>
            <span class="w3-right">
                <button class="w3-btn w3-round w3-left w3-theme-d5"
                    onclick="closeDialog()">x</button>
            </span>
        </header>

        <div class="content w3-theme-l3">
            <div class="entries">
            {{range $id, $ent := .Entries}}
                {{if ne $ent.Type "mask"}} 
                    {{$f:=$ent.FormatInput $value $first}}
                    <label class="w3-padding">
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
                            value="{{$f.Value}}"/>
                        {{$ent.Label}}
                    </label>
                {{end}}
            {{end}}
            </div>
            <div class="entries doc w3-theme-l5">
                {{template "pop-doc" $doc}}
            </div>
        </div>

        <footer class="w3-container w3-theme-l1">
            <button
                type="button"
                onclick="closeDialog()"
                class="w3-btn w3-round w3-left w3-margin w3-theme-d5">
                Cancel
            </button>
            <button
                type="submit"
                class="w3-btn w3-round w3-right w3-margin w3-theme-d5">
                Apply
            </button>
        </footer>
    </form>
</div> 
{{end}}

{{define "error"}} 
    An error has occurred. Sad. 
    {{.}}
{{end}}

<!-- status -->
{{define "status"}}
<h5>Status:</h5>
{{end}}
