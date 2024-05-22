package logger

import (
	"fmt"
	"os"
	"runtime/debug"
	"strings"
	"sync"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
	"github.com/spf13/viper"
)

// TODOs
// - timestamp with gray color like in rust
// - create spans like in rust with context
// - wrap fields in curly braces

const (
	colorBlack = iota + 30
	colorRed
	colorGreen
	colorYellow
	colorBlue
	colorMagenta
	colorCyan
	colorWhite

	colorBold   = 1
	colorItalic = 3
)

var once sync.Once

func GetLogLevel() zerolog.Level {
	logLevel := viper.GetString("log")
	switch {
	case logLevel == "debug":
		return zerolog.DebugLevel
	case logLevel == "trace":
		return zerolog.TraceLevel
	}

	return zerolog.InfoLevel
}

func Get() {
	// ensure it runs once even it Get() is called many times
	once.Do(func() {
		zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
		zerolog.TimeFieldFormat = time.RFC3339Nano

		// TODO procution and developement logger based on ENV
		output := zerolog.ConsoleWriter{
			Out:        os.Stdout,
			TimeFormat: time.RFC3339,
			FieldsExclude: []string{
				"user_agent",
				"git_revision",
				"go_version",
			},
		}

		output.FormatLevel = func(i interface{}) string {
			if s, ok := i.(string); ok {
				var color int
				switch i {
				case "info":
					color = colorGreen
				case "debug":
					color = colorBlue
				case "trace":
					color = colorMagenta
				case "error":
					color = colorRed
				}

				return fmt.Sprintf("\x1b[%d;%dm%-6s\x1b[0m", color, colorBold, strings.ToUpper(s))
			}
			return "FMT_ERROR"
		}

		output.FormatFieldName = func(i interface{}) string {
			return fmt.Sprintf("\x1b[%d;%dm%s\x1b[0m=", colorYellow, colorItalic, i)
		}

		output.FormatTimestamp = func(i interface{}) string {
			if s, ok := i.(string); ok {
				parsedTime, _ := time.Parse(time.RFC3339Nano, s)
				return fmt.Sprintf("\x1b[37m%s\x1b[0m", parsedTime.UTC().Format("2006-01-02T15:04:05.999999Z"))
			}
			return "FMT_ERROR"
		}

		var gitRevision string
		buildInfo, ok := debug.ReadBuildInfo()
		if ok {
			for _, v := range buildInfo.Settings {
				if v.Key == "vcs.revision" {
					gitRevision = v.Value
					break
				}
			}
		}

		log.Logger = zerolog.New(output).
			Level(GetLogLevel()).
			With().
			Caller().
			Timestamp().
			Str("git_revision", gitRevision).
			Str("go_version", buildInfo.GoVersion).
			Logger()
	})
}
