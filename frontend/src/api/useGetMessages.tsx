// import axios, { AxiosError } from 'axios'
import { useChatStore, useCompanionStore, useErrorStore } from '../store/store'
import { useNavigate } from 'react-router-dom'
import axios, { AxiosError } from 'axios'

const useGetMessages = () => {
  const nav = useNavigate()
  const setErrorContent = useErrorStore(state => state.setErrorContent)
  const setChatData = useChatStore(state => state.setChatData)
  const companion = useCompanionStore(state => state.companion)

	const getMessages = async () => {
		setErrorContent('')
    try{
      const res = await axios.get(`https://150.241.82.68/api/chat/get-messages/${companion}`, {withCredentials: true,})
      setChatData(res.data)
    } catch(err) {
      // если истек токен
      if((err as AxiosError).status === 401){
        setErrorContent((err as AxiosError).message)
        setTimeout(() => {
					localStorage.removeItem('registered')
					nav('/signup')
				}, 1000);
      } else{
        console.error(err)
        setErrorContent((err as AxiosError).message)
      }
    }
  }

	return {getMessages}
}

export default useGetMessages