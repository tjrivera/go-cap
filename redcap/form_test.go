package redcap

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

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
