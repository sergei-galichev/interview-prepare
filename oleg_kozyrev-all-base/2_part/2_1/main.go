package main

import (
	"log/slog"
	"strconv"
	"strings"
	"time"
)

func main() {
	builder := strings.Builder{}

	start := time.Now()

	for i := 0; i < 100_000; i++ {
		//builder.WriteString(fmt.Sprintf("%d", i))
		builder.WriteString(strconv.Itoa(i))
	}

	slog.Info("Time execution", "duration", time.Since(start))
}
