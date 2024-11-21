import { useNavigate } from 'react-router-dom'
import { useCompanionStore, useErrorStore } from '../store/store'
import axios, { AxiosError } from 'axios'
import { FieldValues } from 'react-hook-form'

const useFindUser = () => {
	const nav = useNavigate()
	const setErrorContent = useErrorStore(store => store.setErrorContent)
	const setCompanion = useCompanionStore(store => store.setCompanion)

	const findUser = async (data: FieldValues) => {
		setErrorContent('')
		try {
			await axios.get(`https://150.241.82.68/api/user/check/${data.findUserByName}`)
			setCompanion(data.findUserByName)
			nav('/messanger')
		} catch(err) {
			// если истек токен
      if((err as AxiosError).status === 401){
        localStorage.removeItem('registered')
        nav('/signup')
        setErrorContent((err as AxiosError).message)
      } else{
        console.error(err)
        setErrorContent((err as AxiosError).message)
      }
		}
	}

	return {findUser}
}

export default useFindUser