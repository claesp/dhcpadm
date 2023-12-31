// Code generated by qtc from "style.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line templates/style.qtpl:1
package templates

//line templates/style.qtpl:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line templates/style.qtpl:1
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line templates/style.qtpl:2
type Style interface {
//line templates/style.qtpl:2
	Body() string
//line templates/style.qtpl:2
	StreamBody(qw422016 *qt422016.Writer)
//line templates/style.qtpl:2
	WriteBody(qq422016 qtio422016.Writer)
//line templates/style.qtpl:2
}

//line templates/style.qtpl:7
func StreamStyleTemplate(qw422016 *qt422016.Writer, s Style) {
//line templates/style.qtpl:7
	qw422016.N().S(`
`)
//line templates/style.qtpl:8
	s.StreamBody(qw422016)
//line templates/style.qtpl:8
	qw422016.N().S(`
`)
//line templates/style.qtpl:9
}

//line templates/style.qtpl:9
func WriteStyleTemplate(qq422016 qtio422016.Writer, s Style) {
//line templates/style.qtpl:9
	qw422016 := qt422016.AcquireWriter(qq422016)
//line templates/style.qtpl:9
	StreamStyleTemplate(qw422016, s)
//line templates/style.qtpl:9
	qt422016.ReleaseWriter(qw422016)
//line templates/style.qtpl:9
}

//line templates/style.qtpl:9
func StyleTemplate(s Style) string {
//line templates/style.qtpl:9
	qb422016 := qt422016.AcquireByteBuffer()
//line templates/style.qtpl:9
	WriteStyleTemplate(qb422016, s)
//line templates/style.qtpl:9
	qs422016 := string(qb422016.B)
//line templates/style.qtpl:9
	qt422016.ReleaseByteBuffer(qb422016)
//line templates/style.qtpl:9
	return qs422016
//line templates/style.qtpl:9
}

//line templates/style.qtpl:11
type BaseStyle struct{}

//line templates/style.qtpl:12
func (p *BaseStyle) StreamBody(qw422016 *qt422016.Writer) {
//line templates/style.qtpl:12
	qw422016.N().S(`
body {
	margin: 0;
	padding: 0;
}

header {
	float: left;
	width: 100%;
}

nav {
	float: left;
}

main {
	float: left;
}

footer {
	float: left;
	width: 100%;
}
`)
//line templates/style.qtpl:35
}

//line templates/style.qtpl:35
func (p *BaseStyle) WriteBody(qq422016 qtio422016.Writer) {
//line templates/style.qtpl:35
	qw422016 := qt422016.AcquireWriter(qq422016)
//line templates/style.qtpl:35
	p.StreamBody(qw422016)
//line templates/style.qtpl:35
	qt422016.ReleaseWriter(qw422016)
//line templates/style.qtpl:35
}

//line templates/style.qtpl:35
func (p *BaseStyle) Body() string {
//line templates/style.qtpl:35
	qb422016 := qt422016.AcquireByteBuffer()
//line templates/style.qtpl:35
	p.WriteBody(qb422016)
//line templates/style.qtpl:35
	qs422016 := string(qb422016.B)
//line templates/style.qtpl:35
	qt422016.ReleaseByteBuffer(qb422016)
//line templates/style.qtpl:35
	return qs422016
//line templates/style.qtpl:35
}
