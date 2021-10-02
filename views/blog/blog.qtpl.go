// Code generated by qtc from "blog.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

package blog

import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

func (blog *Blog) StreamTitle(qw422016 *qt422016.Writer) {
	qw422016.N().S(`Articles`)
}

func (blog *Blog) WriteTitle(qq422016 qtio422016.Writer) {
	qw422016 := qt422016.AcquireWriter(qq422016)
	blog.StreamTitle(qw422016)
	qt422016.ReleaseWriter(qw422016)
}

func (blog *Blog) Title() string {
	qb422016 := qt422016.AcquireByteBuffer()
	blog.WriteTitle(qb422016)
	qs422016 := string(qb422016.B)
	qt422016.ReleaseByteBuffer(qb422016)
	return qs422016
}

func (blog *Blog) StreamBody(qw422016 *qt422016.Writer) {
	qw422016.N().S(`
`)
	blog.tags.StreamTags(qw422016)
	qw422016.N().S(`
`)
	blog.articles.StreamArticles(qw422016)
}

func (blog *Blog) WriteBody(qq422016 qtio422016.Writer) {
	qw422016 := qt422016.AcquireWriter(qq422016)
	blog.StreamBody(qw422016)
	qt422016.ReleaseWriter(qw422016)
}

func (blog *Blog) Body() string {
	qb422016 := qt422016.AcquireByteBuffer()
	blog.WriteBody(qb422016)
	qs422016 := string(qb422016.B)
	qt422016.ReleaseByteBuffer(qb422016)
	return qs422016
}
