import { useNavigate } from 'react-router-dom'
import { useErrorStore, useUserNameStore } from '../store/store'
import axios, { AxiosError } from 'axios'

const useFindUser = () => {
	const nav = useNavigate();
  const setUserName = useUserNameStore(state => state.setUserName);
  const setErrorContent = useErrorStore(state => state.setErrorContent);

	// eslint-disable-next-line @typescript-eslint/no-explicit-any
	const findUserApi = async (data: any) => {
		setErrorContent('')
    try{
      const res = await axios.get(`http://150.241.82.68/api/chat/get-messages/${data}`, {withCredentials: true,})
      console.log(res.data)
      setUserName(data.findUserByName)
      nav('/messanger')
    } catch(err) {
      console.error(err)
      if((err as AxiosError).status === 401){
        localStorage.removeItem('registered')
        nav('/signup')
      } else{
        setErrorContent((err as AxiosError).message)
      }
    }
  }

	return {findUserApi}
}

export default useFindUser