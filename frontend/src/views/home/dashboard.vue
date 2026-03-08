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

		<router-view />
	</DashboardLayout>
</template>

<style scoped></style>