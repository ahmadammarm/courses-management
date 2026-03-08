<script setup lang="ts">

import { ref, reactive, watch } from "vue";
import { useRouter, useRoute } from "vue-router";
import { useUpdateCourse, useCourse } from "@/composables/use-courses";

interface ValidationErrors {
	[key: string]: string;
}

const router = useRouter();
const route = useRoute();
const courseId = route.params.id as string;

const title = ref<string>('');
const description = ref<string>('');
const content = ref<string>('');
const slug = ref<string>('');
const status = ref<'draft' | 'published'>('draft');

const errors = reactive<ValidationErrors>({});
const touched = reactive<Record<string, boolean>>({});

const { data: courseData, isLoading: isCourseLoading } = useCourse(courseId);
const { mutate: updateCourse, isPending } = useUpdateCourse(courseId);

// Watch for data changes and populate form
watch(courseData, (newData) => {
	if (newData?.data) {
		const course = newData.data;
		title.value = course.title;
		description.value = course.description;
		content.value = course.content;
		slug.value = course.slug;
		status.value = course.status;
	}
}, { immediate: true });

// Validation rules
const validateField = (field: string, value: string) => {
	switch (field) {
		case 'title':
			if (!value.trim()) {
				errors.title = 'Title is required';
			} else if (value.length < 5) {
				errors.title = 'Title must be at least 5 characters';
			} else if (value.length > 100) {
				errors.title = 'Title must not exceed 100 characters';
			} else {
				delete errors.title;
			}
			break;
		case 'description':
			if (!value.trim()) {
				errors.description = 'Description is required';
			} else if (value.length < 10) {
				errors.description = 'Description must be at least 10 characters';
			} else if (value.length > 500) {
				errors.description = 'Description must not exceed 500 characters';
			} else {
				delete errors.description;
			}
			break;
		case 'content':
			if (!value.trim()) {
				errors.content = 'Content is required';
			} else if (value.length < 20) {
				errors.content = 'Content must be at least 20 characters';
			} else {
				delete errors.content;
			}
			break;
	}
};

const handleBlur = (field: string) => {
	touched[field] = true;
	validateField(field, field === 'title' ? title.value : field === 'description' ? description.value : content.value);
};

const validateAll = () => {
	validateField('title', title.value);
	validateField('description', description.value);
	validateField('content', content.value);
	touched.title = true;
	touched.description = true;
	touched.content = true;
	return Object.keys(errors).length === 0;
};

const handleSubmit = (e: Event) => {
	e.preventDefault();

	touched.title = true;
	touched.description = true;
	touched.content = true;

	if (!validateAll()) {
		return;
	}

	updateCourse(
		{
			title: title.value,
			slug: slug.value,
			description: description.value,
			content: content.value,
			status: status.value,
		},
		{
			onError: (err: any) => {
				if (err.response?.data?.errors) {
					Object.assign(errors, err.response.data.errors);
				} else {
					errors.general = err.response?.data?.message || 'Failed to update course. Please try again.';
				}
			},
		}
	);
};

const getTitleCharCount = () => title.value.length;
const getDescCharCount = () => description.value.length;
const getContentCharCount = () => content.value.length;
</script>

