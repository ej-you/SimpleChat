import { useNavigate } from 'react-router-dom'
import Auth from './Auth'

const SignUp = () => {
	const nav = useNavigate()
	const handleSubmit = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault()
    nav('/finduser')
  }
	return (
		<Auth handleSubmit={handleSubmit}/>
	)
}

export default SignUp