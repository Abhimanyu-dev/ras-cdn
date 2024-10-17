package utils

import (
	"fmt"
	"strings"
)

func GenerateName(uuid, filename, rid string) string {
	ridstr := fmt.Sprintf("%04v", rid)
	parts := strings.Split(filename, ".")
	parts[0] = "SPO-IITK_" + ridstr + "_" + parts[0] + "_" + uuid

	return strings.Join(parts, ".")
}

func GetRollNo(filename string) string {
	parts := strings.Split(filename, "_")
	newFileName := parts[2] + ".pdf"
	return newFileName
}