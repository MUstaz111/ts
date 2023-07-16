package handlers

import (
	"regexp"
	"strings"
)

func FindMaxSubstring(s string) string {
	var maxSubstring string
	var currentSubstring string

	for i := 0; i < len(s); i++ {
		if index := strings.IndexByte(currentSubstring, s[i]); index != -1 {
			if len(currentSubstring) > len(maxSubstring) {
				maxSubstring = currentSubstring
			}
			currentSubstring = currentSubstring[index+1:]
		}
		currentSubstring += string(s[i])
	}

	if len(currentSubstring) > len(maxSubstring) {
		maxSubstring = currentSubstring
	}

	return maxSubstring
}

func FindEmails(text string) []string {
	// Регулярное выражение для поиска email
	regex := regexp.MustCompile(`\b[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,}\b`)

	emails := regex.FindAllString(text, -1)

	return emails
}

func FindIINs(text string) []string {
	// Регулярное выражение для поиска ИИН
	regex := regexp.MustCompile(`\b\d{12}\b`)

	iins := regex.FindAllString(text, -1)

	return iins
}

func FindIdentifiers(str string) []string {
	// Регулярное выражение для поиска идентификаторов
	regex := regexp.MustCompile(`\b[a-zA-Z_][a-zA-Z0-9_]*\b`)

	// Приведение строки к нижнему регистру для нечувствительного к регистру поиска
	lowercaseStr := strings.ToLower(str)

	// Поиск всех совпадений с регулярным выражением
	matches := regex.FindAllString(lowercaseStr, -1)

	// Удаление повторяющихся идентификаторов
	uniqueMatches := make(map[string]bool)
	for _, match := range matches {
		uniqueMatches[match] = true
	}

	// Формирование списка найденных идентификаторов
	identifiers := make([]string, 0, len(uniqueMatches))
	for identifier := range uniqueMatches {
		identifiers = append(identifiers, identifier)
	}

	return identifiers
}
