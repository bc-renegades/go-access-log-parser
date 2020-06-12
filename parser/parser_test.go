package parser

import (
	"io"
	"reflect"
	"strings"
	"testing"
	"time"
)

func TestParse(t *testing.T) {

	dateTest, _ := time.Parse("02/Jan/2006:15:04:05 -0700", "12/Jun/2020:04:23:50 -0300")

	type args struct {
		file io.Reader
	}

	tests := []struct {
		name    string
		args    args
		want    Logs
		wantErr bool
	}{
		{
			name: "Test",
			args: args{
				file: strings.NewReader("Hello World\n World Hello\n asdada"),
			},
			want:    Logs{},
			wantErr: true,
		},
		{
			name: "Test line apache",
			args: args{
				file: strings.NewReader(`10.196.138.122:54274 - "-|-" [12/Jun/2020:04:23:50 -0300] "GET /test.com/transactions/99501553?_url=/transactions/99501553 HTTP/1.1" 401 72 "-" "-" Location:["-"] 0 632330 -- TLSv1.2 ECDHE-RSA-AES256-GCM-SHA384`),
			},
			want: Logs{
				{
					IP:         "10.196.138.122",
					Port:       54274,
					Date:       dateTest,
					Resource:   "/test.com/transactions/99501553?_url=/transactions/99501553",
					Method:     "GET",
					Protocol:   "HTTP/1.1",
					StatusCode: 401,
				},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Parse(tt.args.file)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parse() got = %v, want %v", got, tt.want)
			}
		})
	}
}
