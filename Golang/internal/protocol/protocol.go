package protocol

import (
	"ELP/internal/randomwalk"
)

type Request struct {
	NumWalks    int
	DurationSec int
	Graph       string
}

type NodeProb struct {
	Node int64
	Prob float64
}

type Response struct {
	DurationSec int
	StepsMono   int64
	StepsMulti  int64
	Speedup     float64
	TopNodes    []randomwalk.NodeProb
}
