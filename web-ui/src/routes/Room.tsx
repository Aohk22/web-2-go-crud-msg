import { useState, useEffect } from 'react';
import { useParams } from 'react-router';
import { RoomUsers, RoomMessages, MessageForm } from '../components';
import { fetchRoomMessages, fetchRoomUsers, roomExists } from '../lib/utils';
import type { Message, User } from '../types/schemas';
import './styles/Room.css'

export default function Room() {
	const [exist, setExist] = useState<boolean>(false);
	const [messages, setMessages] = useState<Message[]>([]);
	const [users, setUsers] = useState<User[]>([]);
	const { id: roomId } = useParams();

	useEffect(() => {
		if (!roomId) return;
		roomExists(roomId).then((result) => setExist(result));

		fetchRoomMessages<Message>(roomId)
		.then((messages) => setMessages(messages))
		.catch((err) => {console.log(err)});

		fetchRoomUsers<User>(roomId)
		.then((users) => setUsers(users))
		.catch((err) => {console.log(err)});
	}, [roomId]);

	if (!exist || roomId === undefined) {
		return <p>room doesn't exist, or room id undefined</p>;
	}

	return (
		<>
			<div className='lists'>
				<RoomMessages 
					users={users} 
					messages={messages} 
				/>
				<RoomUsers users={users} />
			</div>
			<MessageForm roomId={roomId} />
		</>
	);
}
