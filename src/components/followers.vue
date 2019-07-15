<template>
	<div>
		<h1>Followers</h1>
		<button @click="getFollowers">getFollowers</button>
		<div v-for="follower in followers">
			<img :src="follower.profile_image_url">	
			{{ follower.name }}<br/><br/></div>
	</div>
</template>
<script>
import axios from 'axios'
	
const base = "http://localhost:7777/followers/"

export default{
	name:"main",
	data(){
		return {
			followers:[],
		}
	},
	methods:{
		getFollowers(){
			axios.get(base)
				.then(response => {
					this.followers = response.data.users
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
