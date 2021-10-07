package repository

import (
	"fmt"
	"strconv"
	"strings"
)

func getMultiBinds(lenargs int, startFrom int) (string, int) {
	if lenargs < 1 {
		return "", startFrom
	}
	i := 1
	bindVars := ""
	if startFrom > 0 {
		i = startFrom
		for i <= lenargs {
			bindVars += "$" + strconv.Itoa(i) + ","
			i++
		}
	}

	bindVars = strings.TrimSuffix(bindVars, ",")
	return bindVars, i
}

// bindVarsToWhere helps to add to where statement necessary count of WHERE statements
// Example:
// AND (m.size = $1 AND m.size = $2)
func getWhereMultiBinds(link string, field string, lenargs int, startFrom int) (string, int) {
	if link == "" || field == "" || lenargs < 1 {
		return "", startFrom
	}
	fmt.Println("start binding")
	bindVars := ""
	i := 1
	if startFrom > 0 {
		i = startFrom
	}

	if lenargs > 1 {
		bindVars += "("

		for i <= lenargs {
			bindVars += " " + field + " = " + "$" + strconv.Itoa(i)
			i++
			bindVars += " " + link
		}
		bindVars = strings.TrimSuffix(bindVars, link)
		bindVars += ")"
	} else {
		bindVars += field + " = " + "$" + strconv.Itoa(i)
		i++
	}
	bindVars = " AND " + bindVars
	return bindVars, i
}

func bindVarsToInsert(stmt string, lenargs int, len int) string {
	fmt.Println("start binding")
	bindVars := ""
	i := 1
	for i <= len {
		bindVars += "("
		for j := 0; j < lenargs; j++ {
			bindVars += "$" + strconv.Itoa(i) + ","
			i++
		}
		bindVars = strings.TrimSuffix(bindVars, ",")
		bindVars += "),"
	}
	stmt += bindVars
	stmt = fmt.Sprintf(stmt+bindVars, modelsTable)
	return strings.TrimSuffix(stmt, ",")
}
