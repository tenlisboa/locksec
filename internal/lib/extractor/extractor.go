package extractor

import (
	"errors"
	"regexp"
)

type fileType int

const (
	Json fileType = iota
	Yml
)

var rxjson = `"([^"]+)":\s*"([^"]+)"(?:,|$)`
var rxyaml = `(?i)([A-Za-z_-]+)[:| ]\s*"?([^"\n]+)"?\s*$`

var errMatch = errors.New("line does not match the pattern")

type extractor struct {
	ftype fileType
}

func New(ftype fileType) *extractor {
	return &extractor{
		ftype: ftype,
	}
}

func (e *extractor) ExtractLine(line string) (string, any, error) {
	re := regexp.MustCompile(e.getRex())

	match := re.FindStringSubmatch(line)

	if len(match) < 3 {
		return "", nil, errMatch
	}

	key, value := match[1], match[2]

	return key, value, nil
}

func (e *extractor) getRex() string {
	switch e.ftype {
	case Json:
		return rxjson
	case Yml:
		return rxyaml
	default:
		return ""
	}
}

func ToFType(index int) fileType {
	return fileType(index)
}
