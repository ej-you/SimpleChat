// useErrorStore
export interface IErrors {
	errorContent: string;
	setErrorContent: (newErrorContent: string) => void;
}

// useChatStore
interface IMessage {
	content: string
	createdAt: string
	sender: { id?: string; username: string }
}

interface IChatData {
	id?: string
	messages: IMessage[]
	users?: object[]
}

export interface IChatStore {
	chatData: IChatData | null
	setChatData: (newChatData: IChatData) => void
	addMessage: (newMessage: IMessage) => void
}