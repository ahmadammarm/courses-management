<script setup lang="ts">

import { ref, reactive } from "vue";
import { useRouter } from "vue-router";
import { useLogin } from "@/composables/use-auth";
import Cookies from "js-cookie";

interface ValidationErrors {
	[key: string]: string;
}

const router = useRouter();

const { mutate, isPending } = useLogin();

const email = ref<string>('');
const password = ref<string>('');

const errors = reactive<ValidationeErrors>({});

const handleLogin = (e: Event) => {
	e.preventDefault();

	mutate({ email: email.value, password: password.value }, {
		onSuccess: (data: any) => {

			Cookies.set('token', data.data.token)
			Cookies.set('user', JSON.stringify({
				id: data.data.id,
				name: data.data.name,
				username: data.data.username,
				email: data.data.email,
				expertise: data.data.expertise,
			}))

			router.push("/dashboard");
		},
		onError: (error: any) => {
			Object.assign(errors, error.response.data.errors);
		},
	});
}

</script>

<template>
	<div class="flex justify-center mt-20">
		<div class="w-full max-w-md">
			<div class="bg-white rounded-2xl shadow-md p-8">
				<h4 class="text-2xl font-bold text-center mb-4">LOGIN</h4>
				<hr class="mb-4" />
				<div v-if="errors.Error" class="bg-red-100 text-red-700 p-3 rounded-2xl mb-4">
					Username or Password is incorrect
				</div>
				<form @submit="handleLogin">
					<div class="mb-4">
						<label class="block mb-2 font-bold text-gray-700">Email</label>
						<input v-model="email" type="text"
							class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500"
							placeholder="Email" />
						<div v-if="errors.Email" class="bg-red-100 text-red-700 p-3 rounded-2xl mt-2">
							{{ errors.Email }}
						</div>
					</div>

					<div class="mb-6">
						<label class="block mb-2 font-bold text-gray-700">Password</label>
						<input v-model="password" type="password"
							class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500"
							placeholder="Password" />
						<div v-if="errors.Password" class="bg-red-100 text-red-700 p-3 rounded-2xl mt-2">
							{{ errors.Password }}
						</div>
					</div>

					<button type="submit"
						class="w-full bg-blue-600 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded-2xl disabled:opacity-50"
						:disabled="isPending">
						{{ isPending ? 'Loading...' : 'LOGIN' }}
					</button>
				</form>
			</div>
		</div>
	</div>
</template>

<style scoped></style>