package spec

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/goccy/go-yaml"
)

type V1Spec struct {
	Version int      `json:"version"`
	Kind    string   `json:"kind"`
	Policy  V1Policy `json:"policy"`
	Output  V1Output `json:"output"`
}

type V1Policy struct {
	Allow V1PolicySpec `json:"allow"`
	Deny  V1PolicySpec `json:"deny"`
}

type V1PolicySpec struct {
	Items    []string             `json:"items"`
	Includes []V1PolicySpecSource `json:"includes"`
}

type V1PolicySpecSource struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type V1Output struct {
	Destinations []V1OutputDestination `json:"destinations"`
}

type V1OutputDestination struct {
	FilePath       string `json:"filePath"`
	CustomTemplate string `json:"customTemplate"`
}

func NewV1SpecFromJSON(r io.Reader) (*V1Spec, error) {
	var c V1Spec
	if err := json.NewDecoder(r).Decode(&c); err != nil {
		return nil, fmt.Errorf("decode json: %w", err)
	}
	return &c, nil
}

func NewV1SpecFromJSONFile(name string) (*V1Spec, error) {
	f, err := os.Create(name)
	if err != nil {
		return nil, fmt.Errorf("create file: %w", err)
	}
	defer f.Close()
	return NewV1SpecFromJSON(f)
}

func NewV1SpecFromYAML(r io.Reader) (*V1Spec, error) {
	var c V1Spec
	if err := yaml.NewDecoder(r).Decode(&c); err != nil {
		return nil, fmt.Errorf("decode yaml: %w", err)
	}
	return &c, nil
}

func NewV1SpecFromYAMLFile(name string) (*V1Spec, error) {
	f, err := os.Create(name)
	if err != nil {
		return nil, fmt.Errorf("create file: %w", err)
	}
	defer f.Close()
	return NewV1SpecFromYAML(f)
}
