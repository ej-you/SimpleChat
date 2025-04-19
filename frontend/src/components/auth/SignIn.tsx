import AuthApi from '../../api/AuthApi'

const SignIn = () => {
	return (
		<AuthApi apiUrl='https://fredcv.ru:8091/simple-chat/api/user/login'/>
	)
}

export default SignIn