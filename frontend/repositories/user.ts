import type { UserData, MachineData } from '~/lib/types';
import { BACKEND_HOST } from '~/lib/config';

export type GetUserData = Pick<UserData, 'id' | 'admin' | 'rfid'> & {
  data: Pick<UserData, 'username' | 'email' | 'password'>;
}

// Public API
export const getRfidUser = (id: string) =>
  $fetch<{ userId: number }>(`${BACKEND_HOST}/rfids/${id}/user`);

// User API
export const getUser = (id: number) =>
  $fetch<GetUserData>(`${BACKEND_HOST}/users/${id}`);

export const login = (credentials: { email: string; password: string }) =>
  $fetch<UserData>(`${BACKEND_HOST}/auth/login`, {
    method: 'POST',
    body: credentials,
  });

// Admin API
export const adminLogin = (credentials: { email: string; password: string }) =>
  $fetch<UserData>(`${BACKEND_HOST}/admin/auth/login`, {
    method: 'POST',
    body: credentials,
  });

export const getAllUsers = () =>
  $fetch<GetUserData[]>(`${BACKEND_HOST}/admin/users`);

export const createUser = (user: Omit<UserData, 'id'>) =>
  $fetch<UserData>(`${BACKEND_HOST}/admin/users`, {
    method: 'POST',
    body: user,
  });

export const updateUser = (id: number, user: Partial<UserData>) =>
  $fetch<UserData>(`${BACKEND_HOST}/admin/users/${id}`, {
    method: 'PUT',
    body: user,
  });

export const updateUserRfids = (id: number, rfid: string) =>
  $fetch<UserData>(`${BACKEND_HOST}/admin/users/${id}/rfids/${rfid}`, {
    method: 'PUT',
  });
