import { useNotifyStore } from '../../store/store'

const Error = () => {
	const notifyContent = useNotifyStore(state => state.notifyContent)
	
	return (
		notifyContent &&
		<div className="box absolute left-1/2 -translate-x-1/2 top-2 w-fit bg-primary px-2 rounded-lg text-center z-10">
			<p className='font-normal'>{notifyContent}</p>
		</div>
	)
}

export default Error