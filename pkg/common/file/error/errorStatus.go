package error

const (
	statusFileNotFount              = 1
	statusFileGeneral               = 2
	statusFileNotExist              = 3
	statusFileCopy                  = 4
	statusOpenFile                  = 5
	statusReadFile                  = 6
	statusWriteFile                 = 7
	statusUnmarshalFile             = 8
	statusMarshalFile               = 9
	statusCreateFile                = 10
	statusOpenDir                   = 11
	statusListDir                   = 12
	statusCreateDir                 = 13
	statusCopyDir                   = 14
	statusFileNotIsDir              = 15
	statusCloseFile                 = 16
	statusModifyFile                = 17
	statusCurrentDirectoryPathWrong = 18
	statusUnknown                   = 19
)

var statusText = map[int]string{
	statusFileNotFount:              "file not fount.",
	statusFileGeneral:               "file general error.",
	statusFileNotExist:              "file not exist.",
	statusFileCopy:                  "file copy error.",
	statusOpenFile:                  "Open file error.",
	statusReadFile:                  "Read file error.",
	statusWriteFile:                 "Write file error.",
	statusUnmarshalFile:             "Unmarshal file error.",
	statusMarshalFile:               "Marshal file error.",
	statusCreateFile:                "Create file error.",
	statusOpenDir:                   "Open dir error.",
	statusListDir:                   "List dir error.",
	statusCreateDir:                 "Create dir error.",
	statusCopyDir:                   "Copy dir error.",
	statusFileNotIsDir:              "File not is dir error.",
	statusCloseFile:                 "Close file error.",
	statusModifyFile:                "Modify file error.",
	statusCurrentDirectoryPathWrong: "Current directory path wrong.",
	statusUnknown:                   "Error unknown.",
}

// statusTextFunc returns a text for the General status code.
// It returns the empty string if the code is unknown.
func statusTextFunc(code int) string {
	return statusText[code]
}
