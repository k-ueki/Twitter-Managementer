import Vue from 'vue'
import VueRouter from 'vue-router'

import Index from '@/index.vue'
import Followers from '@/followers.vue'
import HelloWorld from '@/components/HelloWorld.vue'

Vue.use(VueRouter)

const router = new VueRouter({
	mode:'history',
	routes:[
		{
			path:'',
			name:'Index',
			component:Index
		},
		{
			path:'/followers/',
			name:"Followers",
			component:Followers
		}
	]
})

export default router
