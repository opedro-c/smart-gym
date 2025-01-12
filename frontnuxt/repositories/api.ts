import type { UserData, MachineData } from '~/types';

const backendHost = process.env.BACKEND_HOST || 'http://localhost:8080';

export type GetUserData = Pick<UserData, 'id' | 'admin' | 'rfid'> & {
  data: Pick<UserData, 'username' | 'email' | 'password'>;
}

// Public API
export const getRfidUser = (id: string) =>
  $fetch<{ userId: number }>(`${backendHost}/rfids/${id}/user`);

// User API
export const getUser = (id: number) =>
  $fetch<GetUserData>(`${backendHost}/users/${id}`);

export const getMachines = () =>
  $fetch<MachineData[]>(`${backendHost}/machines`);

export const login = (credentials: { email: string; password: string }) =>
  $fetch<UserData>(`${backendHost}/auth/login`, {
    method: 'POST',
    body: credentials,
  });

// Admin API
export const adminLogin = (credentials: { email: string; password: string }) =>
  $fetch<{ token: string }>(`${backendHost}/admin/auth/login`, {
    method: 'POST',
    body: credentials,
  });

export const getAllUsers = () =>
  $fetch<GetUserData[]>(`${backendHost}/admin/users`);

export const createUser = (user: Omit<UserData, 'id'>) =>
  $fetch<UserData>(`${backendHost}/admin/users`, {
    method: 'POST',
    body: user,
  });

export const updateUser = (id: number, user: Partial<UserData>) =>
  $fetch<UserData>(`${backendHost}/admin/users/${id}`, {
    method: 'PUT',
    body: user,
  });

export const updateUserRfids = (id: string, rfid: string) =>
  $fetch<UserData>(`${backendHost}/admin/users/${id}/rfids/${rfid}`, {
    method: 'PUT',
  });

export const createMachine = (machine: MachineData) =>
  $fetch<MachineData>(`${backendHost}/admin/machines`, {
    method: 'POST',
    body: machine,
  });

export const updateMachine = (id: number, machine: Partial<MachineData>) =>
  $fetch<MachineData>(`${backendHost}/admin/machines/${id}`, {
    method: 'PUT',
    body: machine,
  });

export const deleteMachine = (id: string) =>
  $fetch<void>(`${backendHost}/admin/machines/${id}`, {
    method: 'DELETE',
  });

export default {
  getRfidUser,
  getUser,
  getMachines,
  login,
  adminLogin,
  getAllUsers,
  createUser,
  updateUser,
  updateUserRfids,
  createMachine,
  updateMachine,
  deleteMachine,
};
