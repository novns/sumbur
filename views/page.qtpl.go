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

func StreamPage(qw422016 *qt422016.Writer) {
	qw422016.N().S(`<!DOCTYPE html>

<html lang="en">

<head>
<meta charset="UTF-8">
</head>

<body>
sumbur
</body>

</html>
`)
}

func WritePage(qq422016 qtio422016.Writer) {
	qw422016 := qt422016.AcquireWriter(qq422016)
	StreamPage(qw422016)
	qt422016.ReleaseWriter(qw422016)
}

func Page() string {
	qb422016 := qt422016.AcquireByteBuffer()
	WritePage(qb422016)
	qs422016 := string(qb422016.B)
	qt422016.ReleaseByteBuffer(qb422016)
	return qs422016
}
