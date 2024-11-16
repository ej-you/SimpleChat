import { useEffect } from 'react'
import { Link, useNavigate } from 'react-router-dom'
import { useChatStore } from '../../store/store'
import Message from './message/message'

interface IMessage {
	content: string
	createdAt: string
	sender: object
}

interface IChat {
	messages: IMessage[]
}

const Messanger = () => {
	const nav = useNavigate()
	const chat = useChatStore(state => state.chatData) as IChat
	
	useEffect(() =>{
    if(!localStorage.getItem('registered')){
      nav('/signin')
    }
  },[nav])
	
	return (
		<div className="h-screen flex flex-col py-10 px-60">
    <header className='flex items-center justify-center relative'>
        <Link to='/' className='absolute left-0 text-primary underline cursor-pointer font-bold'>Back</Link>
        <h1 className='text-title text-xl font-bold'>Username</h1>
    </header>
    <div className="main flex flex-col flex-grow py-10 gap-4 overflow-y-scroll">

			{chat.messages.map((el: IMessage, index: number) => <Message key={index} el={el} />)}
			
			{/* <div className="flex items-center gap-4 flex-row-reverse">
				<div className='bg-background-400 max-w-screen-xl break-words flex-wrap py-3.5 px-4 rounded-l-xl rounded-t-xl'>
					<p className='text-primary text-base font-light'>Hello</p>
				</div>
				<p className='text-subtitle-gray text-sm'>20:17</p>
			</div> */}

    </div>
    <footer className='flex flex-col gap-4 background-400'>
			<hr className='w-full border-background-400'/>
			<div className="flex gap-4">
				<input className='w-full text-subtitle-gray placeholder:text-subtitle-gray font-bold bg-background-400 appearance-none py-3 px-4 rounded-xl border-subtitle-gray outline-none' type="text" name="" id="" placeholder='Type here...' />
				<div className="bg-primary rounded-xl flex justify-center h-full aspect-[1/1]"><img className='w-6' src="../../../public/uil_message.svg" alt="" /></div>
			</div>
    </footer>
</div>
	)
}

export default Messanger