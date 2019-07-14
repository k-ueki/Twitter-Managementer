<template>
  <div id="app">
	<router-view></router-view>
  </div>
</template>

<script>
import HelloWorld from './components/HelloWorld.vue'
import Main from './main.vue'

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
		Main
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
