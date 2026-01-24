package constants

const (
	ErrNoCommands     = "no commands provided"
	ErrUnknownCommand = "unknown command: %s"
	ErrInvalidArgs    = "invalid arguments"
	ErrConfigLoad     = "failed to load configuration"
	ErrFileOperation  = "file operation failed"
	ErrNetwork        = "network operation failed"
)

const (
	CmdTrain    = "train"
	CmdDownload = "download"
	CmdPredict  = "predict"
)

const (
	ExitSuccess = iota
	ExitError
	ExitInvalidArgs
	ExitConfigError
	ExitFileError
	ExitNetworkError
)
