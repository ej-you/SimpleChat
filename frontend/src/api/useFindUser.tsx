import { useNavigate } from 'react-router-dom'
import { useErrorStore } from '../store/store'
import axios from 'axios'
import { FieldValues } from 'react-hook-form'
import UseError from '../hooks/useError'

const useFindUser = () => {
	const nav = useNavigate()
	const setErrorContent = useErrorStore(store => store.setErrorContent)
	const {handleError} = UseError()

	const findUser = async (data: FieldValues) => {
		setErrorContent('')
		try {
			const res = await axios.get(`https://web-server/api/chat/with/${data.findUserByName}`, {withCredentials: true,})
			nav(`/messanger/${res.data.id}`)
		} catch(err) {
			handleError(err)
		}
	}

	return {findUser}
}

export default useFindUser