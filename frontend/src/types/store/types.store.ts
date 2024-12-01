// useErrorStore
export interface IErrors {
	errorContent: string;
	setErrorContent: (newErrorContent: string) => void;
}

// useNotifyStore
export interface INotify {
	notifyContent: string;
	setNotifyContent: (newNotifyContent: string) => void;
}

// UseCompanionStore
export interface ICompanion {
	companion: string
	setCompanion: (newCompanion: string) => void
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