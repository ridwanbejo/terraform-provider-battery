// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package string

import (
	"errors"
	"strings"
	"regexp"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func Title(str string) string {
	tempStr := strings.Replace(strings.Replace(str, "_", " ", -1), "-", " ", -1)
	caser := cases.Title(language.English)
	titleStr := caser.String(tempStr)

	return titleStr
}

func PascalCase(str string) string {
	titleStr := Title(str)
	pascalStr := strings.Replace(titleStr, " ", "", -1)

	return pascalStr
}

func CamelCase(str string) (string, error) {
	if len(str) == 0 {
        return "", errors.New("string length minimum is 1")
    } else {
    	tempStr := PascalCase(str)
		lowercasedFirstChar := strings.ToLower(string(tempStr[0]))
		camelCaseStr := lowercasedFirstChar + tempStr[1:]

		return camelCaseStr, nil
    }
}

func SnakeCase(str string) string {
	tempStr := strings.ToLower(str)
	snakeCased := strings.Replace(strings.Replace(tempStr, "-", " ", -1), " ", "_", -1)

	return snakeCased
}

func KebabCase(str string) string {
	tempStr := strings.ToLower(str)
	kebabCased := strings.Replace(strings.Replace(tempStr, "_", " ", -1), " ", "-", -1)

	return kebabCased
}

func Capitalize(str string) (string, error) {
	if len(str) == 0 {
        return "", errors.New("string length minimum is 1")
    } else {
		uppercasedFirstChar := strings.ToUpper(string(str[0]))
		capitalizedStr := uppercasedFirstChar + str[1:]

		return capitalizedStr, nil
	}
}

func WordsCount(str string) int {
	tempWords := strings.Split(str, " ")
	wordsCount := len(tempWords)

	return wordsCount
}

func AddPrefix(str, prefix string) string {
	tempStr := prefix + str

	return tempStr
}

func AddSuffix(str, suffix string) string {
	tempStr := str + suffix

	return tempStr
}

func IsAlpha(str string) bool {
	checkPattern := regexp.MustCompile(`^[A-Za-z]+$`).MatchString

	return checkPattern(str)
}

func IsAlnum(str string) bool {
	checkPattern := regexp.MustCompile(`^[A-Za-z1-9]+$`).MatchString

	return checkPattern(str)
}

func IsNumeric(str string) bool {
	checkPattern := regexp.MustCompile(`^[0-9]+$`).MatchString

	return checkPattern(str)
}

func InsertAt(str string, substr string, index int) (string, error) {
	if (index < 0) || (index > len(str)) {
        return "", errors.New("index for InsertAt must be between 0 and max index of string")
    } else {
    	newStr := str[:index] + substr + str[index:]

		return newStr, nil
    }
}
