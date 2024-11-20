import { useChatStore, useErrorStore } from '../store/store'
import chatData from '../test_data/testData'

const useGetMessages = () => {
  const setErrorContent = useErrorStore(state => state.setErrorContent)
  const setChatData = useChatStore(state => state.setChatData)

	const getMessages = async () => {
		setErrorContent('')
    setChatData(chatData)
  }

	return {getMessages}
}

export default useGetMessages