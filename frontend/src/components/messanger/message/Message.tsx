import React from 'react'
import { IMessageProps } from '../../../types/props/types.props'

const Message:React.FC<IMessageProps> = ({el}) => {
	const dateString = el.createdAt
	const dateObject = new Date(dateString)
	const currentUser = localStorage.getItem('registered')

	// Получение часов, минут и секунд
	const hours = dateObject.getHours()
	const minutes = dateObject.getMinutes()
	const seconds = dateObject.getSeconds()

	// Формат времени в строку
	const timeString = `${String(hours).padStart(2, '0')}:${String(minutes).padStart(2, '0')}:${String(seconds).padStart(2, '0')}`;

	return (
		<>
		{/* <div className={`flex items-center gap-4 ${el.sender.username === localStorage.getItem('registered') && 'flex-row-reverse'}`}>
			<div className={`bg-background-400 max-w-screen-xl  py-3.5 px-4 ${el.sender.username === localStorage.getItem('registered') ? 'rounded-l-xl' : 'rounded-r-xl'} rounded-t-xl`}>
				<p className={`${el.sender.username === localStorage.getItem('registered') ? 'text-primary' : 'text-white'} text-base font-light text`}>
						{el.content}
				</p>
			</div>
		<p className='text-subtitle-gray text-sm text-right'>{timeString}</p>
		</div> */}
		<div className={`flex items-center gap-4 ${el.sender.username === currentUser && 'flex-row-reverse'}`}>
			<div className={`bg-background-400 width py-3.5 px-4 flex flex-col ${el.sender.username === localStorage.getItem('registered') ? 'rounded-l-xl' : 'rounded-r-xl'} rounded-t-xl`}>
				<p className={`${el.sender.username === currentUser ? 'text-primary' : 'text-white'} text-base font-light text relative`}>
						{el.content}
				</p>
			<p className='text-subtitle-gray text-sm text-right'>{timeString}</p>
			</div>
		</div>
		</>
	)
}

export default Message