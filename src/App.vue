<template>
  <div id="app">
	<router-view></router-view>
  </div>
</template>

<script>
import HelloWorld from './components/HelloWorld.vue'
import Index from './components/index.vue'
import Followers from './components/followers.vue'
import Timeline from './components/timeline.vue'

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
		HelloWorld,
		Timeline
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
