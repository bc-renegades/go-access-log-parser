package parser

import (
	"regexp"
	"strconv"
	"time"

	"github.com/pkg/errors"
)

const (
	ipKey = iota + 1
	portKey
	dateKey
	methodKey
	resourceKey
	protocolKey
	statusCodeKey
)

type line struct {
	value []string
}

func newLine(value string, exp *regexp.Regexp) (line, error) {
	var results = exp.FindAllStringSubmatch(value, -1)
	if len(results) <= 0 {
		return line{}, errors.New("nenhum resultado encontrado")
	}

	return line{value: results[0]}, nil
}

func (l line) ip() IP {
	return IP(l.value[ipKey])
}

func (l line) port() (Port, error) {
	if l.value[portKey] == "" {
		return 0, nil
	}
	port, err := strconv.ParseUint(l.value[portKey], 10, 32)
	if err != nil {
		return Port(0), errors.Wrap(err, "error ao converter porta")
	}

	return Port(port), nil
}

func (l line) date() (time.Time, error) {
	date, err := time.Parse("02/Jan/2006:15:04:05 -0700", l.value[dateKey])
	if err != nil {
		return time.Time{}, errors.Wrap(err, "error ao converter data")
	}

	return date, nil
}

func (l line) method() HTTPMethod {
	return HTTPMethod(l.value[methodKey])
}

func (l line) resource() Resource {
	return Resource(l.value[resourceKey])
}

func (l line) protocol() HTTPProtocol {
	return HTTPProtocol(l.value[protocolKey])
}

func (l line) statusCode() (HTTPStatusCode, error) {
	statusCode, err := strconv.ParseUint(l.value[statusCodeKey], 10, 16)
	if err != nil {
		return HTTPStatusCode(0), errors.Wrap(err, "error ao converter status code")
	}

	return HTTPStatusCode(statusCode), nil
}
