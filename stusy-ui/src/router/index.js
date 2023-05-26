import {createRouter, createWebHistory} from 'vue-router';
import HomeView from '../views/HomeView.vue';

const routes = [
    {
        path: '/home',
        name: 'Домашнаяя',
        meta: {layout: 'HomeLayout'},
        component: HomeView,
    },
    {
        path: '/landing',
        name: 'Главная',
        meta: {layout: 'EmptyLayout'},
        component: () => import('../views/MainView.vue'),
    },
    {
        path: '/',
        alias : '/auth',
        name: 'Вход',
        meta: {layout: 'EmptyLayout'},
        component: () => import('../views/AuthView.vue'),
    },
    {
        path: '/profile',
        name: 'Профиль',
        meta: {layout: 'HomeLayout'},
        component: () => import('../views/ProfileView'),
    },
    {
        path: '/courses',
        name: 'Курсы',
        meta: {layout: 'HomeLayout'},
        component: () => import('../views/CourseView.vue'),
    },
    {
        path: '/creatingCourses',
        name: 'Создание курса',
        meta: {layout: 'HomeLayout'},
        component: () => import('../views/CreatingCourses'),
    },
    {
        path: '/creatingTest',
        name: 'Создание теста',
        meta: {layout: 'HomeLayout'},
        component: () => import('../views/CreatingTest'),
    },
    {
        path: '/mycourses',
        name: 'Мои курсы',
        meta: {layout: 'HomeLayout'},
        component: () => import('../views/MyCourses'),
    },
    {
        path: '/:pathMatch(.*)*',
        name: 'Не найдено',
        meta: {layout: 'EmptyLayout'},
        component: () => import('../views/404View.vue'),
    },
];

const router = createRouter({
    history: createWebHistory(process.env.BASE_URL),
    routes,
});

router.beforeEach((to, from, next) => {
    document.title = to.name;
    next();
});

export default router;
