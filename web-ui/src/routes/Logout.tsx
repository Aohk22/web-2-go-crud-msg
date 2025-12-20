import { Navigate } from "react-router";

export default function Logout() {
	localStorage.removeItem('jwtToken');
	return (
		<Navigate to='/login' replace />
	);
}
