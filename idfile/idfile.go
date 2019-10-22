package idfile

import (
	"fmt"
	"log"
	"os"
)

const (
	MAX_BYTES_TO_READ = 512
)

type fileSig struct {
	minSize      int
	prefix       string
	content      string
	contentStart int
	contentSize  int
	fileType     string
}

var fileSignatures = []fileSig{
	{45, "\x7FELF", "", 0, 0, "Elf"},
	{16, "!<arch>\n", "", 0, 0, "ar"},
	{32, "\x89PNG\x0d\x0a\x1a\x0a", "", 0, 0, "PNG"},
	{16, "GIF87a", "", 0, 0, "GIF"},
	{16, "GIF89a", "", 0, 0, "GIF"},
	{32, "\xff\xd8", "", 0, 0, "JPEG"},
	{32, "\xca\xfe\xba\xbe", "", 0, 0, "Java-class"},
	{32, "dex\n", "", 0, 0, "Android-dex"},
	{500, "", "ustar", 257, 5, "tar"},
	{32, "PK\x03\x04", "", 0, 0, "Zip"},
	{32, "BZh", "", 0, 0, "bzip2"},
	{16, "\x1f\x8b", "", 0, 0, "gzip"},
	{32, "", "\xfa\xed\xfe", 1, 3, "Mach-O"},
	{36, "OggS\x00\x02", "", 0, 0, "Ogg-data"},
	{32, "RIF", "WAVEfmt ", 8, 8, "WAV"},
	{16, "\x00\x01\x00\x00", "", 0, 0, "font"},
	{16, "ttcf\x00", "", 0, 0, "font"},
	{16, "<?php", "", 0, 0, "PHP"},
	{512, "-----BEGIN CERTIFICATE-----", "", 0, 0, "PEM"},
	{50, "BM", "\x00\x00\x00\x00", 6, 4, "BMP"},
	{112, "MZ", "\x50\x45\x00\x00", 60, 4, "MS-exe"},
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

func FindFileType(fileName string) string {
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
			return fs.fileType
		}
	}
	return "UNKNOWN"
}
