package file_test

import (
	"reflect"
	"testing"
	"time"
	
	. "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/suite"
	
	errorAVA "github.com/ver13/ava/pkg/common/error"
	fileAVA "github.com/ver13/ava/pkg/common/file"
)

type fileSuite struct {
	suite.Suite

	file fileAVA.FileI
}

func TestFileInit(t *testing.T) {
	suite.Run(t, new(fileSuite))
}

func (r *fileSuite) BeforeTest() {
	r.T().Log("BeforeTest")
}

func (r *fileSuite) AfterTest() {
	r.T().Log("AfterTest")
}

func (r *fileSuite) SetupSuite() {
	r.T().Log("SetupSuite")
	r.file = fileAVA.NewFile()
}

func (r *fileSuite) SetupTest() {
	r.T().Log("SetupTest")
}

func (r *fileSuite) TearDownSuite() {
	r.T().Log("TearDownSuite")
}

func (r *fileSuite) TearDownTest() {
	r.T().Log("TearDownTest")
}

func (r *fileSuite) TestFile_FileAppendBytes() {
	Convey("file append bytes ", r.T(), func(t *testing.T) {
		type args struct {
			filename string
			data     []byte
		}
		tests := []struct {
			name string
			args args
			want *errorAVA.Error
		}{
			// TODO: Add test cases.
		}
		for _, tt := range tests {
			r.T().Run(tt.name, func(t *testing.T) {
				f := &fileAVA.File{}
				if got := f.FileAppendBytes(tt.args.filename, tt.args.data); !reflect.DeepEqual(got, tt.want) {
					r.T().Errorf("FileAppendBytes() = %v, want %v", got, tt.want)
				}
			})
		}
	})
}

func (r *fileSuite) TestFile_FileAppendPrintf() {
	Convey("file append printf ", r.T(), func(t *testing.T) {
		type args struct {
			filename string
			format   string
			args     []interface{}
		}
		tests := []struct {
			name string
			args args
			want errorAVA.Error
		}{
			// TODO: Add test cases.
		}
		for _, tt := range tests {
			r.T().Run(tt.name, func(t *testing.T) {
				f := &fileAVA.File{}
				if got := f.FileAppendPrintf(tt.args.filename, tt.args.format, tt.args.args...); !reflect.DeepEqual(got, tt.want) {
					r.T().Errorf("FileAppendPrintf() = %v, want %v", got, tt.want)
				}
			})
		}
	})
}

func (r *fileSuite) TestFile_FileAppendString() {
	Convey("file append string ", r.T(), func(t *testing.T) {
		type args struct {
			filename string
			data     string
		}
		tests := []struct {
			name string
			args args
			want errorAVA.ErrorI
		}{
			// TODO: Add test cases.
		}
		for _, tt := range tests {
			r.T().Run(tt.name, func(t *testing.T) {
				f := &fileAVA.File{}
				if got := f.FileAppendString(tt.args.filename, tt.args.data); !reflect.DeepEqual(got, tt.want) {
					r.T().Errorf("FileAppendString() = %v, want %v", got, tt.want)
				}
			})
		}
	})
}

func (r *fileSuite) TestFile_FileCRC64() {
	Convey("file CRC64 ", r.T(), func(t *testing.T) {
		type args struct {
			filenameOrURL string
		}
		tests := []struct {
			name  string
			args  args
			want  uint64
			want1 errorAVA.ErrorI
		}{
			// TODO: Add test cases.
		}
		for _, tt := range tests {
			r.T().Run(tt.name, func(t *testing.T) {
				f := &fileAVA.File{}
				got, got1 := f.FileCRC64(tt.args.filenameOrURL)
				if got != tt.want {
					r.T().Errorf("FileCRC64() got = %v, want %v", got, tt.want)
				}
				if !reflect.DeepEqual(got1, tt.want1) {
					r.T().Errorf("FileCRC64() got1 = %v, want %v", got1, tt.want1)
				}
			})
		}
	})
}

