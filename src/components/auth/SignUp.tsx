import { useNavigate } from 'react-router-dom'
import Auth from './Auth'
import axios from 'axios'
import { FieldValues } from 'react-hook-form'

const SignUp = () => {
	const nav = useNavigate()

	// const onSubmit = (data: FieldValues): void => {
	// 	axios.post('http://150.241.82.68/user/register', data)
	// 	.then(res => {
	// 		localStorage.setItem('registered', '')
	// 		console.log(res)
	// 		nav('/')
	// 	})
	// 	.catch(err => console.error(err))
  // }
	
	const onSubmit = async (data: FieldValues) => {
		try {
			const res = await axios.post('http://150.241.82.68/user/register', data)
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

export default SignUp