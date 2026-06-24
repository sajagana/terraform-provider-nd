// Copyright (c) 2026 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"text/template"

	"gopkg.in/yaml.v3"
)

// defsManifest models generator/defs/defs.yaml, which is the source list of
// resource YAML files that can produce documentation pages.
type defsManifest struct {
	Files []string `yaml:"files"`
}

// definitionFile models the top-level shape of each resource definition YAML.
type definitionFile struct {
	Resource ResourceDef `yaml:"resource"`
}

// ResourceDef is the resource-level YAML data used by the page frontmatter,
// title, description, and resource type name sections.
type ResourceDef struct {
	Name               string         `yaml:"name"`
	ResourceName       string         `yaml:"-"`
	Title              string         `yaml:"resource_title"`
	DocCategory        string         `yaml:"doc_category"`
	Category           string         `yaml:"category"`
	GenerateTFResource bool           `yaml:"generate_tf_resource"`
	Description        string         `yaml:"description"`
	ResourceNote       string         `yaml:"resource_note"`
	ImportNote         string         `yaml:"import_note"`
	APIEndpoint        string         `yaml:"api_endpoint"`
	APIReference       string         `yaml:"api_reference"`
	UILocation         string         `yaml:"ui_location"`
	Attributes         []AttributeDef `yaml:"attributes"`
}

// AttributeDef is the attribute-level YAML data used by the Required,
// Optional, and Read-Only schema sections.
type AttributeDef struct {
	ModelName    string         `yaml:"model_name"`
	TFName       string         `yaml:"tf_name"`
	Description  string         `yaml:"description"`
	Type         string         `yaml:"type"`
	Mandatory    bool           `yaml:"mandatory"`
	Optional     bool           `yaml:"optional"`
	Computed     bool           `yaml:"computed"`
	Sensitive    bool           `yaml:"sensitive"`
	PayloadHide  bool           `yaml:"payload_hide"`
	TFHide       bool           `yaml:"tf_hide"`
	DefaultValue any            `yaml:"default_value"`
	MinInt       *int64         `yaml:"min_int"`
	MaxInt       *int64         `yaml:"max_int"`
	MinFloat     *float64       `yaml:"min_float"`
	MaxFloat     *float64       `yaml:"max_float"`
	Attributes   []AttributeDef `yaml:"attributes"`
	ValidValues  []string       `yaml:"valid_values"`
}

// TemplateData is the complete data contract passed into resource doc templates.
type TemplateData struct {
	Resource           ResourceDef
	ExampleHCL         string
	ExampleSource      string
	ExamplesURL        string
	RequiredAttributes []AttributeDef
	OptionalAttributes []AttributeDef
	ReadOnlyAttributes []AttributeDef
}

// -----------------------------------------------------------------------------
// Command Orchestration
// -----------------------------------------------------------------------------

// main reads CLI flags and starts resource documentation generation.
func main() {
	defsPath := flag.String("defs", "generator/defs/defs.yaml", "path to defs manifest")
	templatesDir := flag.String("templates", "templates", "path to templates directory")
	outDir := flag.String("out", "docs/resources", "output directory for rendered resource docs")
	providerPrefix := flag.String("provider-prefix", "nd", "Terraform provider prefix for resource names")
	examplesDir := flag.String("examples-dir", "examples/resources", "path to resource examples directory")
	flag.Parse()

	if err := run(*defsPath, *templatesDir, *outDir, *providerPrefix, *examplesDir); err != nil {
		fmt.Fprintf(os.Stderr, "resource-docgen: %v\n", err)
		os.Exit(1)
	}
}

// run coordinates the end-to-end flow: manifest -> YAML definition -> template
// data -> rendered resource Markdown file.
func run(defsPath, templatesDir, outDir, providerPrefix, examplesDir string) error {
	manifest, err := readManifest(defsPath)
	if err != nil {
		return err
	}

	if err := os.MkdirAll(outDir, 0755); err != nil {
		return fmt.Errorf("create output directory %q: %w", outDir, err)
	}

	defsDir := filepath.Dir(defsPath)
	for _, name := range manifest.Files {
		defPath := filepath.Join(defsDir, name)
		data, err := readDefinition(defPath, providerPrefix, examplesDir)
		if err != nil {
			return err
		}
		if data.Resource.Name == "" || !data.Resource.GenerateTFResource {
			continue
		}

		tmplPath, err := selectTemplatePath(templatesDir, data.Resource.Name)
		if err != nil {
			return err
		}
		rendered, err := renderTemplate(tmplPath, data)
		if err != nil {
			return err
		}

		outPath := filepath.Join(outDir, data.Resource.Name+".md")
		if err := os.WriteFile(outPath, []byte(rendered), 0644); err != nil {
			return fmt.Errorf("write %q: %w", outPath, err)
		}
		fmt.Printf("generated %s from %s using %s\n", outPath, defPath, tmplPath)
	}

	return nil
}

