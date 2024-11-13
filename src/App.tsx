import { useNavigate } from 'react-router-dom'
import { useEffect } from 'react'
// import axios from 'axios'

function App() {
  const nav = useNavigate()

  useEffect(() =>{
    if(!localStorage.getItem('registered')){
      nav('/signin')
    }
  },[nav])

  const onSubmit = async (e:React.FormEvent<HTMLFormElement>) => {
    e.preventDefault()
    
    // try{
    //   const res = await axios.get('150.241.82.68/api/chat/restricted')
    //   console.log(res)
    //   nav('/messanger')
    // } catch(err) {
    //   console.error(err)
    // }
  }

  return (
    <div className="flex flex-col justify-center h-screen items-center text-center gap-12">
      <div className="flex flex-col gap-2">
        <h1 className='text-title text-xl font-bold'>Find user for chatting</h1>
				<h2 className='text-subtitle-gray font-bold'>type user login</h2>
      </div>

      <form action="" className='flex flex-col w-60 gap-3.5' onSubmit={(e) => {onSubmit(e)}}>
        <div className="relative">
          <input type="text" id="floating_outlined" className="block w-full text-subtitle-gray font-bold bg-transparent appearance-none py-2.5 px-4 rounded-xl border-2  border-subtitle-gray outline-none" autoComplete="off"/>
          <label htmlFor="floating_outlined" className="absolute text-sm text-subtitle-gray duration-300 transform -translate-y-4  top-2.5 z-10 origin-[0] bg-background-800 px-0.5 start-3">Login</label>
        </div>
        <input type="submit" className='bg-primary text-background-800 font-bold py-3.5 px-4 rounded-xl outline-none cursor-pointer' value='Find'/>
      </form>

    </div>
  )
}

export default App