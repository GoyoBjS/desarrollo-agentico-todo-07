# Cambios documentados

Cada caso que cierra el workflow `adw` deja aquí su documento de cierre: qué se construyó, cómo se usa, cómo se probó y qué conviene saber para trabajar sobre ello. Lo escribe la skill `document` en la rama del caso, y viaja en el PR con el código.

Este índice es condicional: cada entrada dice cuándo conviene cargar su documento. Se llega a él desde [`docs/README.md`](../../README.md): quien planifica, como las skills `feature`, `bug` y `chore`, lee solo los documentos cuyas condiciones encajan con el encargo.

Un documento de cierre cuenta qué pasó en su caso, y no se reescribe cuando otro caso cambia el mismo código. Lo vigente está en `docs/knowledge/` y `docs/procedures/`.

## Formato de las entradas

```markdown
- [<tipo>-<N>-<slug>.md](<tipo>-<N>-<slug>.md): <qué cambió, en una línea>
  - Cuando se trabaje con <rutas, acciones o plantillas concretas>
  - Cuando se modifique <modelo, campo o regla relacionados>
  - Cuando se investigue <un problema concreto de esa área>
```

Entre tres y cinco condiciones por documento, y las entradas en orden por número de issue. Si un documento ya tiene entrada, se actualizan sus condiciones en lugar de añadir otra.

## Índice

- [feature-1-fecha-limite-en-las-tareas.md](feature-1-fecha-limite-en-las-tareas.md): las tareas tienen una fecha límite opcional, solo día, visible en la lista y editable
  - Cuando se trabaje con `models/task.go`, la columna `tasks.due_date` o sus migraciones
  - Cuando se modifique `TasksUpdate` en `actions/tasks.go` o el formulario de `templates/tasks/edit.plush.html`
  - Cuando se cambie cómo se muestra una tarea en `templates/tasks/index.plush.html` o sus estilos (`.task-body`, `.task-due`)
  - Cuando se añadan datos a las tareas o se ordenen, filtren o avisen por fecha
  - Cuando se investigue una fecha límite que se desplaza un día o se rechaza con «La fecha límite no es válida.»
