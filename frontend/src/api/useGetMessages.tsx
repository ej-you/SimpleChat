// import axios, { AxiosError } from 'axios'
import { useChatStore, useErrorStore } from '../store/store'
import { useNavigate, useParams } from 'react-router-dom'
import axios, { AxiosError } from 'axios'

const useGetMessages = () => {
  const nav = useNavigate()
  const {id} = useParams()
  const setErrorContent = useErrorStore(state => state.setErrorContent)
  const setChatData = useChatStore(state => state.setChatData)

	const getMessages = async () => {
		setErrorContent('')
    try{
      const res = await axios.get(`https://150.241.82.68/api/chat/${id}`, {withCredentials: true,})
      setChatData(res.data)
    } catch(err) {
      console.error(err)
      setErrorContent((err as AxiosError).message)
      // если истек токен
      if((err as AxiosError).status === 401){
        setTimeout(() => {
					localStorage.removeItem('registered')
					nav('/signin')
				}, 1000);
      }
    }
  }

	return {getMessages}
}

export default useGetMessages