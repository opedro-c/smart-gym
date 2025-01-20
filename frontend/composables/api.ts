import * as MachineRepository from '~/repositories/machine'
import * as UserRepository from '~/repositories/user'
import * as WorkoutRepository from '~/repositories/workout'

export const useApi = () => {
  return {
    ...MachineRepository,
    ...UserRepository,
    ...WorkoutRepository
  }
}