package redcap

import (
    "testing"
    "github.com/stretchr/testify/assert"
)


var project = NewRedcapProject("https://redcap.vanderbilt.edu/api/", "8E66DB6844D58E990075AFB51658A002", true)
// Uninitialized Project
var project_uninit = NewRedcapProject("https://redcap.vanderbilt.edu/api/", "8E66DB6844D58E990075AFB51658A002", false)

func TestProjectForms(t *testing.T) {

    assert.Equal(t, len(project.Forms), 3)
}

func TestFormToSQL(t *testing.T) {
    s:= `
CREATE TABLE testing
(
	foo_score text,
	bar_score text
);`
    assert.Equal(t, s, project.Forms["testing"].ToSQL("postgres"))
}

func TestProjectToSQL(t *testing.T) {
    // Probably fails due to inconsistent sorting of tables
    s:= `
CREATE TABLE demographics
(
	study_id text,
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
	matcheck1 text,
	matcheck2 text,
	matcheck3 text
);
CREATE TABLE testing
(
	foo_score text,
	bar_score text
);
CREATE TABLE imaging
(
	image_path text
);`
    assert.Equal(t, s, project.ToSQL("postgres"))
}

func TestProjectFieldLabels(t *testing.T){
    assert.Equal(t, len(project.Field_labels), 17)

}
