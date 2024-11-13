import { BrowserRouter, Route, Routes } from 'react-router-dom'
import SignIn from './components/auth/SignIn'
import App from './App'
import SignUp from './components/auth/SignUp'
import Messanger from './components/messanger/Messanger'

const Router = () => {
	return (
		<BrowserRouter future={{ v7_startTransition: true, v7_relativeSplatPath: true  }}>
			<Routes>
				<Route path='/' element={<App />}></Route>
				<Route path='/signin' element={<SignIn />}/>
				<Route path='/signup' element={<SignUp />}/>
				<Route path='/messanger' element={<Messanger />}/>
			</Routes>
		</BrowserRouter>
	)
}

export default Router