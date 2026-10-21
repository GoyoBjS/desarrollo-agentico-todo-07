# Fecha límite en las tareas

**Issue:** #1 · **Tipo:** feature

## Resumen
Cada tarea puede tener una fecha límite opcional, de un día y sin hora. Se ve en la lista debajo del título y se pone, cambia o quita desde la pantalla de edición.

## Qué se construyó
- Columna `due_date` (tipo `date`, opcional) en `tasks`; las tareas existentes quedan sin fecha.
- La lista muestra «Fecha límite: DD/MM/AAAA» debajo del título, solo si la tarea la tiene.
- La edición tiene un campo «Fecha límite» (`<input type="date">`) con la ayuda «Opcional. Déjala vacía si la tarea no tiene fecha límite.».
- Vaciar el campo quita la fecha; un valor que no es una fecha da `422` con «La fecha límite no es válida.» y no cambia la tarea.
- La ayuda del título pasa a «Hasta 200 caracteres.»: ya no es cierto que los demás datos no cambien.

## Implementación técnica
| Fichero | Papel |
| --- | --- |
| `migrations/20261021000000_add_due_date_to_tasks.{up,down}.fizz` | Añade y quita la columna `due_date` (`date`, nula) |
| `models/task.go` | Campo `DueDate *time.Time`, `SetDueDate`, `DueDateValue` (AAAA-MM-DD), `DueDateText` (DD/MM/AAAA) y la validación |
| `actions/tasks.go` | `TasksUpdate` lee `due_date` y llama a `SetDueDate` |
| `templates/tasks/index.plush.html`, `edit.plush.html` | Fecha en la fila (con `<time>`) y campo de edición con su ayuda y error |
| `public/assets/app.css` | Estilos de `.task-body`, `.task-due` y `input[type="date"]`; la hoja sube a `?v=4` |
| `e2e/fixtures.ts`, `e2e/tareas/poner-fecha-limite.spec.ts` | La fixture admite `dueDate`; test de navegador nuevo |

Cambios que importan:
- `SetDueDate` guarda el texto recibido en un campo interno (`dueDateInput`); `Validate` rechaza un texto no vacío que no se pudo interpretar. Así la validación sigue en el modelo.
- La fecha es un día sin zona horaria: se interpreta y se formatea con `2006-01-02`, sin desfase.
- Crear una tarea no admite fecha (el formulario de alta no tiene el campo); solo la edición la cambia.
- El nombre accesible de los botones de la fila no incluye la fecha.
- Se acepta cualquier día, también pasado.

## Cómo se usa
1. En la lista, pulsa «Editar» en una tarea.
2. Elige un día en «Fecha límite» y pulsa «Guardar cambios».
3. Vuelves a la lista con «Tarea actualizada.» y la fecha bajo el título.
4. Para quitarla, vacía el campo y guarda.

## Configuración
Una migración nueva: `soda migrate` (o arrancar con las migraciones al día). No añade variables ni ajustes.

## Cómo se probó
Acta vigente: [validado, intento 2 de 3](https://github.com/GoyoBjS/desarrollo-agentico-todo-07/issues/1#issuecomment-6320317362). El intento 1 se [rechazó](https://github.com/GoyoBjS/desarrollo-agentico-todo-07/issues/1#issuecomment-6320085034) porque la ayuda del título y la documentación seguían diciendo que los demás datos no cambian.

| Afirmación | Resultado |
| --- | --- |
| La lista muestra la fecha; las tareas sin fecha, incluidas las existentes, no muestran nada | Validado |
| Desde la edición se pone, cambia, conserva (al cambiar solo el título) y quita la fecha | Validado |
| El servidor rechaza `2026-02-30` y `mañana` con `422` sin cambiar la tarea | Validado |
| Solo se pide, guarda y muestra el día; la base guarda `date` sin desfase | Validado |
| Completar y reabrir con JavaScript conserva la fecha y los nombres accesibles | Validado |
| Tests de Go y 6 de 6 de navegador, `gofmt`, `go vet`, migraciones, `go build` y `docker build` | Validado |
| Fidelidad a los criterios, documentación, textos de la interfaz y capas | Validado por jueces |

Para repetirlo (ver [`docs/procedures/development.md`](../../procedures/development.md)):
```bash
go test -p 1 ./...
npx playwright test
npx playwright test e2e/tareas/poner-fecha-limite.spec.ts
```

## Notas
- Fuera de alcance: avisos de tareas vencidas, orden por fecha y hora.
- Pendiente menor del acta: los métodos de formato de la fecha usan receptor por valor y los demás por puntero, y `dueDateInput` lleva una etiqueta `db:"-"` que sobra (`models/task.go`).
- Observaciones del acta: crear una tarea enviando `due_date` en JSON podría fijarla, igual que ya pasa con `completed`; tras una fecha no válida, el campo vuelve vacío en lugar de conservar lo enviado.
- La documentación viva ya está actualizada: `docs/knowledge/features/tareas.md` y `docs/knowledge/overview.md`.
