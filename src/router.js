import Vue from 'vue'
import VueRouter from 'vue-router'

import Index from '@/components/index.vue'
import Followers from '@/components/followers.vue'
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
