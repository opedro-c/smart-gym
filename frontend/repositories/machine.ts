import type { MachineData } from '~/lib/types';
import { BACKEND_HOST } from '~/lib/config';

export const getMachines = () =>
  $fetch<MachineData[]>(`${BACKEND_HOST}/machines`);

export const createMachine = (machine: MachineData) =>
  $fetch<MachineData>(`${BACKEND_HOST}/admin/machines`, {
    method: 'POST',
    body: machine,
  });

export const updateMachine = (id: number, machine: Partial<MachineData>) =>
  $fetch<MachineData>(`${BACKEND_HOST}/admin/machines/${id}`, {
    method: 'PUT',
    body: machine,
  });

export const deleteMachine = (id: string) =>
  $fetch<void>(`${BACKEND_HOST}/admin/machines/${id}`, {
    method: 'DELETE',
  });
