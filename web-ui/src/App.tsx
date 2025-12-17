import { BrowserRouter, Navigate, Route, Routes } from 'react-router-dom';
import './App.css';
import Login from './components/Login';
import Home from './components/Home';
import Logout from './components/Logout';
import Room from './components/Room';
import { checkJwt } from './lib/utils.ts';

export default function App() {
	const loggedIn = checkJwt();
	return (
		<BrowserRouter>
		<Routes>
			<Route path='/login' element={<Login />} /> 
			<Route path='/logout' element={<Logout />} />
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

