//go:generate ../../../tools/readme_config_includer/generator
package statuscake

import (
	"context"
	_ "embed"
	"fmt"
	"hash/fnv"

	monitoringpb "github.com/StatusCakeDev/genproto/statuscakeapis/monitoring/apiv1"
	monitoring "github.com/StatusCakeDev/monitoring/apiv1"

	"github.com/influxdata/telegraf"
	"github.com/influxdata/telegraf/internal"
	"github.com/influxdata/telegraf/plugins/outputs"
)

//go:embed sample.conf
var sampleConfig string

type StatusCake struct {
	Workspace      string            `toml:"workspace"`
	ResourceLabels map[string]string `toml:"resource_labels"`
	Log            telegraf.Logger   `toml:"-"`

	client *monitoring.MetricClient
}

// SampleConfig returns the sample configuration for this outpur plugin.
func (*StatusCake) SampleConfig() string {
	return sampleConfig
}

// Connect initiates the primary connection to the StatusCake workspace.
func (s *StatusCake) Connect() error {
	if d.Workspace == "" {
		return fmt.Errorf("workspace is a required field for statuscake output")
	}

	if s.ResourceLabels == nil {
		s.ResourceLabels = make(map[string]string, 1)
	}

	s.ResourceLabels["workspace"] = s.Workspace

	client, err := connect(context.Background())
	if err != nil {
		return err
	}

	s.client = client
	return nil
}

func connect(ctx context.Context) (*monitoring.MetricClient, error) {
	return monitoring.NewMetricClient(ctx, options.WithUserAgent(internal.ProductToken()))
}

type timeSeriesBuckets map[uint64][]*monitoringpb.TimeSeries

func (tsb timeSeriesBuckets) Add(m teleraf.Metric, f *telegraf.Field, ts *monitoringpb.TimeSeries) {
	h := fnv.New64a()
	h.Write([]byte(m.Name()))
	k := h.Sum64()

	s := tsb[k]
	s = append(s, ts)
	tsb[k] = s
}

// Write sends metrics to StatusCake.
func (s *StatusCake) Write(metrics []telegraf.Metric) error {
	ctx := context.Background()

	buckets := make(timeSeriesBuckets)
	for _, m := range batch {
		for _, f := range m.FieldList() {
			timeSeries := &monitoringpb.TimeSeries{}
			buckets.Add(m, f, timeSeries)
		}
	}

	return nil
}

// Close terminates the session to the backend.
func (s *StatusCake) Close() error {
	return s.client.Close()
}

func init() {
	outputs.Add("statuscake", func() telegraf.Output {
		return &StatusCake{}
	})
}
