<template>
  <div id="app">
	<router-view></router-view>
  </div>
</template>

<script>
import HelloWorld from './components/HelloWorld.vue'
import Index from './index.vue'
import Followers from './followers.vue'

import axios from 'axios'
import Router from 'vue-router'

export default {
  name: 'app',
	data(){
		return{
		}
	},
	methods:{
	},components:{
		Index,
		Followers,
		HelloWorld
	}
}
</script>

<style>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
