import { useState, useEffect } from 'react';
import { useParams } from 'react-router';
import { RoomUsers, RoomMessages, MessageForm } from '../components';
import type { Message, User } from '../types/schemas';
import './Room.css'
import { fetchRoomMessages, fetchRoomUsers } from '../lib/utils';

const API_LINK = import.meta.env.VITE_API_URL;

async function roomExists(id: string): Promise<boolean> {
	try {
        const res = await fetch(`${API_LINK}/room/${id}`, {
            method: 'GET',
            headers: {
                'Authorization': `Bearer ${localStorage.getItem('jwtToken')}`,
            },
        });
        return res.ok;
    } catch (err) {
        console.log(err);
        return false;
    }
}

export default function Room() {
	const [exist, setExist] = useState<boolean>(false);
	const [messages, setMessages] = useState<Message[]>([]);
	const [users, setUsers] = useState<User[]>([]);
	const [refresh, setRefresh] = useState(0);
	const { id: roomId } = useParams();

	function refreshMessages() {
		setRefresh(r => r + 1)
	}

	useEffect(() => {
		if (!roomId) return;
		roomExists(roomId).then((exist) => setExist(exist));

		fetchRoomMessages<Message>(roomId)
		.then((messages) => setMessages(messages))
		.catch((err) => {console.log(err)});

		fetchRoomUsers<User>(roomId)
		.then((users) => setUsers(users))
		.catch((err) => {console.log(err)});
	}, [roomId]);

	useEffect(() => {
		if (!roomId) return;
		fetchRoomMessages<Message>(roomId)
		.then((messages) => setMessages(messages))
		.catch((err) => {console.log(err)});
	}, [refresh])

	if (!exist || roomId === undefined) {
		return <p>room doesn't exist, or room id undefined</p>;
	}

	return (
		<div className='contentBlock'>
			<div className='contentList'>
				<div className='messageList'>
					<RoomMessages 
						users={users}
						messages={messages} 
					/>
				</div>
				<div className='userList'>
					<RoomUsers users={users} />
				</div>
			</div>
			<div className='messageForm'>
				<MessageForm 
					roomId={roomId}
					onSent={refreshMessages}
				/>
			</div>
		</div>
	);
}
