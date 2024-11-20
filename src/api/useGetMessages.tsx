// import axios, { AxiosError } from 'axios'
import { useChatStore, useErrorStore } from '../store/store'
import chatData from '../test_data/testData'
// import { useNavigate } from 'react-router-dom'

const useGetMessages = () => {
  // const nav = useNavigate()
  const setErrorContent = useErrorStore(state => state.setErrorContent)
  const setChatData = useChatStore(state => state.setChatData)

	// eslint-disable-next-line @typescript-eslint/no-explicit-any
	const getMessages = async (data: any) => {
    setChatData(chatData)
    console.log(data)
		setErrorContent('')

    // try{
    //   const res = await axios.get(`http://150.241.82.68/api/chat/get-messages/${data}`, {withCredentials: true,})
    //   console.log(res.data)
    //   nav('/messanger')
    // } catch(err) {
    //   if((err as AxiosError).status === 401){
    //     // localStorage.removeItem('registered')
    //     // nav('/signup')
    //     setErrorContent((err as AxiosError).message)
    //   } else{
    //     console.error(err)
    //     setErrorContent((err as AxiosError).message)
    //   }
    // }
  }

	return {getMessages}
}

export default useGetMessages