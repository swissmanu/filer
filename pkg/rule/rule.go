package rule

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
	"olymp.alabor.me/dev/git/swissmanu/filer/pkg/conf"
)

type Rules struct {
	Rules []Rule `json:"rules"`
}

type Rule struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Actions     []Action `json:"actions"`
}

type Action struct {
	Type   string `json:"type"`
	Target string `json:"target"`
}

func ReadRules(config conf.Specification) (*Rules, error) {
	data, err := ioutil.ReadFile(config.RulesPath)
	if err != nil {
		return nil, err
	}

	c := &Rules{}
	err = yaml.Unmarshal(data, c)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func FindRule(rules []Rule, ruleName string) (*Rule, error) {
	for _, r := range rules {
		if r.Name == ruleName {
			return &r, nil
		}
	}

	return nil, errors.New("Could not find rule named " + ruleName)
}

func ApplyRule(rule *Rule, sourcePath string, config *conf.Specification) error {
	for _, a := range rule.Actions {
		err := executeAction(a, sourcePath, config)
		if err != nil {
			return err
		}
	}
	return nil
}

func executeAction(action Action, sourcePath string, config *conf.Specification) error {
	if action.Type == "move" {
		targetPath := filepath.Join(config.DataPath, action.Target, filepath.Base(sourcePath))

		os.MkdirAll(filepath.Dir(targetPath), os.ModePerm)

		err := moveFile(sourcePath, targetPath)
		if err != nil {
			return err
		}
	} else {
		return errors.New("Cannot execute unknown action " + action.Type + " to " + sourcePath)
	}
	return nil
}

func moveFile(sourcePath, destPath string) error {
	// https://gist.github.com/var23rav/23ae5d0d4d830aff886c3c970b8f6c6b
	// because os.Rename does not work with volume mounts
	inputFile, err := os.Open(sourcePath)
	if err != nil {
		return fmt.Errorf("Couldn't open source file: %s", err)
	}
	outputFile, err := os.Create(destPath)
	if err != nil {
		inputFile.Close()
		return fmt.Errorf("Couldn't open dest file: %s", err)
	}
	defer outputFile.Close()
	_, err = io.Copy(outputFile, inputFile)
	inputFile.Close()
	if err != nil {
		return fmt.Errorf("Writing to output file failed: %s", err)
	}
	// The copy was successful, so now delete the original file
	err = os.Remove(sourcePath)
	if err != nil {
		return fmt.Errorf("Failed removing original file: %s", err)
	}
	return nil
}
