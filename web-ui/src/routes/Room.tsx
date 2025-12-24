import { useState, useEffect } from 'react'
import { useParams } from 'react-router'
import { RoomUsers, RoomMessages, MessageForm } from '../components'
import { fetchRoomMessages, fetchRoomUsers, roomExists } from '../lib/utils'
import useWs from '../hooks/socket'
import type { Message, User } from '../types/schemas'
import './styles/Room.css'

// not good code :()
export default function Room() {
	const [exist, setExist] = useState<boolean>(false)
	const [messages, setMessages] = useState<Message[]>([])
	const [users, setUsers] = useState<User[]>([])
	const [ready, val, send] = useWs('ws://api.lububu.lan:8080/ws')
	const { id: roomId } = useParams()

	useEffect(() => {
		if (!roomId) return
		console.log('checking room existance')
		roomExists(roomId).then((result) => setExist(result))

		console.log('fetching messages')
		fetchRoomMessages<Message>(roomId)
		.then((messages) => setMessages(messages))
		.catch((err) => {console.log(err)})

		console.log('fetching users')
		fetchRoomUsers<User>(roomId)
		.then((users) => setUsers(users))
		.catch((err) => {console.log(err)})
	}, [roomId])

	useEffect(() => {
		if (ready && val) {
			console.log('got from socket:', val)
			const json = JSON.parse(val)
			const msg: Message = { 
				Id: json.id, 
				Content: json.content,
				UserId: json.userId,
				Time: json.time,
			}
			setMessages(prev => [msg, ...prev])
		}
	}, [ready, val])

	if (!exist || roomId === undefined) {
		return <p>room doesn't exist, or room id undefined</p>
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
			<MessageForm 
				roomId={roomId} 
				send={send} 
			/>
		</>
	)
}
