package accounting

import "strconv"

type EventType int

const (
	TransferAdd EventType = iota
	TransferDel
	Progress
)

var Event func(t EventType, stat *StatsInfo)

func init() {
	Event = func(t EventType, stat *StatsInfo) {}
}
func (s *StatsInfo) TransferringCount() int {
	return s.transferring.count()
}
func (s *StatsInfo) Group() string {
	return s.group
}
func (s *StatsInfo) JobId() (int64, error) {
	if len(s.group) <= 4 {
		return 0, nil
	}
	jobId, err := strconv.Atoi(s.group[4:])
	if err != nil {
		return 0, err
	}
	return int64(jobId), nil
}
