package filter

import "github.com/bmizerany/assert"
import "testing"
import "bytes"

func TestBooleanFilter(t *testing.T) {
	in := bytes.NewBufferString(`
gauges.api-1.memory.percent|80|1421164536386
gauges.api-2.memory.percent|30|1421164536386
gauges.ingestion-2.memory.percent|50|1421164536386
gauges.api-3.memory.percent|50|1421164536386
gauges.cdn-1.memory.percent|50|1421164536386
gauges.cdn-2.memory.percent|50|1421164536386
gauges.ingestion-1.memory.percent|50|1421164536386`)

	out := bytes.NewBuffer(nil)

	conf := Config{
		`^gauges\.api-*`: true,
		`^gauges\.cdn-*`: true,
	}

	err := Filter(in, out, conf)
	assert.Equal(t, nil, err)

	exp := `gauges.api-1.memory.percent|80|1421164536386
gauges.api-2.memory.percent|30|1421164536386
gauges.api-3.memory.percent|50|1421164536386
gauges.cdn-1.memory.percent|50|1421164536386
gauges.cdn-2.memory.percent|50|1421164536386
`

	assert.Equal(t, exp, string(out.Bytes()))
}

func TestStringMapping(t *testing.T) {
	in := bytes.NewBufferString(`
gauges.api-1.memory.percent|80|1421164536386
gauges.api-2.memory.percent|30|1421164536386
gauges.ingestion-2.memory.percent|50|1421164536386
gauges.api-3.memory.percent|50|1421164536386
gauges.cdn-1.memory.percent|50|1421164536386
gauges.cdn-2.memory.percent|50|1421164536386
counts.app.signups|10|1421164536386
gauges.ingestion-1.memory.percent|50|1421164536386`)

	out := bytes.NewBuffer(nil)

	conf := Config{
		`^counts\.app`:             true,
		`^gauges\.(api-\d+)`:       "$1",
		`^gauges\.cdn-(\d+)\.(.*)`: "content-$1-$2",
	}

	err := Filter(in, out, conf)
	assert.Equal(t, nil, err)

	exp := `api-1|80|1421164536386
api-2|30|1421164536386
api-3|50|1421164536386
content-1-memory.percent|50|1421164536386
content-2-memory.percent|50|1421164536386
counts.app.signups|10|1421164536386
`

	assert.Equal(t, exp, string(out.Bytes()))
}