// -----------------------------------------------------------------------------
// Source YAML Loading
// -----------------------------------------------------------------------------

// readManifest loads defs.yaml so only explicitly listed definition files are
// considered for documentation generation.
func readManifest(path string) (*defsManifest, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read defs manifest %q: %w", path, err)
	}

	var manifest defsManifest
	if err := yaml.Unmarshal(data, &manifest); err != nil {
		return nil, fmt.Errorf("parse defs manifest %q: %w", path, err)
	}
	return &manifest, nil
}

// readDefinition parses one resource YAML file and builds all data required by
// the document sections in the output template.
func readDefinition(path, providerPrefix, examplesDir string) (TemplateData, error) {
	rawYAML, err := os.ReadFile(path)
	if err != nil {
		return TemplateData{}, fmt.Errorf("read definition %q: %w", path, err)
	}

	var def definitionFile
	if err := yaml.Unmarshal(rawYAML, &def); err != nil {
		return TemplateData{}, fmt.Errorf("parse definition %q: %w", path, err)
	}
	if def.Resource.Name == "" || !def.Resource.GenerateTFResource {
		return TemplateData{Resource: def.Resource}, nil
	}

	def.Resource.ResourceName = providerPrefix + "_" + def.Resource.Name
	if def.Resource.Title == "" {
		def.Resource.Title = titleFromName(def.Resource.Name)
	}
	if def.Resource.DocCategory == "" {
		def.Resource.DocCategory = def.Resource.Category
	}

	required, optional, readOnly := groupAttributes(def.Resource.Attributes)
	exampleHCL, exampleSource, err := readExampleHCL(def.Resource, examplesDir)
	if err != nil {
		return TemplateData{}, err
	}
	return TemplateData{
		Resource:           def.Resource,
		ExampleHCL:         exampleHCL,
		ExampleSource:      exampleSource,
		ExamplesURL:        examplesURL(def.Resource.ResourceName),
		RequiredAttributes: required,
		OptionalAttributes: optional,
		ReadOnlyAttributes: readOnly,
	}, nil
}

// groupAttributes prepares the Required Attributes, Optional Attributes, and
// Read-Only Attributes sections.
func groupAttributes(attrs []AttributeDef) (required, optional, readOnly []AttributeDef) {
	for _, attr := range attrs {
		if skipAttribute(attr) {
			continue
		}
		switch {
		case attr.Mandatory:
			required = append(required, attr)
		case attr.Optional:
			optional = append(optional, attr)
		case attr.Computed || attr.PayloadHide:
			readOnly = append(readOnly, attr)
		}
	}
	return required, optional, readOnly
}

// selectTemplatePath returns the resource-specific template when it exists,
// otherwise it falls back to the generic resource template.
func selectTemplatePath(templatesDir, resourceName string) (string, error) {
	resourceTemplatePath := filepath.Join(templatesDir, "resources", resourceName+".md.tmpl")
	if _, err := os.Stat(resourceTemplatePath); err == nil {
		return resourceTemplatePath, nil
	} else if !os.IsNotExist(err) {
		return "", fmt.Errorf("check resource template %q: %w", resourceTemplatePath, err)
	}

	return filepath.Join(templatesDir, "resources.md.tmpl"), nil
}

// renderTemplate executes the selected Go template with the YAML-derived
// TemplateData.
func renderTemplate(path string, data TemplateData) (string, error) {
	funcs := template.FuncMap{
		"renderAttribute": renderAttributeMarkdown,
	}

	tmpl, err := template.New(filepath.Base(path)).Funcs(funcs).ParseFiles(path)
	if err != nil {
		return "", fmt.Errorf("parse template %q: %w", path, err)
	}

	var out bytes.Buffer
	if err := tmpl.Execute(&out, data); err != nil {
		return "", fmt.Errorf("render template %q for %q: %w", path, data.Resource.Name, err)
	}
	return out.String(), nil
}

// renderAttributeMarkdown renders a schema attribute with only the doc fields
// we want to expose: Terraform name, model name, type, default, valid values,
// allowed range, and nested child attributes.
func renderAttributeMarkdown(attr AttributeDef) string {
	var b strings.Builder
	writeAttributeMarkdown(&b, attr, 0)
	return strings.TrimRight(b.String(), "\n")
}

