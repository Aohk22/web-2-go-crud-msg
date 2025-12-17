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
			const rooms = await res.json();
			setRooms(rooms);
		}
		getRooms();
	}, []);

	const listItems = rooms.map(({Id, Time, Name}) => {
		return (
			<li key={Id}>
				<span>{Id}</span> - <span>{Name}</span> - <span>{Time}</span>
			</li>
		)
	});

	return (
		<ul className='roomCards'>{listItems}</ul>
	);
}

export default function Home() {
	return (
		<div className="homeBlock">
			<h2>rooms</h2>
			<RoomCards />
		</div>
	);
}
