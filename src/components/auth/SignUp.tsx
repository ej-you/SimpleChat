import { useNavigate } from 'react-router-dom'
import Auth from './Auth'
import axios from 'axios'
import { FieldValues } from 'react-hook-form'
import {useErrorStore} from '../../store/store'

const SignUp = () => {
	const nav = useNavigate()
	const setErrorContent = useErrorStore(state => state.setErrorContent)

	const onSubmit = async (data: FieldValues) => {
		setErrorContent('')
		try {
			const res = await axios.post('http://150.241.82.68/api/user/register', data)
			localStorage.setItem('registered', 'true')
			console.log(res)
			nav('/')
		} catch(err) {
			console.error(err)
			setErrorContent((err as Error).message)
		}
  }

	return (
		<Auth onSubmit={onSubmit} />
	)
}	

export default SignUp