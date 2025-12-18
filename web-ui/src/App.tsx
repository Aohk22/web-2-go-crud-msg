import { BrowserRouter, Navigate, Route, Routes } from 'react-router';
import { Login, Logout, Home, Room } from './routes';
import { checkJwt } from './lib/utils.ts';
import { useState } from 'react';
import './routes/Styles.css'

export default function App() {
	const [loggedIn, setLoggedIn] = useState(checkJwt());
	return (
		<BrowserRouter>
		<Routes>
			<Route path='/login' element={<Login onLogin={()=>setLoggedIn(true)} />} /> 
			<Route path='/logout' element={<Logout onLogout={()=>setLoggedIn(false)} />} />
			<Route path='/' element={
				loggedIn ? <Home /> : <Navigate to='/login' replace />
			} />
			<Route path='/room/:id' element={
				loggedIn ? <Room /> : <Navigate to='/login' replace />
			} />
			<Route path='*' element={<Navigate to='/' replace />} />
		</Routes>
		</BrowserRouter>
	);
}

