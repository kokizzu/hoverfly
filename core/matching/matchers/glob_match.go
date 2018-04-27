package matchers

import glob "github.com/ryanuber/go-glob"

func GlobMatch(match interface{}, toMatch string) bool {
	matchString, ok := match.(string)
	if !ok {
		return false
	}

	return glob.Glob(matchString, toMatch)
}
