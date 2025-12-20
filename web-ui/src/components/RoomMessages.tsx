import type { Message, User } from '../types/schemas';

export default function RoomMessages(
	{ 
		users, messages
	} : { 
		users: User[], 
		messages: Message[],
	}
) {
	const userMap = Object.fromEntries(users.map(u => [u.Id, u.Name]));

	const messageList = messages.map((m) => {
		return (
			<li key={m.Id}>
			{userMap[m.UserId]}: {m.Content}
			</li>
		)
	});

	return (
		<ul>
		{messageList}
		</ul>
	)
}
