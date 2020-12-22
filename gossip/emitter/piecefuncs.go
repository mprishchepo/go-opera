package emitter

import (
	"math"

	"github.com/Fantom-foundation/go-opera/gossip/emitter/piecefunc"
)

var (
	// confirmingEmitIntervalF is a piecewise function for validator confirming internal depending on a stake amount before him
	confirmingEmitIntervalF = []piecefunc.Dot{
		{
			X: 0,
			Y: 1.0 * piecefunc.DecimalUnit,
		},
		{
			X: 0.33 * piecefunc.DecimalUnit,
			Y: 1.05 * piecefunc.DecimalUnit,
		},
		{
			X: 0.66 * piecefunc.DecimalUnit,
			Y: 1.20 * piecefunc.DecimalUnit,
		},
		{
			X: 0.8 * piecefunc.DecimalUnit,
			Y: 1.5 * piecefunc.DecimalUnit,
		},
		{ // validators >0.8 emit confirming events very rarely
			X: 1.0 * piecefunc.DecimalUnit,
			Y: 1000.0 * piecefunc.DecimalUnit,
		},
	}
	// eventMetricF is a piecewise function for validator's event metric diff depending on a number of newly observed events
	scalarUpdMetricF = []piecefunc.Dot{
		{
			X: 0,
			Y: 0,
		},
		{ // first observed event gives a major metric diff
			X: 1.0 * piecefunc.DecimalUnit,
			Y: 0.66 * piecefunc.DecimalUnit,
		},
		{ // second observed event gives a minor diff
			X: 2.0 * piecefunc.DecimalUnit,
			Y: 0.8 * piecefunc.DecimalUnit,
		},
		{ // other observed event give only a subtle diff
			X: 8.0 * piecefunc.DecimalUnit,
			Y: 0.99 * piecefunc.DecimalUnit,
		},
		{
			X: math.MaxUint32 * piecefunc.DecimalUnit,
			Y: 1.0 * piecefunc.DecimalUnit,
		},
	}
	// eventMetricF is a piecewise function for event metric adjustment depending on a non-adjusted event metric
	eventMetricF = []piecefunc.Dot{
		{ // event metric is never zero
			X: 0,
			Y: 0.005 * piecefunc.DecimalUnit,
		},
		{ // if metric is below ~0.2, then validator shouldn't emit event unless waited very long
			X: 0.2 * piecefunc.DecimalUnit,
			Y: 0.05 * piecefunc.DecimalUnit,
		},
		{
			X: 0.3 * piecefunc.DecimalUnit,
			Y: 0.22 * piecefunc.DecimalUnit,
		},
		{ // ~0.3-0.5 is an optimal metric to emit an event
			X: 0.4 * piecefunc.DecimalUnit,
			Y: 0.45 * piecefunc.DecimalUnit,
		},
		{
			X: 1.0 * piecefunc.DecimalUnit,
			Y: 1.0 * piecefunc.DecimalUnit,
		},
		{ // event metric is never above 1.0
			X: math.MaxUint32 * piecefunc.DecimalUnit,
			Y: 1.0 * piecefunc.DecimalUnit,
		},
	}
)
