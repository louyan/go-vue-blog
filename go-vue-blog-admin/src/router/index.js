import Vue from 'vue'
import VueRouter from 'vue-router'

Vue.use(VueRouter)

// import HelloWorld from '@/components/HelloWorld.vue'
import Home from '@/pages/Home';
import Login from '@/pages/Login';
import PostList from '@/pages/PostList';
import Category from '@/pages/Category';
import Tag from '@/pages/Tag';
import Edit from '@/pages/Edit';

const router = new VueRouter({
    routes:[
        {
            path:'/',
            name: 'Home',
            component: Home,
            meta:{
                paths:[
                    {
                        name:'首页',
                        link:''
                    }
                ],
                requireAuth:true
            }
        },
        {
            path: '/login',
            name: 'Login',
            component: Login
        },
        {
            path:'/postlist',
            name: 'PostList',
            component: PostList,
            meta:{
                paths:[
                    {
                        name:'首页',
                        link:'/'
                    },
                    {
                        name:'文章列表',
                        link:'/'
                    }
                ],
                requireAuth:true
            }
        },
        {
            path: '/edit',
            name: 'Add',
            component: Edit,
            meta: {
                paths: [
                    {
                        name: '首页',
                        link: '/'
                    },
                    {
                        name: '新增文章',
                        link: ''
                    }
                ],
                requireAuth: true
            },
            children: [{
                path: ':id',
                name: 'Edit',
                component: Edit,
                meta: {
                    paths: [
                        {
                            name: '首页',
                            link: '/'
                        },
                        {
                            name: '编辑文章',
                            link: ''
                        }
                    ],
                    requireAuth: true
                }
            }]
        },
        {
            path: '/category',
            name: 'Category',
            component: Category,
            meta: {
                paths: [
                    {
                        name: '首页',
                        link: '/'
                    },
                    {
                        name: '分类管理',
                        link: ''
                    }
                ],
                requireAuth: true
            }
        },
        {
            path: '/tag',
            name: 'Tag',
            component: Tag,
            meta: {
                paths: [
                    {
                        name: '首页',
                        link: '/'
                    },
                    {
                        name: '标签管理',
                        link: ''
                    }
                ],
                requireAuth: true
            }
        },

    ],
});

// router.beforeEach((to, from, next) => {
//     if (to.meta.requireAuth) {
//         const token = localStorage.getItem('BLOG_TOKEN');
//         if (token && token !== 'null') {
//             next();
//         } else {
//             next('/login');
//         }
//     } else {
//         next();
//     }
// });

export default router;