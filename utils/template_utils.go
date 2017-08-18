package utils

const FSDiffOutput = `
-----{{.DiffType}}-----

These entries have been added to {{.Image1}}:{{if not .Diff.Adds}} None{{else}}
	{{range .Diff.Adds}}{{print .}}
	{{end}}{{end}}

These entries have been deleted from {{.Image1}}:{{if not .Diff.Dels}} None{{else}}
	{{range .Diff.Dels}}{{print .}}
	{{end}}{{end}}

These entries have been changed between {{.Image1}} and {{.Image2}}:{{if not .Diff.Mods}} None{{else}}
	{{range .Diff.Mods}}{{print .}}
	{{end}}{{end}}
`

const SingleVersionDiffOutput = `
-----{{.DiffType}}-----

Packages found only in {{.Image1}}:{{if not .Diff.Packages1}} None{{else}}
NAME	VERSION	SIZE{{range $name, $value := .Diff.Packages1}}{{"\n"}}{{print "-"}}{{$name}}	{{$value.Version}}	{{$value.Size}}B{{end}}{{end}}

Packages found only in {{.Image2}}:{{if not .Diff.Packages2}} None{{else}}
NAME	VERSION	SIZE{{range $name, $value := .Diff.Packages2}}{{"\n"}}{{print "-"}}{{$name}}	{{$value.Version}}	{{$value.Size}}B{{end}}{{end}}

Version differences:{{if not .Diff.InfoDiff}} None{{else}}
PACKAGE	IMAGE1 ({{.Image1}})	IMAGE2 ({{.Image2}}){{range .Diff.InfoDiff}}{{"\n"}}{{print "-"}}{{.Package}}	{{.Info1.Version}}, {{.Info1.Size}}B	{{.Info2.Version}}, {{.Info2.Size}}B{{end}}{{end}}
`

const MultiVersionDiffOutput = `
-----{{.DiffType}}-----

Packages found only in {{.Image1}}:{{if not .Diff.Packages1}} None{{else}}
NAME	VERSION	SIZE{{range $name, $value := .Diff.Packages1}}{{"\n"}}{{print "-"}}{{$name}}	{{range $key, $info := $value}}{{$info.Version}}	{{$info.Size}}B{{end}}{{end}}{{end}}

Packages found only in {{.Image2}}:{{if not .Diff.Packages2}} None{{else}}
NAME	VERSION	SIZE{{range $name, $value := .Diff.Packages2}}{{"\n"}}{{print "-"}}{{$name}}	{{range $key, $info := $value}}{{$info.Version}}	{{$info.Size}}B{{end}}{{end}}{{end}}

Version differences:{{if not .Diff.InfoDiff}} None{{else}}
PACKAGE	IMAGE1 ({{.Image1}})	IMAGE2 ({{.Image2}}){{range .Diff.InfoDiff}}{{"\n"}}{{print "-"}}{{.Package}}	{{range .Info1}}{{.Version}}, {{.Size}}B{{end}}	{{range .Info2}}{{.Version}}, {{.Size}}B{{end}}{{end}}{{end}}
`

const HistoryDiffOutput = `
-----{{.DiffType}}-----

Docker history lines found only in {{.Image1}}:{{if not .Diff.Adds}} None{{else}}{{block "list" .Diff.Adds}}{{"\n"}}{{range .}}{{print "-" .}}{{"\n"}}{{end}}{{end}}{{end}}

Docker history lines found only in {{.Image2}}:{{if not .Diff.Dels}} None{{else}}{{block "list2" .Diff.Dels}}{{"\n"}}{{range .}}{{print "-" .}}{{"\n"}}{{end}}{{end}}{{end}}
`

const ListAnalysisOutput = `
-----{{.AnalyzeType}}-----

Analysis for {{.Image}}:{{if not .Analysis}} None{{else}}{{block "list" .Analysis}}{{"\n"}}{{range .}}{{print "-" .}}{{"\n"}}{{end}}{{end}}{{end}}
`

const MultiVersionPackageOutput = `
-----{{.AnalyzeType}}-----

Packages found in {{.Image}}:{{if not .Analysis}} None{{else}}
NAME	VERSION	SIZE{{range $name, $value := .Analysis}}{{"\n"}}{{print "-"}}{{$name}}	{{range $key, $info := $value}}{{$info.Version}}	{{$info.Size}}B{{end}}{{end}}{{end}}
`

const SingleVersionPackageOutput = `
-----{{.AnalyzeType}}-----

Packages found in {{.Image}}:{{if not .Analysis}} None{{else}}
NAME	VERSION	SIZE{{range $name, $value := .Analysis}}{{"\n"}}{{print "-"}}{{$name}}	{{$value.Version}}	{{$value.Size}}B{{end}}{{end}}
`
