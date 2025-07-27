package repository

import (
	"fmt"
	"strings"
)

func generatePlaceholders(chunkSize, totalNumbers int) string {
	if chunkSize <= 0 || totalNumbers <= 0 {
		return ""
	}

	var builder strings.Builder

	builder.WriteString("(")
	for i := 1; i <= totalNumbers; i++ {
		if i > 1 {
			if (i-1)%chunkSize == 0 {
				builder.WriteString("),(")
			} else {
				builder.WriteString(", ")
			}
		}
		builder.WriteString(fmt.Sprintf("$%d", i))
	}
	builder.WriteString(")")

	return builder.String()
}
