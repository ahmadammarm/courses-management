import Cookies from 'js-cookie';
import Api from './api';

function parseJwt(token: string): any | null {
	try {
		const base64Url = token.split('.')[1];
		const base64 = base64Url.replace(/-/g, '+').replace(/_/g, '/');
		const jsonPayload = decodeURIComponent(atob(base64).split('').map(c => '%' + ('00' + c.charCodeAt(0).toString(16)).slice(-2)).join(''));
		return JSON.parse(jsonPayload);
	} catch (e) {
		return null;
	}
}

export function isTokenExpired(token: string): boolean {
	const payload = parseJwt(token);
	if (!payload || !payload.exp) return true;
	const now = Math.floor(Date.now() / 1000);
	return payload.exp <= now;
}

let logoutTimeout: number | undefined;

export function setupSession() {
	const token = Cookies.get('token');

	if (token) {
		if (isTokenExpired(token)) {
			Cookies.remove('token');
			Cookies.remove('user');
			window.location.replace('/login');
		} else {
			Api.defaults.headers.common['Authorization'] = `Bearer ${token}`;

			const payload = parseJwt(token);
			if (payload && payload.exp) {
				const msUntilExp = (payload.exp * 1000) - Date.now();
				if (msUntilExp > 0) {
					logoutTimeout = window.setTimeout(() => {
						Cookies.remove('token');
						Cookies.remove('user');
						window.location.replace('/login');
					}, msUntilExp);
				}
			}
		}
	}

	// Attach request interceptor to always set Authorization header from cookie
	Api.interceptors.request.use((config) => {
		const t = Cookies.get('token');
		if (t) {
			config.headers = config.headers || {};
			(config.headers as any)['Authorization'] = `Bearer ${t}`;
		}
		return config;
	});

	// Logout on 401 responses
	Api.interceptors.response.use(
		response => response,
		(error) => {
			if (error?.response?.status === 401) {
				Cookies.remove('token');
				Cookies.remove('user');
				window.location.replace('/login');
			}
			return Promise.reject(error);
		}
	);
}

export function clearSessionTimers() {
	if (logoutTimeout) {
		clearTimeout(logoutTimeout);
		logoutTimeout = undefined;
	}
}
