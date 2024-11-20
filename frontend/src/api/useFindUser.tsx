import { useNavigate } from 'react-router-dom'
import { useErrorStore } from '../store/store'
// import axios, { AxiosError } from 'axios'

const useFindUser = () => {
	const nav = useNavigate()
	const setErrorContent = useErrorStore(store => store.setErrorContent)

	const findUser = async () => {
		setErrorContent('')
		nav('/messanger')
		// try {
		// 	await axios.get('')
		// 	nav('/messanger')
		// } catch(err) {
		// 	setErrorContent((err as AxiosError).message)
		// }
	}

	return {findUser}
}

export default useFindUser