package grammars

import (
	"errors"
	"fmt"
	"path/filepath"

	"github.com/steve-care-software/grammars/domain/grammars/blocks/lines/tokens/elements/references"
)

type repositoryMemory struct {
	grammars map[string]map[uint]Grammar
}

func createRepositoryMemory(
	grammarsList map[string]Grammar,
) Repository {
	mp := map[string]map[uint]Grammar{}
	for path, oneGrammar := range grammarsList {
		if _, ok := mp[path]; !ok {
			mp[path] = map[uint]Grammar{}
		}

		version := oneGrammar.Version()
		mp[path][version] = oneGrammar
	}

	out := repositoryMemory{
		grammars: mp,
	}

	return &out
}

// Retrieve retrieves a memory repository
func (app *repositoryMemory) Retrieve(reference references.Reference) (Grammar, error) {
	path := filepath.Join(reference.Path()...)
	version := reference.Version()
	if versionGrammar, ok := app.grammars[path]; ok {
		if ins, ok := versionGrammar[version]; ok {
			return ins, nil
		}

		str := fmt.Sprintf("the version (%d) of the provided grammar path (%s) could not be found", version, path)
		return nil, errors.New(str)
	}

	str := fmt.Sprintf("the grammar (path: %s) could not be found", path)
	return nil, errors.New(str)
}
