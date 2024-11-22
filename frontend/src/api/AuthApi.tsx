import { useNavigate } from 'react-router-dom'
import { FieldValues } from 'react-hook-form'
import { useErrorStore } from '../store/store'
import Auth from '../components/auth/Auth'
import { useEffect } from 'react'
import { IAuthApiProps } from '../types/props/types.props'
import axios, { AxiosError } from 'axios'


const AuthApi:React.FC<IAuthApiProps> = ({ apiUrl }) => {
  const nav = useNavigate()
  const setErrorContent = useErrorStore(state => state.setErrorContent)

  useEffect(() => {
    localStorage.removeItem('registered')
    setErrorContent('')
  }, [setErrorContent])

  const onSubmit = async (data: FieldValues) => {
    setErrorContent('')
    try {
      await axios.post(apiUrl, data, {withCredentials: true,})
      localStorage.setItem('registered', data.username)
      nav('/')
    } catch (err) {
      // если истек токен
      if((err as AxiosError).status === 401) {
        setErrorContent((err as AxiosError).message)
        localStorage.removeItem('registered')
        nav('/signup')
      } else{
        console.error(err)
        setErrorContent((err as AxiosError).message)
      }
    }
  }

  return (
    <Auth onSubmit={onSubmit} />
  )
}

export default AuthApi