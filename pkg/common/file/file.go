package file

import (
	"bytes"
	"compress/flate"
	"compress/zlib"
	"crypto/md5"
	"encoding/csv"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"hash/crc64"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	// "strconv"
	"strings"
	"time"

	errorAVA "github.com/ver13/ava/pkg/common/error"
	errorFileAVA "github.com/ver13/ava/pkg/common/file/error"
)

type File struct {
}

func NewFile() FileI {
	return &File{}
}

func (f *File) FileGetBytes(filenameOrURL string, timeout ...time.Duration) ([]byte, *errorAVA.Error) {
	if strings.Contains(filenameOrURL, "://") {
		if strings.Index(filenameOrURL, "file://") == 0 {
			filenameOrURL = filenameOrURL[len("file://"):]
		} else {
			client := http.DefaultClient
			if len(timeout) > 0 {
				client = &http.Client{Timeout: timeout[0]}
			}
			r, err := client.Get(filenameOrURL)
			if err != nil {
				return nil, errorFileAVA.ReadFile(err, filenameOrURL)
			}
			defer r.Body.Close()
			if r.StatusCode < 200 || r.StatusCode > 299 {
				return nil, errorFileAVA.Unknown(fmt.Errorf("%d: %s", r.StatusCode, http.StatusText(r.StatusCode)), filenameOrURL)
			}
			data, errRead := ioutil.ReadAll(r.Body)
			if errRead != nil {
				return nil, errorFileAVA.ReadFile(errRead, fmt.Sprintf("Body: %s", r.Body))
			}
			return data, nil
		}
	}
	data, errRead := ioutil.ReadFile(filenameOrURL)
	if errRead != nil {
		return nil, errorFileAVA.ReadFile(errRead, filenameOrURL)
	}
	return data, nil
}

func (f *File) FileSetBytes(filename string, data []byte) *errorAVA.Error {
	if err := ioutil.WriteFile(filename, data, 0660); err != nil {
		return errorFileAVA.WriteFile(err, filename)
	}
	return nil
}

func (f *File) FileAppendBytes(filename string, data []byte) *errorAVA.Error {
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0660)
	if err != nil {
		return errorFileAVA.OpenFile(err, filename)
	}
	defer file.Close()
	_, err = file.Write(data)
	return errorFileAVA.WriteFile(err, filename)
}

