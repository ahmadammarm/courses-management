<script setup lang="ts">

import { useLogout } from "@/composables/use-auth";
import DashboardLayout from '@/components/DashboardLayout.vue'

const logout = useLogout();

</script>

<template>
	<DashboardLayout>
		<template #actions>
			<button @click="logout" class="px-4 py-2 bg-red-500 text-white rounded-lg hover:bg-red-600 transition duration-300">
				Logout
			</button>
		</template>

		<h1 class="text-3xl font-bold mb-4">Welcome to the Dashboard</h1>
		<p class="text-lg mb-6">This is your dashboard where you can manage your courses and view your progress.</p>
	</DashboardLayout>
</template>

<style scoped></style>