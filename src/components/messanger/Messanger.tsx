import { Link } from 'react-router-dom'

const Messanger = () => {
	return (
		<div className="container h-screen flex flex-col py-5 px-10">
    <header className='flex items-center justify-center relative'>
        <Link to='/finduser' className='absolute left-0 text-primary underline cursor-pointer text-base font-semibold'>Back</Link>
        <h1 className='text-title text-xl font-semibold'>Igor Romashun</h1>
    </header>
    <div className="main flex-grow py-2">
			<div>
				<p>Message</p>
			</div>
    </div>
    <footer className='flex flex-col gap-3 background-400'>
			<hr className='w-full border-background-400'/>
			<input className='w-full block text-sm text-subtitle-gray font-semibold bg-background-400 appearance-none py-2 px-4 rounded-xl outline-none ' type="text" name="" id="" />
    </footer>
</div>
	)
}

export default Messanger