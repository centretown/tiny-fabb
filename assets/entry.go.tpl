{{define "list"}}
    <div class="w3-responsive">
        <table class="w3-table" style="margin-top:.5rem">
        {{range $id, $form := .}}
            {{range $i, $ent := $form.Entries}}
            <tr>
                <td>
                    <a href="javascript:openEdit('{{$id}}');"
                        title="Edit current setting.">
                        {{$ent.Code}}: {{$ent.Label}}
                    </a>
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
            name="settings-form"
            method="post"
            action="/apply/{{$first.ID}}/">

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
                        <p></p>
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

{{define "commands-list"}}
<div class="w3-responsive">
  <table class="w3-table w3-striped">
    {{range $code, $val := .Commands}}
    <tr>
      <td>
        <a href="{{$val.LinkHelp}}" target="_blank"
          >{{$code}}</a
        >
      </td>
      <td>
        <a href="#">{{$val.Label}}</a>
      </td>
    </tr>
    {{end}}
  </table>
</div>
{{end}}

<!-- status -->
{{define "status"}}
<h5>Status:</h5>
{{end}}