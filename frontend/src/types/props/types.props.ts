import { FieldValues } from 'react-hook-form'

// Auth
export interface IAuthProps {
  onSubmit: (data: FieldValues) => void
}

// AuthApi
export interface IAuthApiProps {
	apiUrl: string
}

// Message
export interface IMessageProps {
	el: {content: string, createdAt: string, sender: {username: string} }
}