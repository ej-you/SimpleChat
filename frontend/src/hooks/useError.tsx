import { AxiosError } from 'axios'
import { useNavigate } from 'react-router-dom'
import { useErrorStore } from '../store/store'

const UseError = () => {
	const nav = useNavigate()
  const setErrorContent = useErrorStore(state => state.setErrorContent)

	const handleError = (err: unknown) => {
		const error = (err as { response: { data: { errors: string } } }).response.data.errors
		const errorMessages = Object.values(error)
		setErrorContent(errorMessages[0])
		
		// если истек токен
		if((err as AxiosError).status === 401){
			setTimeout(() => {
				localStorage.removeItem('registered')
				nav('/signin')
			}, 1000)
		}
	}

	return {handleError}

}

export default UseError