package parsinglogfiles

import (
	"fmt"
	"regexp"
)

func IsValidLine(text string) bool {
	ok, err := regexp.MatchString(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`, text)
	if err != nil {
		return false
	}
	return ok
}
func SplitLogLine(text string) []string {
	return regexp.MustCompile(`<>|<\*>|<~~~>|<=>|<-\*~\*->`).Split(text, -1)
}
func CountQuotedPasswords(lines []string) int {
	r := regexp.MustCompile(`(?i)".*password.*"`)
	var count int
	for i := range lines {
		matches := r.FindAllString(lines[i], -1)
		count += len(matches)
	}
	return count
}
func RemoveEndOfLineText(text string) string {
	return regexp.MustCompile(`end-of-line\d+`).ReplaceAllString(text, "")
}
func TagWithUserName(lines []string) []string {
	var result []string
	for _, log := range lines {
		match := regexp.MustCompile(`\bUser\s+([A-Za-z0-9_]+)\b`).FindStringSubmatch(log)
		if match != nil {
			username := match[1]
			result = append(result, fmt.Sprintf("[USR] %s %s", username, log))
		} else {
			result = append(result, log)
		}
	}
	return result
}
