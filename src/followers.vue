<template>
	<div>
		<h1>Followers</h1>
		<button @click="getTweets">getTweets</button>
	</div>
</template>
<script>
import axios from 'axios'
	
const base = "http://localhost:7777/followers/"

export default{
	name:"main",
	data(){
		return {
			tweets:[],
		}
	},
	methods:{
		getTweets(){
			axios.get(base)
				.then(response => {
					this.tweets = response.data
					console.log(response.data)
				}).catch(error => {
					console.log(error)
				})
		}
	}
}

</script>
<style>

</style>
