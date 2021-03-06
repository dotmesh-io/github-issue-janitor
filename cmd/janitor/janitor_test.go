package main

import (
	"testing"
)

func testParseBodyForIssueLinks(t *testing.T) {

	sampleBody := `As discussed in the Engineering Banner meeting, we want to tame the board: https://docs.google.com/document/d/1XUJRJ1upx3vn0jyu1JNQL-6lfgp4YMC_fZZ_tPJfVHA/edit#bookmark=id.ggjx2k8m2k4x

Here's the actions we made issues of to make this happen:

- [x] https://github.com/dotmesh-io/github-issue-janitor/issues/3 - make the janitor maintain the labels in every repo, so we can easily add a "theme" label (same colour as epics, but a different name)
- [ ] https://github.com/dotmesh-io/github-issue-janitor/issues/6 - make the janitor allow issues to exist "outside" of a board as long as they're mentioned in an epic or theme issue.
- [ ] https://github.com/dotmesh-io/github-issue-janitor/issues/7 - automate cleanup of done issues from project boards
- [ ] #346 - with that groundwork in place, we tidy up the Engineering board.`

	il := parseBodyForIssueLinks(sampleBody, "testrepo")
	t.Errorf("Got %#v", il)
}
