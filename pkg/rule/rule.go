package rule

import (
	"errors"
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
	Name    string   `json:"name"`
	Actions []Action `json:"actions"`
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

		err := os.Rename(sourcePath, targetPath)
		if err != nil {
			return err
		}
	} else {
		return errors.New("Cannot apply unknown action " + action.Type + " to " + sourcePath)
	}
	return nil
}
