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
	tempNavInPage = `{{if .ZeroPage}}<a href="{relocatable: /index.html}">[</a> Indice |{{else}}<a href="{relocatable: /index.html}">[ Indice </a>|{{end}}
{{range .NavDet}}{{if .IsSelected}}  {{.PageIx}}  {{if .IsLast}} <a href="{relocatable: /index_{{.PageIx}}.html}"> ]</a> {{else}} | {{end}}{{else}}<a href="{relocatable: /index_{{.PageIx}}.html}"> {{.PageIx}} {{if .IsLast}}]{{end}}</a>{{if .IsLast}}{{else}}|{{end}}{{end}}{{end}}

`
)
