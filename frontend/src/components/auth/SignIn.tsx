import AuthApi from '../../api/AuthApi'

const SignIn = () => {
	return (
		<AuthApi apiUrl='http://backend:8000/api/user/login'/>
	)
}

export default SignIn