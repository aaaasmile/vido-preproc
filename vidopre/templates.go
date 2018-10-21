package vidopre

const (
	tempIndex0Page = `---
	title: Home
	in_menu: true
	routed_title: Invido
	directory_name: invido
	sort_info: 0
---

`

	tempIndexOtherPages = `---
title: Home
in_menu: false
directory_name: invido
---

`
	tempNavInPage = `{{if .ZeroPage}}<a href="{relocatable: /index.html}">[</a> Indice |{{end}}{{if .FirstPage}}<a href="{relocatable: /index.html}">[ Indice </a> | 01 |{{end}}
{{range .NavDet}}
{{if .IsSelected}}|{{.PageIx}}|{{else}}<a href="{relocatable: /index_{{.PageIx}}.html}"> {{.PageIx}} </a>{{end}}
{{end}}
`
)