func (r *fileSuite) TestFile_FileCopy() {
	Convey("file copy ", r.T(), func(t *testing.T) {
		type args struct {
			source string
			dest   string
		}
		tests := []struct {
			name string
			args args
			want errorAVA.ErrorI
		}{
			// TODO: Add test cases.
		}
		for _, tt := range tests {
			r.T().Run(tt.name, func(t *testing.T) {
				f := &fileAVA.File{}
				if got := f.FileCopy(tt.args.source, tt.args.dest); !reflect.DeepEqual(got, tt.want) {
					r.T().Errorf("FileCopy() = %v, want %v", got, tt.want)
				}
			})
		}
	})
}

func (r *fileSuite) TestFile_FileCopyDir() {
	Convey("file copy dir ", r.T(), func(t *testing.T) {
		type args struct {
			source string
			dest   string
		}
		tests := []struct {
			name string
			args args
			want errorAVA.ErrorI
		}{
			// TODO: Add test cases.
		}
		for _, tt := range tests {
			r.T().Run(tt.name, func(t *testing.T) {
				f := &fileAVA.File{}
				if got := f.FileCopyDir(tt.args.source, tt.args.dest); !reflect.DeepEqual(got, tt.want) {
					r.T().Errorf("FileCopyDir() = %v, want %v", got, tt.want)
				}
			})
		}
	})
}

func (r *fileSuite) TestFile_FileExists() {
	Convey("file exists ", r.T(), func(t *testing.T) {
		type args struct {
			filename string
		}
		tests := []struct {
			name string
			args args
			want *errorAVA.Error
		}{
			// TODO: Add test cases.
		}
		for _, tt := range tests {
			r.T().Run(tt.name, func(t *testing.T) {
				f := &fileAVA.File{}
				if got := f.FileExists(tt.args.filename); got != tt.want {
					r.T().Errorf("FileExists() = %v, want %v", got, tt.want)
				}
			})
		}
	})
}

func (r *fileSuite) TestFile_FileFind() {
	Convey("file find ", r.T(), func(t *testing.T) {
		type args struct {
			searchDirs []string
			filenames  []string
		}
		tests := []struct {
			name         string
			args         args
			wantFilePath string
			wantFound    bool
		}{
			// TODO: Add test cases.
		}
		for _, tt := range tests {
			r.T().Run(tt.name, func(t *testing.T) {
				f := &fileAVA.File{}
				gotFilePath, gotFound := f.FileFind(tt.args.searchDirs, tt.args.filenames...)
				if gotFilePath != tt.wantFilePath {
					r.T().Errorf("FileFind() gotFilePath = %v, want %v", gotFilePath, tt.wantFilePath)
				}
				if gotFound != tt.wantFound {
					r.T().Errorf("FileFind() gotFound = %v, want %v", gotFound, tt.wantFound)
				}
			})
		}
	})
}

func (r *fileSuite) TestFile_FileFindModified() {
	Convey("file find modified ", r.T(), func(t *testing.T) {
		type args struct {
			searchDirs []string
			filenames  []string
		}
		tests := []struct {
			name         string
			args         args
			wantFilePath string
			wantFound    bool
			wantModified time.Time
		}{
			// TODO: Add test cases.
		}
		for _, tt := range tests {
			r.T().Run(tt.name, func(t *testing.T) {
				f := &fileAVA.File{}
				gotFilePath, gotFound, gotModified := f.FileFindModified(tt.args.searchDirs, tt.args.filenames...)
				if gotFilePath != tt.wantFilePath {
					r.T().Errorf("FileFindModified() gotFilePath = %v, want %v", gotFilePath, tt.wantFilePath)
				}
				if gotFound != tt.wantFound {
					r.T().Errorf("FileFindModified() gotFound = %v, want %v", gotFound, tt.wantFound)
				}
				if !reflect.DeepEqual(gotModified, tt.wantModified) {
					r.T().Errorf("FileFindModified() gotModified = %v, want %v", gotModified, tt.wantModified)
				}
			})
		}
	})
}

