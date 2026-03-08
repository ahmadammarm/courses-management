<script setup lang="ts">
import { ref, watch } from "vue";
import { useRouter, useRoute } from "vue-router";
import { useCourse, useDeleteCourse } from "@/composables/use-courses";

const router = useRouter();
const route = useRoute();
const courseId = route.params.id as string;

const { data: courseData, isLoading, error, refetch } = useCourse(courseId);
const { mutate: deleteCourse, isPending: isDeleting } = useDeleteCourse();

const showDeleteModal = ref(false);
const isDeleted = ref(false);

const handleDelete = () => {
	// Optimistic update - mark as deleted immediately
	isDeleted.value = true;

	deleteCourse(courseId, {
		onSuccess: () => {
			showDeleteModal.value = false;
			// Redirect after successful delete
			router.push('/dashboard/courses');
		},
		onError: (err: any) => {
			// Revert if failed
			isDeleted.value = false;
			alert(err.response?.data?.message || 'Failed to delete course');
		}
	});
};

const formatDate = (dateString: string) => {
	if (!dateString) return '-';
	return new Date(dateString).toLocaleDateString('en-US', {
		year: 'numeric',
		month: 'long',
		day: 'numeric',
		hour: '2-digit',
		minute: '2-digit'
	});
};
</script>

