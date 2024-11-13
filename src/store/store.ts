import { create } from 'zustand'

interface IErrors {
	errorContent: string
	setErrorContent: (newErrorContent: string) => void
}

export const useErrorStore = create<IErrors>(set => ({
	errorContent: '',
	setErrorContent: (newErrorContent) => set(() => ({ errorContent: newErrorContent }))
}))