import { useNavigate } from 'react-router-dom'
import Auth from './Auth'
import { FieldValues } from 'react-hook-form'
import axios from 'axios'
import { useState } from 'react'

const SignIn = () => {
  const nav = useNavigate()
  const [errName, setErrName] = useState<string>('')

  const onSubmit = async (data: FieldValues) => {
    setErrName('')
		try {
			const res = await axios.post('http://150.241.82.68/api/user/login', data)
			localStorage.setItem('registered', 'true')
			console.log(res)
			nav('/')
		} catch(err) {
			console.error(err)
      setErrName((err as Error).message)
		}
  }

	return (
		<Auth onSubmit={onSubmit} errName={errName}/>
	)
}

export default SignIn