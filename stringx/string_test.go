package stringx

import (
	"reflect"
	"regexp"
	"testing"
)

func TestString_Render(t *testing.T) {
	type fields struct {
		s       string
		pattern *regexp.Regexp
	}
	type args struct {
		data interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *String
		wantErr bool
	}{
		// Add test cases.
		{
			name: "test1",
			fields: fields{
				s:       "hello {{.name}}",
				pattern: nil,
			},
			args: args{
				data: map[string]interface{}{
					"name": "world",
				},
			},
			want:    New("hello world"),
			wantErr: false,
		},
		{
			name: "test2",
			fields: fields{
				s:       "hello {{.}}",
				pattern: nil,
			},
			args: args{
				data: map[string]interface{}{},
			},
			want:    New("hello map[]"),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &String{
				s:       tt.fields.s,
				pattern: tt.fields.pattern,
			}
			got, err := s.Render(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("String.Render() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("String.Render() = %v, want %v", got, tt.want)
			}
		})
	}
}
