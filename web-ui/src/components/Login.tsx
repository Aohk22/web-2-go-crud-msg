import z from 'zod';
import './Styles.css'
import { checkJwt } from '../lib/utils.ts'
import { Navigate } from 'react-router-dom';

const FormInput = z.object({
	username: z.string(),
	password: z.string(),
})

export default function Login() {
	async function login(formData: FormData) {
		'use server';
		const details = FormInput.parse({
			username: formData.get('username'),
			password: formData.get('password'),
		});
	
		const res = await fetch(`${import.meta.env.VITE_API_URL}/login`, {
			method: "POST",
			headers: { "Content-Type": "application/json" },
			body: JSON.stringify(details),
		});

		const jwt = await res.text();
		localStorage.setItem('jwtToken', jwt);
	}
	
	if (checkJwt()) {
		return <Navigate to='/' replace />
	} else {
		return (
			<div className='loginBlock'>
				<h2>Login</h2>
				<form className='loginForm' action={login}>
					<input type='text' name='username' placeholder='username'></input>
					<input type='password' name='password' placeholder='password'></input>
					<button type='submit'>confirm</button>
				</form>
			</div>
		);
	}

}
