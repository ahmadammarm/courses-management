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

export interface Course {
	id: string;
	title: string;
	slug: string;
	description: string;
	content: string;
	status: 'draft' | 'published';
	instructor_id: string;
	created_at: string;
	updated_at: string;
}

export interface CourseWithInstructor {
	id: string;
	title: string;
	description: string;
	slug: string;
	content: string;
	status: string;
	created_at: string;
	updated_at: string;
	instructor_id: string;
	instructor_name: string;
	instructor_username: string;
}

export interface CreateCourseRequest {
	title: string;
	description: string;
	content: string;
	status: 'draft' | 'published';
}

export interface UpdateCourseRequest {
	title: string;
	slug: string;
	description: string;
	content: string;
	status: 'draft' | 'published';
}

export interface UpdateCourseRequest {
	title: string;
	slug: string;
	description: string;
	content: string;
	status: 'draft' | 'published';
}