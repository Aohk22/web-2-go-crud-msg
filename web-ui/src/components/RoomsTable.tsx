import { useState, useEffect } from 'react';

export default function RoomsTable() {
	const [rooms, setRooms] = useState([]);

	useEffect(() => { 
		async function getRooms() {
			'use server';
			const res = await fetch(`${import.meta.env.VITE_API_URL}/rooms`, {
				method: 'GET',
				headers: {
					'Authorization': `Bearer ${localStorage.getItem('jwtToken')}`,
				}
			})

			if (!res.ok) {
				console.log(`fetching rooms failed: ${await res.text()}`);
				return;
			}

			const rooms = await res.json();
			setRooms(rooms);
		}
		getRooms();
	}, []);

	const listItems = rooms.map(({Id, Time, Name}) => {
		var d = new Date(Time);
		return (
			<tr key={Id}>
				<td><a href={`/room/${Id}`}>{Name}</a></td>
				<td>{d.toLocaleDateString()}</td>
			</tr>
		)
	});

	return (
		<table className='roomsTable'>
		<thead>
			<tr>
				<th>room name</th>
				<th>create date</th>
			</tr>
		</thead>
		<tbody>
			{listItems}
		</tbody>
		</table>
	);
}
