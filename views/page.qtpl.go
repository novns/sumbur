// Code generated by qtc from "page.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

package views

import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

type HTML interface {
	Title() string
	StreamTitle(qw422016 *qt422016.Writer)
	WriteTitle(qq422016 qtio422016.Writer)
	TitleAdd() string
	StreamTitleAdd(qw422016 *qt422016.Writer)
	WriteTitleAdd(qq422016 qtio422016.Writer)
	Body() string
	StreamBody(qw422016 *qt422016.Writer)
	WriteBody(qq422016 qtio422016.Writer)
}

type BasePage struct{}

func (page *BasePage) StreamTitle(qw422016 *qt422016.Writer) {
}

func (page *BasePage) WriteTitle(qq422016 qtio422016.Writer) {
	qw422016 := qt422016.AcquireWriter(qq422016)
	page.StreamTitle(qw422016)
	qt422016.ReleaseWriter(qw422016)
}

func (page *BasePage) Title() string {
	qb422016 := qt422016.AcquireByteBuffer()
	page.WriteTitle(qb422016)
	qs422016 := string(qb422016.B)
	qt422016.ReleaseByteBuffer(qb422016)
	return qs422016
}

func (page *BasePage) StreamTitleAdd(qw422016 *qt422016.Writer) {
}

func (page *BasePage) WriteTitleAdd(qq422016 qtio422016.Writer) {
	qw422016 := qt422016.AcquireWriter(qq422016)
	page.StreamTitleAdd(qw422016)
	qt422016.ReleaseWriter(qw422016)
}

func (page *BasePage) TitleAdd() string {
	qb422016 := qt422016.AcquireByteBuffer()
	page.WriteTitleAdd(qb422016)
	qs422016 := string(qb422016.B)
	qt422016.ReleaseByteBuffer(qb422016)
	return qs422016
}

func (page *BasePage) StreamBody(qw422016 *qt422016.Writer) {
}

func (page *BasePage) WriteBody(qq422016 qtio422016.Writer) {
	qw422016 := qt422016.AcquireWriter(qq422016)
	page.StreamBody(qw422016)
	qt422016.ReleaseWriter(qw422016)
}

func (page *BasePage) Body() string {
	qb422016 := qt422016.AcquireByteBuffer()
	page.WriteBody(qb422016)
	qs422016 := string(qb422016.B)
	qt422016.ReleaseByteBuffer(qb422016)
	return qs422016
}

func StreamPage(qw422016 *qt422016.Writer, page HTML) {
	qw422016.N().S(`<!DOCTYPE html>

<html lang="en">


<head>

<meta charset="UTF-8">
<meta name="viewport" content="width=device-width, initial-scale=1">

<title>`)
	page.StreamTitle(qw422016)
	page.StreamTitleAdd(qw422016)
	qw422016.N().S(`</title>

<link href="https://cdn.jsdelivr.net/npm/bootstrap@5.1.2/dist/css/bootstrap.min.css"
    rel="stylesheet" integrity="sha384-uWxY/CJNBR+1zjPWmfnSnVxwRheevXITnMqoEIeG1LJrdI0GlVs/9cVSyPYXdcSF" crossorigin="anonymous">

</head>


<body>


<header class="container container-fluid navbar navbar-dark bg-secondary mb-3 py-1">

<div class="navbar-brand">Sumbur demo</div>

<form id="auth-form" action="/auth" method="post" class="d-flex">
`)
	if AuthState {
		qw422016.N().S(`<input type="submit" value="Logout" class="btn btn-danger">
`)
	} else {
		qw422016.N().S(`<input type="password" name="password" size="12" placeholder="Password" class="form-control me-2">
<input type="submit" value="Login" class="btn btn-primary">
`)
	}
	qw422016.N().S(`</form>

</header>


<main class="container">

<h2>`)
	page.StreamTitle(qw422016)
	page.StreamTitleAdd(qw422016)
	qw422016.N().S(`</h2>
`)
	page.StreamBody(qw422016)
	qw422016.N().S(`
</main>


<footer class="container">
</footer>


<script src="https://cdn.jsdelivr.net/npm/jquery@3.6.0/dist/jquery.min.js"
    integrity="sha256-/xUj+3OJU5yExlq6GSYGSHk7tPXikynS7ogEvDej/m4=" crossorigin="anonymous"></script>

</body>


</html>
`)
}

func WritePage(qq422016 qtio422016.Writer, page HTML) {
	qw422016 := qt422016.AcquireWriter(qq422016)
	StreamPage(qw422016, page)
	qt422016.ReleaseWriter(qw422016)
}

func Page(page HTML) string {
	qb422016 := qt422016.AcquireByteBuffer()
	WritePage(qb422016, page)
	qs422016 := string(qb422016.B)
	qt422016.ReleaseByteBuffer(qb422016)
	return qs422016
}
