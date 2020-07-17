package parser

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
	"time"
)

const parseLogRegex = `(?m)^(?P<ip>\S*):?(?P<port>\S*).*\[(?P<date>.*)\]\s"(?P<method>\S*)\s(?P<resource>\S*)\s([^"]*)"\s(?P<status>\S*)\s(?P<bytes>\S*)\s"([^"]*)"\s"([^"]*)"`

type (
	IP             string
	Port           uint
	Resource       string
	HTTPMethod     string
	HTTPProtocol   string
	HTTPStatusCode uint16
)

type Log struct {
	IP         IP
	Port       Port
	Date       time.Time
	Resource   Resource
	Method     HTTPMethod
	Protocol   HTTPProtocol
	StatusCode HTTPStatusCode
}

func NewLog(
	IP IP,
	port Port,
	date time.Time,
	resource Resource,
	method HTTPMethod,
	protocol HTTPProtocol,
	statusCode HTTPStatusCode,
) Log {
	return Log{
		IP:         IP,
		Port:       port,
		Date:       date,
		Resource:   resource,
		Method:     method,
		Protocol:   protocol,
		StatusCode: statusCode,
	}
}

type Logs []Log

func Parse(reader io.Reader) (Logs, error) {
	var (
		scanner = bufio.NewScanner(reader)
		logs    = Logs{}
	)

	expression, err := regexp.Compile(parseLogRegex)
	if err != nil {
		return []Log{}, err
	}

	for scanner.Scan() {
		if !expression.Match(scanner.Bytes()) {
			fmt.Println(scanner.Text())
			continue
		}

		line, err := newLine(scanner.Text(), expression)
		if err != nil {
			return []Log{}, err
		}

		port, err := line.port()
		if err != nil {
			return []Log{}, err
		}

		date, err := line.date()
		if err != nil {
			return []Log{}, err
		}

		statusCode, err := line.statusCode()
		if err != nil {
			return []Log{}, err
		}

		var log = NewLog(
			line.ip(),
			port,
			date,
			line.resource(),
			line.method(),
			line.protocol(),
			statusCode,
		)

		logs = append(logs, log)
	}

	return logs, nil
}
