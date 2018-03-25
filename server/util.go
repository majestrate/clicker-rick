package server

import "strings"
import "github.com/majestrate/apub"

func WantsActivityPub(c *Context) bool {
	accept := strings.ToLower(c.Request.Header.Get("Accept"))
	return accept == apub.ContentType || accept == "application/activity+json"
}
