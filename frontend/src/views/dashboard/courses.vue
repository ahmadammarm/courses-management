<script setup lang="ts">
import { onMounted } from "vue";
import { useInstructorCourses } from "@/composables/use-courses";

const { data, isLoading, error, refetch } = useInstructorCourses();

// Force refetch on mount to get fresh data
onMounted(() => {
	refetch();
});
</script>

<template>
	<div>
		<div class="flex items-center justify-between mb-6">
			<h1 class="text-3xl font-bold text-gray-900">My Courses</h1>
			<router-link
				to="/dashboard/create"
				class="px-4 py-2 bg-indigo-600 text-white rounded-lg hover:bg-indigo-700 transition-colors flex items-center gap-2"
			>
				<svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
					<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4"/>
				</svg>
				Create Course
			</router-link>
		</div>

		<div v-if="isLoading" class="bg-white rounded-xl shadow-sm border border-gray-200 p-8 text-center">
			<div class="animate-pulse">
				<div class="h-4 bg-gray-200 rounded w-1/4 mx-auto mb-4"></div>
				<div class="h-20 bg-gray-200 rounded"></div>
			</div>
		</div>

		<div v-else-if="error" class="bg-white rounded-xl shadow-sm border border-gray-200 p-8 text-center">
			<p class="text-red-500 mb-4">Error loading courses</p>
			<button @click="refetch()" class="text-indigo-600 hover:text-indigo-700 font-medium">
				Try again
			</button>
		</div>

		<div v-else-if="!data?.data?.length" class="bg-white rounded-xl shadow-sm border border-gray-200 p-8 text-center">
			<div class="w-16 h-16 bg-gray-100 rounded-full flex items-center justify-center mx-auto mb-4">
				<svg class="w-8 h-8 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
					<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6.253v13m0-13C10.832 5.477 9.246 5 7.5 5S4.168 5.477 3 6.253v13C4.168 18.477 5.754 18 7.5 18s3.332.477 4.5 1.253m0-13C13.168 5.477 14.754 5 16.5 5c1.747 0 3.332.477 4.5 1.253v13C19.832 18.477 18.247 18 16.5 18c-1.746 0-3.332.477-4.5 1.253"/>
				</svg>
			</div>
			<h3 class="text-lg font-semibold text-gray-900 mb-2">No courses yet</h3>
			<p class="text-gray-500 mb-4">Create your first course to get started</p>
			<router-link
				to="/dashboard/create"
				class="inline-flex items-center gap-2 px-4 py-2 bg-indigo-600 text-white rounded-lg hover:bg-indigo-700 transition-colors"
			>
				<svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
					<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4"/>
				</svg>
				Create Your First Course
			</router-link>
		</div>

		<div v-else class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
			<div
				v-for="course in data?.data"
				:key="course.id"
				class="bg-white rounded-xl shadow-sm border border-gray-200 overflow-hidden hover:shadow-md transition-shadow"
			>
				<div class="h-32 bg-gradient-to-r from-indigo-500 to-purple-600 flex items-center justify-center">
					<svg class="w-12 h-12 text-white/50" fill="none" stroke="currentColor" viewBox="0 0 24 24">
						<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6.253v13m0-13C10.832 5.477 9.246 5 7.5 5S4.168 5.477 3 6.253v13C4.168 18.477 5.754 18 7.5 18s3.332.477 4.5 1.253m0-13C13.168 5.477 14.754 5 16.5 5c1.747 0 3.332.477 4.5 1.253v13C19.832 18.477 18.247 18 16.5 18c-1.746 0-3.332.477-4.5 1.253"/>
					</svg>
				</div>
				<div class="p-5">
					<div class="flex items-center justify-between mb-2">
						<span
							:class="[
								'px-2.5 py-0.5 rounded-full text-xs font-medium',
								course.status === 'published'
									? 'bg-green-100 text-green-700'
									: 'bg-yellow-100 text-yellow-700'
							]"
						>
							{{ course.status }}
						</span>
						<span class="text-xs text-gray-400">
							{{ new Date(course.created_at).toLocaleDateString() }}
						</span>
					</div>
					<h3 class="font-semibold text-gray-900 mb-2 line-clamp-1">{{ course.title }}</h3>
					<p class="text-sm text-gray-500 line-clamp-2 mb-4">{{ course.description }}</p>
					<div class="flex items-center gap-2 pt-4 border-t border-gray-100">
						<router-link :to="`/dashboard/courses/${course.id}/edit`" class="flex-1 text-sm text-indigo-600 hover:text-indigo-700 font-medium text-center">
							Edit
						</router-link>
						<router-link :to="`/dashboard/courses/${course.id}`" class="flex-1 text-sm text-gray-500 hover:text-gray-600 font-medium text-center">
							View
						</router-link>
					</div>
				</div>
			</div>
		</div>
	</div>
</template>
