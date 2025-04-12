package analyzer

type ImageQualityResult struct {
	IsBlurred    bool   `json:"is_blured"`
	BlurScore    string `json:"blur_score"`
	QualityScore string `json:"quality_score"`
	HasError     bool   `json:"has_error"`
	Filename     string `json:"filename"`
	Dimensions   string `json:"dimensions"`
	Channels     int    `json:"channels"`
	Message      string `json:"message"`
	Processed    string `json:"processed"`
}
