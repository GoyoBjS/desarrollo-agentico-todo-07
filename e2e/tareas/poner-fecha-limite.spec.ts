import { test, expect } from '../fixtures';

test.describe('Tareas', () => {
  test('poner una fecha límite', async ({ page, tasks }) => {
    await tasks([{ title: 'Pagar el alquiler' }, { title: 'Revisar el presupuesto', dueDate: '2026-11-20' }]);
    await page.reload();

    // 1. La lista muestra la fecha de la tarea que la tiene y nada en la otra
    const alquiler = page.locator('.task', { hasText: 'Pagar el alquiler' });
    const presupuesto = page.locator('.task', { hasText: 'Revisar el presupuesto' });
    await expect(presupuesto.locator('.task-due')).toHaveText('Fecha límite: 20/11/2026');
    await expect(alquiler.locator('.task-due')).toHaveCount(0);

    // 2. Edita «Pagar el alquiler» y le pone una fecha
    await page.getByRole('link', { name: 'Editar tarea: Pagar el alquiler' }).click();
    await page.getByLabel('Fecha límite').fill('2026-11-05');
    await page.getByRole('button', { name: 'Guardar cambios' }).click();

    // 3. Vuelve a la lista con la fecha, en el mismo orden
    await expect(page.getByRole('status')).toHaveText('Tarea actualizada.');
    await expect(alquiler.locator('.task-due')).toHaveText('Fecha límite: 05/11/2026');
    await expect(page.locator('.task-title')).toHaveText(['Pagar el alquiler', 'Revisar el presupuesto']);

    // 4. La fecha sigue tras recargar
    await page.reload();
    await expect(alquiler.locator('.task-due')).toHaveText('Fecha límite: 05/11/2026');
  });
});
