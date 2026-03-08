<script setup lang="ts">
import { useAuthUser } from "@/composables/use-auth";

const user = useAuthUser();
</script>

<template>
	<div>
		<h1 class="text-3xl font-bold text-gray-900 mb-6">My Profile</h1>

		<div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
			<div class="flex items-center gap-6">
				<div class="w-20 h-20 bg-indigo-100 rounded-full flex items-center justify-center">
					<span class="text-2xl font-bold text-indigo-600">
						{{ user?.name?.charAt(0).toUpperCase() || 'U' }}
					</span>
				</div>
				<div>
					<h2 class="text-xl font-semibold text-gray-900">{{ user?.name || 'User' }}</h2>
					<p class="text-gray-500">@{{ user?.username || 'username' }}</p>
				</div>
			</div>

			<div class="mt-6 pt-6 border-t border-gray-100 grid grid-cols-1 md:grid-cols-2 gap-6">
				<div>
					<label class="block text-sm font-medium text-gray-500 mb-1">Email</label>
					<p class="text-gray-900">{{ user?.email || '-' }}</p>
				</div>
				<div>
					<label class="block text-sm font-medium text-gray-500 mb-1">Expertise</label>
					<p class="text-gray-900">{{ user?.expertise || '-' }}</p>
				</div>
			</div>
		</div>
	</div>
</template>
