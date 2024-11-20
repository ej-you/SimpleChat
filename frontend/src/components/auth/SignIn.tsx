import AuthApi from '../../api/AuthApi'

const SignIn = () => {
	return (
		<AuthApi apiUrl='http://150.241.82.68/api/user/login'/>
	)
}

export default SignIn