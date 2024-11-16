import { useNavigate } from 'react-router-dom'
import { FieldValues } from 'react-hook-form'
import axios, { AxiosError } from 'axios'
import { useErrorStore } from '../store/store'
import Auth from '../components/auth/Auth'
import { useEffect } from 'react'

interface IProps {
	apiUrl: string
}

const AuthApi:React.FC<IProps> = ({ apiUrl }) => {
  const nav = useNavigate()
  const setErrorContent = useErrorStore(state => state.setErrorContent)

  useEffect(() => {
    localStorage.removeItem('registered')
    setErrorContent('')
  }, [setErrorContent])

  const onSubmit = async (data: FieldValues) => {
    setErrorContent('')
    try {
      const res = await axios.post(apiUrl, data)
      localStorage.setItem('registered', data.username)
      console.log(res)
      nav('/')
    } catch (err) {
      console.error(err)  
      setErrorContent((err as AxiosError).message)
    }
  }

  return (
    <Auth onSubmit={onSubmit} />
  )
}

export default AuthApi