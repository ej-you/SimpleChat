import { useNavigate } from 'react-router-dom'
import Auth from './Auth'

const SignIn = () => {
  const nav = useNavigate()

  const handleSubmit = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault()
    localStorage.setItem('token', 'Sfef32DesfF3t')
    nav('/')
  }

	return (
		<Auth handleSubmit={handleSubmit}/>
	)
}

export default SignIn