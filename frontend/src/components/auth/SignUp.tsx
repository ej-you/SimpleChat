import AuthApi from '../../api/AuthApi'

const SignUp = () => {
  return (
    <AuthApi apiUrl='https://fredcv.ru:8091/api/user/register'/>
  )
}

export default SignUp