package filters

import "github.com/mishankov/logman"

// LevelFilter implements Filter interface.
type LevelFilter struct {
	level logman.LogLevel
}

// NewLevelFilter creates a new LevelFilter.
func NewLevelFilter(level logman.LogLevel) LevelFilter {
	return LevelFilter{level: level}
}

// Filter returns true for messages that are on the same or higher log level than LevelFilter.level.
func (lf LevelFilter) Filter(level logman.LogLevel, _, _ string) bool {
	return level >= lf.level
}
