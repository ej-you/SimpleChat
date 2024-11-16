import { create } from 'zustand'
import defaultChatData from '../test_data/testData'

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

interface IChat {
	chatData: object
}

export const useChatStore = create<IChat>(() => ({
	chatData: defaultChatData
}))