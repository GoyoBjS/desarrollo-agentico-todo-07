package models

import (
	"fmt"
	"strings"
	"time"

	"github.com/gobuffalo/pop/v6"
	"github.com/gobuffalo/validate/v3"
	"github.com/gofrs/uuid"
)

// TitleMaxLength es el número máximo de caracteres permitido en el título.
const TitleMaxLength = 200

// DueDateLayout es el formato de la fecha límite en los formularios: solo el día.
const DueDateLayout = "2006-01-02"

// Task representa una tarea de la lista.
type Task struct {
	ID        uuid.UUID  `json:"id" db:"id" form:"-"`
	Title     string     `json:"title" db:"title" form:"title"`
	Completed bool       `json:"completed" db:"completed" form:"-"`
	DueDate   *time.Time `json:"due_date" db:"due_date" form:"-"`
	CreatedAt time.Time  `json:"created_at" db:"created_at" form:"-"`
	UpdatedAt time.Time  `json:"updated_at" db:"updated_at" form:"-"`

	// dueDateInput guarda el valor recibido por SetDueDate para que Validate lo compruebe.
	dueDateInput string `db:"-"`
}

// Tasks es una colección de Task.
type Tasks []Task

// Validate se ejecuta en cada llamada a los métodos "pop.Validate*".
func (t *Task) Validate(tx *pop.Connection) (*validate.Errors, error) {
	verrs := validate.NewErrors()

	title := strings.TrimSpace(t.Title)
	if title == "" {
		verrs.Add("title", "El título es obligatorio.")
	} else if len([]rune(title)) > TitleMaxLength {
		verrs.Add("title", fmt.Sprintf("El título no puede superar los %d caracteres.", TitleMaxLength))
	}

	if t.dueDateInput != "" && t.DueDate == nil {
		verrs.Add("due_date", "La fecha límite no es válida.")
	}

	return verrs, nil
}

// SetDueDate interpreta el valor del formulario (AAAA-MM-DD) como fecha límite.
// Un valor vacío quita la fecha; uno que no es una fecha la deja vacía y Validate lo rechaza.
func (t *Task) SetDueDate(value string) {
	value = strings.TrimSpace(value)
	t.dueDateInput = value
	t.DueDate = nil
	if value == "" {
		return
	}
	if day, err := time.Parse(DueDateLayout, value); err == nil {
		t.DueDate = &day
	}
}

// DueDateValue devuelve el día de la fecha límite como AAAA-MM-DD, o "" si no tiene.
func (t Task) DueDateValue() string {
	if t.DueDate == nil {
		return ""
	}
	return t.DueDate.Format(DueDateLayout)
}

// DueDateText devuelve el día de la fecha límite como DD/MM/AAAA, o "" si no tiene.
func (t Task) DueDateText() string {
	if t.DueDate == nil {
		return ""
	}
	return t.DueDate.Format("02/01/2006")
}
