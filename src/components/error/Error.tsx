import { useErrorStore } from '../../store/store'

const Error = () => {
	const errorContent = useErrorStore(state => state.errorContent)
	
	return (
		errorContent &&
		<div className="box absolute left-1/2 -translate-x-1/2 top-2 w-fit bg-title px-2 rounded-lg">
			<p className='font-normal'>{errorContent}</p>
		</div>
	)
}

export default Error