func writeAttributeMarkdown(b *strings.Builder, attr AttributeDef, depth int) {
	if skipAttribute(attr) {
		return
	}

	indent := strings.Repeat("    ", depth)
	detailIndent := strings.Repeat("    ", depth+1)
	fmt.Fprintf(b, "%s* `%s` (%s) - (%s) %s\n", indent, attr.TFName, attr.ModelName, attributeTypeLabel(attr), attr.Description)
	if !valueIsEmpty(attr.DefaultValue) {
		fmt.Fprintf(b, "%s* Default Value: `%s`\n", detailIndent, markdownValue(attr.DefaultValue))
	}
	if len(attr.ValidValues) > 0 {
		fmt.Fprintf(b, "%s* Valid Values: %s.\n", detailIndent, markdownValues(attr.ValidValues))
	}
	if allowedRange := allowedRange(attr); allowedRange != "" {
		fmt.Fprintf(b, "%s* Allowed Range: `%s`\n", detailIndent, allowedRange)
	}
	for _, child := range attr.Attributes {
		writeAttributeMarkdown(b, child, depth+1)
	}
}

func skipAttribute(attr AttributeDef) bool {
	return attr.TFName == "" || attr.TFHide
}

func attributeTypeLabel(attr AttributeDef) string {
	if attr.Sensitive {
		return "Sensitive, " + attr.Type
	}
	return attr.Type
}

// -----------------------------------------------------------------------------
// Example Usage Section
// -----------------------------------------------------------------------------

// readExampleHCL loads the required resource example file for the Example Usage
// section.
func readExampleHCL(resource ResourceDef, examplesDir string) (string, string, error) {
	examplePath := filepath.Join(examplesDir, resource.ResourceName, "resource.tf")
	data, err := os.ReadFile(examplePath)
	if err != nil {
		return "", "", fmt.Errorf("read example file %q: %w", examplePath, err)
	}
	return strings.TrimSpace(string(data)), examplePath, nil
}

// titleFromName turns a YAML resource name into a readable title for docs.
func titleFromName(name string) string {
	parts := strings.Split(name, "_")
	for i, part := range parts {
		if part == "" {
			continue
		}
		parts[i] = strings.ToUpper(part[:1]) + part[1:]
	}
	return strings.Join(parts, " ")
}

// examplesURL returns the public examples folder URL for the Terraform type.
func examplesURL(typeName string) string {
	return "https://github.com/CiscoDevNet/terraform-provider-nd/tree/main/examples/resources/" + typeName
}

func markdownValues(values []string) string {
	var b strings.Builder
	for i, value := range values {
		if i > 0 {
			b.WriteString(", or ")
		}
		fmt.Fprintf(&b, "`%s`", value)
	}
	return b.String()
}

// allowedRange renders numeric ranges from the generator's existing min/max
// YAML fields.
func allowedRange(attr AttributeDef) string {
	if attr.MinInt != nil || attr.MaxInt != nil {
		return rangeText(formatOptionalInt(attr.MinInt), formatOptionalInt(attr.MaxInt))
	}
	if attr.MinFloat != nil || attr.MaxFloat != nil {
		return rangeText(formatOptionalFloat(attr.MinFloat), formatOptionalFloat(attr.MaxFloat))
	}
	return ""
}

func rangeText(minValue, maxValue string) string {
	switch {
	case minValue != "" && maxValue != "":
		return minValue + " to " + maxValue
	case minValue != "":
		return minValue + " or higher"
	case maxValue != "":
		return "up to " + maxValue
	default:
		return ""
	}
}

func formatOptionalInt(value *int64) string {
	if value == nil {
		return ""
	}
	return strconv.FormatInt(*value, 10)
}

func formatOptionalFloat(value *float64) string {
	if value == nil {
		return ""
	}
	return strconv.FormatFloat(*value, 'f', -1, 64)
}

// -----------------------------------------------------------------------------
// Markdown Value Formatting Helpers
// -----------------------------------------------------------------------------

// markdownValue renders a YAML scalar as inline Markdown code content.
func markdownValue(value any) string {
	switch v := value.(type) {
	case nil:
		return ""
	case bool:
		return strconv.FormatBool(v)
	case int:
		return strconv.Itoa(v)
	case int64:
		return strconv.FormatInt(v, 10)
	case float64:
		return strconv.FormatFloat(v, 'f', -1, 64)
	case string:
		return strings.TrimSpace(v)
	default:
		return fmt.Sprintf("%v", v)
	}
}

// valueIsEmpty reports whether a YAML value should be treated as absent for
// documentation rendering.
func valueIsEmpty(value any) bool {
	if value == nil {
		return true
	}
	if s, ok := value.(string); ok {
		return strings.TrimSpace(s) == ""
	}
	return false
}
