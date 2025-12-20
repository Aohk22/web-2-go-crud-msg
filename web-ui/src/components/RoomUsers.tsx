import type { User } from "../types/schemas";

export default function RoomUsers(
	{ 
		users
	} : { 
		users: User[]
	}
) {
	const userList = users.map((user) => {
		return <li key={user.Id}>{user.Name}</li>
	});

	return (
		<ul>
		{userList}
		</ul>
	);
}
