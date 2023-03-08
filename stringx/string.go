package stringx

import (
	"bytes"
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
	"text/template"

	"github.com/iancoleman/strcase"
)

type String struct {
	s       string
	pattern *regexp.Regexp
}

func New(s string) *String {
	return &String{s: s}
}

func (s *String) String() string {
	return s.s
}

func (s *String) Len() int {
	return len(s.s)
}

func (s *String) IsEmpty() bool {
	return s.Len() == 0
}

func (s *String) IsBlank() bool {
	return s.Trim().IsEmpty()
}

func (s *String) Trim() *String {
	s.s = strings.TrimSpace(s.s)
	return s
}

func (s *String) TrimLeft() *String {
	s.s = strings.TrimLeft(s.s, " ")
	return s
}

func (s *String) TrimRight() *String {
	s.s = strings.TrimRight(s.s, " ")
	return s
}

func (s *String) TrimPrefix(prefix string) *String {
	s.s = strings.TrimPrefix(s.s, prefix)
	return s
}

func (s *String) TrimSuffix(suffix string) *String {
	s.s = strings.TrimSuffix(s.s, suffix)
	return s
}

func (s *String) TrimLeftFunc(f func(rune) bool) *String {
	s.s = strings.TrimLeftFunc(s.s, f)
	return s
}

func (s *String) TrimRightFunc(f func(rune) bool) *String {
	s.s = strings.TrimRightFunc(s.s, f)
	return s
}

func (s *String) TrimFunc(f func(rune) bool) *String {
	s.s = strings.TrimFunc(s.s, f)
	return s
}

// ReplaceAll replaces all occurrences of old with new.
func (s *String) ReplaceAll(old, new string) *String {
	s.s = strings.ReplaceAll(s.s, old, new)
	return s
}

// Replace replaces the first n occurrences of old with new.
func (s *String) Replace(old, new string, n int) *String {
	s.s = strings.Replace(s.s, old, new, n)
	return s
}

// ReplaceAllFunc replaces all occurrences of old with new.
func (s *String) ReplaceAllFunc(old string, replacer func(string) string) *String {
	s.s = strings.ReplaceAll(s.s, old, replacer(old))
	return s
}

// ReplaceFunc replaces the first n occurrences of old with new.
func (s *String) ReplaceFunc(old string, replacer func(string) string, n int) *String {
	s.s = strings.Replace(s.s, old, replacer(old), n)
	return s
}

func (s *String) ToLower() *String {
	s.s = strings.ToLower(s.s)
	return s
}

func (s *String) ToUpper() *String {
	s.s = strings.ToUpper(s.s)
	return s
}

// | ToSnake(s)                      | any_kind_of_string |
// | ToScreamingSnake(s)             | ANY_KIND_OF_STRING |
// | ToKebab(s)                      | any-kind-of-string |
// | ToScreamingKebab(s)             | ANY-KIND-OF-STRING |
// | ToDelimited(s, '.')             | any.kind.of.string |
// | ToScreamingDelimited(s, '.')    | ANY.KIND.OF.STRING |
// | ToCamel(s)                      | AnyKindOfString    |
// | ToLowerCamel(s)                 | anyKindOfString    |

func (s *String) ToTitle() *String {
	s.s = strings.ToTitle(s.s)
	return s
}

func (s *String) ToLowerCamelCase() *String {
	s.s = strcase.ToLowerCamel(s.s)
	return s
}

func (s *String) ToSnakeCase() *String {
	s.s = strcase.ToSnake(s.s)
	return s
}

func (s *String) ToScreamingSnakeCase() *String {
	s.s = strcase.ToScreamingSnake(s.s)
	return s
}

func (s *String) ToKebabCase() *String {
	s.s = strcase.ToKebab(s.s)
	return s
}

func (s *String) ToScreamingKebabCase() *String {
	s.s = strcase.ToScreamingKebab(s.s)
	return s
}

func (s *String) ToDelimitedCase(delimiter uint8) *String {
	s.s = strcase.ToDelimited(s.s, delimiter)
	return s
}

func (s *String) ToPascalCase() *String {
	s.s = strcase.ToCamel(s.s)
	return s
}

// WithRegex
func (s *String) WithRegex(regex *regexp.Regexp) *String {
	s.pattern = regex
	return s
}

func (s *String) WithRegexString(regex string) *String {
	s.pattern = regexp.MustCompile(regex)
	return s
}

