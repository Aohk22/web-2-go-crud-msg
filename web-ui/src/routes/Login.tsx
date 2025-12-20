import z from 'zod';
import { checkLogin } from '../lib/utils.ts'
import { Navigate, useNavigate } from 'react-router';

const FormInput = z.object({
	username: z.string(),
	password: z.string(),
})

export default function Login() {
	const navigate = useNavigate();

	function login(formData: FormData) {
		const details = FormInput.parse({
			username: formData.get('username'),
			password: formData.get('password'),
		});
	
		fetch(`${import.meta.env.VITE_API_URL}/login`, {
			method: 'POST',
			headers: { 'Content-Type': 'application/json' },
			body: JSON.stringify(details),
		})
		.then((res) => {
			if (!res.ok) {
				throw new Error('login failed');
			}
			return res.text()
		})
		.then((jwt) => {
			localStorage.setItem('jwtToken', jwt);
			navigate('/')
		})
		.catch((err) => {
			console.log(err);
		})

	}
	
	if (checkLogin()) {
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
