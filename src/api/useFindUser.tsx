import { useNavigate } from 'react-router-dom'
// import { useErrorStore } from '../store/store'

const useFindUser = () => {
	const nav = useNavigate()
	// const setErrorContent = useErrorStore(store => store.setErrorContent)

	const findUser = async () => {
		nav('/messanger')
		
		// setErrorContent('')
		// try {
		// 	await axios.post('', 'username')
		// 	nav('/messanger')
		// } catch(err) {
		// 	setErrorContent((err as AxiosError).message)
		// }
	}

	return {findUser}
}

export default useFindUser