import { useNavigate } from 'react-router-dom'
import Auth from './Auth'
import { FieldValues } from 'react-hook-form'
import axios from 'axios'

const SignIn = () => {
  const nav = useNavigate()

  const onSubmit = async (data: FieldValues) => {
		try {
			const res = await axios.post('http://150.241.82.68/user/login', data)
			localStorage.setItem('registered', '')
			console.log(res)
			nav('/')
		} catch(err) {
			console.error(err)
		}
  }

	return (
		<Auth onSubmit={onSubmit}/>
	)
}

export default SignIn