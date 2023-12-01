// custom_processor.go
package hw_processor

import (
	"context"
	"fmt"
	"go.opentelemetry.io/otel/api/metric"
)

// HelloWorldProcessor is a custom OpenTelemetry processor that processes metrics without adding new labels.
type HelloWorldProcessor struct{}

// ProcessStart is called when processing starts.
func (p *HelloWorldProcessor) ProcessStart(ctx context.Context, record metric.Record) (context.Context, metric.Int64Number, metric.Float64Number) {
	fmt.Println("Processing metrics: Hello, World! Start")

	// You can perform any processing on the metrics if needed

	return ctx, metric.NewInt64Number(0), metric.NewFloat64Number(0)
}

// ProcessEnd is called when processing ends.
func (p *HelloWorldProcessor) ProcessEnd(ctx context.Context, startCtx context.Context, record metric.Record) {
	fmt.Println("Processing metrics: Hello, World! End")

	// You can perform any processing on the metrics if needed
}
