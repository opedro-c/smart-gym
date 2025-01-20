import { z } from "zod"

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


export const WorkoutDataSchema = z.object({
    started_at: z.coerce.date(),
    finished_at: z.coerce.date(),
    origin_id: z.string(),
    user_id: z.number(),
    data: z.object({
        weight: z.number(),
    })
})

export type WorkoutData = z.infer<typeof WorkoutDataSchema>;