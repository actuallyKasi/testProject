package api

import "github.com/akhenakh/statgo"

type CpuMetricsResponseObject struct{
	statgo.CPUStats
}

type MemMetricsResponseObject struct{
	Total int
	Free int
	Used int
	Cache int
	SwapTotal int
	SwapUsed int
	SwapFree int
}

