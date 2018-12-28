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
	tempNewPost = `h2. {{.Title}}

p(data). {{.Date}}

{{.Content}}`

	tempHtmlBase = `{{define "base"}}
<html>
<head>
    {{template "header" .}}
</head>
<body>
    {{template "body" .}}
    {{template "footer" .}}
</body>
</html>
{{end}}`

	tempHtmlIndex = `{{ define "header"}}
<meta charset="utf-8">
<meta http-equiv="X-UA-Compatible" content="IE=edge,chrome=1">
<title>Invido site pre processor</title>
<meta name="description" content="">
<meta name="viewport" content="width=device-width, initial-scale=1">

{{end}} 

{{ define "body"}}
<div class="ui container">
  {{block "main" .}}{{end}}

	{{end}} 

	{{define "footer"}}
	<div class="ui vertical footer segment">
			<div class="ui container">
				<div class="ui stackable  divided equal ten stackable grid">
					<div class="five wide column">
						<h4 class="ui header">Version</h4>
						<p>Software build {{.Buildnr}}</p>
					</div>
					<div class="seven wide column">
						<h4 class="ui header">Info</h4>
						<p><i class="copyright icon"></i> 2018 by Invido.it</p>
					</div>
				</div>
			</div>
		</div>
</div>
<!-- Put here script to beloaded-->

{{end}}`

	tempHtmlEditPost = `{{define "main"}}

<div class="ui attached message">
  <div class="header">
    Edit Post
  </div>
</div>
<form class="ui form attached fluid segment" action="/save-post/" method="POST">
  <div class="field">
    <label>Title</label>
    <input placeholder="TitlePost" name="titlepost" type="text" size="35" value='{{printf "%s" .TitlePost}}'>
  </div>
  <div class="field">
    <label>Post</label>
    <textarea rows="25" cols="110" name="contentpost" placeholder="ContentPost">{{printf "%s" .ContentPost}}</textarea>
  </div>
  <div>
		<input class="ui blue submit button" type="submit"  value="Save">
		<a class="ui button" href="/create-page-index/">Create Index  Pages</a>
		<a class="ui button" href="/exec-webgen/">Start Webgen</a>
  </div>
</form>
<div>
	<a class="ui button" href="{{printf "%s" .WebgenOutIndexFile}}" target="_blank">Navigate to webgen out</a>
</div>
<div>
	<p>
		{{printf "%s" .LastMessage}}
	</p>
</div>
{{end}}`
)
