import { useState, useEffect } from 'react';

function RoomCards() {
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

export default function Home() {
	return (
		<div className='homeBlock'>
			<span className='homeHead'>
				<h2>rooms</h2>
				<sub><a href='/logout'>logout</a></sub>
			</span>
			<RoomCards />
		</div>
	);
}
