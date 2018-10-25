package designers

import "github.com/go-modules-by-example-staging//home/gopher/scratchpad/repos/goinfo-maj-branch/contributors"

func FullNames() []string {
	var res []string
	for _, p := range contributors.Details() {
		switch p.FullName {
		case "Rob Pike", "Ken Thompson", "Robert Griesemer":
			res = append(res, p.FullName)
		}
	}
	return res
}
