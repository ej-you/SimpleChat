import { create } from 'zustand';

interface IErrors {
    errorContent: string;
    setErrorContent: (newErrorContent: string) => void;
}

export const useErrorStore = create<IErrors>(set => ({
    errorContent: '',
    setErrorContent: (newErrorContent) => set(() => ({ errorContent: newErrorContent })),
}));

interface IMessage {
	content: string
	createdAt?: string
	sender?: { id?: string; username?: string }
}

interface IChatData {
	messages: IMessage[];
	id?: string;
	users?: object[];
}

interface IChatStore {
	chatData: IChatData | null;
	setChatData: (newChatData: IChatData) => void;
	addMessage: (newMessage: IMessage) => void;
}

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
					createdAt: new Date().toISOString(),
					sender: { 
						username: localStorage.getItem('registered') || ''
					}
				}
			]
		} 
		: {
			messages: [
				{
				content: newMessage.content,
				createdAt: new Date().toISOString(),
				sender: { 
					username: localStorage.getItem('registered') || ''
				}
			}],
		}
	}))
}));