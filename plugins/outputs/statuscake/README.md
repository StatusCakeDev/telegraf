# StatusCake Output Plugin

This plugin writes to the [StatusCake Metrics API][metrics] and requires
authentication be managed using [application default credentials][adc].

## Global configuration options <!-- @/docs/includes/plugin_config.md -->

In addition to the plugin-specific configuration settings, plugins support
additional global and plugin configuration settings. These settings are used to
modify metrics, tags, and field or create aliases and configure ordering, etc.
See the [CONFIGURATION.md][CONFIGURATION.md] for more details.

[CONFIGURATION.md]: ../../../docs/CONFIGURATION.md

## Configuration

```toml @sample.conf
# Configuration for StatusCake to send metrics to.
[[outputs.statuscake]]
  ## StatusCake workspace
  workspace = "default"

  ## Additional resource labels
  # [outputs.statuscake.resource_labels]
  #   node_id = "$HOSTNAME"
  #   location = "london"
```

## Metrics

WIP

[metrics]: https://developers.statuscake.com/guides/metrics/#submit-metrics
[adc]: https://developers.statuscake.com/guides/application-default-credentials
