{{/* Copyright (c) 2021 Dave Marsh. See LICENSE. */}}

{{define "pop-detail"}}
    <div class="doc">
    <header class="doc-title">{{.Link}}: {{.Title}}</header>
    {{range $idx, $text := .Text}}
        <p class="">{{$text}}</p>
    {{end}}
    {{range $idx, $doc := .Subs}}
        {{template "pop-detail" $doc}}
    {{end}}
    </div>
{{end}}

{{define "pop-doc"}}
    {{template "pop-detail" .}}
{{end}}

{{define "check"}}
{{$f:=.}}
{{$ent:=.Entry}}
<div class="entries">
    <label class="w3-padding">
        <input
            name="{{$f.Name}}"
            class="{{$f.Class}}"
            type="{{$ent.Type}}"
            {{if $f.HasChecked}}checked{{end}}
            value="{{$f.Value}}"/>
        {{$f.Entry.Label}}
    </label>
</div>
{{end}}

{{define "text-digits"}}
{{$f:=.}}
{{$ent:=.Entry}}
<div class="entries">
    <div class="w3-padding">
        {{$ent.Label}}</div>
    <input
        name="{{$f.Name}}"
        class="{{$f.Class}} w3-margin"
        type="{{$ent.Type}}"
        {{if $f.HasRange}}min="{{$ent.Min}}"
            max="{{$ent.Max}}"{{end}}
        {{if $f.HasStep}}step="{{$ent.Step}}"{{end}}
        value="{{$f.Value}}"/>
</div>
{{end}}

{{/* accepts Identifier */}}
{{define "edit"}} 
{{$frm := .Form}}
{{$value := $frm.Value}}
{{$first := index $frm.Entries 0}}
<div class="entries">
    {{range $id, $ent := $frm.Entries}}
    {{if eq $ent.Type "mask" "cmd"}} 
    {{else}}
        {{$f:=$ent.FormatInput $value $first}}
        {{if eq $ent.Type "checkbox" "radio"}}
            {{template "check" $f}}
        {{else}}
            {{template "text-digits" $f}}
        {{end}}
    {{end}}
    {{end}}
</div>
{{end}}

{{/* accepts Identifier */}}
{{define "flex-form"}}
{{$ent := index .Form.Entries 0}}
{{$val := .Form.Value}}
<div class="w3-card w3-theme-d4 flexitem">
    <button 
        class="w3-block w3-btn w3-theme-d4 w3-left-align"
        onclick="toggleForm('{{.View}}','{{$ent.ID}}','{{.FormID}}');">
        <i class="bi {{.Icon}}"></i>
        <label class="">{{$ent.Label}}</label>
        {{if ne $ent.Type "cmd"}}
            <span id="{{.FormID}}-value" class="w3-right">{{$val}}</span>
        {{end}}
    </button>
    {{$doc := $ent.FindDoc $ent.Code}}
    <div id="{{.FormID}}" class="w3-hide">
        <button class="w3-block w3-theme-d3"
            onclick="toggleShow('{{.FormID}}-doc')">
            <i class="bi bi-file-earmark-arrow-down"></i>
            <label class="">More Information...</label>
        </button>
        <div id="{{.FormID}}-doc" class="flexcontent w3-theme-l4 w3-hide">
            {{template "pop-doc" $doc}}
        </div>

        <form id="{{.FormID}}-form" name="{{.Form.ID}}"
            class="w3-theme-l3 flexcontent">
            {{template "edit" .}}
        </form>

        <button class="w3-block w3-theme-d3"
            onclick="submitForm('{{.View}}','{{$ent.ID}}','{{.FormID}}');">
            <i class="bi bi-asterisk"></i>
            <label class="w3-padding-small">Update</label>
        </button>

        <div class="w3-theme-l2 flexcontent">
            {{range $is, $s := .Results}}
            <div class="response">
                {{$s}}
            </div>
            {{end}}
        </div>
    </div>
</div>
{{end}}

{{define "commands"}}
{{$gctl:=.}}
{{$view:="commands"}}
{{$icon:="bi-command"}}
{{$ctlID := $gctl.ID}}
{{$forms := $gctl.ViewForms $view}}

<div class="flexview">
    {{range $id, $form := $forms}}
        {{$ent := index $form.Entries 0}}
        {{$rl := $ent.ResponseList $form.Value}}
        {{$results := $gctl.FormatResponseList $rl}}
        {{$ident := $form.Identify $ctlID $view $icon $results}}
        {{template "flex-form" $ident}}
    {{end}}
</div>
{{end}}

{{define "settings"}}
{{$gctl:=.}}
{{$view:="settings"}}
{{$icon:="bi-gear"}}
{{$ctlID := $gctl.ID}}
{{$forms := $gctl.ViewForms $view}}

<div class="flexview">
    {{range $id, $form := $forms}}
        {{$ent := index $form.Entries 0}}
        {{$results := $ent.ResponseList $form.Value}}
        {{/* {{$results := $gctl.FormatResponseList $rl}} */}}
        {{$ident := $form.Identify $ctlID $view $icon $results}}
        {{template "flex-form" $ident}}
    {{end}}
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
