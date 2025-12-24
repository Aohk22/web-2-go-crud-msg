import { getUserId } from '../lib/utils'
import './styles/MessageForm.css'

// auto expand text area
document.querySelectorAll("textarea").forEach(function(textarea) {
	textarea.style.height = textarea.scrollHeight + "px"
	textarea.style.overflowY = "hidden"

	textarea.addEventListener("input", function() {
		this.style.height = "auto"
		this.style.height = this.scrollHeight + "px"
	})
})

export default function MessageForm(
	{ 
		roomId, send
	} : {
		roomId: string,
		send: Function,
	}
) {
	const userId = getUserId()

	function handleFormKeyPress(e: React.KeyboardEvent<HTMLTextAreaElement>) {
		if (e.key === 'Enter' && !e.shiftKey) {
			if (e.currentTarget.value.trim().length > 0) {
				const form = e.currentTarget.form
				if (!form) {
					console.log('current target not a form')
				} else {
					form.requestSubmit()
				}
			}
			e.preventDefault()
		}
	}

	async function sendMessage(formData: FormData) {
		if (typeof formData.get('content') != 'string') {
			console.log('message form has wrong type')
			return
		}
		if (formData.get('content').trim().length <= 0) {
			console.log('content cannot be empty')
			return
		}

		formData.append('userId', userId)
		formData.append('roomId', roomId)
		formData.append('time', `${Math.ceil(Date.now()/1000)}`)
		const json = JSON.stringify(Object.fromEntries(formData.entries()))
		send(json)
	}

	return (
		<form 
			id='messageForm'
			className='messageForm'
			action={sendMessage}
		>
			<textarea 
				id='messageFormContent'
				className='messageFormContent'
				onKeyDown={handleFormKeyPress}
				form='messageForm'
				name='content'
			>
			</textarea>
			<button id='messageFormSubmit' type='submit'>send</button>
		</form>
	)
}
