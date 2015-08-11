package redcap

import (
	"fmt"
	"strings"
)

type RedcapForm struct {
	Name       string
	Fields     []RedcapField
	Unique_key RedcapField
}

func (form *RedcapForm) String() string {
	return fmt.Sprintf("REDCAP Form: %s", form.Name)
}

func (form *RedcapForm) containsField(field RedcapField) bool {
	for _, f := range form.Fields {
		if f.Field_name == field.Field_name {
			return true
		}
	}
	return false
}

func (form *RedcapForm) addFieldToForm(field RedcapField) {
	form.Fields = append(form.Fields, field)
}

// ToSQL generates PostgreSQL flavored DDL. go-cap makes no attempt
// to infer data types within REDCap forms and will defer to TEXT as
// its default datatype.
func (form *RedcapForm) ToSQL(db string) string {

	if db == "postgres" {

		s := fmt.Sprintf("\nCREATE TABLE %s\n(\n", form.Name)
		s += fmt.Sprintf("\t%s text,\n", form.Unique_key.Field_name)
		s += fmt.Sprintf("\tredcap_event_name text,\n")

		for _, field := range form.Fields {
			// Handle checkbox fields
			if (len(field.Select_choices_or_calculations) > 0) && field.Field_type == "checkbox" {
				for _, choice := range field.Select_choices_or_calculations {
					s += fmt.Sprintf("\t%s___%d %s,\n", field.Field_name, choice.Id, "text")
				}
			} else {
				// Suppress study ID if it is in the form
				if field.Field_name == form.Unique_key.Field_name {
					continue
				} else {
					s += fmt.Sprintf("\t%s %s,\n", field.Field_name, "text")
				}
			}
		}

		s += fmt.Sprintf("\tform_status text,\n")
		s = strings.TrimRight(s, ",\n") + "\n);"

		return s
	} else {
		fmt.Printf("The provided SQL dialect (%s) is not supported.", db)
		return ""
	}

}