func (s *String) IsMatch() bool {
	return s.pattern.MatchString(s.s)
}

func (s *String) FindAllString() []string {
	return s.pattern.FindAllString(s.s, -1)
}

func (s *String) FindAllStringIndex() [][]int {
	return s.pattern.FindAllStringIndex(s.s, -1)
}

func (s *String) FindAllStringSubmatch() [][]string {
	return s.pattern.FindAllStringSubmatch(s.s, -1)
}

// FindAllStringSubmatchIndex returns a slice of all successive matches of the Regexp and the index pairs for the matches.
func (s *String) FindAllStringSubmatchIndex() [][]int {
	return s.pattern.FindAllStringSubmatchIndex(s.s, -1)
}

// ReplaceAllStringFunc replaces all matches of the Regexp with the return value of function repl.
func (s *String) ReplaceAllStringFunc(repl func(string) string) *String {
	s.s = s.pattern.ReplaceAllStringFunc(s.s, repl)
	return s
}

// ReplaceAllString replaces all matches of the Regexp with the replacement string repl.
func (s *String) ReplaceAllString(repl string) *String {
	s.s = s.pattern.ReplaceAllString(s.s, repl)
	return s
}

// ReplaceAllLiteralString replaces all matches of the Regexp with the replacement string repl.
func (s *String) ReplaceAllLiteralString(repl string) *String {
	s.s = s.pattern.ReplaceAllLiteralString(s.s, repl)
	return s
}

// Split splits the string s around matches of the Regexp.
func (s *String) Split(sep string) []string {
	return strings.Split(s.s, sep)
}

// Equal compares two strings.
func (s *String) Equal(s2 String) bool {
	return s.s == s2.s
}

// EqualString compares two strings.
func (s *String) EqualString(s2 string) bool {
	return s.s == s2
}

// Append appends a string.
func (s *String) Append(s2 String) *String {
	s.s += s2.s
	return s
}

// AppendString appends a string.
func (s *String) AppendString(s2 string) *String {
	s.s += s2
	return s
}

// Prepend prepends a string.
func (s *String) Prepend(s2 String) *String {
	s.s = s2.s + s.s
	return s
}

// PrependString prepends a string.
func (s *String) PrependString(s2 string) *String {
	s.s = s2 + s.s
	return s
}

// Insert inserts a string at the specified index.
func (s *String) Insert(index int, s2 String) *String {
	s.s = s.s[:index] + s2.s + s.s[index:]
	return s
}

// InsertString inserts a string at the specified index.
func (s *String) InsertString(index int, s2 string) *String {
	s.s = s.s[:index] + s2 + s.s[index:]
	return s
}

// Remove removes the specified number of characters from the specified index.
func (s *String) Remove(index int, length int) *String {
	s.s = s.s[:index] + s.s[index+length:]
	return s
}

// RemoveAll removes all characters.
func (s *String) RemoveAll() *String {
	s.s = ""
	return s
}

// Bytes returns a slice of bytes.
func (s *String) Bytes() []byte {
	return []byte(s.s)
}

// JsonUnmarshal unmarshals the JSON-encoded data and stores the result in the value pointed to by v.
func (s *String) JsonUnmarshal(v interface{}) error {
	return json.Unmarshal(s.Bytes(), v)
}

// FindIndex returns the index of the first instance of substr in s, or -1 if substr is not present in s.
func (s *String) FindIndex(sub string) int {
	return strings.Index(s.s, sub)
}

// FindLastIndex returns the index of the last instance of substr in s, or -1 if substr is not present in s.
func (s *String) FindLastIndex(sub string) int {
	return strings.LastIndex(s.s, sub)
}

// Padding adds characters to a string until the specified length is reached.
func (s *String) Padding(length int, padStr string) *String {
	if s.Len() >= length {
		return s
	}

	padCount := length - s.Len()
	for i := 0; i < padCount; i++ {
		s.s += padStr
	}
	return s
}

// Format formats according to a format specifier and returns the string as a value that satisfies valuer.
func (s *String) Format(args ...interface{}) *String {
	s.s = fmt.Sprintf(s.s, args...)
	return s
}

// Render renders a template.
func (s *String) Render(data interface{}) (*String, error) {
	t := template.Must(template.New("t").Parse(s.s))
	var buf bytes.Buffer
	if err := t.Execute(&buf, data); err != nil {
		return nil, err
	}

	s.s = buf.String()
	return s, nil
}