func (r *fileSuite) TestFile_FileGetBytes() {
	Convey("file get bytes ", r.T(), func(t *testing.T) {
		type args struct {
			filenameOrURL string
			timeout       []time.Duration
		}
		tests := []struct {
			name  string
			args  args
			want  []byte
			want1 errorAVA.ErrorI
		}{
			// TODO: Add test cases.
		}
		for _, tt := range tests {
			r.T().Run(tt.name, func(t *testing.T) {
				f := &fileAVA.File{}
				got, got1 := f.FileGetBytes(tt.args.filenameOrURL, tt.args.timeout...)
				if !reflect.DeepEqual(got, tt.want) {
					r.T().Errorf("FileGetBytes() got = %v, want %v", got, tt.want)
				}
				if !reflect.DeepEqual(got1, tt.want1) {
					r.T().Errorf("FileGetBytes() got1 = %v, want %v", got1, tt.want1)
				}
			})
		}
	})
}

func (r *fileSuite) TestFile_FileGetCSV() {
	Convey("file get CSV ", r.T(), func(t *testing.T) {
		type args struct {
			filenameOrURL string
			timeout       []time.Duration
		}
		tests := []struct {
			name  string
			args  args
			want  [][]string
			want1 errorAVA.ErrorI
		}{
			// TODO: Add test cases.
		}
		for _, tt := range tests {
			r.T().Run(tt.name, func(t *testing.T) {
				f := &fileAVA.File{}
				got, got1 := f.FileGetCSV(tt.args.filenameOrURL, tt.args.timeout...)
				if !reflect.DeepEqual(got, tt.want) {
					r.T().Errorf("FileGetCSV() got = %v, want %v", got, tt.want)
				}
				if !reflect.DeepEqual(got1, tt.want1) {
					r.T().Errorf("FileGetCSV() got1 = %v, want %v", got1, tt.want1)
				}
			})
		}
	})
}

func (r *fileSuite) TestFile_FileGetGz() {
	Convey("file get Gz ", r.T(), func(t *testing.T) {
		type args struct {
			filenameOrURL string
		}
		tests := []struct {
			name  string
			args  args
			want  []byte
			want1 errorAVA.ErrorI
		}{
			// TODO: Add test cases.
		}
		for _, tt := range tests {
			r.T().Run(tt.name, func(t *testing.T) {
				f := &fileAVA.File{}
				got, got1 := f.FileGetGz(tt.args.filenameOrURL)
				if !reflect.DeepEqual(got, tt.want) {
					r.T().Errorf("FileGetGz() got = %v, want %v", got, tt.want)
				}
				if !reflect.DeepEqual(got1, tt.want1) {
					r.T().Errorf("FileGetGz() got1 = %v, want %v", got1, tt.want1)
				}
			})
		}
	})
}

func (r *fileSuite) TestFile_FileGetInflate() {
	Convey("file get inflate ", r.T(), func(t *testing.T) {
		type args struct {
			filenameOrURL string
		}
		tests := []struct {
			name  string
			args  args
			want  []byte
			want1 errorAVA.ErrorI
		}{
			// TODO: Add test cases.
		}
		for _, tt := range tests {
			r.T().Run(tt.name, func(t *testing.T) {
				f := &fileAVA.File{}
				got, got1 := f.FileGetInflate(tt.args.filenameOrURL)
				if !reflect.DeepEqual(got, tt.want) {
					r.T().Errorf("FileGetInflate() got = %v, want %v", got, tt.want)
				}
				if !reflect.DeepEqual(got1, tt.want1) {
					r.T().Errorf("FileGetInflate() got1 = %v, want %v", got1, tt.want1)
				}
			})
		}
	})
}

func (r *fileSuite) TestFile_FileGetJSON() {
	Convey("file get JSON ", r.T(), func(t *testing.T) {
		type args struct {
			filenameOrURL string
			timeout       []time.Duration
		}
		tests := []struct {
			name       string
			args       args
			wantResult interface{}
			wantErr    errorAVA.ErrorI
		}{
			// TODO: Add test cases.
		}
		for _, tt := range tests {
			r.T().Run(tt.name, func(t *testing.T) {
				f := &fileAVA.File{}
				gotResult, gotErr := f.FileGetJSON(tt.args.filenameOrURL, tt.args.timeout...)
				if !reflect.DeepEqual(gotResult, tt.wantResult) {
					r.T().Errorf("FileGetJSON() gotResult = %v, want %v", gotResult, tt.wantResult)
				}
				if !reflect.DeepEqual(gotErr, tt.wantErr) {
					r.T().Errorf("FileGetJSON() gotErr = %v, want %v", gotErr, tt.wantErr)
				}
			})
		}
	})
}

