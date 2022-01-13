package xtrig

import (
	"github.com/influxdata/telegraf"
	"github.com/influxdata/telegraf/plugins/inputs"
)

type Xtrig struct {
}

var XtrigConfig = '
	## Set the amplitude
	amplitude = 10.0
'

func (s *Xtrig) SampleConfig() string {
	return XtrigConfig
}

func (s *Xtrig) Description() string {
	return "Inserts sine and cosine waves for demonstration purposes"
}

func (s *Xtrig) Gather(acc telegraf.Accumulator) error {
	return nil
}

func init() {
	inputs.Add("xtrig", func() telegraf.Input { return &Xtrig{} })
}
