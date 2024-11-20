import { Link } from 'react-router-dom'
import { useForm } from 'react-hook-form'
import Error from '../error/Error'
import { IAuthProps } from '../../types/props/types.props'

const Auth: React.FC<IAuthProps> = ({onSubmit}) => {
  const {register, handleSubmit, formState: { errors }} = useForm()

	return (
    <>
    <Error />
		<div className="flex flex-col justify-center h-screen items-center text-center gap-12">
      <div className="flex flex-col gap-2">
        <h1 className='text-title text-xl font-bold'>Messanger</h1>

        <h2 className='text-subtitle-gray text-base font-bold'>
					{location.pathname === '/signin' ? 
					<>
					Sign in or <Link to='/signup' className='text-subtitle-purple underline'>sign up</Link>
					</>
					:
					<>
					<Link to='/signin' className='text-subtitle-purple underline'>Sign in</Link> or sign up
					</>
					}
				</h2>

      </div>

      <form action="" className='flex flex-col w-60 gap-4' onSubmit={handleSubmit(onSubmit)}>
        <div className="flex flex-col gap-6">

          <div className="">
            <div className="relative">
              <input {...register('username', {required: true})} type="text" id="floating_outlined" className="block w-full text-subtitle-gray bg-transparent font-bold appearance-none py-2.5 px-4 rounded-xl border-2  outline-none" autoComplete='off'/>
              <label htmlFor="floating_outlined" className="absolute text-sm text-subtitle-gray duration-300 transform -translate-y-4  top-2.5 z-10 origin-[0] bg-background-800 px-0.5 start-3">Login</label>
            </div>
            {errors.username && <p className='text-sm text-error'>Enter login</p>}
          </div>

          <div className="">
            <div className="relative">
              <input {...register('password', {required: true})} type="text" id="floating_outlined" className='block w-full text-subtitle-gray bg-transparent font-bold appearance-none py-2.5 px-4 rounded-xl border-2 outline-none' autoComplete='off'/>
              <label htmlFor="floating_outlined" className="absolute text-sm text-subtitle-gray duration-300 transform -translate-y-4  top-2.5 z-10 origin-[0] bg-background-800 px-0.5 start-3">Password</label>
            </div>
            {errors.password && <p className='text-sm text-error'>Enter password</p>}
          </div>

        </div>
        <input type="submit" className='bg-primary text-background-800 font-bold py-3.5 px-4 rounded-xl outline-none cursor-pointer' value={location.pathname === '/signin' ? 'Sign in' : 'Sign up'}/>
      </form>
      
    </div>
    </>
	)
}

export default Auth