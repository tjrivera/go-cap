package redcap

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var project = NewRedcapProject("https://redcap.vanderbilt.edu/api/", "8E66DB6844D58E990075AFB51658A002", true)

// Uninitialized Project
var project_uninit = NewRedcapProject("https://redcap.vanderbilt.edu/api/", "8E66DB6844D58E990075AFB51658A002", false)

func TestProjectForms(t *testing.T) {
	assert.Equal(t, len(project.Forms), 3)
}

func TestFormToSQL(t *testing.T) {
	s := `
CREATE TABLE testing
(
	study_id text,
	redcap_event_name text,
	foo_score text,
	bar_score text,
	form_status text
);`
	assert.Equal(t, s, project.Forms["testing"].ToSQL("postgres"))
}

func TestProjectToSQL(t *testing.T) {
	// Probably fails due to inconsistent sorting of tables
	s := `
CREATE TABLE demographics
(
	study_id text,
	redcap_event_name text,
	first_name text,
	last_name text,
	dob text,
	sex text,
	address text,
	phone_number text,
	file text,
	matrix1 text,
	matrix2 text,
	matrix3 text,
	matcheck1___1 text,
	matcheck1___2 text,
	matcheck1___3 text,
	matcheck2___1 text,
	matcheck2___2 text,
	matcheck2___3 text,
	matcheck3___1 text,
	matcheck3___2 text,
	matcheck3___3 text,
	form_status text
);
CREATE TABLE testing
(
	study_id text,
	redcap_event_name text,
	foo_score text,
	bar_score text,
	form_status text
);
CREATE TABLE imaging
(
	study_id text,
	redcap_event_name text,
	image_path text,
	form_status text
);`
	assert.Equal(t, s, project.ToSQL("postgres"))
}

func TestProjectFieldLabels(t *testing.T) {
	assert.Equal(t, len(project.Field_labels), 17)

}
