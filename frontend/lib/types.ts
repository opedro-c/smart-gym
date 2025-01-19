export type UserData = {
    id: number
    admin: boolean
    username: string,
    email: string
    password: string
    rfid: string
}

export type MachineData = {
    id: number,
    name: string,
    origin_id: string
}

export type MachineStatus = {
    origin_id: string
    status: boolean
}