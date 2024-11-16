import { useChatStore } from '../store/store'
import chatData from '../test_data/testData'

const useGetMessages = () => {
  // const setErrorContent = useErrorStore(state => state.setErrorContent);
  const setChatData = useChatStore(state => state.setChatData)

	// eslint-disable-next-line @typescript-eslint/no-explicit-any
	const getMessages = async (data: any) => {
    setChatData(chatData)
    console.log(data)

		// setErrorContent('')
    // try{
    //   const res = await axios.get(`http://150.241.82.68/api/chat/get-messages/${data}`, {withCredentials: true,})
    //   console.log(res.data)
    //   setUserName(data.findUserByName)
    //   setErrorContent('')
    //   nav('/messanger')
    // } catch(err) {
    //   console.error(err)
    //   if((err as AxiosError).status === 401){
    //     localStorage.removeItem('registered')
    //     nav('/signup')
    //   } else{
    //     setErrorContent((err as AxiosError).message)
    //   }
    // }

  }

	return {getMessages}
}

export default useGetMessages