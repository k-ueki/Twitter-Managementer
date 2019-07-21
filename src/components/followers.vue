<template>
	<div>
		<h1>Followers</h1>
		<button @click="registerFollowers">RegisterFollowers</button>
		<button @click="getFollowers">GetFollowers'List</button>
		<div v-for="follower in followers">
			<img :src="follower.profile_image_url">	
			{{ follower.name }}<br/>{{ follower }}<br/><br/></div>
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
		registerFollowers(){
			var params = new URLSearchParams();
			params.append("mode","register");
			axios.post(base,params)
				.then(response => {
					this.followers = response.data.users
					console.log(response.data)
				}).catch(error => {
					console.log(error)
				})
		},
		getFollowers(){
			axios.get(base)
				.then(resp=>{
					this.followers = resp.data.users
				}).catch(err=>{
					console.log(err)
				})
		}
	}
}

</script>
<style>

</style>
