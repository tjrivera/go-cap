package redcap

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProjectToSQL(t *testing.T) {
	var project = NewRedcapProject("https://redcap.chop.edu/api/", "71F3F25FC3BCF2232E27298E7AFBC636", true)
	s := `
CREATE TABLE baseline_visit_data
(
	study_id text,
	redcap_event_name text,
	specify_mood text,
	meds___1 text,
	meds___2 text,
	meds___3 text,
	meds___4 text,
	meds___5 text,
	height text,
	weight text,
	comments text,
	prealb_b text,
	creat_b text,
	chol_b text,
	transferrin_b text,
	ibd_flag text,
	general_ibd text,
	chrons text,
	ulcerative_colitis text,
	colonoscopy text,
	colonoscopy_date text,
	form_status text
);
CREATE TABLE demographics
(
	study_id text,
	redcap_event_name text,
	date_enrolled text,
	ethnicity text,
	race text,
	sex text,
	given_birth text,
	num_children text,
	form_status text
);
CREATE TABLE meal_description_form
(
	study_id text,
	redcap_event_name text,
	meal_description text,
	types_of_food text,
	healthy text,
	form_status text
);`
	assert.Equal(t, s, project.ToSQL("postgres"))
}

func TestProjectFieldLabels(t *testing.T) {
	var project = NewRedcapProject("https://redcap.chop.edu/api/", "71F3F25FC3BCF2232E27298E7AFBC636", true)
	assert.Equal(t, len(project.Field_labels), 25)
}

func TestProjectUniqueKey(t *testing.T) {
	var project = NewRedcapProject("https://redcap.chop.edu/api/", "71F3F25FC3BCF2232E27298E7AFBC636", true)
	f := RedcapField{
		Branching_logic:                            "",
		Custom_alignment:                           "",
		Field_label:                                "Study ID",
		Field_name:                                 "study_id",
		Field_note:                                 "",
		Field_type:                                 "text",
		Form_name:                                  "demographics",
		Identifier:                                 "",
		Matrix_group_name:                          "",
		Matrix_ranking:                             "",
		Question_number:                            "",
		Required_field:                             false,
		Section_header:                             "",
		Choices:                                    []RedcapFieldChoice(nil),
		Text_validation_max:                        "",
		Text_validation_min:                        "",
		Text_validation_type_or_show_slider_number: ""}
	assert.Equal(t, project.Unique_key, f)
}

func TestExportRecords(t *testing.T) {
	var project = NewRedcapProject("https://redcap.chop.edu/api/", "71F3F25FC3BCF2232E27298E7AFBC636", true)
	var params = ExportParameters{
		Records: []string{"S891XSB0XD1NKRPF:4G0AZ5XFH"},
		Format:  "json"}
	var s string
	s = `[{"study_id":"S891XSB0XD1NKRPF:4G0AZ5XFH","redcap_event_name":"enrollment_arm_1","date_enrolled":"2015-09-01","ethnicity":"1","race":"4","sex":"1","given_birth":"","num_children":"","demographics_complete":"0","specify_mood":"","meds___1":"","meds___2":"","meds___3":"","meds___4":"","meds___5":"","height":"","weight":"","comments":"","prealb_b":"","creat_b":"","chol_b":"","transferrin_b":"","ibd_flag":"","general_ibd":"","chrons":"","ulcerative_colitis":"","colonoscopy":"","colonoscopy_date":"","baseline_visit_data_complete":"","meal_description":"","types_of_food":"","healthy":"","meal_description_form_complete":""}]`
	res := string(project.ExportRecords(params))
	assert.Equal(t, s, res)
}

func TestGetRecords(t *testing.T) {
	var project = NewRedcapProject("https://redcap.chop.edu/api/", "71F3F25FC3BCF2232E27298E7AFBC636", true)
	records := project.GetRecords()
	assert.True(t, len(records) > 0)
}
