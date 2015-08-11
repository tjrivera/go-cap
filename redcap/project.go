package redcap

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type RedcapProject struct {
	Forms        map[string]*RedcapForm
	Token        string
	Url          string
	Field_labels []string
	Unique_key   RedcapField
	Version      string
	Metadata     []RedcapField
	//Longitudinal attributes
	Events    []RedcapEvent
	Arm_names []string
	Arm_nums  []int
}

type ExportParameters struct {
	Records                []string
	Fields                 []string
	Forms                  []string
	Events                 []string
	RawOrLabel             string `default:"raw"`
	EventName              string `default:"label"`
	Format                 string `default:"json"`
	ExportSurveyFields     bool
	ExportDataAccessGroups bool
	ExportCheckboxLabel    bool
}

// containsForm is a convenience method to check whether a project contains
// a form.
func (project *RedcapProject) containsForm(form string) bool {
	for _, a := range project.Forms {
		if a.Name == form {
			return true
		}
	}
	return false
}

func (project *RedcapProject) GetMetadata() []RedcapField {
	if project.Metadata != nil {
		return project.Metadata
	}

	var fields []RedcapField

	res, err := http.PostForm(project.Url,
		url.Values{
			"token":   {project.Token},
			"content": {"metadata"},
			"format":  {"json"}})

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("Error reading REDCap body: %s", err)
	}

	err = json.Unmarshal(body, &fields)
	if err != nil {
		log.Fatalf("Error parsing REDCap metadata: %s", err)
	}

	// The first field should always be our unique key
	project.Unique_key = fields[0]
	project.Metadata = fields

	return fields
}

func (project *RedcapProject) GetFieldLabels() []string {

	var field_labels []string

	fields := project.GetMetadata()
	for _, field := range fields {
		field_labels = append(field_labels, field.Field_label)
	}

	project.Field_labels = field_labels

	return field_labels
}

func (project *RedcapProject) GetForms() map[string]*RedcapForm {

	project.Forms = make(map[string]*RedcapForm)
	var fields []RedcapField
	fields = project.GetMetadata()

	// Get unique list of forms from metadata
	for _, field := range fields {
		if !project.containsForm(field.Form_name) {
			f := RedcapForm{Name: field.Form_name, Unique_key: project.Unique_key}
			project.Forms[field.Form_name] = &f
			project.Forms[field.Form_name].addFieldToForm(field)
		} else {
			project.Forms[field.Form_name].addFieldToForm(field)
		}
		fmt.Println(project.Forms[field.Form_name])
	}

	return project.Forms
}

// ExportRecords creates a request to REDCap's API for record-type content
func (project *RedcapProject) ExportRecords(p ExportParameters) []byte {

	// Set default parameters
	if p.RawOrLabel == "" {
		p.RawOrLabel = "raw"
	}
	if p.EventName == "" {
		p.EventName = "label"
	}
	if p.Format == "" {
		p.Format = "json"
	}

	// Prepare array fields for query
	records := strings.Join(p.Records, ",")
	fields := strings.Join(p.Fields, ",")
	events := strings.Join(p.Events, ",")
	forms := strings.Join(p.Forms, ",")

	res, err := http.PostForm(project.Url,
		url.Values{
			// Required Parameters
			"token":   {project.Token},
			"content": {"record"},
			// Optional Parameters
			"records":                {records},
			"fields":                 {fields},
			"forms":                  {forms},
			"events":                 {events},
			"rawOrLabel":             {p.RawOrLabel},
			"eventName":              {p.EventName},
			"format":                 {p.Format},
			"exportSurveyFields":     {strconv.FormatBool(p.ExportSurveyFields)},
			"exportDataAccessGroups": {strconv.FormatBool(p.ExportDataAccessGroups)},
			"exportCheckboxLabel":    {strconv.FormatBool(p.ExportCheckboxLabel)}},
	)

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("Error retrieving records")
	}

	return body
}

func (project *RedcapProject) ToSQL(db string) string {
	s := ""
	for _, form := range project.Forms {
		s += form.ToSQL(db)
	}
	return s
}

// Initialize a RedcapProject instance with metadata
func (project *RedcapProject) initialize() {
	project.GetMetadata()
	project.GetFieldLabels()
	project.GetForms()
}

func NewRedcapProject(url string, token string, initialize bool) *RedcapProject {
	project := RedcapProject{
		Url:   url,
		Token: token,
	}

	if initialize {
		project.initialize()
	}

	return &project
}
