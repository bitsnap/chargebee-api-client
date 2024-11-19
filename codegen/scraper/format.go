package scraper

import (
	"fmt"
	"slices"
	"strings"

	. "github.com/bitsnap/go-util"
)

func formatTag(snakeCaseName string, required bool) string {
	r := ""
	if required {
		r = ` validate:"required"`
	}

	return fmt.Sprintf("`json:\"%v\"%v`", snakeCaseName, r)
}

var boolFields = []string{
	"metered",
}

func formatEnumName(snakeCaseName string, attributeType string) string {
	if strings.HasPrefix(snakeCaseName, "is_") || strings.HasPrefix(snakeCaseName, "enabled_") {
		return "bool"
	}

	if slices.Contains(boolFields, snakeCaseName) {
		return "bool"
	}

	if attributeType == "enumerated string" {
		return "enums." + SnakeToCamel(snakeCaseName, true) + "Enum"
	}

	return attributeType
}

func formatSimpleType(attributeType string) (ok bool, res string) {
	if attributeType == "string" {
		return true, "string"
	}

	if attributeType == "long" {
		return true, "int64"
	}

	if attributeType == "integer" {
		return true, "int"
	}

	if attributeType == "boolean" {
		return true, "bool"
	}

	if attributeType == "float" {
		return true, "float32"
	}

	if attributeType == "double" {
		return true, "float64"
	}

	if attributeType == "bigdecimal" {
		return true, "*decimal.Decimal"
	}

	if attributeType == "jsonarray" {
		return true, "[]string"
	}

	if attributeType == "jsonobject" {
		return true, "map[string]interface{}"
	}

	return false, ""
}

func formatType(attributeType string) string {
	if attributeType == "timestamp(UTC) in seconds" {
		return "uint64"
	}

	if attributeType == "in cents" {
		return "uint64"
	}

	if ok, t := formatSimpleType(attributeType); ok {
		return t
	}

	if strings.HasPrefix(attributeType, "list of ") {
		subTypeName := strings.Replace(attributeType, "list of ", "", 1)
		if ok, t := formatSimpleType(subTypeName); ok == true {
			return "[]" + t
		}

		if ValidModel(subTypeName) {
			return "[]" + SnakeToCamel(subTypeName, true)
		}

		return "[]models." + SnakeToCamel(subTypeName, true)
	}

	return attributeType
}
