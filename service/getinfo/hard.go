package getinfo

import (
	"context"

	"github.com/gogf/gf/v2/os/glog"
	"github.com/jaypipes/ghw"
	"github.com/jaypipes/ghw/pkg/block"
	"github.com/jaypipes/ghw/pkg/cpu"
	"github.com/jaypipes/ghw/pkg/memory"
	"github.com/jaypipes/ghw/pkg/topology"
)

type hardinfo struct {
	Cpu      *cpu.Info      `json:"cpu,omitempty"`
	Memory   *memory.Info   `json:"memory,omitempty"`
	Block    *block.Info    `json:"block,omitempty"`
	Topology *topology.Info `json:"topology,omitempty"`
}

func Hard(ctx context.Context) (interface{}, error) {
	cpu, err := ghw.CPU()
	if err != nil {
		glog.Errorf(ctx, "Error getting CPU info: %v", err)
	}

	memory, err := ghw.Memory()
	if err != nil {
		glog.Errorf(ctx, "Error getting memory info: %v", err)
	}

	block, err := ghw.Block()
	if err != nil {
		glog.Errorf(ctx, "Error getting block info: %v", err)
	}

	topology, err := ghw.Topology()
	if err != nil {
		glog.Errorf(ctx, "Error getting block info: %v", err)
	}

	hardinfo := hardinfo{
		Cpu:      cpu,
		Memory:   memory,
		Block:    block,
		Topology: topology,
	}

	return hardinfo, nil
}
