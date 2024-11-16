import { create } from 'zustand'

interface IErrors {
	errorContent: string
	setErrorContent: (newErrorContent: string) => void
}

export const useErrorStore = create<IErrors>(set => ({
	errorContent: '',
	setErrorContent: (newErrorContent) => set(() => ({ errorContent: newErrorContent }))
}))

interface IChat {
	chatData: object | null
	setChatData: (newChatData:object) => void 
}

export const useChatStore = create<IChat>(set => ({
	chatData: null,
	setChatData: (newChatData) => set(() => ({ chatData: newChatData}))
}))