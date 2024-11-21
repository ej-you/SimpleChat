import React from 'react'
import { IMessageProps } from '../../../types/props/types.props'

const Message:React.FC<IMessageProps> = ({el}) => {
	return (
		<>
		<div className={`flex items-center gap-4 ${el.sender.username === localStorage.getItem('registered') && 'flex-row-reverse'}`}>
			<div className={`bg-background-400 max-w-screen-xl break-all flex-wrap py-3.5 px-4 ${el.sender.username === localStorage.getItem('registered') ? 'rounded-l-xl' : 'rounded-r-xl'} rounded-t-xl`}>
				<p className={`${el.sender.username === localStorage.getItem('registered') ? 'text-primary' : 'text-white'} text-base font-light`}>{el.content}</p>
			</div>
			<p className='text-subtitle-gray text-sm text-right'>{el.createdAt}</p>
		</div>
		</>
	)
}

export default Message