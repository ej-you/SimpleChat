import { create } from 'zustand';
import { IChatStore, ICompanion, IErrors, INotify } from '../types/store/types.store'

export const useErrorStore = create<IErrors>(set => ({
    errorContent: '',
    setErrorContent: (newErrorContent) => set(() => ({ errorContent: newErrorContent })),
}));

export const useNotifyStore = create<INotify>(set => ({
	notifyContent: '',
	setNotifyContent: (newNotifyContent) => set(() => ({ notifyContent: newNotifyContent }))
}))

export const useCompanionStore = create<ICompanion>(set => ({
	companion: '',
	setCompanion: (newCompanion) => set(() => ({ companion: newCompanion }))
}))

export const useChatStore = create<IChatStore>(set => ({
	chatData: null,
	setChatData: (newChatData) => set(() => ({ chatData: newChatData })),
	addMessage: (newMessage) => set(state => ({
		chatData: state.chatData ? {
			...state.chatData,
			messages: [
				...state.chatData.messages,
				{
					content: newMessage.content,
					createdAt: newMessage.createdAt,
					sender: newMessage.sender 
				}
			]
		} 
		: {
			messages: [
				{
					content: newMessage.content,
					createdAt: newMessage.createdAt,
					sender: newMessage.sender
				}
			]
		}
	}))
}))