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
			const res = await axios.get(`https://150.241.82.68/api/chat/with/${data.findUserByName}`, {withCredentials: true,})
			nav(`/messanger/${res.data.id}`)
		} catch(err) {
			// если истек токен
      if((err as AxiosError).status === 401){
        setErrorContent((err as AxiosError).message)
				setTimeout(() => {
					localStorage.removeItem('registered')
					nav('/signin')
				}, 1000)
      } else{
        console.error(err)
        setErrorContent((err as AxiosError).message)
      }
		}
	}

	return {findUser}
}

export default useFindUser