package converter

import (
	"encoding/json"
	"os"

	"gopkg.in/yaml.v3"
)

func JSONToYAML(inputFile string, outputFile string) error {

	data, err := os.ReadFile(inputFile)

	if err != nil {
		return err
	}

	var obj interface{}

	err = json.Unmarshal(data, &obj)

	if err != nil {
		return err
	}

	yamlData, err := yaml.Marshal(obj)

	if err != nil {
		return err
	}

	err = os.WriteFile(outputFile, yamlData, 0644)

	if err != nil {
		return err
	}

	return nil
}

func YAMLToJSON(inputFile string, outputFile string) error {

	data, err := os.ReadFile(inputFile)

	if err != nil {
		return err
	}

	var obj interface{}

	err = yaml.Unmarshal(data, &obj)

	if err != nil {
		return err
	}

	jsonData, err := json.MarshalIndent(obj, "", "  ")

	if err != nil {
		return err
	}

	err = os.WriteFile(outputFile, jsonData, 0644)

	if err != nil {
		return err
	}

	return nil
}
