import Auth from './Auth'

const SignUp = () => {
	const handleSubmit = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault()
    console.log('sign up')
  }
	return (
		<Auth handleSubmit={handleSubmit}/>
	)
}

export default SignUp