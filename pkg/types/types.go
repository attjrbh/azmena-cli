package types

type Duration struct {
	Seconds int
	Minutes int
	Hours   int
}

type DurationInfo struct {
	OkPathsCount  int
	FailedPaths   []string
	TotalDuration float64
}

type FileStatus struct {
	Ok       bool
	Path     string
	Duration float64
}
