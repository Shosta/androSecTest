// +build prod

package config

// Production ready configuration const based on the build tags.
const (
	SecurityAssessmentRootDir = "/home/shosta/android/security"
	IsDebug                   = false
)

var IsVerboseLogRequired = false
