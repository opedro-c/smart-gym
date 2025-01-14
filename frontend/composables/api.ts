import * as MachineRepository from '~/repositories/machine'
import * as UserRepository from '~/repositories/user'

export const useApi = () => {
  return {
    ...MachineRepository,
    ...UserRepository
  }
}