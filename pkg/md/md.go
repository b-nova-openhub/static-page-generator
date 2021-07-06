package md

import (
	"github.com/b-nova-openhub/stapagen/pkg/config"
	"github.com/b-nova-openhub/stapagen/pkg/util"
	"github.com/gomarkdown/markdown"
)

func ConvertBodyToMarkdown(s string) string {
	return string(markdown.ToHTML([]byte(getBody(s)), nil, nil))
}

func getBody(s string) string {
	return util.SubstringAfter(s, "</"+config.AppConfig.ContentDelimiterTag+">")
}
