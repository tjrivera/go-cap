package redcap

import (
	"encoding/json"
	"log"
	"strconv"
	"strings"
)

type RedcapField struct {
	Branching_logic                            string
	Custom_alignment                           string
	Field_label                                string
	Field_name                                 string
	Field_note                                 string
	Field_type                                 string
	Form_name                                  string
	Identifier                                 string
	Matrix_group_name                          string
	Matrix_ranking                             string
	Question_number                            string
	Required_field                             bool
	Section_header                             string
	Choices                                    []RedcapFieldChoice
	Calculations                               string
	Text_validation_max                        string
	Text_validation_min                        string
	Text_validation_type_or_show_slider_number string
	Value                                      string
}

type RedcapFieldChoice struct {
	Id    int
	Label string
}

func (field *RedcapField) UnmarshalJSON(raw []byte) error {

	var tmp = make(map[string]interface{})
	err := json.Unmarshal(raw, &tmp)
	if err != nil {
		log.Fatal(err)
	}

	field.Branching_logic = tmp["branching_logic"].(string)
	field.Custom_alignment = tmp["custom_alignment"].(string)
	field.Field_label = tmp["field_label"].(string)
	field.Field_name = tmp["field_name"].(string)
	field.Field_note = tmp["field_note"].(string)
	field.Field_type = tmp["field_type"].(string)
	field.Form_name = tmp["form_name"].(string)
	field.Identifier = tmp["identifier"].(string)
	field.Matrix_group_name = tmp["matrix_group_name"].(string)
	field.Matrix_ranking = tmp["matrix_ranking"].(string)
	field.Question_number = tmp["question_number"].(string)
	field.Section_header = tmp["section_header"].(string)
	field.Text_validation_max = tmp["text_validation_max"].(string)
	field.Text_validation_min = tmp["text_validation_min"].(string)
	field.Text_validation_type_or_show_slider_number = tmp["text_validation_type_or_show_slider_number"].(string)

	// Marshal Redcap Choices
	var choices []RedcapFieldChoice
	if tmp["field_type"] == "checkbox" || tmp["field_type"] == "dropdown" || tmp["field_type"] == "radio" {
		for _, choice := range strings.Split(tmp["select_choices_or_calculations"].(string), "|") {
			choice := strings.TrimSpace(choice)
			if choice != "" {
				s := strings.Split(choice, ",")
				id, err := strconv.Atoi(s[0])

				if err != nil {
					log.Fatal("[go-cap] error marshalling redcap choices: ", err)
				}
				label := s[1]
				choices = append(choices, RedcapFieldChoice{id, label})
			}
		}
	} else if tmp["field_type"] == "calc" {
		field.Calculations = tmp["select_choices_or_calculations"].(string)
	}

	field.Choices = choices

	// Marshal Required Field flag
	if tmp["required_field"].(string) == "Y" {
		field.Required_field = true
	} else {
		field.Required_field = false
	}

	return nil

}
