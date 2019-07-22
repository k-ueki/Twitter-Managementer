<template>
	<div>
		<h1>Followers</h1>
		<button @click="followedStatus">FollowedStatus</button>
		<button @click="getFollowers">GetFollowers'List</button>

		<div class="dispAreaWrapper">
			<div class="news">
				<h3>New!</h3>
				<div v-for="newfollower in newfollowers">
					<!--
					{{newfollower}}
					-->
					<img :src="newfollower.ProfileImageUrl">	
					<a v-bind:href="'https://twitter.com/' + newfollower.ScreenName">{{ newfollower.Name }}</a><br/>
					<br/>
					<button class="">+</button>
					<button class="">x</button>
				</div>
			</div>
			<div class="byes">
				<h3>Bye!</h3>
				<div v-for="byefollower in byefollowers">
					<!--
					{{byefollower}}
					-->
					<img :src="byefollower.ProfileImageUrl">	
					<a v-bind:href="'https://twitter.com/' + byefollower.ScreenName">{{ byefollower.Name }}</a><br/>
					<br/>
					<button class="">+</button>
					<button class="">x</button>
				</div>
			</div>
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
			newfollowers:[],
			byefollowers:[],
		}
	},
	methods:{
		followedStatus(){
			var params = new URLSearchParams();
			params.append("mode","register");
			axios.post(base,params)
				.then(response => {
					this.newfollowers = response.data[0].Users
					this.byefollowers = response.data[1].Users
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
.dispAreaWrapper{

}
.news{
	width:40%;
	margin-left:10%;
	float:left;	
	border:2px black solid;
}
.byes{
	width:40%;
	margin-right:10%;
	float:right;
	border:2px black solid;
}
</style>
