import React from 'react'

interface IProps {
	el: {content: string, createdAt: string, sender: object}
}

const Message:React.FC<IProps> = ({el}) => {
	return (
		<div className="flex items-center gap-4">
			<div className='bg-background-400 max-w-screen-xl break-words flex-wrap py-3.5 px-4 rounded-r-xl rounded-t-xl'>
				<p className='text-white text-base font-light'>{el.content}</p>
			</div>
			<p className='text-subtitle-gray text-sm'>{el.createdAt}</p>
		</div>
	)
}

export default Message