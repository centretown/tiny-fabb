	<html>
    <head>
        <style>
            .camera {
                padding: 0px;
                margin: 0;
                margin-block-start: 0;
                margin-block-end: 0;
                margin-inline-start: 0;
                margin-inline-end: 0;
            }

            .camera img {
                display: block;
                border-radius: 4px;
                margin-top: 8px;
            }
        </style>
    </head>
    <body>
    <figure class="camera">
        <div id="stream-camera0" class="image-container">
            <img id="camera0-stream" src="/camera0/mjpeg">
        </div>
    </figure>
    <figure class="camera">
        <div id="stream-camera1" class="image-container">
            <img id="camera1-stream" src="/camera1/mjpeg">
        </div>
    </figure>
</body>
</html>


{{define "list"}}{
    {{range $id, $form := .}}
        {{range $i, $ent := $form.Entries}}
            {{$v := $ent.Value $form.Value}} "{{$ent.ID}}": 
            {{if eq $ent.Type "text"}}
                "{{$v}}"
            {{else}}
                {{$v}}
            {{end}},
        {{end}}
    {{end}}
}
{{end}}
