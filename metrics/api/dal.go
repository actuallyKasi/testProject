package api
import (
	"github.com/akhenakh/statgo"
)
var s *statgo.Stat
func init(){
	s = statgo.NewStat()
}
func GetCpuMetrics() (*statgo.CPUStats ) {
	c := s.CPUStats()
	return c
}

func GetMemMetrics() *statgo.MemStats{
	m := s.MemStats()
	return m
}

func GetDiskIOMetrics() []*statgo.DiskIOStats{
	d := s.DiskIOStats()
	return d
}

func GetNwIOMetrics() []*statgo.NetIOStats{
	n := s.NetIOStats()
	return n
}