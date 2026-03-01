import { createRouter, createWebHistory, type RouteRecordRaw } from "vue-router";

const routes: RouteRecordRaw[] = [
	{
		path: "/",
		name: "Home",
		component: () => import("../views/home/index.vue"),
	},
	{
		path: "/login",
		name: "Login",
		component: () => import("../views/auth/login.vue"),
	},
	{
		path: "/register",
		name: "Register",
		component: () => import("../views/auth/register.vue"),
	},
];

const router = createRouter({
	history: createWebHistory(),
	routes,
});

export default router;