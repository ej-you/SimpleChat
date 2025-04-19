import AuthApi from '../../api/AuthApi'

const SignUp = () => {
  return (
    <AuthApi apiUrl='http://backend:8000/api/user/register'/>
  )
}

export default SignUp