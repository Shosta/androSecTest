// +build prod

package config

// Production ready configuration const based on the build tags.
const (
	SecurityAssessmentRootDir = "/home/androSecTest-Results"
	IsDebug                   = false
)

var IsVerboseLogRequired = false
