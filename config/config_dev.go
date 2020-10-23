// +build dev

package config

// Development ready configuration const based on the build tags.
const (
	SecurityAssessmentRootDir = "/home/androSecTest-Results"
	IsDebug                   = true
)

var IsVerboseLogRequired = true
