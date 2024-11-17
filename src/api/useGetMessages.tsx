import { useChatStore } from '../store/store'
import chatData from '../test_data/testData'

const useGetMessages = () => {
  // const setErrorContent = useErrorStore(state => state.setErrorContent);
  const setChatData = useChatStore(state => state.setChatData)

	// eslint-disable-next-line @typescript-eslint/no-explicit-any, @typescript-eslint/no-unused-vars
	const getMessages = async (_data: any) => {
    setChatData(chatData)

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