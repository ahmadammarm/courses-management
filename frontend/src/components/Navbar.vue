<script setup lang="ts">

import { useRouter } from "vue-router";


const router = useRouter();

const handleRedirect = (path: string) => {
	router.push(path);
}
</script>

<template>
	<header class="bg-white shadow-md">
		<div class="max-w-6xl mx-auto px-6 py-4 flex items-center justify-between">
			<div class="flex items-center gap-4">
				<div class="text-2xl font-extrabold text-indigo-600">Courses</div>
				<nav class="hidden md:flex gap-4 text-gray-600">
					<a href="#" class="hover:text-indigo-600">Features</a>
					<a href="#" class="hover:text-indigo-600">Pricing</a>
					<a href="#" class="hover:text-indigo-600">About</a>
				</nav>
			</div>

			<div class="flex items-center gap-3">
				<button @click="handleRedirect('/login')"
					class="px-4 py-2 rounded-lg bg-indigo-600 text-white hover:bg-indigo-700 transition">Login</button>
				<button @click="handleRedirect('/register')"
					class="px-4 py-2 rounded-lg border border-indigo-600 text-indigo-600 hover:bg-indigo-50 transition">Register</button>
			</div>
		</div>
	</header>
</template>



<style scoped></style>
