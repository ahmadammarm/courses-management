<template>
  <div class="min-h-screen flex bg-gray-50">
    <aside class="w-72 bg-white border-r hidden md:block">
      <div class="p-6 border-b">
        <div class="text-xl font-bold text-indigo-600">Courses</div>
        <div class="text-sm text-gray-500">Dashboard</div>
      </div>
      <nav class="p-4">
        <ul class="space-y-2">
          <li>
            <router-link to="/dashboard" class="flex items-center gap-3 px-4 py-2 rounded-lg hover:bg-indigo-50">
              <span class="font-medium text-gray-700">Overview</span>
            </router-link>
          </li>
          <li>
            <router-link to="/dashboard/courses" class="flex items-center gap-3 px-4 py-2 rounded-lg hover:bg-indigo-50">
              <span class="font-medium text-gray-700">My Courses</span>
            </router-link>
          </li>
          <li>
            <router-link to="/dashboard/create" class="flex items-center gap-3 px-4 py-2 rounded-lg hover:bg-indigo-50">
              <span class="font-medium text-gray-700">Create Course</span>
            </router-link>
          </li>
          <li>
            <router-link to="/dashboard/profile" class="flex items-center gap-3 px-4 py-2 rounded-lg hover:bg-indigo-50">
              <span class="font-medium text-gray-700">Profile</span>
            </router-link>
          </li>
        </ul>
      </nav>
    </aside>

    <div class="flex-1">
      <div class="bg-white p-4 border-b flex items-center justify-between">
        <div class="flex items-center gap-3">
          <button class="md:hidden px-3 py-1 bg-gray-100 rounded-lg">Menu</button>
          <h2 class="text-lg font-semibold">Dashboard</h2>
        </div>
        <div>
          <slot name="actions" />
        </div>
      </div>

      <main class="p-6">
        <slot />
      </main>
    </div>
  </div>
</template>

<script setup lang="ts"></script>

<style scoped></style>
