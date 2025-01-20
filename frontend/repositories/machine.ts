import type { MachineData, MachineStatus } from '~/lib/types';
import { BACKEND_HOST, STATUS_HOST } from '~/lib/config';

export const getMachines = () =>
  $fetch<MachineData[]>(`${BACKEND_HOST}/machines`);

export const getActiveMachines = () =>
  $fetch<MachineStatus[]>(`${STATUS_HOST}/status`);

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
