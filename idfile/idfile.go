package idfile

import (
	"fmt"
	"log"
	"os"
)

const (
	MAX_BYTES_TO_READ = 512
)

type fileHandler func(string) string 

type fileSig struct {
	minSize      int
	prefix       string
	content      string
	contentStart int
	contentSize  int
	fileType     int
}
const (
    ELF = 1
    LIB = 2
    PNG = 3
    GIF = 4
    JPEG = 5
    JAVA_CLASS = 6
    ANDROID_EXE = 7
    TAR = 8
    ZIP = 9
    BZIP = 10
    MAC_EXE = 11
    OGG_DATA = 12
    WAV_AUDIO = 13
    FONT_FILE = 14
    PHP_SOURCE = 15
    DIGITAL_CERT = 16
    BMP = 17
    MS_EXE = 18
    GENERIC = 19
    GZIP = 20
)

var fileSignatures = []fileSig{
	{45, "\x7FELF", "", 0, 0, ELF},
	{16, "!<arch>\n", "", 0, 0, LIB },
	{32, "\x89PNG\x0d\x0a\x1a\x0a", "", 0, 0, PNG},
	{16, "GIF87a", "", 0, 0, GIF},
	{16, "GIF89a", "", 0, 0, GIF},
	{32, "\xff\xd8", "", 0, 0, JPEG},
	{32, "\xca\xfe\xba\xbe", "", 0, 0, JAVA_CLASS},
	{32, "dex\n", "", 0, 0, ANDROID_EXE"},
	{500, "", "ustar", 257, 5, TAR},
	{32, "PK\x03\x04", "", 0, 0, ZIP},
	{32, "BZh", "", 0, 0, BZIP},
	{16, "\x1f\x8b", "", 0, 0, GZIP},
	{32, "", "\xfa\xed\xfe", 1, 3, MAC_EXE},
	{36, "OggS\x00\x02", "", 0, 0, OGG_DATA},
	{32, "RIF", "WAVEfmt ", 8, 8, WAV_AUDIO},
	{16, "\x00\x01\x00\x00", "", 0, 0, FONT_FILE},
	{16, "ttcf\x00", "", 0, 0, FONT_FILE},
	{16, "<?php", "", 0, 0, PHP_SOURCE},
	{512, "-----BEGIN CERTIFICATE-----", "", 0, 0, DIGITAL_CERT},
	{50, "BM", "\x00\x00\x00\x00", 6, 4, BMP},
	{112, "MZ", "\x50\x45\x00\x00", 60, 4, MS_EXE},
	{0, "", "", 0, 0, GENERIC},
}

//{16, "BC\xc0\xde", "", 0, 0, "LLVM IR bitcode"},

func prefixMatch(s []byte, prefix string) bool {
	if len(prefix) == 0 {
		return true
	}
	return len(s) >= len(prefix) && byteCompare(s[:len(prefix)], prefix)
}

func contentMatch(s []byte, content string, contentStart int, contentSize int) bool {
	if len(content) == 0 {
		return true
	}
	return byteCompare(s[contentStart:contentStart+contentSize], content)
}

func byteCompare(a []byte, b string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range []byte(b) {
		if v != a[i] {
			return false
		}
	}
	return true
}

func FindFileType(fileName string) fileSig {
	file, _ := os.OpenFile(fileName, os.O_RDONLY, 0666)
	defer file.Close()

	var contentByte = make([]byte, MAX_BYTES_TO_READ)
	numByte, _ := file.Read(contentByte)
	contentByte = contentByte[:numByte]

	for _, fs := range fileSignatures {
                //fmt.Println(fs)
		if numByte > fs.minSize &&
			prefixMatch(contentByte, fs.prefix) &&
			contentMatch(contentByte, fs.content, fs.contentStart, fs.contentSize) {
			return fs
		}
	}
	return nil
}

type TypeViolation struct {
    ViolationFound bool
    BinaryContentInTextFile bool
    ProgramContentInImageFile bool
    Detail string
}


type LocationViolation struct {
    ViolationFound bool `json:"violations,omitempty"`
    FileInWrongLocation bool `json:"wrong_location,omitempty"`
    FileHidden  bool    `json:"hidden,omitempty"`
    Detail string       json:"detail"`
}

type AttribViolation struct {
    ViolationFound bool
    PermissionIncorrect bool
    Detail string
}

type ContentViolation struct {
    ViolationFound bool
    ContentMalicious bool
    SizeViolation bool
    ModifiedSystemFile bool
    Detail string
}

type FileAnalysis struct {
    FileName string
    ViolationCount int
    ViolationsInfo map[string] string
}


func AnalyzeFile(fileName string) string {

    fs := FindFileType(fileName)

    switch fs.fileType {
        case ELF:
        case LIB:
        case PNG:
        case GIF:
        case JPEG:
        case JAVA_CLASS:
        case ANDROID_EXE:
        case TAR:
        case ZIP:
        case BZIP:
        case MAC_EXE:
        case OGG_DATA:
        case WAV_AUDIO:
        case FONT_FILE:
        case PHP_SOURCE:
        case DIGITAL_CERT:
        case BMP:
        case MS_EXE:
        case GENERIC:
    }

    if fs == nil {
        return json.Marshal("")
    }

    fileAnalysis := fs.flHdr(fs.fileType)

    fileAnalysisJSON, err := json.Marshal(fileAnalysis)

    if  err != nil {
        log.Errorf("Error on Marshaling fileAnalysis for %v \n", fileName)
    }

    return string(fileAnalysisJSON)

}

type SourceViolation struct {
    NewLinesPercentage  int
    CrypticCodePresent  bool
    MaliciousContentInsert bool
    Detail string
}

func sourceFileAnalyzer(fs, fileName string) (bool, string)  {
    var violationFound  bool

    size, lines, err := getSizeLineCount(fileName)

    if err != nil {
        log.Errorf("Unable to ID file %v \n", fileName)
        return false
    }

    expectedLineRangeStart :=  (size/80) * 80/100
    expectedLineRangeEnd    :=  (lines - expectedLineRangeStart )+ lines 

    if lines < expectedLineRangeStart || line >
        expectedLineRangeEnd {
            sv.NewLinesPercentage = true
            sv.Detail = fmt.Sprintf(
                    "newline violation lines [%v] vs size [%v] \n", 
                    lines, size)
        }




}
