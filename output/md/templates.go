package md

import (
	"time"

	"github.com/jessevdk/go-assets"
)

var _Assetsac44302fb6150a621aa9d04a0350aac972bf7e18 = "# {{ .Table.Name }}\n\n## Description\n\n{{ .Table.Comment }}\n\n## Columns\n\n| Name | Type | Default | Nullable | Children | Parents | Comment |\n| ---- | ---- | ------- | -------- | -------- | ------- | ------- |\n{{- range $i, $c := .Table.Columns }}\n| {{ $c.Name }} | {{ $c.Type }} | {{ $c.Default.String }} | {{ $c.Nullable }} | {{ range $ii, $r := $c.ChildRelations -}}[{{ $r.Table.Name }}]({{ $r.Table.Name }}.md) {{ end }} | {{ range $ii, $r := $c.ParentRelations -}}[{{ $r.ParentTable.Name }}]({{ $r.ParentTable.Name }}.md) {{ end }} | {{ $c.Comment }} |\n{{- end }}\n\n## Constraits\n\n| Name | Type | Def |\n| ---- | ---- | --- |\n{{- range $i, $c := .Table.Constraits }}\n| {{ $c.Name }} | {{ $c.Type }} | {{ $c.Def }} |\n{{- end }}\n\n## Indexes\n\n| Name | Def |\n| ---- | --- |\n{{- range $i, $idx := .Table.Indexes }}\n| {{ $idx.Name }} | {{ $idx.Def }} |\n{{- end }}\n\n---\n\n> Generated by [tbls](https://github.com/k1LoW/tbls)"
var _Assets43889384df1c6f74d764c29d91b9d5637eb46061 = "# {{ .Schema.Name }}\n\n## Tables\n\n| Name | Columns | Comment | Type |\n| ---- | ------- | ------- | ---- |\n{{- range $i, $t := .Schema.Tables }}\n| [{{ $t.Name }}]({{ $t.Name }}.md) | {{ len $t.Columns }} | {{ $t.Comment }} | {{ $t.Type }} |\n{{- end }}\n\n---\n\n> Generated by [tbls](https://github.com/k1LoW/tbls)"

// Assets returns go-assets FileSystem
var Assets = assets.NewFileSystem(map[string][]string{"/": []string{"table.md.tmpl", "index.md.tmpl"}}, map[string]*assets.File{
	"/": &assets.File{
		Path:     "/",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1526867533, 1526867533600958719),
		Data:     nil,
	}, "/table.md.tmpl": &assets.File{
		Path:     "/table.md.tmpl",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1526799512, 1526799512117449435),
		Data:     []byte(_Assetsac44302fb6150a621aa9d04a0350aac972bf7e18),
	}, "/index.md.tmpl": &assets.File{
		Path:     "/index.md.tmpl",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1526810631, 1526810631217635259),
		Data:     []byte(_Assets43889384df1c6f74d764c29d91b9d5637eb46061),
	}}, "")