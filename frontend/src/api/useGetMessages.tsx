// import axios, { AxiosError } from 'axios'
import { FieldValues } from 'react-hook-form'
import { useChatStore, useErrorStore } from '../store/store'
import { useNavigate } from 'react-router-dom'
import axios, { AxiosError } from 'axios'

const useGetMessages = () => {
  const nav = useNavigate()
  const setErrorContent = useErrorStore(state => state.setErrorContent)
  const setChatData = useChatStore(state => state.setChatData)

	const getMessages = async (data: FieldValues) => {
		setErrorContent('')
    try{
      const res = await axios.get(`http://150.241.82.68/api/chat/get-messages/${data}`, {withCredentials: true,})
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