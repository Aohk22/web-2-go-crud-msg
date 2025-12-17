import { useState, useEffect } from 'react';
import { useNavigate, useParams } from 'react-router-dom';

const API_LINK = import.meta.env.VITE_API_URL;

interface User {
	Id: number, 
	Name: string 
};

interface Message {
	Id: number,
	Time: string,
	Content: string,
	UserId: number,
}

async function fetchRoomData<T>(id: string, path: string): Promise<T[]> {
	'use server';
	const res = await fetch(`${API_LINK}/room/${id}/${path}`, {
		method: 'GET',
		headers: { 
			'Authorization': `Bearer ${localStorage.getItem('jwtToken')}`,
		},
	})

	if (!res.ok) {
		throw new Error(`failed to fetch: ${await res.text()}`);
	}

	return res.json();
}

async function roomExists(id: string): Promise<boolean> {
	const res = await fetch(`${API_LINK}/room/${id}`, {
		method: 'GET',
		headers: { 
			'Authorization': `Bearer ${localStorage.getItem('jwtToken')}`,
		},
	})
	return res.ok
}

export default function Room() {
	const [exist, setExist] = useState<boolean>(false);
	const [users, setUsers] = useState<User[]>([]);
	const [messages, setMessages] = useState<Message[]>([]);
	const { id } = useParams();
	const navigate = useNavigate();

	useEffect(() => {
		if (!id) return;
		roomExists(id).then(setExist);
	}, [id]);

	useEffect(() => {
		if (!id || exist === false) return;
		Promise.all([
			fetchRoomData<User>(id, 'users'),
			fetchRoomData<Message>(id, 'messages'),
		])
		.then(([users, messages]) => {
			setUsers(users);
			setMessages(messages);
		})
		.catch(() => navigate('/login'))
	}, [id, exist, navigate]);

	if (!exist || users === null) {
		return <p>room doesn't exist</p>;
	}

	const userMap = Object.fromEntries(
		users.map(u => [u.Id, u.Name])
	);

	const userList = users.map((user) => {
		return <li key={user.Id}>{user.Name}</li>
	});

	const messageList = messages.map((message) => {
		return (
			<li key={message.Id}>
			{userMap[message.UserId]}: {message.Content}
			</li>
		)
	});

	return (
		<div className='contentBlock'>
			<div className='messageList'>
				<ul>
				{messageList}
				</ul>
			</div>
			<div className='userList'>
				<ul>
				{userList}
				</ul>
			</div>
		</div>
	);
}
