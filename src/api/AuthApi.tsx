import { useNavigate } from 'react-router-dom'
import { FieldValues } from 'react-hook-form'
import axios from 'axios'
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
    setErrorContent('')
  }, [setErrorContent])

  const onSubmit = async (data: FieldValues) => {
    setErrorContent('')
    try {
      const res = await axios.post(apiUrl, data)
      localStorage.setItem('registered', 'true')
      console.log(res)
      nav('/')
    } catch (err) {
      console.error(err)
      setErrorContent((err as Error).message)
    }
  }

  return (
    <Auth onSubmit={onSubmit} />
  )
}

export default AuthApi