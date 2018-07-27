package variables

// Attacks folder names to have them gathered in one single place.
const (
	AttacksDir            = "/attacks"
	SourcePackageDir      = "/sourcePackage"
	UnzippedPackageDir    = "/unzippedPackage"
	DisassemblePackageDir = "/disassemblePackage"
	DecompiledPackageDir  = "/decompiledPackage"
	LeakagesDir           = "/leakages"
	DebuggablePackageDir  = "/debuggablePackage"
	InsecureBackupDir     = "/insecureBackup"
	InsecureLoggingDir    = "/insecureLogging"
	InsecureStorageDir    = "/insecureStorage"
)

// Color const to display color on the terminal command.
const (
	Header    = "\033[95m"
	Blue      = "\033[94m"
	Green     = "\033[92m"
	Orange    = "\033[93m"
	Red       = "\033[91m"
	Endc      = "\033[0m"
	Bold      = "\033[1m"
	Underline = "\033[4m"
)
