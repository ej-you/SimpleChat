import AuthApi from '../../api/AuthApi'

const SignIn = () => {
	return (
		<AuthApi apiUrl='https://web-server/api/user/login'/>
	)
}

export default SignIn