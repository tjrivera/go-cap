package redcap

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProjectForms(t *testing.T) {
	var project = NewRedcapProject("https://redcap.chop.edu/api/", "71F3F25FC3BCF2232E27298E7AFBC636", true)
	assert.Equal(t, len(project.Forms), 3)
}

func TestFormToSQL(t *testing.T) {
	var project = NewRedcapProject("https://redcap.chop.edu/api/", "71F3F25FC3BCF2232E27298E7AFBC636", true)
	s := `
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
);`
	assert.Equal(t, s, project.Forms["demographics"].ToSQL("postgres"))
}
