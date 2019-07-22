<template>
	<div>
		<h1>Followers</h1>
		<button @click="followedStatus">FollowedStatus</button>
		<button @click="getFollowers">GetFollowers'List</button>
		<div v-for="follower in followers">
			<img :src="follower.profile_image_url">	
			<a v-bind:href="'https://twitter.com/' + follower.screen_name">{{ follower.name }}</a><br/>
			<!--
			{{follower}}<br/><br/>
			-->
			<br/>
		</div>
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
		followedStatus(){
			var params = new URLSearchParams();
			params.append("mode","register");
			axios.post(base,params)
				.then(response => {
					this.followers = response.data
					//this.followers = response.data
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
