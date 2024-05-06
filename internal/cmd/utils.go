package cmd

import (
	"regexp"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// - styling

func StyleHeaders(tpl string) string {
	cobra.AddTemplateFunc("HeadingStyle", color.New(color.BgWhite, color.FgBlack).SprintFunc())
	cobra.AddTemplateFunc("UsageStyle", color.New(color.BgGreen, color.FgBlack).SprintFunc())

	tpl = strings.NewReplacer(
		`Usage:`, `{{UsageStyle " Usage "}}`,
		`Aliases:`, `{{HeadingStyle " Aliases "}}`,
		`Available Commands:`, `{{HeadingStyle " Commands "}}`,
		`Global Flags:`, `{{HeadingStyle " Global Flags "}}`,
	).Replace(tpl)

	re := regexp.MustCompile(`(?m)^Flags:\s*$`)
	return re.ReplaceAllLiteralString(tpl, `{{HeadingStyle " Flags "}}`)
}

func StyleCmds(tpl string) string {
	cmds := color.New(color.FgCyan)
	cobra.AddTemplateFunc("NameStyle", cmds.SprintFunc())
	cobra.AddTemplateFunc("sum", func(a, b int) int {
		return a + b
	})

	re := regexp.MustCompile(`(?i){{\s*rpad\s+.Name\s+.NamePadding\s*}}`)
	tpl = re.ReplaceAllLiteralString(tpl, "{{rpad (NameStyle .Name) (sum .NamePadding 10)}}")

	// - exec
	cen := color.New(color.FgGreen)
	cobra.AddTemplateFunc("CommandPathStyle", cen.SprintFunc())
	cobra.AddTemplateFunc("UseLineStyle", func(s string) string {
		spl := strings.SplitN(s, " ", 2)
		spl[0] = cen.Sprint(spl[0])
		return strings.Join(spl, " ")
	})

	re = regexp.MustCompile(`(?i){{\s*.CommandPath\s*}}`)
	tpl = re.ReplaceAllLiteralString(tpl, "{{CommandPathStyle .CommandPath}}")

	re = regexp.MustCompile(`(?i){{\s*.UseLine\s*}}`)
	return re.ReplaceAllLiteralString(tpl, "{{UseLineStyle .UseLine}}")
}

func StyleFlags(tpl string) string {
	cf := color.New(color.FgCyan)

	// - styling short and full flags (-f, --flag)
	cobra.AddTemplateFunc("FlagStyle", func(s string) string {
		lines := strings.Split(s, "\n")
		for k := range lines {
			re := regexp.MustCompile(`(--?\S+)`)
			for _, flag := range re.FindAllString(lines[k], 2) {
				lines[k] = strings.Replace(lines[k], flag, cf.Sprint(flag), 1)
			}
		}
		s = strings.Join(lines, "\n")

		return s
	})

	// patch usage template
	re := regexp.MustCompile(`(?i)(\.(InheritedFlags|LocalFlags)\.FlagUsages)`)
	return re.ReplaceAllString(tpl, "FlagStyle $1")
}

func ApplyStyle(cmd *cobra.Command) {
	tpl := cmd.UsageTemplate()
	tpl = StyleHeaders(tpl)
	tpl = StyleCmds(tpl)
	tpl = StyleFlags(tpl)
	cmd.SetUsageTemplate(tpl)
}
