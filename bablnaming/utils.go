package bablnaming

import (
	"regexp"
	"strings"
)

func RequestPathToTopic(method string) string {
	return strings.Replace(method[1:], "/", ".", 1)
}

func TopicToModule(topic string) string {
	parts := strings.Split(topic, ".")
	owner := parts[1]
	camel := parts[2]

	name := regexp.MustCompile(`[A-Z]+`).ReplaceAllStringFunc(camel, func(m string) string {
		return "-" + strings.ToLower(m)
	})
	name = strings.TrimLeft(name, "-")

	return owner + "/" + name
}

func ModuleToTopic(module string, meta bool) string {
	function := "IO"
	if meta {
		function = "meta"
	}

	parts := strings.Split(module, "/")
	owner := parts[0]
	name := strings.Replace(strings.Title(parts[1]), "-", "", -1)

	return "babl." + owner + "." + name + "." + function
}
