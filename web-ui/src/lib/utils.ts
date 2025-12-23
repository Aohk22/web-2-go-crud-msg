const API = import.meta.env.VITE_API_URL;

function parseJwt(token: string) {
	var base64Url = token.split('.')[1];
	var base64 = base64Url.replace(/-/g, '+').replace(/_/g, '/');
	var jsonPayload = decodeURIComponent(window.atob(base64).split('').map(function(c) {
		return '%' + ('00' + c.charCodeAt(0).toString(16)).slice(-2);
	}).join(''));
	return JSON.parse(jsonPayload);
}

export function checkLogin(): boolean {
	const token = localStorage.getItem('jwtToken');
	if (token === null) {
		return false;
	}
	try {
		const jwtJson = parseJwt(token);
		const expireTime = jwtJson.exp * 1000;
		var d = new Date();
		return (d.getTime() < expireTime);
	} catch {
		console.log('could not parse jawt token');
		return false;
	}
}

export async function fetchRoomMessages<Message>(id: string): Promise<Message[]> {
	'use server';
	const jawtToken = localStorage.getItem('jwtToken');
	return fetch(`${API}/room/${id}/messages`, {
		method: 'GET', 
		headers: { 'Authorization': `Bearer ${jawtToken}`}
	}).then((res) => {
		if (!res.ok) {
			throw new Error('fetch room messages not ok');
		}
		return res.json();
	}).then((json) => { 
		return json
	}).catch((err) => { console.log(err) })
}

export async function fetchRoomUsers<User>(id: string): Promise<User[]> {
	'use server';
	const jawtToken = localStorage.getItem('jwtToken');
	return fetch(`${API}/room/${id}/users`, {
		method: 'GET',
		headers: { 'Authorization': `Bearer ${jawtToken}` },
	}).then((res) => {
		if (!res.ok) {
			throw new Error('fetch room users not ok');
		}
		return res.json();
	}).then((json) => { 
		return json
	}).catch((err) => { console.log(err) })
}

export async function fetchOldMessages<Message>(rid: string, time: string): Promise<Message[]> {
	'use server';
	const jawtToken = localStorage.getItem('jwtToken');
	return fetch(`${API}/room/${rid}/messages`, {
		method: 'POST',
		headers: { 'Authorization': `Bearer ${jawtToken}` },
		body: JSON.stringify({ 'time': time })
	}).then((res) => {
		if (!res.ok) {
			throw new Error('fetch old message not ok');
		}
		return res.json();
	}).then((json) => {
		return json
	}).catch((err) => { console.log(err) })
}

export function getUserId(): string {
	const token = localStorage.getItem('jwtToken');
	if (token === null) {
		console.log('getUserId(): token not found');
		return "";
	}
	try {
		const jwtJson = parseJwt(token);
		return jwtJson.uid;
	} catch {
		console.log('getUserId(): could not parse jawt token');
		return "";
	}
}

export async function roomExists(id: string): Promise<boolean> {
	try {
        const res = await fetch(`${API}/room/${id}`, {
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
