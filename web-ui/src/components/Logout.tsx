import { Navigate } from "react-router-dom";

export default function Logout() {
	localStorage.removeItem('jwtToken');
	return (
		<Navigate to='/login' replace />
	);
}
