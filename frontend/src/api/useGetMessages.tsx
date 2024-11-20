// import axios, { AxiosError } from 'axios'
import { useChatStore, useErrorStore } from '../store/store'
import { useNavigate } from 'react-router-dom'
import axios, { AxiosError } from 'axios'

const useGetMessages = () => {
  const nav = useNavigate()
  const setErrorContent = useErrorStore(state => state.setErrorContent)
  const setChatData = useChatStore(state => state.setChatData)

	const getMessages = async (nickname: string) => {
		setErrorContent('')
    try{
      const res = await axios.get(`https://150.241.82.68/api/chat/get-messages/${nickname}`, {withCredentials: true,})
      setChatData(res.data)
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

	return {getMessages}
}

export default useGetMessages