func (f *File) FileGetString(filenameOrURL string, timeout ...time.Duration) (string, *errorAVA.Error) {
	bytes, err := f.FileGetBytes(filenameOrURL, timeout...)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func (f *File) FileSetString(filename string, data string) *errorAVA.Error {
	return f.FileSetBytes(filename, []byte(data))
}

func (f *File) FileAppendString(filename string, data string) *errorAVA.Error {
	return f.FileAppendBytes(filename, []byte(data))
}

func (f *File) FileGetJSON(filenameOrURL string, timeout ...time.Duration) (result interface{}, err *errorAVA.Error) {
	err = f.FileUnmarshallJSON(filenameOrURL, &result, timeout...)
	return result, err
}

func (f *File) FileUnmarshallJSON(filenameOrURL string, result interface{}, timeout ...time.Duration) *errorAVA.Error {
	data, err := f.FileGetBytes(filenameOrURL, timeout...)
	if err != nil {
		return err
	}
	if errUnmarshal := json.Unmarshal(data, result); errUnmarshal != nil {
		return errorFileAVA.UnmarshalFile(errUnmarshal, fmt.Sprintf("Type JSON. Data: %s", data))
	}
	return nil
}

func (f *File) FileSetJSON(filename string, data interface{}) *errorAVA.Error {
	bytes, err := json.Marshal(data)
	if err != nil {
		return errorFileAVA.MarshalFile(err, fmt.Sprintf("Type JSON. file: %s Data: %s", filename, data))
	}
	return f.FileSetBytes(filename, bytes)
}

func (f *File) FileSetJSONIndent(filename string, data interface{}, indent string) *errorAVA.Error {
	bytes, err := json.MarshalIndent(data, "", indent)
	if err != nil {
		return errorFileAVA.MarshalFile(err, fmt.Sprintf("Type JSONIndent. file: %s Data: %s", filename, data))
	}
	return f.FileSetBytes(filename, bytes)
}

func (f *File) FileGetXML(filenameOrURL string, timeout ...time.Duration) (result interface{}, err *errorAVA.Error) {
	err = f.FileUnmarshallXML(filenameOrURL, &result, timeout...)
	return result, err
}

func (f *File) FileUnmarshallXML(filenameOrURL string, result interface{}, timeout ...time.Duration) *errorAVA.Error {
	data, err := f.FileGetBytes(filenameOrURL, timeout...)
	if err != nil {
		return err
	}
	if errUnmarshal := xml.Unmarshal(data, result); errUnmarshal != nil {
		return errorFileAVA.UnmarshalFile(errUnmarshal, fmt.Sprintf("Type XML. Data: %s", data))
	}
	return nil
}

func (f *File) FileSetXML(filename string, data interface{}) *errorAVA.Error {
	bytes, err := xml.Marshal(data)
	if err != nil {
		return errorFileAVA.MarshalFile(err, fmt.Sprintf("Type XML. file: %s Data: %s", filename, data))
	}
	return f.FileSetBytes(filename, bytes)
}

func (f *File) FileGetCSV(filenameOrURL string, timeout ...time.Duration) ([][]string, *errorAVA.Error) {
	data, err := f.FileGetBytes(filenameOrURL, timeout...)
	if err != nil {
		return nil, err
	}
	reader := csv.NewReader(bytes.NewBuffer(data))
	dataRead, errRead := reader.ReadAll()
	if errRead != nil {
		return nil, errorFileAVA.ReadFile(errRead, filenameOrURL)
	}
	return dataRead, nil
}

func (f *File) FileSetCSV(filename string, records [][]string) *errorAVA.Error {
	file, err := os.Create(filename)
	if err != nil {
		return errorFileAVA.CreateFile(err, filename)
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	if err := writer.WriteAll(records); err != nil {
		return errorFileAVA.WriteFile(err, filename)
	}
	return nil
}

// FileGetLines returns a string slice with the text lines of filenameOrURL.
// The lines can be separated by \n or \r\n.
func (f *File) FileGetLines(filenameOrURL string, timeout ...time.Duration) (lines []string, err *errorAVA.Error) {
	data, err := f.FileGetBytes(filenameOrURL, timeout...)
	if err != nil {
		return nil, err
	}

	lastR := -1
	lastN := -1

	for i, c := range data {
		if c == '\r' {
			l := string(data[lastN+1 : i])
			lines = append(lines, l)
			lastR = i
		}
		if c == '\n' {
			if i != lastR+1 {
				l := string(data[lastN+1 : i])
				lines = append(lines, l)
			}
			lastN = i
		}
	}
	l := string(data[lastN+1:])
	lines = append(lines, l)

	return lines, nil
}

func (f *File) FileSetLines(filename string, lines []string) *errorAVA.Error {
	return f.FileSetString(filename, strings.Join(lines, "\n"))
}

// FileGetNonEmptyLines returns a string slice with the non empty text lines of filenameOrURL.
// The lines can be separated by \n or \r\n.
func (f *File) FileGetNonEmptyLines(filenameOrURL string, timeout ...time.Duration) (lines []string, err *errorAVA.Error) {
	data, err := f.FileGetBytes(filenameOrURL, timeout...)
	if err != nil {
		return nil, err
	}

	lastR := -1
	lastN := -1

	for i, c := range data {
		if c == '\r' {
			l := string(data[lastN+1 : i])
			if l != "" {
				lines = append(lines, l)
			}
			lastR = i
		}
		if c == '\n' {
			if i != lastR+1 {
				l := string(data[lastN+1 : i])
				if l != "" {
					lines = append(lines, l)
				}
			}
			lastN = i
		}
	}
	l := string(data[lastN+1:])
	if l != "" {
		lines = append(lines, l)
	}

	return lines, nil
}

// FileGetLastLine reads the last line from a file.
// In case of a network file, the whole file is read.
// In case of a local file, the last 64kb are read, so if the last line is longer than 64kb it is not returned completely.
// The first optional timeout is used for network files only.
func (f *File) FileGetLastLine(filenameOrURL string, timeout ...time.Duration) (line string, err *errorAVA.Error) {
	if strings.Index(filenameOrURL, "file://") == 0 {
		return f.FileGetLastLine(filenameOrURL[len("file://"):])
	}

	var data []byte

	if strings.Contains(filenameOrURL, "://") {
		data, err = f.FileGetBytes(filenameOrURL, timeout...)
		if err != nil {
			return "", err
		}
	} else {
		file, err := os.Open(filenameOrURL)
		if err != nil {
			return "", errorFileAVA.OpenFile(err, filenameOrURL)
		}
		defer file.Close()
		info, err := file.Stat()
		if err != nil {
			return "", errorFileAVA.ReadFile(err, filenameOrURL)
		}
		if start := info.Size() - 64*1024; start > 0 {
			file.Seek(start, io.SeekStart)
		}
		data, err = ioutil.ReadAll(file)
		if err != nil {
			return "", errorFileAVA.ReadFile(err, filenameOrURL)
		}
	}

	pos := bytes.LastIndex(data, []byte{'\n'})
	return string(data[pos+1:]), nil
}

// FileTimeModified returns the modified time of a file, or the zero time value in case of an error.
func (f *File) FileTimeModified(filename string) time.Time {
	info, err := os.Stat(filename)
	if err != nil {
		return time.Time{}
	}
	return info.ModTime()
}

func (f *File) FileExists(filename string) *errorAVA.Error {
	_, err := os.Stat(filename)
	if err != nil {
		errorFileAVA.FileNotFount(err, filename)
	}
	return nil
}

func (f *File) FileIsDir(dirname string) bool {
	info, err := os.Stat(dirname)
	return err == nil && info.IsDir()
}

func (f *File) FileFind(searchDirs []string, filenames ...string) (filePath string, found bool) {
	for _, dir := range searchDirs {
		for _, filename := range filenames {
			filePath = filepath.Join(dir, filename)
			if f.FileExists(filePath) != nil {
				return filePath, true
			}
		}
	}
	return "", false
}

func (f *File) FileFindModified(searchDirs []string, filenames ...string) (filePath string, found bool, modified time.Time) {
	for _, dir := range searchDirs {
		for _, filename := range filenames {
			filePath = filepath.Join(dir, filename)
			if t := f.FileTimeModified(filePath); !t.IsZero() {
				return filePath, true, t
			}
		}
	}
	return "", false, time.Time{}
}

func (f *File) FileTouch(filename string) *errorAVA.Error {
	if f.FileExists(filename) != nil {
		now := time.Now()
		if err := os.Chtimes(filename, now, now); err != nil {
			return errorFileAVA.ModifyFile(err, filename)
		}
		return nil
	}
	file, err := os.Create(filename)
	if err != nil {
		return errorFileAVA.CreateFile(err, filename)
	}
	if err := file.Close(); err != nil {
		return errorFileAVA.CloseFile(err, filename)
	}
	return nil
}

func (f *File) FileMD5String(filenameOrURL string) (string, *errorAVA.Error) {
	sum, err := f.FileMD5Bytes(filenameOrURL)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", sum), nil
}

func (f *File) FileMD5Bytes(filenameOrURL string) ([]byte, *errorAVA.Error) {
	data, err := f.FileGetBytes(filenameOrURL)
	if err != nil {
		return nil, err
	}
	hash := md5.New()
	_, errCopy := io.Copy(hash, bytes.NewBuffer(data))
	if errCopy != nil {
		return nil, errorFileAVA.FileCopy(errCopy, fmt.Sprintf("file: %s. Hash: %s.", filenameOrURL, hash))
	}
	return hash.Sum(nil), nil
}

var crc64Table *crc64.Table

func (f *File) FileCRC64(filenameOrURL string) (uint64, *errorAVA.Error) {
	data, err := f.FileGetBytes(filenameOrURL)
	if err != nil {
		return 0, err
	}
	if crc64Table == nil {
		crc64Table = crc64.MakeTable(crc64.ECMA)
	}
	return crc64.Checksum(data, crc64Table), nil
}

func (f *File) FileGetInflate(filenameOrURL string) ([]byte, *errorAVA.Error) {
	data, err := f.FileGetBytes(filenameOrURL)
	if err != nil {
		return nil, err
	}
	reader := flate.NewReader(bytes.NewBuffer(data))
	defer reader.Close()
	data, errRead := ioutil.ReadAll(reader)
	if errRead != nil {
		return nil, errorFileAVA.ReadFile(errRead, filenameOrURL)
	}
	return data, nil
}

func (f *File) FileGetGz(filenameOrURL string) ([]byte, *errorAVA.Error) {
	data, err := f.FileGetBytes(filenameOrURL)
	if err != nil {
		return nil, err
	}
	reader, errZlib := zlib.NewReader(bytes.NewBuffer(data))
	if errZlib != nil {
		return nil, errorFileAVA.ReadFile(errZlib, filenameOrURL)
	}
	defer reader.Close()
	data, errRead := ioutil.ReadAll(reader)
	if errRead != nil {
		return nil, errorFileAVA.ReadFile(errRead, filenameOrURL)
	}
	return data, nil
}

// FileSize returns the size of a file or zero in case of an error.
func (f *File) FileSize(filename string) int64 {
	info, err := os.Stat(filename)
	if err != nil {
		return 0
	}
	return info.Size()
}

func (f *File) FilePrintf(filename, format string, args ...interface{}) *errorAVA.Error {
	file, err := os.OpenFile(filename, os.O_WRONLY, 0660)
	if err == nil {
		_, err = fmt.Fprintf(file, format, args...)
		file.Close()
	}
	return errorFileAVA.OpenFile(err, filename)
}

func (f *File) FileAppendPrintf(filename, format string, args ...interface{}) *errorAVA.Error {
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0660)
	if err == nil {
		_, err = fmt.Fprintf(file, format, args...)
		file.Close()
	}
	return errorFileAVA.OpenFile(err, filename)
}

func (f *File) FileScanf(filename, format string, args ...interface{}) *errorAVA.Error {
	file, err := os.OpenFile(filename, os.O_RDONLY, 0660)
	if err == nil {
		_, err = fmt.Fscanf(file, format, args...)
		file.Close()
	}
	return errorFileAVA.OpenFile(err, filename)
}

func (f *File) ListDir(dir string) ([]string, *errorAVA.Error) {
	d, err := os.Open(dir)
	if err != nil {
		return nil, errorFileAVA.OpenDir(err, dir)
	}
	defer d.Close()
	data, errRead := d.Readdirnames(-1)
	if errRead != nil {
		return nil, errorFileAVA.ListDir(errRead, dir)
	}
	return data, nil
}

func (f *File) ListDirFiles(dir string) ([]string, *errorAVA.Error) {
	d, err := os.Open(dir)
	if err != nil {
		return nil, errorFileAVA.OpenDir(err, dir)
	}
	defer d.Close()
	fileInfos, err := d.Readdir(-1)
	if err != nil {
		return nil, errorFileAVA.ListDir(err, dir)
	}
	result := make([]string, 0, len(fileInfos))
	for i := range fileInfos {
		if !fileInfos[i].IsDir() {
			result = append(result, fileInfos[i].Name())
		}
	}
	return result, nil
}

func (f *File) ListDirDirectories(dir string) ([]string, *errorAVA.Error) {
	d, err := os.Open(dir)
	if err != nil {
		return nil, errorFileAVA.OpenDir(err, dir)
	}
	defer d.Close()
	fileInfos, err := d.Readdir(-1)
	if err != nil {
		return nil, errorFileAVA.ListDir(err, dir)
	}
	result := make([]string, 0, len(fileInfos))
	for i := range fileInfos {
		if fileInfos[i].IsDir() {
			result = append(result, fileInfos[i].Name())
		}
	}
	return result, nil
}

// FileCopy copies file source to destination dest.
// Based on Jaybill McCarthy's code which can be found at http://jayblog.jaybill.com/post/id/26
func (f *File) FileCopy(source string, dest string) *errorAVA.Error {
	sourceFile, err := os.Open(source)
	if err != nil {
		return errorFileAVA.OpenDir(err, source)
	}
	defer sourceFile.Close()
	destFile, err := os.Create(dest)
	if err != nil {
		return errorFileAVA.CreateDir(err, dest)
	}
	defer destFile.Close()
	_, err = io.Copy(destFile, sourceFile)
	if err == nil {
		si, err := os.Stat(source)
		if err == nil {
			err = os.Chmod(dest, si.Mode())
		}
	}
	return errorFileAVA.CopyDir(err, fmt.Sprintf("Source: %s, Dest: %s", source, dest))
}

// FileCopyDir recursively copies a directory tree, attempting to preserve permissions.
// Source directory must exist, destination directory must *not* exist.
func (f *File) FileCopyDir(source string, dest string) *errorAVA.Error {
	// get properties of source dir
	fileInfo, err := os.Stat(source)
	if err != nil {
		return errorFileAVA.FileNotFount(err, source)
	}
	if !fileInfo.IsDir() {
		return errorFileAVA.FileNotIsDir(nil, source)
	}
	// ensure dest dir does not already exist
	_, err = os.Open(dest)
	if !os.IsNotExist(err) {
		return errorFileAVA.FileNotFount(err, dest)
	}
	// create dest dir
	err = os.MkdirAll(dest, fileInfo.Mode())
	if err != nil {
		return errorFileAVA.CreateDir(err, dest)
	}
	entries, errFile := ioutil.ReadDir(source)
	if errFile != nil {
		return errorFileAVA.ListDir(errFile, source)
	}
	for _, entry := range entries {
		sourcePath := filepath.Join(source, entry.Name())
		destinationPath := filepath.Join(dest, entry.Name())
		if entry.IsDir() {
			if errCopy := f.FileCopyDir(sourcePath, destinationPath); errCopy != nil {
				return errCopy
			}
		} else {
			// perform copy
			if errCopy := f.FileCopy(sourcePath, destinationPath); errCopy != nil {
				return errCopy
			}
		}
	}
	return nil
}

// Abs returns an absolute representation of path.
// If the path is not absolute it will be joined with the current working directory to turn it into an absolute path.
// The absolute path name for a given file is not guaranteed to be unique.
// Abs calls Clean on the result.
func (f *File) Abs(path string) (string, *errorAVA.Error) {
	result, err := filepath.Abs(path)
	if err != nil {
		return result, errorFileAVA.CurrentDirectoryPathWrong(err, path)
	}
	return result, nil
}