func (r *fileSuite) TestFile_FileGetLastLine() {
	Convey("file get last line ", r.T(), func(t *testing.T) {
		type args struct {
			filenameOrURL string
			timeout       []time.Duration
		}
		tests := []struct {
			name     string
			args     args
			wantLine string
			wantErr  errorAVA.ErrorI
		}{
			// TODO: Add test cases.
		}
		for _, tt := range tests {
			r.T().Run(tt.name, func(t *testing.T) {
				f := &fileAVA.File{}
				gotLine, gotErr := f.FileGetLastLine(tt.args.filenameOrURL, tt.args.timeout...)
				if gotLine != tt.wantLine {
					r.T().Errorf("FileGetLastLine() gotLine = %v, want %v", gotLine, tt.wantLine)
				}
				if !reflect.DeepEqual(gotErr, tt.wantErr) {
					r.T().Errorf("FileGetLastLine() gotErr = %v, want %v", gotErr, tt.wantErr)
				}
			})
		}
	})
}

func (r *fileSuite) TestFile_FileGetLines() {
	Convey("file get lines ", r.T(), func(t *testing.T) {
		type args struct {
			filenameOrURL string
			timeout       []time.Duration
		}
		tests := []struct {
			name      string
			args      args
			wantLines []string
			wantErr   errorAVA.ErrorI
		}{
			// TODO: Add test cases.
		}
		for _, tt := range tests {
			r.T().Run(tt.name, func(t *testing.T) {
				f := &fileAVA.File{}
				gotLines, gotErr := f.FileGetLines(tt.args.filenameOrURL, tt.args.timeout...)
				if !reflect.DeepEqual(gotLines, tt.wantLines) {
					r.T().Errorf("FileGetLines() gotLines = %v, want %v", gotLines, tt.wantLines)
				}
				if !reflect.DeepEqual(gotErr, tt.wantErr) {
					r.T().Errorf("FileGetLines() gotErr = %v, want %v", gotErr, tt.wantErr)
				}
			})
		}
	})
}

func (r *fileSuite) TestFile_FileGetNonEmptyLines() {
	Convey("file get non empty lines ", r.T(), func(t *testing.T) {
		type args struct {
			filenameOrURL string
			timeout       []time.Duration
		}
		tests := []struct {
			name      string
			args      args
			wantLines []string
			wantErr   errorAVA.ErrorI
		}{
			// TODO: Add test cases.
		}
		for _, tt := range tests {
			r.T().Run(tt.name, func(t *testing.T) {
				f := &fileAVA.File{}
				gotLines, gotErr := f.FileGetNonEmptyLines(tt.args.filenameOrURL, tt.args.timeout...)
				if !reflect.DeepEqual(gotLines, tt.wantLines) {
					r.T().Errorf("FileGetNonEmptyLines() gotLines = %v, want %v", gotLines, tt.wantLines)
				}
				if !reflect.DeepEqual(gotErr, tt.wantErr) {
					r.T().Errorf("FileGetNonEmptyLines() gotErr = %v, want %v", gotErr, tt.wantErr)
				}
			})
		}
	})
}

func (r *fileSuite) TestFile_FileGetString() {
	Convey("file get string ", r.T(), func(t *testing.T) {
		type args struct {
			filenameOrURL string
			timeout       []time.Duration
		}
		tests := []struct {
			name  string
			args  args
			want  string
			want1 errorAVA.ErrorI
		}{
			// TODO: Add test cases.
		}
		for _, tt := range tests {
			r.T().Run(tt.name, func(t *testing.T) {
				f := &fileAVA.File{}
				got, got1 := f.FileGetString(tt.args.filenameOrURL, tt.args.timeout...)
				if got != tt.want {
					r.T().Errorf("FileGetString() got = %v, want %v", got, tt.want)
				}
				if !reflect.DeepEqual(got1, tt.want1) {
					r.T().Errorf("FileGetString() got1 = %v, want %v", got1, tt.want1)
				}
			})
		}
	})
}

