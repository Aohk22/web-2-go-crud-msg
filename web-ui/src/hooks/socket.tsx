import { useEffect, useRef, useState, } from 'react'
import type { RefObject } from 'react'

// singleton / multiple ws, not sure what to use yet

export default function useWs(url: string): [boolean, any, Function]  {
	const [isReady, setIsReady] = useState<boolean>(false)
	const [val, setVal] = useState()

	const ws = useRef<RefObject<WebSocket>>(null)

	useEffect(() => {
		const s = new WebSocket(url, ['wamp', `${localStorage.getItem('jwtToken')}`])

		s.onopen = () => {
			console.log('connected to socket', url)
			setIsReady(true)
		}

		s.onclose = () => {
			console.log('disconnected from socket', url)
			setIsReady(false)
		}

		s.onmessage = (e) => {
			console.log('message to send:', e.data)
			setVal(e.data)
		}

		ws.current = s

		return () => {
			if (s.readyState === 1) {
				s.close()
			} else {
				s.addEventListener('open', () => {
					s.close()
				})
			}
		}
	}, [])

	return [isReady, val, ws.current?.send.bind(ws.current)]
}
