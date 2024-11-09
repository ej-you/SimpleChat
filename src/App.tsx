function App() {
  return (
    <div className="container flex flex-col justify-center h-screen items-center text-center gap-10">
      <div className="flex flex-col gap-2">
        <h1 className='text-title text-xl font-semibold'>Messanger</h1>
        <h2 className='text-subtitle-gray text-base font-semibold'>Sign in or <a className='text-subtitle-purple underline'>sign up</a></h2>
      </div>
      <form action="" className='flex flex-col w-60 gap-4'>
        <div className="relative">
          <input type="text" id="floating_outlined" className="block w-full text-sm text-subtitle-gray font-semibold bg-transparent appearance-none py-2 px-4 rounded-xl border-2  border-subtitle-gray outline-none " placeholder=" " />
          <label htmlFor="floating_outlined" className="absolute text-sm text-subtitle-gray duration-300 transform -translate-y-4 scale-75 top-2 z-10 origin-[0] bg-background-800 px-2 start-2">Login</label>
        </div>
        <div className="relative">
          <input type="text" id="floating_outlined" className="block w-full text-sm text-subtitle-gray font-semibold bg-transparent appearance-none py-2 px-4 rounded-xl border-2  border-subtitle-gray outline-none " placeholder=" " />
          <label htmlFor="floating_outlined" className="absolute text-sm text-subtitle-gray duration-300 transform -translate-y-4 scale-75 top-2 z-10 origin-[0] bg-background-800 px-2 start-2">Password</label>
        </div>
        <input type="submit" className='bg-primary text-background-800 text-sm font-semibold py-2 px-4 rounded-xl outline-none' value='Sign in'/>
      </form>
    </div>
  )
}

export default App