func (r *fileSuite) TestFile_FileGetXML() {
	Convey("file get XML ", r.T(), func(t *testing.T) {
		type args struct {
			filenameOrURL string
			timeout       []time.Duration
		}
		tests := []struct {
			name       string
			args       args
			wantResult interface{}
			wantErr    errorAVA.ErrorI
		}{
			// TODO: Add test cases.
		}
		for _, tt := range tests {
			r.T().Run(tt.name, func(t *testing.T) {
				f := &fileAVA.File{}
				gotResult, gotErr := f.FileGetXML(tt.args.filenameOrURL, tt.args.timeout...)
				if !reflect.DeepEqual(gotResult, tt.wantResult) {
					r.T().Errorf("FileGetXML() gotResult = %v, want %v", gotResult, tt.wantResult)
				}
				if !reflect.DeepEqual(gotErr, tt.wantErr) {
					r.T().Errorf("FileGetXML() gotErr = %v, want %v", gotErr, tt.wantErr)
				}
			})
		}
	})
}

func (r *fileSuite) TestFile_FileIsDir() {
	Convey("file is dir ", r.T(), func(t *testing.T) {
		type args struct {
			dirname string
		}
		tests := []struct {
			name string
			args args
			want bool
		}{
			// TODO: Add test cases.
		}
		for _, tt := range tests {
			r.T().Run(tt.name, func(t *testing.T) {
				f := &fileAVA.File{}
				if got := f.FileIsDir(tt.args.dirname); got != tt.want {
					r.T().Errorf("FileIsDir() = %v, want %v", got, tt.want)
				}
			})
		}
	})
}

func (r *fileSuite) TestFile_FileMD5Bytes() {
	Convey("file MD5 bytes ", r.T(), func(t *testing.T) {
		type args struct {
			filenameOrURL string
		}
		tests := []struct {
			name  string
			args  args
			want  []byte
			want1 errorAVA.ErrorI
		}{
			// TODO: Add test cases.
		}
		for _, tt := range tests {
			r.T().Run(tt.name, func(t *testing.T) {
				f := &fileAVA.File{}
				got, got1 := f.FileMD5Bytes(tt.args.filenameOrURL)
				if !reflect.DeepEqual(got, tt.want) {
					r.T().Errorf("FileMD5Bytes() got = %v, want %v", got, tt.want)
				}
				if !reflect.DeepEqual(got1, tt.want1) {
					r.T().Errorf("FileMD5Bytes() got1 = %v, want %v", got1, tt.want1)
				}
			})
		}
	})
}

func (r *fileSuite) TestFile_FileMD5String() {
	Convey("file MD5 string ", r.T(), func(t *testing.T) {
		type args struct {
			filenameOrURL string
		}
		tests := []struct {
			name  string
			args  args
			want  string
			want1 errorAVA.ErrorI
		}{
			// TODO: Add test cases.
		}
		for _, tt := range tests {
			r.T().Run(tt.name, func(t *testing.T) {
				f := &fileAVA.File{}
				got, got1 := f.FileMD5String(tt.args.filenameOrURL)
				if got != tt.want {
					r.T().Errorf("FileMD5String() got = %v, want %v", got, tt.want)
				}
				if !reflect.DeepEqual(got1, tt.want1) {
					r.T().Errorf("FileMD5String() got1 = %v, want %v", got1, tt.want1)
				}
			})
		}
	})
}

func (r *fileSuite) TestFile_FilePrintf() {
	Convey("file printf ", r.T(), func(t *testing.T) {
		type args struct {
			filename string
			format   string
			args     []interface{}
		}
		tests := []struct {
			name string
			args args
			want errorAVA.ErrorI
		}{
			// TODO: Add test cases.
		}
		for _, tt := range tests {
			r.T().Run(tt.name, func(t *testing.T) {
				f := &fileAVA.File{}
				if got := f.FilePrintf(tt.args.filename, tt.args.format, tt.args.args...); !reflect.DeepEqual(got, tt.want) {
					r.T().Errorf("FilePrintf() = %v, want %v", got, tt.want)
				}
			})
		}
	})
}

