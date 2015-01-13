
# filter

 Statsite filter command, allowing you to whitelist and rename metrics based on a simple JSON configuration file containing regular expression rules. This allows you to send specific metrics to a given service.

## Examples

Sample input:

```
gauges.api-1.memory.percent|80|1421164536386
gauges.api-2.memory.percent|30|1421164536386
gauges.ingestion-2.memory.percent|50|1421164536386
gauges.api-3.memory.percent|50|1421164536386
gauges.cdn-1.memory.percent|50|1421164536386
gauges.cdn-2.memory.percent|50|1421164536386
counts.app.signups|10|1421164536386
gauges.ingestion-1.memory.percent|50|1421164536386
```

Strip all `counts.`, `gauges.`, and `.timers` prefixes:

```json
{
  "^(?:counts|gauges|timers)\\.(.*)": "$1"
}
```

Outputs:

```
api-1.memory.percent|80|1421164536386
api-2.memory.percent|30|1421164536386
ingestion-2.memory.percent|50|1421164536386
api-3.memory.percent|50|1421164536386
cdn-1.memory.percent|50|1421164536386
cdn-2.memory.percent|50|1421164536386
app.signups|10|1421164536386
ingestion-1.memory.percent|50|1421164536386
```

Include all counters as-is, but strip `gauges.` prefix from api metrics:

```json
{
  "^gauges\\.(api-.*)": "$1",
  "^counts\\.": true
}
```

Outputs:

```
api-1.memory.percent|80|1421164536386
api-2.memory.percent|30|1421164536386
api-3.memory.percent|50|1421164536386
counts.app.signups|10|1421164536386
```

Clean up some metric names for services like Stathat which have minimal realestate:

```json
{
  "^gauges\\.(api-.*)\\.(memory|disk)\\.percent": "$1.$2"
}
```

Outputs:

```
api-1.memory|80|1421164536386
api-2.memory|30|1421164536386
api-3.memory|50|1421164536386
```

# License

MIT