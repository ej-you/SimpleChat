import { useNavigate } from 'react-router-dom'
import { FieldValues } from 'react-hook-form'
import { useErrorStore } from '../store/store'
import Auth from '../components/auth/Auth'
import { useEffect } from 'react'
import { IAuthApiProps } from '../types/props/types.props'
import axios from 'axios'

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
      const error = (err as { response: { data: { errors: string } } }).response.data.errors
      const errorMessages = Object.values(error)
      setErrorContent(errorMessages[0])
    }
  }

  return (
    <Auth onSubmit={onSubmit} />
  )
}

export default AuthApi