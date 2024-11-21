import { useNavigate } from 'react-router-dom'
import { FieldValues } from 'react-hook-form'
import { useErrorStore } from '../store/store'
import Auth from '../components/auth/Auth'
import { useEffect } from 'react'
import { IAuthApiProps } from '../types/props/types.props'


const AuthApi:React.FC<IAuthApiProps> = ({ apiUrl }) => {
  console.log(apiUrl)
  const nav = useNavigate()
  const setErrorContent = useErrorStore(state => state.setErrorContent)

  useEffect(() => {
    localStorage.removeItem('registered')
    setErrorContent('')
  }, [setErrorContent])

  const onSubmit = async (data: FieldValues) => {
    setErrorContent('')
    localStorage.setItem('registered', data.username)
    nav('/')
  }

  return (
    <Auth onSubmit={onSubmit} />
  )
}

export default AuthApi