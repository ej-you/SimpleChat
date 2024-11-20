import { useNavigate } from 'react-router-dom'
import { useErrorStore } from '../store/store'

const useFindUser = () => {
	const nav = useNavigate()
	const setErrorContent = useErrorStore(store => store.setErrorContent)

	const findUser = async () => {
		setErrorContent('')
		nav('/messanger')
	}

	return {findUser}
}

export default useFindUser