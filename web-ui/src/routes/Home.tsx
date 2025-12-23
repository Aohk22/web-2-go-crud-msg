import { RoomsTable } from '../components';
import './styles/Home.css'

export default function Home() {
	return (
		<>
			<span className='homeHead'>
				<h2>rooms</h2>
				<sub><a href='/logout'>logout</a></sub>
			</span>
			<RoomsTable />
		</>
	);
}
