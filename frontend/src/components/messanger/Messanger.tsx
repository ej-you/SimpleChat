import { useEffect, useRef } from 'react'
import { Link, useNavigate } from 'react-router-dom'
import useGetMessages from '../../api/useGetMessages'
import { useChatStore } from '../../store/store'
import Error from '../error/Error'
import Message from './message/Message'
import Footer from './footer/Footer'
import { IChat, IMessage } from '../../types/messanger/types.messanger'
import { io } from 'socket.io-client'
const socket = io('wss://150.241.82.68/api/messanger')

const Messanger = () => {
	const nav = useNavigate()
	const nickname = localStorage.getItem('registered') as string
	const chatRef = useRef<HTMLDivElement>(null)

	useEffect(() => {
		if (!nickname) {
			nav('/signin')
		}
	}, [nav, nickname])

	const { getMessages } = useGetMessages()
	useEffect(() => {
		getMessages()
		// eslint-disable-next-line react-hooks/exhaustive-deps
	}, [])

	// получение сообщений
	useEffect(() => {
		socket.on('receive_message', (newMessage) => {
			// useChatStore.getState().addMessage(newMessage)
			console.log(newMessage)
		})

		return () => {
			socket.off('receive_message')
		}
	}, [])

	const chat = useChatStore(state => state.chatData) as IChat
	const companion = chat && chat.users.filter(el => el.username !== nickname)

	// Скролл чата
	useEffect(() => {
		if (chatRef.current) {
			chatRef.current.scrollTop = chatRef.current.scrollHeight;
    }
	}, [chat])

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

				<main ref={chatRef} className='main flex flex-col flex-grow py-10 gap-4 overflow-y-scroll'>
					{chat &&
					chat.messages.map((el: IMessage, index: number) => (<Message key={index} el={el} />))
					}
				</main>

				<Footer socket={socket} />

			</div>
		</>
	)
}

export default Messanger