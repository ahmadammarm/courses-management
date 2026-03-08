import { createRouter, createWebHistory, type RouteRecordRaw } from "vue-router";
import Cookies from 'js-cookie';
import { isTokenExpired } from '@/lib/session';

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
	{
		path: "/dashboard",
		name: "Dashboard",
		component: () => import("../views/home/dashboard.vue"),
		children: [
			{
				path: "",
				name: "DashboardOverview",
				component: () => import("../views/dashboard/overview.vue"),
			},
			{
				path: "courses",
				name: "MyCourses",
				component: () => import("../views/dashboard/courses.vue"),
			},
			{
				path: "create",
				name: "CreateCourse",
				component: () => import("../views/dashboard/create-course.vue"),
			},
			{
				path: "courses/:id/edit",
				name: "EditCourse",
				component: () => import("../views/dashboard/edit-course.vue"),
			},
			{
				path: "courses/:id",
				name: "ViewCourse",
				component: () => import("../views/dashboard/view-course.vue"),
			},
			{
				path: "profile",
				name: "Profile",
				component: () => import("../views/dashboard/profile.vue"),
			},
		],
	},
];

const router = createRouter({
	history: createWebHistory(),
	routes,
});

// Navigation guard: protect /dashboard and its subpaths when not authenticated
router.beforeEach((to, next) => {
	const token = Cookies.get('token');
	const loggedIn = !!token && !isTokenExpired(token);

	// Protect dashboard routes (including nested routes)
	if (to.path.startsWith('/dashboard')) {
		if (!loggedIn) {
			Cookies.remove('token');
			Cookies.remove('user');
			return next({ name: 'Login' });
		}
	}

	// Prevent accessing login/register when already authenticated
	if (to.path === '/login' || to.path === '/register') {
		if (loggedIn) {
			return next({ name: 'Dashboard' });
		}
	}

	return next();
});

export default router;