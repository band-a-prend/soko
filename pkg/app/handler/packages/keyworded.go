// Used to show recently keyworded versions

package packages

import (
	"net/http"
	"soko/pkg/app/handler/feeds"
)

// Keyworded renders a template containing
// a list of 50 recently keyworded versions
func Keyworded(w http.ResponseWriter, r *http.Request) {
	keywordedVersions := GetKeywordedVersions(50)
	RenderPackageTemplates("changedVersions", "changedVersions", "changedVersionRow", GetFuncMap(), CreateFeedData("Keyworded", keywordedVersions), w)
}

func KeywordedFeed(w http.ResponseWriter, r *http.Request) {
	keywordedVersions := GetKeywordedVersions(1000)
	feeds.Changes(GetTextFuncMap(), CreateFeedData("Keyworded", keywordedVersions), w)
}
