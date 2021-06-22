{{define "list"}}{
    {{range $id, $form := .}}{{range $i, $ent := $form.Entries}}{{$v := $ent.Value $form.Value}} "{{$ent.ID}}": {{if eq $ent.Type "text"}}"{{$v}}"{{else}}{{$v}}{{end}},
    {{end}}
    {{end}}
}
{{end}}

{{define "doc"}}{{.Link}}: {{.Title}}
{{range $idx, $text := .Text}}{{$text}}
{{end}}
{{range $idx, $doc := .Subs}}{{template "doc" $doc}}{{end}}{{end}}

{{define "pop-detail"}}<h4>{{.Link}}: {{.Title}}</h4><ul>{{range $idx, $text := .Text}}<li>{{$text}}</li>{{end}}</ul>{{range $idx, $doc := .Subs}}{{template "pop-detail" $doc}}{{end}}{{end}}

{{define "pop-doc"}}
<div id="pop-doc" class="w3-dropdown-content w3-border w3-container pop">
{{template "pop-detail" .}}
</div>{{end}}

