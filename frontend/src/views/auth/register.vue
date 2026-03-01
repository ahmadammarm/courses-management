<script setup lang="ts">

import { ref, reactive } from "vue";

import { useRouter } from "vue-router";
import { useRegister } from "@/composables/use-auth";


interface ValidationErrors {
	[key: string]: string;
}

const router = useRouter();

const name = ref<string>('');
const username = ref<string>('');
const email = ref<string>('');
const password = ref<string>('');
const expertise = ref<string>('');

const errors = reactive<ValidationErrors>({});

const { mutate, isPending } = useRegister();

const handleRegister = (e: Event) => {
	e.preventDefault();

	mutate(
		{
			name: name.value,
			username: username.value,
			email: email.value,
			password: password.value,
			expertise: expertise.value,
		},
		{
			onSuccess: () => {

				alert("Registration successful! Please login to continue.");
				router.push("/login");
			},
			onError: (error: any) => {

				Object.assign(errors, error.response.data.errors);
			},
		}
	);
};
</script>

<template>
	<div class="flex justify-center items-center min-h-screen">
		<div class="w-full max-w-md">
			<div class="bg-white rounded-lg shadow-md p-8">
				<h4 class="text-2xl font-bold text-center mb-6">REGISTER</h4>
				<hr class="mb-6" />
				<form @submit.prevent="handleRegister">
					<div class="grid grid-cols-2 gap-4 mb-4">
						<div>
							<label class="block text-sm font-bold mb-2">Full Name</label>
							<input v-model="name" type="text"
								class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500"
								placeholder="Full Name" />
							<div v-if="errors.Name" class="text-red-600 text-sm mt-2 p-2 bg-red-50 rounded-lg">
								{{ errors.Name }}
							</div>
						</div>
						<div>
							<label class="block text-sm font-bold mb-2">Username</label>
							<input v-model="username" type="text"
								class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500"
								placeholder="Username" />
							<div v-if="errors.Username" class="text-red-600 text-sm mt-2 p-2 bg-red-50 rounded-lg">
								{{ errors.Username }}
							</div>
						</div>
					</div>

					<div class="grid grid-cols-2 gap-4 mb-6">
						<div>
							<label class="block text-sm font-bold mb-2">Email address</label>
							<input v-model="email" type="email"
								class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500"
								placeholder="Email Address" />
							<div v-if="errors.Email" class="text-red-600 text-sm mt-2 p-2 bg-red-50 rounded-lg">
								{{ errors.Email }}
							</div>
						</div>
						<div>
							<label class="block text-sm font-bold mb-2">Password</label>
							<input v-model="password" type="password"
								class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500"
								placeholder="Password" />
							<div v-if="errors.Password" class="text-red-600 text-sm mt-2 p-2 bg-red-50 rounded-lg">
								{{ errors.Password }}
							</div>
						</div>
					</div>

					<div class="grid grid-cols-1 gap-4 mb-6">
						<div>
							<label class="block text-sm font-bold mb-2">Expertise</label>
							<input v-model="expertise" type="text"
								class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500"
								placeholder="Expertise" />
							<div v-if="errors.Expertise" class="text-red-600 text-sm mt-2 p-2 bg-red-50 rounded-lg">
								{{ errors.Expertise }}
							</div>
						</div>
					</div>

					<button type="submit"
						class="w-full bg-blue-600 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded-lg disabled:opacity-50 disabled:cursor-not-allowed transition-colors"
						:disabled="isPending">
						{{ isPending ? 'Loading...' : 'REGISTER' }}
					</button>
				</form>
			</div>
		</div>
	</div>
</template>
