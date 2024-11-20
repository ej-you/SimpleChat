// messanger
export interface IMessage {
	content: string
	createdAt: string
	sender: { username: string }
}

interface IUsers {
	id: string
	username: string
}

export interface IChat {
	messages: IMessage[]
	users: IUsers[]
}