<template>
	<div class="max-w-4xl mx-auto">
		<!-- Header -->
		<div class="mb-8">
			<button
				@click="router.push('/dashboard/courses')"
				class="flex items-center gap-2 text-gray-600 hover:text-gray-900 mb-4 transition-colors"
			>
				<svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
					<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 19l-7-7 7-7"/>
				</svg>
				Back to Courses
			</button>
			<h1 class="text-3xl font-bold text-gray-900">Edit Course</h1>
			<p class="mt-2 text-gray-600">Update the details of your course.</p>
		</div>

		<!-- Loading State -->
		<div v-if="isCourseLoading" class="bg-white rounded-xl shadow-sm border border-gray-200 p-8 text-center">
			<div class="animate-pulse">
				<div class="h-4 bg-gray-200 rounded w-1/4 mx-auto mb-4"></div>
				<div class="h-20 bg-gray-200 rounded"></div>
			</div>
		</div>

		<!-- Form Card -->
		<div v-else class="bg-white rounded-xl shadow-sm border border-gray-200 overflow-hidden">
			<form @submit="handleSubmit" class="p-8">
				<!-- General Error -->
				<div v-if="errors.general" class="mb-6 p-4 bg-red-50 border border-red-200 rounded-lg">
					<p class="text-red-600 text-sm">{{ errors.general }}</p>
				</div>

				<!-- Title Field -->
				<div class="mb-6">
					<label for="title" class="block text-sm font-semibold text-gray-700 mb-2">
						Course Title <span class="text-red-500">*</span>
					</label>
					<input
						v-model="title"
						type="text"
						id="title"
						@blur="handleBlur('title')"
						:class="[
							'w-full px-4 py-3 rounded-lg border transition-all duration-200 focus:outline-none focus:ring-2',
							errors.title && touched.title
								? 'border-red-300 focus:ring-red-200 focus:border-red-400 bg-red-50'
								: 'border-gray-300 focus:ring-indigo-200 focus:border-indigo-400 hover:border-gray-400'
						]"
						placeholder="e.g., Introduction to Web Development"
					/>
					<div class="flex justify-between items-center mt-2">
						<div v-if="errors.title && touched.title" class="text-red-600 text-sm flex items-center gap-1">
							<svg class="w-4 h-4" fill="currentColor" viewBox="0 0 20 20">
								<path fill-rule="evenodd" d="M18 10a8 8 0 11-16 0 8 8 0 0116 0zm-7 4a1 1 0 11-2 0 1 1 0 012 0zm-1-9a1 1 0 00-1 1v4a1 1 0 102 0V6a1 1 0 00-1-1z" clip-rule="evenodd"/>
							</svg>
							{{ errors.title }}
						</div>
						<div v-else class="text-gray-400 text-sm"></div>
						<span class="text-sm" :class="[getTitleCharCount() > 100 ? 'text-red-500' : 'text-gray-400']">
							{{ getTitleCharCount() }}/100
						</span>
					</div>
				</div>

				<!-- Description Field -->
				<div class="mb-6">
					<label for="description" class="block text-sm font-semibold text-gray-700 mb-2">
						Description <span class="text-red-500">*</span>
					</label>
					<textarea
						v-model="description"
						id="description"
						rows="4"
						@blur="handleBlur('description')"
						:class="[
							'w-full px-4 py-3 rounded-lg border transition-all duration-200 focus:outline-none focus:ring-2 resize-none',
							errors.description && touched.description
								? 'border-red-300 focus:ring-red-200 focus:border-red-400 bg-red-50'
								: 'border-gray-300 focus:ring-indigo-200 focus:border-indigo-400 hover:border-gray-400'
						]"
						placeholder="Provide a brief description of what students will learn in this course..."
					></textarea>
					<div class="flex justify-between items-center mt-2">
						<div v-if="errors.description && touched.description" class="text-red-600 text-sm flex items-center gap-1">
							<svg class="w-4 h-4" fill="currentColor" viewBox="0 0 20 20">
								<path fill-rule="evenodd" d="M18 10a8 8 0 11-16 0 8 8 0 0116 0zm-7 4a1 1 0 11-2 0 1 1 0 012 0zm-1-9a1 1 0 00-1 1v4a1 1 0 102 0V6a1 1 0 00-1-1z" clip-rule="evenodd"/>
							</svg>
							{{ errors.description }}
						</div>
						<div v-else class="text-gray-400 text-sm"></div>
						<span class="text-sm" :class="getDescCharCount() > 500 ? 'text-red-500' : 'text-gray-400'">
							{{ getDescCharCount() }}/500
						</span>
					</div>
				</div>

				<!-- Content Field -->
				<div class="mb-6">
					<label for="content" class="block text-sm font-semibold text-gray-700 mb-2">
						Course Content <span class="text-red-500">*</span>
					</label>
					<textarea
						v-model="content"
						id="content"
						rows="8"
						@blur="handleBlur('content')"
						:class="[
							'w-full px-4 py-3 rounded-lg border transition-all duration-200 focus:outline-none focus:ring-2 resize-none font-mono text-sm',
							errors.content && touched.content
								? 'border-red-300 focus:ring-red-200 focus:border-red-400 bg-red-50'
								: 'border-gray-300 focus:ring-indigo-200 focus:border-indigo-400 hover:border-gray-400'
						]"
						placeholder="Enter the full course content, syllabus, or lesson details..."
					></textarea>
					<div class="flex justify-between items-center mt-2">
						<div v-if="errors.content && touched.content" class="text-red-600 text-sm flex items-center gap-1">
							<svg class="w-4 h-4" fill="currentColor" viewBox="0 0 20 20">
								<path fill-rule="evenodd" d="M18 10a8 8 0 11-16 0 8 8 0 0116 0zm-7 4a1 1 0 11-2 0 1 1 0 012 0zm-1-9a1 1 0 00-1 1v4a1 1 0 102 0V6a1 1 0 00-1-1z" clip-rule="evenodd"/>
							</svg>
							{{ errors.content }}
						</div>
						<div v-else class="text-gray-400 text-sm"></div>
						<span class="text-sm text-gray-400">
							{{ getContentCharCount() }} characters
						</span>
					</div>
				</div>

				<!-- Status Field -->
				<div class="mb-8">
					<label class="block text-sm font-semibold text-gray-700 mb-3">
						Course Status <span class="text-red-500">*</span>
					</label>
					<div class="grid grid-cols-2 gap-4">
						<label
							:class="[
								'cursor-pointer relative flex items-center p-4 rounded-lg border-2 transition-all duration-200',
								status === 'draft'
									? 'border-indigo-500 bg-indigo-50'
									: 'border-gray-200 hover:border-gray-300 hover:bg-gray-50'
							]"
						>
							<input
								type="radio"
								v-model="status"
								value="draft"
								class="sr-only"
							/>
							<div class="flex items-center gap-3">
								<div :class="[
									'w-5 h-5 rounded-full border-2 flex items-center justify-center',
									status === 'draft' ? 'border-indigo-500' : 'border-gray-300'
								]">
									<div v-if="status === 'draft'" class="w-2.5 h-2.5 rounded-full bg-indigo-500"></div>
								</div>
								<div>
									<span class="font-semibold text-gray-900">Draft</span>
									<p class="text-sm text-gray-500">Save as draft and edit later</p>
								</div>
							</div>
						</label>
						<label
							:class="[
								'cursor-pointer relative flex items-center p-4 rounded-lg border-2 transition-all duration-200',
								status === 'published'
									? 'border-green-500 bg-green-50'
									: 'border-gray-200 hover:border-gray-300 hover:bg-gray-50'
							]"
						>
							<input
								type="radio"
								v-model="status"
								value="published"
								class="sr-only"
							/>
							<div class="flex items-center gap-3">
								<div :class="[
									'w-5 h-5 rounded-full border-2 flex items-center justify-center',
									status === 'published' ? 'border-green-500' : 'border-gray-300'
								]">
									<div v-if="status === 'published'" class="w-2.5 h-2.5 rounded-full bg-green-500"></div>
								</div>
								<div>
									<span class="font-semibold text-gray-900">Published</span>
									<p class="text-sm text-gray-500">Publish immediately for students</p>
								</div>
							</div>
						</label>
					</div>
				</div>

				<!-- Action Buttons -->
				<div class="flex items-center justify-end gap-4 pt-6 border-t border-gray-100">
					<button
						type="button"
						@click="router.push('/dashboard/courses')"
						class="px-6 py-3 text-gray-700 font-medium rounded-lg border border-gray-300 hover:bg-gray-50 transition-colors duration-200"
					>
						Cancel
					</button>
					<button
						type="submit"
						:disabled="isPending"
						class="px-8 py-3 bg-indigo-600 text-white font-medium rounded-lg hover:bg-indigo-700 disabled:opacity-50 disabled:cursor-not-allowed transition-colors duration-200 flex items-center gap-2"
					>
						<svg v-if="isPending" class="animate-spin w-5 h-5" fill="none" viewBox="0 0 24 24">
							<circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
							<path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
						</svg>
						{{ isPending ? 'Updating...' : 'Update Course' }}
					</button>
				</div>
			</form>
		</div>
	</div>
</template>
