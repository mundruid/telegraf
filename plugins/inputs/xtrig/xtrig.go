package xtrig

import (
	"github.com/influxdata/telegraf"
	"github.com/influxdata/telegraf/plugins/inputs"
	"math"
)

type Xtrig struct {
	x         float64
	Amplitude float64
}

var XtrigConfig = `
	## Set the amplitude
	amplitude = 10.0
`

func (s *Xtrig) SampleConfig() string {
	return XtrigConfig
}

func (s *Xtrig) Description() string {
	return "Inserts sine and cosine waves for demonstration purposes"
}

func (s *Xtrig) Gather(acc telegraf.Accumulator) error {
	sinner := math.Sin((s.x*math.Pi)/5.0) * s.Amplitude
	cosinner := math.Cos((s.x*math.Pi)/5.0) * s.Amplitude

	fields := make(map[string]interface{})
	fields["sine"] = sinner
	fields["cosine"] = cosinner

	tags := make(map[string]string)

	s.x += 1.0
	// acc: instantiation of accumulator
	acc.AddFields("trig", fields, tags)

	return nil
}

func init() {
	inputs.Add("xtrig", func() telegraf.Input { return &Xtrig{x: 0.0} })
}
