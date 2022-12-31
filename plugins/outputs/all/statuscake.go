//go:build !custom || outputs || outputs.statuscake

package all

import _ "github.com/influxdata/telegraf/plugins/outputs/statuscake" // register plugin