func (r *fileSuite) TestFile_FileScanf() {
	Convey("file scanf ", r.T(), func(t *testing.T) {
		type args struct {
			filename string
			format   string
			args     []interface{}
		}
		tests := []struct {
			name string
			args args
			want errorAVA.ErrorI
		}{
			// TODO: Add test cases.
		}
		for _, tt := range tests {
			r.T().Run(tt.name, func(t *testing.T) {
				f := &fileAVA.File{}
				if got := f.FileScanf(tt.args.filename, tt.args.format, tt.args.args...); !reflect.DeepEqual(got, tt.want) {
					r.T().Errorf("FileScanf() = %v, want %v", got, tt.want)
				}
			})
		}
	})
}

func (r *fileSuite) TestFile_FileSetBytes() {
	Convey("file set bytes ", r.T(), func(t *testing.T) {
		type args struct {
			filename string
			data     []byte
		}
		tests := []struct {
			name string
			args args
			want errorAVA.ErrorI
		}{
			// TODO: Add test cases.
		}
		for _, tt := range tests {
			r.T().Run(tt.name, func(t *testing.T) {
				f := &fileAVA.File{}
				if got := f.FileSetBytes(tt.args.filename, tt.args.data); !reflect.DeepEqual(got, tt.want) {
					r.T().Errorf("FileSetBytes() = %v, want %v", got, tt.want)
				}
			})
		}
	})
}

func (r *fileSuite) TestFile_FileSetCSV() {
	Convey("file set CSV ", r.T(), func(t *testing.T) {
		type args struct {
			filename string
			records  [][]string
		}
		tests := []struct {
			name string
			args args
			want errorAVA.ErrorI
		}{
			// TODO: Add test cases.
		}
		for _, tt := range tests {
			r.T().Run(tt.name, func(t *testing.T) {
				f := &fileAVA.File{}
				if got := f.FileSetCSV(tt.args.filename, tt.args.records); !reflect.DeepEqual(got, tt.want) {
					r.T().Errorf("FileSetCSV() = %v, want %v", got, tt.want)
				}
			})
		}
	})
}

func (r *fileSuite) TestFile_FileSetJSON() {
	Convey("file set JSON ", r.T(), func(t *testing.T) {
		type args struct {
			filename string
			data     interface{}
		}
		tests := []struct {
			name string
			args args
			want errorAVA.ErrorI
		}{
			// TODO: Add test cases.
		}
		for _, tt := range tests {
			r.T().Run(tt.name, func(t *testing.T) {
				f := &fileAVA.File{}
				if got := f.FileSetJSON(tt.args.filename, tt.args.data); !reflect.DeepEqual(got, tt.want) {
					r.T().Errorf("FileSetJSON() = %v, want %v", got, tt.want)
				}
			})
		}
	})
}

func (r *fileSuite) TestFile_FileSetJSONIndent() {
	Convey("file set JSON indent ", r.T(), func(t *testing.T) {
		type args struct {
			filename string
			data     interface{}
			indent   string
		}
		tests := []struct {
			name string
			args args
			want errorAVA.ErrorI
		}{
			// TODO: Add test cases.
		}
		for _, tt := range tests {
			r.T().Run(tt.name, func(t *testing.T) {
				f := &fileAVA.File{}
				if got := f.FileSetJSONIndent(tt.args.filename, tt.args.data, tt.args.indent); !reflect.DeepEqual(got, tt.want) {
					r.T().Errorf("FileSetJSONIndent() = %v, want %v", got, tt.want)
				}
			})
		}
	})
}

func (r *fileSuite) TestFile_FileSetLines() {
	Convey("file set lines ", r.T(), func(t *testing.T) {
		type args struct {
			filename string
			lines    []string
		}
		tests := []struct {
			name string
			args args
			want errorAVA.ErrorI
		}{
			// TODO: Add test cases.
		}
		for _, tt := range tests {
			r.T().Run(tt.name, func(t *testing.T) {
				f := &fileAVA.File{}
				if got := f.FileSetLines(tt.args.filename, tt.args.lines); !reflect.DeepEqual(got, tt.want) {
					r.T().Errorf("FileSetLines() = %v, want %v", got, tt.want)
				}
			})
		}
	})
}

