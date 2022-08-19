package blogrenderer

import "github.com/gomarkdown/markdown"

type Post struct {
	Title       string
	Body        string
	Description string
	Tags        []string
}

type HTMLPost struct {
	Post
	HTMLBody string
}

func (p Post) Render() HTMLPost {
	md := markdown.ToHTML([]byte(p.Body), nil, nil)
	return HTMLPost{p, string(md)}
}
