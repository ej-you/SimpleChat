import Auth from './Auth'

const SignIn = () => {
  const handleSubmit = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault()
    console.log('sing in')
  }
	return (
		<Auth handleSubmit={handleSubmit}/>
	)
}

export default SignIn