func (r *fileSuite) TestFile_FileSetString() {
	Convey("file set string ", r.T(), func(t *testing.T) {
		type args struct {
			filename string
			data     string
		}
		tests := []struct {
			name string
			args args
			want errorAVA.ErrorI
		}{
			// TODO: Add test cases.
		}
		for _, tt := range tests {
			r.T().Run(tt.name, func(t *testing.T) {
				f := &fileAVA.File{}
				if got := f.FileSetString(tt.args.filename, tt.args.data); !reflect.DeepEqual(got, tt.want) {
					r.T().Errorf("FileSetString() = %v, want %v", got, tt.want)
				}
			})
		}
	})
}

func (r *fileSuite) TestFile_FileSetXML() {
	Convey("file set XML ", r.T(), func(t *testing.T) {
		type args struct {
			filename string
			data     interface{}
		}
		tests := []struct {
			name string
			args args
			want errorAVA.ErrorI
		}{
			// TODO: Add test cases.
		}
		for _, tt := range tests {
			r.T().Run(tt.name, func(t *testing.T) {
				f := &fileAVA.File{}
				if got := f.FileSetXML(tt.args.filename, tt.args.data); !reflect.DeepEqual(got, tt.want) {
					r.T().Errorf("FileSetXML() = %v, want %v", got, tt.want)
				}
			})
		}
	})
}

func (r *fileSuite) TestFile_FileSize() {
	Convey("file size ", r.T(), func(t *testing.T) {
		type args struct {
			filename string
		}
		tests := []struct {
			name string
			args args
			want int64
		}{
			// TODO: Add test cases.
		}
		for _, tt := range tests {
			r.T().Run(tt.name, func(t *testing.T) {
				f := &fileAVA.File{}
				if got := f.FileSize(tt.args.filename); got != tt.want {
					r.T().Errorf("FileSize() = %v, want %v", got, tt.want)
				}
			})
		}
	})
}

func (r *fileSuite) TestFile_FileTimeModified() {
	Convey("file time modified ", r.T(), func(t *testing.T) {
		type args struct {
			filename string
		}
		tests := []struct {
			name string
			args args
			want time.Time
		}{
			// TODO: Add test cases.
		}
		for _, tt := range tests {
			r.T().Run(tt.name, func(t *testing.T) {
				f := &fileAVA.File{}
				if got := f.FileTimeModified(tt.args.filename); !reflect.DeepEqual(got, tt.want) {
					r.T().Errorf("FileTimeModified() = %v, want %v", got, tt.want)
				}
			})
		}
	})
}

func (r *fileSuite) TestFile_FileTouch() {
	Convey("file touch ", r.T(), func(t *testing.T) {
		type args struct {
			filename string
		}
		tests := []struct {
			name string
			args args
			want errorAVA.ErrorI
		}{
			// TODO: Add test cases.
		}
		for _, tt := range tests {
			r.T().Run(tt.name, func(t *testing.T) {
				f := &fileAVA.File{}
				if got := f.FileTouch(tt.args.filename); !reflect.DeepEqual(got, tt.want) {
					r.T().Errorf("FileTouch() = %v, want %v", got, tt.want)
				}
			})
		}
	})
}

func (r *fileSuite) TestFile_FileUnmarshallJSON() {
	Convey("file unmarshall JSON ", r.T(), func(t *testing.T) {
		type args struct {
			filenameOrURL string
			result        interface{}
			timeout       []time.Duration
		}
		tests := []struct {
			name string
			args args
			want errorAVA.ErrorI
		}{
			// TODO: Add test cases.
		}
		for _, tt := range tests {
			r.T().Run(tt.name, func(t *testing.T) {
				f := &fileAVA.File{}
				if got := f.FileUnmarshallJSON(tt.args.filenameOrURL, tt.args.result, tt.args.timeout...); !reflect.DeepEqual(got, tt.want) {
					r.T().Errorf("FileUnmarshallJSON() = %v, want %v", got, tt.want)
				}
			})
		}
	})
}

