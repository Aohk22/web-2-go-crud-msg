import { BrowserRouter, Navigate, Route, Routes } from 'react-router';
import { Login, Logout, Home, Room } from './routes';
import { checkLogin } from './lib/utils.ts';
import './routes/Styles.css'
import type { JSX } from 'react';

function RequireAuth({ children } : { children: JSX.Element }) {
	return checkLogin() ? children : <Navigate to='/login' replace />;
}

export default function App() {
	return (
		<BrowserRouter>
		<Routes>
			<Route path='/login' element={<Login />} /> 
			<Route path='/logout' element={<Logout />} />
			<Route path='/' element={
				<RequireAuth children={<Home />} />
			} />
			<Route path='/room/:id' element={
				<RequireAuth children={<Room />} />
			} />
			<Route path='*' element={<Navigate to='/' replace />} />
		</Routes>
		</BrowserRouter>
	);
}

