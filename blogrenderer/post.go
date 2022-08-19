package blogrenderer

import (
	"html/template"
	"strings"

	"github.com/gomarkdown/markdown"
)

type Post struct {
	Title       string
	Body        string
	Description string
	Tags        []string
}

func (p Post) SanitisedTitle() string {
	return strings.ToLower(strings.Replace(p.Title, " ", "-", -1))
}

type HTMLPost struct {
	Post
	HTMLBody string
}

func (p Post) Render() HTMLPost {
	md := template.HTML(markdown.ToHTML([]byte(p.Body), nil, nil))
	return HTMLPost{p, string(md)}
}