func (r *fileSuite) TestFile_FileUnmarshallXML() {
	Convey("file unmarshall XML ", r.T(), func(t *testing.T) {
		type args struct {
			filenameOrURL string
			result        interface{}
			timeout       []time.Duration
		}
		tests := []struct {
			name string
			args args
			want errorAVA.ErrorI
		}{
			// TODO: Add test cases.
		}
		for _, tt := range tests {
			r.T().Run(tt.name, func(t *testing.T) {
				f := &fileAVA.File{}
				if got := f.FileUnmarshallXML(tt.args.filenameOrURL, tt.args.result, tt.args.timeout...); !reflect.DeepEqual(got, tt.want) {
					r.T().Errorf("FileUnmarshallXML() = %v, want %v", got, tt.want)
				}
			})
		}
	})
}

func (r *fileSuite) TestFile_ListDir() {
	Convey("List dir ", r.T(), func(t *testing.T) {
		type args struct {
			dir string
		}
		tests := []struct {
			name  string
			args  args
			want  []string
			want1 errorAVA.ErrorI
		}{
			// TODO: Add test cases.
		}
		for _, tt := range tests {
			r.T().Run(tt.name, func(t *testing.T) {
				f := &fileAVA.File{}
				got, got1 := f.ListDir(tt.args.dir)
				if !reflect.DeepEqual(got, tt.want) {
					r.T().Errorf("ListDir() got = %v, want %v", got, tt.want)
				}
				if !reflect.DeepEqual(got1, tt.want1) {
					r.T().Errorf("ListDir() got1 = %v, want %v", got1, tt.want1)
				}
			})
		}
	})
}

func (r *fileSuite) TestFile_ListDirDirectories() {
	Convey("List dir directories ", r.T(), func(t *testing.T) {
		type args struct {
			dir string
		}
		tests := []struct {
			name  string
			args  args
			want  []string
			want1 errorAVA.ErrorI
		}{
			// TODO: Add test cases.
		}
		for _, tt := range tests {
			r.T().Run(tt.name, func(t *testing.T) {
				f := &fileAVA.File{}
				got, got1 := f.ListDirDirectories(tt.args.dir)
				if !reflect.DeepEqual(got, tt.want) {
					r.T().Errorf("ListDirDirectories() got = %v, want %v", got, tt.want)
				}
				if !reflect.DeepEqual(got1, tt.want1) {
					r.T().Errorf("ListDirDirectories() got1 = %v, want %v", got1, tt.want1)
				}
			})
		}
	})
}

func (r *fileSuite) TestFile_ListDirFiles() {
	Convey("List dir files ", r.T(), func(t *testing.T) {
		type args struct {
			dir string
		}
		tests := []struct {
			name  string
			args  args
			want  []string
			want1 errorAVA.ErrorI
		}{
			// TODO: Add test cases.
		}
		for _, tt := range tests {
			r.T().Run(tt.name, func(t *testing.T) {
				f := &fileAVA.File{}
				got, got1 := f.ListDirFiles(tt.args.dir)
				if !reflect.DeepEqual(got, tt.want) {
					r.T().Errorf("ListDirFiles() got = %v, want %v", got, tt.want)
				}
				if !reflect.DeepEqual(got1, tt.want1) {
					r.T().Errorf("ListDirFiles() got1 = %v, want %v", got1, tt.want1)
				}
			})
		}
	})
}

func (r *fileSuite) TestNewFile() {
	Convey("New file  ", r.T(), func(t *testing.T) {
		tests := []struct {
			name string
			want fileAVA.File
		}{
			// TODO: Add test cases.
		}
		for _, tt := range tests {
			r.T().Run(tt.name, func(t *testing.T) {
				if got := fileAVA.NewFile(); !reflect.DeepEqual(got, tt.want) {
					r.T().Errorf("NewFile() = %v, want %v", got, tt.want)
				}
			})
		}
	})
}
