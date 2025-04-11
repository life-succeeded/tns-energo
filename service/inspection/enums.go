package inspection

type Resolution int

const (
	LimitedResolution Resolution = iota
	StoppedResolution
	ResumedResolution
)

type MethodBy int

const (
	Consumer MethodBy = iota
	Inspector
)

type ReasonType int

const (
	NotIntroduced ReasonType = iota
	ConsumerLimited
	InspectorLimited
	Resumed
)
