<script setup lang="ts">
import { onMounted } from "vue";
import { useInstructorCourses } from "@/composables/use-courses";

const { data, isLoading, error, refetch } = useInstructorCourses();

// Force refetch on mount
onMounted(() => {
	refetch();
});
</script>

<template>
	<div>
		<h1 class="text-3xl font-bold text-gray-900 mb-6">Dashboard Overview</h1>

		<!-- Stats Cards -->
		<div class="grid grid-cols-1 md:grid-cols-3 gap-6 mb-8">
			<div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
				<div class="flex items-center justify-between">
					<div>
						<p class="text-sm font-medium text-gray-500">Total Courses</p>
						<p class="text-3xl font-bold text-gray-900 mt-1">{{ isLoading ? '...' : data?.data?.length || 0 }}</p>
					</div>
					<div class="w-12 h-12 bg-indigo-100 rounded-lg flex items-center justify-center">
						<svg class="w-6 h-6 text-indigo-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
							<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6.253v13m0-13C10.832 5.477 9.246 5 7.5 5S4.168 5.477 3 6.253v13C4.168 18.477 5.754 18 7.5 18s3.332.477 4.5 1.253m0-13C13.168 5.477 14.754 5 16.5 5c1.747 0 3.332.477 4.5 1.253v13C19.832 18.477 18.247 18 16.5 18c-1.746 0-3.332.477-4.5 1.253"/>
						</svg>
					</div>
				</div>
			</div>

			<div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
				<div class="flex items-center justify-between">
					<div>
						<p class="text-sm font-medium text-gray-500">Published</p>
						<p class="text-3xl font-bold text-gray-900 mt-1">
							{{ isLoading ? '...' : data?.data?.filter((c: any) => c.status === 'published').length || 0 }}
						</p>
					</div>
					<div class="w-12 h-12 bg-green-100 rounded-lg flex items-center justify-center">
						<svg class="w-6 h-6 text-green-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
							<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"/>
						</svg>
					</div>
				</div>
			</div>

			<div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
				<div class="flex items-center justify-between">
					<div>
						<p class="text-sm font-medium text-gray-500">Drafts</p>
						<p class="text-3xl font-bold text-gray-900 mt-1">
							{{ isLoading ? '...' : data?.data?.filter((c: any) => c.status === 'draft').length || 0 }}
						</p>
					</div>
					<div class="w-12 h-12 bg-yellow-100 rounded-lg flex items-center justify-center">
						<svg class="w-6 h-6 text-yellow-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
							<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-9.414a2 2 0 112.828 2.828L11.828 15H9v-2.828l8.586-8.586z"/>
						</svg>
					</div>
				</div>
			</div>
		</div>

		<!-- Recent Courses -->
		<div class="bg-white rounded-xl shadow-sm border border-gray-200">
			<div class="p-6 border-b border-gray-200">
				<h2 class="text-lg font-semibold text-gray-900">Recent Courses</h2>
			</div>
			<div v-if="isLoading" class="p-6 text-center text-gray-500">Loading...</div>
			<div v-else-if="error" class="p-6 text-center text-red-500">Error loading courses</div>
			<div v-else-if="!data?.data?.length" class="p-6 text-center text-gray-500">
				<p>No courses yet. Create your first course!</p>
				<router-link to="/dashboard/create" class="text-indigo-600 hover:text-indigo-700 font-medium mt-2 inline-block">
					Create Course →
				</router-link>
			</div>
			<div v-else class="divide-y divide-gray-100">
				<div v-for="course in data?.data?.slice(0, 5)" :key="course.id" class="p-4 hover:bg-gray-50 flex items-center justify-between">
					<div>
						<h3 class="font-medium text-gray-900">{{ course.title }}</h3>
						<p class="text-sm text-gray-500 mt-1">{{ course.description }}</p>
					</div>
					<span
						:class="[
							'px-3 py-1 rounded-full text-xs font-medium',
							course.status === 'published' ? 'bg-green-100 text-green-700' : 'bg-yellow-100 text-yellow-700'
						]"
					>
						{{ course.status }}
					</span>
				</div>
			</div>
		</div>
	</div>
</template>
