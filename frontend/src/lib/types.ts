export interface RegisterRequest {
	name: string;
	username: string;
	email: string;
	password: string;
	expertise: string;
}

export interface LoginRequest {
	email: string;
	password: string;
}

export interface User {
	id: number;
	name: string;
	username: string;
	email: string;
	expertise: string;
}