import type { Message, User } from '../types/schemas';
import './styles/RoomMessages.css'

export default function RoomMessages(
	{ 
		users, messages
	} : { 
		users: User[], 
		messages: Message[],
	}
) {
	const userMap = Object.fromEntries(users.map(u => [u.Id, u.Name]));

	const messageList = messages.map((_val, i, arr) => {
			// const val = arr[arr.length - 1 - i];
			const val = arr[i];
		return (
			<p key={val.Id}>
			{userMap[val.UserId]}: {val.Content}
			</p>
		)
	});

	return (
		<div className='messageList'>
		{messageList}
		</div>
	)
}
