import { create } from 'zustand'

interface IErrors {
	errorContent: string
	setErrorContent: (newErrorContent: string) => void
}

export const useErrorStore = create<IErrors>(set => ({
	errorContent: '',
	setErrorContent: (newErrorContent) => set(() => ({ errorContent: newErrorContent }))
}))

interface IUserName {
	userName: string
	setUserName: (newUserName: string) => void
}

export const useUserNameStore = create<IUserName>(set => ({
	userName: '',
	setUserName: (newUserName) => set(() => ({ userName: newUserName }))
}))