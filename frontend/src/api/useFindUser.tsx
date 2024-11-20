import { useNavigate } from 'react-router-dom'
import { useErrorStore } from '../store/store'
import axios, { AxiosError } from 'axios'
import { FieldValues } from 'react-hook-form'

const useFindUser = () => {
	const nav = useNavigate()
	const setErrorContent = useErrorStore(store => store.setErrorContent)

	const findUser = async (data: FieldValues) => {
		setErrorContent('')
		try {
			await axios.get(`150.241.82.68/api/user/check/${data.findUserByName}`)
			nav('/messanger')
		} catch(err) {
			setErrorContent((err as AxiosError).message)
		}
	}

	return {findUser}
}

export default useFindUser