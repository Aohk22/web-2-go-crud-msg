import { getUserId } from '../lib/utils'

const API = import.meta.env.VITE_API_URL;

export default function MessageForm(
	{ 
		roomId, onSent
	} : {
		roomId: string,
		onSent: Function,
	}
) {
	const userId = getUserId();

	function handleKey(e: React.KeyboardEvent<HTMLTextAreaElement>) {
		if (e.key === 'Enter' && !e.shiftKey) {
			if (e.currentTarget.value.trim().length > 0) {
				e.currentTarget.form?.requestSubmit();
			}
			e.preventDefault();
		}
	}

	async function sendMessage(formData: FormData) {
		'use server';

		formData.append('uid', userId);
		formData.append('rid', roomId);

		const fullForm = {'dataType': 'message', 'data': Object.fromEntries(formData.entries())}
		const jawt = localStorage.getItem('jwtToken');

		fetch(`${API}/room/${roomId}`, {
			method: 'PUT',
			headers: {
				'Authorization': `Bearer ${jawt}`,
				'Content-Type': 'application/json',
			},
			body: JSON.stringify(fullForm)
		})
		.then(() => {
			document.getElementById('messageForm')?.reset();
		})
		.then(() => {
			onSent();
		})
		.catch((err) => console.log(err))
	}

	return (
		<>
		<form 
			id='messageForm'
			className='messageForm'
			action={sendMessage}
		>
			<textarea 
				id='messageFormContent'
				className='messageFormContent'
				onKeyDown={handleKey}
				form='messageForm'
				name='content'
			>
			</textarea>
			<button id='messageFormSubmit' type='submit'>send</button>
		</form>
		</>
	)
}