<template>
	<div class="max-w-4xl mx-auto">
		<!-- Back Button -->
		<button
			@click="router.push('/dashboard/courses')"
			class="flex items-center gap-2 text-gray-600 hover:text-gray-900 mb-6 transition-colors"
		>
			<svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
				<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 19l-7-7 7-7"/>
			</svg>
			Back to Courses
		</button>

		<!-- Loading State -->
		<div v-if="isLoading" class="bg-white rounded-xl shadow-sm border border-gray-200 p-8">
			<div class="animate-pulse">
				<div class="h-8 bg-gray-200 rounded w-1/2 mb-4"></div>
				<div class="h-4 bg-gray-200 rounded w-1/4 mb-2"></div>
				<div class="h-32 bg-gray-200 rounded mt-6"></div>
			</div>
		</div>

		<!-- Error State -->
		<div v-else-if="error" class="bg-white rounded-xl shadow-sm border border-gray-200 p-8 text-center">
			<div class="w-16 h-16 bg-red-100 rounded-full flex items-center justify-center mx-auto mb-4">
				<svg class="w-8 h-8 text-red-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
					<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z"/>
				</svg>
			</div>
			<h3 class="text-lg font-semibold text-gray-900 mb-2">Error Loading Course</h3>
			<p class="text-gray-500 mb-4">The course could not be found or you don't have permission to view it.</p>
			<button
				@click="router.push('/dashboard/courses')"
				class="px-4 py-2 bg-indigo-600 text-white rounded-lg hover:bg-indigo-700 transition-colors"
			>
				Back to Courses
			</button>
		</div>

		<!-- Course Deleted State -->
		<div v-else-if="isDeleted" class="bg-white rounded-xl shadow-sm border border-gray-200 p-8 text-center">
			<div class="w-16 h-16 bg-green-100 rounded-full flex items-center justify-center mx-auto mb-4">
				<svg class="w-8 h-8 text-green-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
					<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7"/>
				</svg>
			</div>
			<h3 class="text-lg font-semibold text-gray-900 mb-2">Course Deleted</h3>
			<p class="text-gray-500 mb-4">Redirecting to courses...</p>
		</div>

		<!-- Course Content -->
		<div v-else-if="courseData?.data" class="space-y-6">
			<!-- Header Card -->
			<div class="bg-white rounded-xl shadow-sm border border-gray-200 overflow-hidden">
				<div class="h-32 bg-gradient-to-r from-indigo-500 to-purple-600"></div>
				<div class="p-6">
					<div class="flex items-start justify-between">
						<div class="flex-1">
							<span
								:class="[
									'px-3 py-1 rounded-full text-xs font-medium mb-3 inline-block',
									courseData.data.status === 'published'
										? 'bg-green-100 text-green-700'
										: 'bg-yellow-100 text-yellow-700'
								]"
							>
								{{ courseData.data.status }}
							</span>
							<h1 class="text-3xl font-bold text-gray-900">{{ courseData.data.title }}</h1>
							<p class="mt-2 text-gray-600">{{ courseData.data.description }}</p>
						</div>
						<div class="flex gap-3 ml-4">
							<router-link
								:to="`/dashboard/courses/${courseId}/edit`"
								class="px-4 py-2 bg-indigo-600 text-white rounded-lg hover:bg-indigo-700 transition-colors flex items-center gap-2"
							>
								<svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
									<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-9.414a2 2 0 112.828 2.828L11.828 15H9v-2.828l8.586-8.586z"/>
								</svg>
								Edit
							</router-link>
							<button
								@click="showDeleteModal = true"
								class="px-4 py-2 bg-red-100 text-red-700 rounded-lg hover:bg-red-200 transition-colors flex items-center gap-2"
							>
								<svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
									<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16"/>
								</svg>
								Delete
							</button>
						</div>
					</div>
				</div>
			</div>

			<!-- Course Details -->
			<div class="grid grid-cols-1 md:grid-cols-3 gap-6">
				<!-- Content -->
				<div class="md:col-span-2 bg-white rounded-xl shadow-sm border border-gray-200 p-6">
					<h2 class="text-lg font-semibold text-gray-900 mb-4">Course Content</h2>
					<div class="prose max-w-none text-gray-600 whitespace-pre-wrap">
						{{ courseData.data.content || 'No content added yet.' }}
					</div>
				</div>

				<!-- Sidebar -->
				<div class="space-y-6">
					<!-- Info Card -->
					<div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
						<h3 class="text-sm font-semibold text-gray-500 uppercase tracking-wider mb-4">Details</h3>
						<div class="space-y-4">
							<div>
								<p class="text-xs text-gray-400">Slug</p>
								<p class="text-sm text-gray-900 font-mono">{{ courseData.data.slug }}</p>
							</div>
							<div>
								<p class="text-xs text-gray-400">Created</p>
								<p class="text-sm text-gray-900">{{ formatDate(courseData.data.created_at) }}</p>
							</div>
							<div>
								<p class="text-xs text-gray-400">Last Updated</p>
								<p class="text-sm text-gray-900">{{ formatDate(courseData.data.updated_at) }}</p>
							</div>
						</div>
					</div>
				</div>
			</div>
		</div>

		<!-- Delete Confirmation Modal -->
		<div v-if="showDeleteModal" class="fixed inset-0 bg-black/50 flex items-center justify-center z-50">
			<div class="bg-white rounded-xl shadow-xl max-w-md w-full mx-4 p-6">
				<div class="flex items-center justify-center w-12 h-12 mx-auto bg-red-100 rounded-full mb-4">
					<svg class="w-6 h-6 text-red-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
						<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16"/>
					</svg>
				</div>
				<h3 class="text-lg font-semibold text-gray-900 text-center mb-2">Delete Course?</h3>
				<p class="text-gray-500 text-center mb-6">
					Are you sure you want to delete this course? This action cannot be undone.
				</p>
				<div class="flex gap-3">
					<button
						@click="showDeleteModal = false"
						class="flex-1 px-4 py-2 border border-gray-300 text-gray-700 rounded-lg hover:bg-gray-50 transition-colors"
					>
						Cancel
					</button>
					<button
						@click="handleDelete"
						:disabled="isDeleting"
						class="flex-1 px-4 py-2 bg-red-600 text-white rounded-lg hover:bg-red-700 disabled:opacity-50 transition-colors"
					>
						{{ isDeleting ? 'Deleting...' : 'Delete' }}
					</button>
				</div>
			</div>
		</div>
	</div>
</template>
