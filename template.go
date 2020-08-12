package main

const pageTemplate = `<!DOCTYPE html>
<html>
<head>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8"/>
<meta name="go-import" content="{{.VanityDomain}}/{{.Name}} git {{.GithubURL}}">
<meta name="go-source" content="{{.VanityDomain}}/{{.Name}} {{.GithubURL}} {{.GithubURL}}/tree/master{/dir} {{.GithubURL}}/blob/master{/dir}/{file}#L{line}">
<meta http-equiv="refresh" content="0; url=https://pkg.go.dev/{{.VanityDomain}}/{{.Name}}/">
</head>
<body>
Nothing to see here; <a href="https://pkg.go.dev/{{.VanityDomain}}/{{.Name}}/">see the package on godoc</a>.
</body>`
