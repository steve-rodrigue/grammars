package grammars

import (
	"errors"
	"fmt"
	"log"
	"path/filepath"

	"github.com/steve-care-software/grammars/domain/grammars/blocks/lines/tokens/elements/references"
)

type repositoryMemory struct {
	grammars map[string]map[uint]Grammar
}

func createRepositoryMemory(
	grammarsList map[string]Grammar,
) Repository {
	out := repositoryMemory{}
	out.Init()
	for path, oneGrammar := range grammarsList {
		err := out.Insert(filepath.SplitList(path), oneGrammar)
		if err != nil {
			log.Printf("there was an error while creating a repositoryMemory: %s", err.Error())
		}
	}

	return &out
}

// Init initializes the repository
func (app *repositoryMemory) Init() error {
	app.grammars = map[string]map[uint]Grammar{}
	return nil
}

// List lists the grammar paths
func (app *repositoryMemory) List() (map[string][]uint, error) {
	output := map[string][]uint{}
	for path, oneGrammarVersion := range app.grammars {
		if _, ok := output[path]; !ok {
			output[path] = []uint{}
		}

		for oneVersion, _ := range oneGrammarVersion {
			output[path] = append(output[path], oneVersion)
		}
	}

	return output, nil
}

// Insert inserts a grammar
func (app *repositoryMemory) Insert(path []string, grammar Grammar) error {
	pathStr := filepath.Join(path...)
	if _, ok := app.grammars[pathStr]; !ok {
		app.grammars[pathStr] = map[uint]Grammar{}
	}

	version := grammar.Version()
	app.grammars[pathStr][version] = grammar
	return nil
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

// Delete deletes a grammar
func (app *repositoryMemory) Delete(reference references.Reference) error {
	path := reference.Path()
	pathStr := filepath.Join(path...)
	if _, ok := app.grammars[pathStr]; !ok {
		str := fmt.Sprintf("there is no grammar at the provided path (%s)", pathStr)
		return errors.New(str)
	}

	version := reference.Version()
	if _, ok := app.grammars[pathStr][version]; !ok {
		str := fmt.Sprintf("there is no grammar at the provided path (%s) for version (%d)", pathStr, version)
		return errors.New(str)
	}

	delete(app.grammars[pathStr], version)
	if len(app.grammars[pathStr]) <= 0 {
		delete(app.grammars, pathStr)
	}

	return nil
}
