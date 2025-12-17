function parseJwt(token: string) {
	var base64Url = token.split('.')[1];
	var base64 = base64Url.replace(/-/g, '+').replace(/_/g, '/');
	var jsonPayload = decodeURIComponent(window.atob(base64).split('').map(function(c) {
		return '%' + ('00' + c.charCodeAt(0).toString(16)).slice(-2);
	}).join(''));
	return JSON.parse(jsonPayload);
}

export function checkJwt(): boolean {
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
