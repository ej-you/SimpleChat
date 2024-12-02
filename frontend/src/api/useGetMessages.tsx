// import axios, { AxiosError } from 'axios'
import { useChatStore, useErrorStore } from '../store/store'
import { useParams } from 'react-router-dom'
import axios from 'axios'
import UseError from '../hooks/useError'

const useGetMessages = () => {
  const {handleError} = UseError()
  const {id} = useParams()
  const setErrorContent = useErrorStore(state => state.setErrorContent)
  const setChatData = useChatStore(state => state.setChatData)

	const getMessages = async () => {
    setErrorContent('')
    try{
      const res = await axios.get(`https://150.241.82.68/api/chat/${id}`, {withCredentials: true,})
      setChatData(res.data)
    } catch(err) {
      handleError(err)
    }
  }

	return {getMessages}
}

export default useGetMessages