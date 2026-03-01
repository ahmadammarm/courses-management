<template>
	<div class="min-h-screen bg-gradient-to-br from-blue-50 to-indigo-100">
		<div class="flex flex-col md:flex-row items-center justify-between p-8 md:p-12 max-w-6xl mx-auto">
			<div class="flex-1 mb-8 md:mb-0">
				<h1 class="text-4xl md:text-5xl font-bold text-gray-900 mb-4">Welcome to Courses</h1>
				<p class="text-lg text-gray-600 mb-6 leading-relaxed">Manage and explore your courses with ease. Join our
					platform to expand your knowledge.</p>
			</div>

			<div class="flex flex-col md:flex-row items-center gap-4 w-full md:w-auto">
				<a href="/login"
					class="w-full md:w-auto px-6 py-3 bg-blue-600 text-white font-semibold rounded-lg hover:bg-blue-700 transition-colors duration-200 text-center">
					Login
				</a>
				<a href="/register"
					class="w-full md:w-auto px-6 py-3 bg-green-600 text-white font-semibold rounded-lg hover:bg-green-700 transition-colors duration-200 text-center">
					Register
				</a>
			</div>
		</div>
	</div>
</template>

<script setup lang="ts">

</script>

<style scoped></style>