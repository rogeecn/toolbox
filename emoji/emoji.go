package emoji

import "regexp"

func Filter(input string) string {
	re := regexp.MustCompile(`[\x{1F600}-\x{1F64F}\x{1F300}-\x{1F5FF}\x{1F680}-\x{1F6FF}\x{2600}-\x{26FF}]`)
	return re.ReplaceAllString(input, "")
}
