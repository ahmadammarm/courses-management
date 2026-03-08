import Api from "@/lib/api"
import type { Course, CourseWithInstructor, CreateCourseRequest, UpdateCourseRequest } from "@/lib/types"
import { useMutation, useQuery, useQueryClient } from "@tanstack/vue-query"
import { useRouter } from 'vue-router';
import Cookies from 'js-cookie';

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

function getInstructorIdFromToken(): string | null {
	const token = Cookies.get('token');
	if (!token) return null;
	const payload = parseJwt(token);
	return payload?.instructor_id || payload?.sub || null;
}

// Composable for creating a course
export const useCreateCourse = () => {
	const router = useRouter();
	const queryClient = useQueryClient();

	return useMutation({
		mutationFn: async (data: CreateCourseRequest) => {
			const response = await Api.post('/courses', data)
			return response.data;
		},
		onSuccess: () => {
			queryClient.invalidateQueries({ queryKey: ['instructor-courses'] });
			queryClient.invalidateQueries({ queryKey: ['dashboard-stats'] });
			router.push('/dashboard/courses');
		}
	})
}

// Composable for updating a course
export const useUpdateCourse = (courseId: string) => {
	const router = useRouter();
	const queryClient = useQueryClient();

	return useMutation({
		mutationFn: async (data: UpdateCourseRequest) => {
			const response = await Api.put(`/courses/${courseId}`, data)
			return response.data;
		},
		onSuccess: () => {
			queryClient.invalidateQueries({ queryKey: ['instructor-courses'] });
			queryClient.invalidateQueries({ queryKey: ['course', courseId] });
			router.push('/dashboard/courses');
		}
	})
}

// Composable for deleting a course with optimistic update
export const useDeleteCourse = () => {
	const router = useRouter();
	const queryClient = useQueryClient();

	return useMutation({
		mutationFn: async (courseId: string) => {
			const response = await Api.delete(`/courses/${courseId}`)
			return response.data;
		},
		// Optimistic update - immediately remove from cache before API call
		onMutate: async (courseId: string) => {
			// Cancel any outgoing refetches
			await queryClient.cancelQueries({ queryKey: ['instructor-courses'] });

			// Snapshot the previous value
			const previousCourses = queryClient.getQueryData(['instructor-courses']);

			// Optimistically remove the course from the list
			queryClient.setQueryData(['instructor-courses'], (old: any) => {
				if (!old?.data) return old;
				return {
					...old,
					data: old.data.filter((course: CourseWithInstructor) => course.id !== courseId)
				};
			});

			return { previousCourses };
		},
		// If mutation fails, rollback to previous value
		onError: (err: any, courseId: string, context: any) => {
			queryClient.setQueryData(['instructor-courses'], context.previousCourses);
		},
		// Always refetch after error or success to ensure sync
		onSettled: () => {
			queryClient.invalidateQueries({ queryKey: ['instructor-courses'] });
			queryClient.invalidateQueries({ queryKey: ['dashboard-stats'] });
		}
	})
}

// Composable for getting all courses by current instructor
export const useInstructorCourses = () => {
	const instructorId = getInstructorIdFromToken();

	return useQuery({
		queryKey: ['instructor-courses'],
		queryFn: async () => {
			const response = await Api.get(`/courses/instructor/${instructorId || 'current'}`);
			return response.data as {
				success: boolean;
				data: CourseWithInstructor[];
			};
		},
		enabled: !!instructorId,
		staleTime: 0,
		refetchOnWindowFocus: true,
	})
}

// Composable for getting a single course by ID
export const useCourse = (courseId: string) => {
	return useQuery({
		queryKey: ['course', courseId],
		queryFn: async () => {
			const response = await Api.get(`/courses/${courseId}`);
			return response.data as {
				success: boolean;
				data: Course;
			};
		},
		enabled: !!courseId,
		staleTime: 0,
		refetchOnWindowFocus: true,
	})
}
