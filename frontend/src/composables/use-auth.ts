import Api from "@/lib/api"
import type { LoginRequest, RegisterRequest, User } from "@/lib/types"
import { useMutation } from "@tanstack/vue-query"
import Cookies from 'js-cookie';
import { useRouter } from 'vue-router';


// Composable for instructoor registration
export const useRegister = () => {
	return useMutation({
		mutationFn: async (data: RegisterRequest) => {
			const response = await Api.post('/auth/register', data)

			return response.data;
		}
	})
}

// composable for instructor login
export const useLogin = () => {
	return useMutation({
		mutationFn: async (data: LoginRequest) => {
			const response = await Api.post('/auth/login', data)
			return response.data;
		}
	})
}


// Composable to get the authenticated user from cookies
export const useAuthUser = (): User | null => {
	const userData = Cookies.get('user');
	if (userData) {
		try {
			return JSON.parse(userData) as User;
		} catch (error) {
			console.error('Error parsing user data from cookies:', error);
			return null;
		}
	}
	return null;
}

// composable to logout the user by removing the cookie
export const useLogout = (): (() => void) => {
	const router = useRouter();

	return () => {
		Cookies.remove('user');
		Cookies.remove('token');

		router.push({ name: 'Login' });
	}
}