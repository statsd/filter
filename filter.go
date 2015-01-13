package filter

import "strconv"
import "regexp"
import "bytes"
import "bufio"
import "io"

var replacer = regexp.MustCompile(`\$\d+`)

type Config map[string]interface{}
type regexps map[*regexp.Regexp]interface{}

// Filter metrics from reader to the given writer.
func Filter(r io.Reader, w io.Writer, conf Config) error {
	s := bufio.NewScanner(r)
	filters := compile(conf)

	for s.Scan() {
		metric := s.Bytes()
		parts := bytes.SplitN(metric, []byte("|"), 2)
		name := parts[0]

		if len(name) == 0 {
			continue
		}

		tail := parts[1]

		for regexp, v := range filters {
			matches := regexp.FindAllSubmatch(name, -1)

			if matches == nil {
				continue
			}

			switch v.(type) {
			case string:
				w.Write(replace([]byte(v.(string)), matches))
				w.Write([]byte("|"))
				w.Write(tail)
				w.Write([]byte("\n"))
			case bool:
				w.Write(metric)
				w.Write([]byte("\n"))
			}
		}
	}

	return s.Err()
}

// Replace placeholders in `b` with the respective match.
func replace(b []byte, matches [][][]byte) []byte {
	return replacer.ReplaceAllFunc(b, func(b []byte) []byte {
		i, err := strconv.Atoi(string(b[1:]))

		if err != nil {
			panic(err)
		}

		return matches[0][i]
	})
}

// Compile map of regexps:substitution.
func compile(conf Config) regexps {
	r := make(regexps)

	for k, v := range conf {
		r[regexp.MustCompile(k)] = v
	}

	return r
}
