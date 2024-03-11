package parsinglogfiles

import "regexp"

func IsValidLine(text string) bool {
	re := regexp.MustCompile(`^\[(?:TRC|DBG|INF|WRN|ERR|FTL)]`)
	return re.MatchString(text)
}

func SplitLogLine(text string) []string {
	re := regexp.MustCompile(`<[~*=-]*>`)
	return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) (count int) {
	re := regexp.MustCompile(`(?i)".*password.*"`)
	for _, line := range lines {
		if re.MatchString(line) {
			count++
		}
	}
	return
}

func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`end-of-line[[:digit:]]*`)
	return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	re := regexp.MustCompile(`User +([^ ]+) `)
	var output []string
	for _, line := range lines {
		res := re.FindStringSubmatch(line)
		if res != nil {
			line = "[USR] " + res[1] + " " + line
		}
		output = append(output, line)
	}
	return output
}
