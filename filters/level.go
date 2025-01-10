package filters

import "github.com/mishankov/logman"

type LevelFilter struct {
	level logman.LogLevel
}

func NewLevelFilter(level logman.LogLevel) LevelFilter {
	return LevelFilter{level: level}
}

func (lf LevelFilter) Filter(level logman.LogLevel, _, _ string) bool {
	return level >= lf.level
}
