import { WORKOUT_HOST } from '~/lib/config';
import type { WorkoutData } from '~/lib/types';


export const getWorkoutData = (user_id:number, origin_id: string, started_at: number, finished_at: number) =>
  $fetch<WorkoutData[]>(`${WORKOUT_HOST}/api/v1/users/${user_id}/origins/${origin_id}/exercises?started_at=${started_at}&finished_at=${finished_at}`);
