package main

import "github.com/segmentio/go-log"
import "github.com/statsd/filter"
import "github.com/tj/docopt"
import "encoding/json"
import "io/ioutil"
import "os"

var Version = "0.0.1"

const Usage = `
  Usage:
    filter --config file
    filter -h | --help
    filter --version

  Options:
    -c, --config file  JSON configuration file
    -h, --help         output help information
    -v, --version      output version

`

func main() {
	args, err := docopt.Parse(Usage, nil, true, Version, false)
	log.Check(err)

	path := args["--config"].(string)
	b, err := ioutil.ReadFile(path)
	log.Check(err)

	var conf filter.Config
	log.Check(json.Unmarshal(b, &conf))

	log.Check(filter.Filter(os.Stdin, os.Stdout, conf))
}
