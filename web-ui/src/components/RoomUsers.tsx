import type { User } from "../types/schemas";
import './styles/RoomUsers.css'

export default function RoomUsers(
	{ 
		users
	} : { 
		users: User[]
	}
) {
	const userList = users.map((user) => {
		return <p key={user.Id}>{user.Name}</p>
	});

	return (
		<div className='userList'>
		{userList}
		</div>
	);
}
