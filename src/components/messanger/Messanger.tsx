import { useEffect } from 'react'
import { Link, useNavigate } from 'react-router-dom'
import useGetMessages from '../../api/useGetMessages'
import { useChatStore } from '../../store/store'
import Error from '../error/Error'
import Message from './message/Message'
import TextField from './TextField'

interface IMessage {
	content: string
	createdAt: string
	sender: { username: string }
}

interface IUsers {
	id: string
	username: string
}

interface IChat {
	messages: IMessage[]
	users: IUsers[]
}

const Messanger = () => {
	const nav = useNavigate()

	useEffect(() => {
		if (!localStorage.getItem('registered')) {
			nav('/signin')
		}
	}, [nav])

	const { getMessages } = useGetMessages()
	getMessages('user3')

	const registeredUser = localStorage.getItem('registered')
	const chat = useChatStore(state => state.chatData) as IChat
	const companion = chat && chat.users.filter(el => el.username !== registeredUser)

	return (
		<>
			<Error />
			<div className='h-screen flex flex-col py-10 px-60'>

				<header className='flex items-center justify-center relative'>
					<Link to='/' className='absolute left-0 text-primary underline cursor-pointer font-bold'>Back</Link>
					{chat && (
					<h1 className='text-title text-xl font-bold'>
						{companion[0] ? companion[0].username : 'Собеседник'}
					</h1>
					)}
				</header>

				<main className='main flex flex-col flex-grow py-10 gap-4 overflow-y-scroll'>
					{chat &&
					chat.messages.map((el: IMessage, index: number) => (<Message key={index} el={el} />))
					}
				</main>

				<footer className='flex flex-col gap-4 background-400'>
					<hr className='w-full border-background-400' />
					<div className='flex gap-4'>
						<TextField />
						<div className='bg-primary rounded-xl flex justify-center px-3 h-fit aspect-[1/1]'>
							<img className='w-6' src='../../../public/uil_message.svg' alt='' />
						</div>
					</div>
				</footer>

			</div>
		</>
	)
}

export default Messanger