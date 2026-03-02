<template>
	<div class="min-h-screen bg-gradient-to-br from-blue-50 to-indigo-100">
		<div>
			<Navbar />
			<div class="flex flex-col md:flex-row items-center justify-between p-8 md:p-12 max-w-6xl mx-auto">
				<div class="flex-1 mb-8 md:mb-0">
					<h1 class="text-4xl md:text-5xl font-bold text-gray-900 mb-4">Welcome to Courses</h1>
					<p class="text-lg text-gray-600 mb-6 leading-relaxed">Manage and explore your courses with ease. Join our
						platform to expand your knowledge.</p>
				</div>
			</div>
		</div>
	</div>
</template>

<script setup lang="ts">
import Navbar from '@/components/Navbar.vue'

</script>

<style scoped></style>