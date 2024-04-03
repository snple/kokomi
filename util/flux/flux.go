package flux

const TemplateBasic = `from(bucket: "{{.bucket}}")
|> range(start: {{.start}}, stop: {{.stop}})
|> filter(fn: (r) => r["_field"] == "{{.id}}")
|> aggregateWindow(every: {{.window}}, fn: {{.fn}}, createEmpty: false)
|> drop(columns: ["_measurement", "_field", "_start", "_stop", "source"])`

const TemplateFill = `from(bucket: "{{.bucket}}")
|> range(start: {{.start}}, stop: {{.stop}})
|> filter(fn: (r) => r["_field"] == "{{.id}}")
|> aggregateWindow(every: {{.window}}, fn: {{.fn}}, createEmpty: true)
|> fill(column: "_value", usePrevious: true)
|> drop(columns: ["_measurement", "_field", "_start", "_stop", "source"])`

const TemplateLinear = `import "interpolate"
from(bucket: "{{.bucket}}")
|> range(start: {{.start}}, stop: {{.stop}})
|> filter(fn: (r) => r["_field"] == "{{.id}}")
|> aggregateWindow(every: {{.window}}, fn: {{.fn}}, createEmpty: false)
|> interpolate.linear(every: {{.window}})
|> drop(columns: ["_measurement", "_field", "_start", "_stop", "source"])`
