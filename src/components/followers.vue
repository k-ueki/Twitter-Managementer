<template>
	<div>
		<h1>Followers</h1>
		<button @click="initFollowers">Init Followers</button>
		<button @click="followedStatus">FollowedStatus</button>
		<button @click="getFollowers">GetFollowers'List</button>

		<div class="dispAreaWrapper">
			<div class="news">
				<h3 style="color:#00FF00;">New!</h3>
				<div v-for="newfollower in newfollowers">
					<!--
					{{newfollower}}
					-->
					<img :src="newfollower.profile_image_url">	
					<a v-bind:href="'https://twitter.com/' + newfollower.screen_name">{{ newfollower.name }}</a><br/>
					<br/>
					<button class="follow" @click="follow">Follow</button>
					<button class="hidden" @click="hidden">Hidden</button>
				</div>
			</div>
			<div class="byes">
				<h3 style="color:red;">Bye!</h3>
				<div v-for="byefollower in byefollowers">
					<!--
					{{byefollower}}
					-->
					<img :src="byefollower.profile_image_url">	
					<a v-bind:href="'https://twitter.com/' + byefollower.screen_name">{{ byefollower.name }}</a><br/>
					<br/>
					<button class="unfollow" @click="unfollow">Unfollow</button>
					<button class="hidden" @click="hidden">Hidden</button>
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
		initFollowers(){
			if(confirm("init your followers for sure?")){
				if(confirm("for sure?")){
					var params = new URLSearchParams();
					params.append("mode","init");
					axios.post(base,params)
						.then(response=>{
										
						}).catch(error=>{

						})
				}
			}
		},
		followedStatus(){
			var params = new URLSearchParams();
			params.append("mode","register");
			axios.post(base,params)
				.then(response => {
					this.newfollowers = response.data[0].users
					this.byefollowers = response.data[1].users
					console.log(response.data)
					console.log(response)
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
		},
		follow(){
			console.log("follow")
		},
		unfollow(){
			if(confirm("Sure?")){
				console.log("unfollow")
			}
		},
		hidden(){
			console.log("hidden")
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
}
.byes{
	width:40%;
	margin-right:10%;
	float:right;
}
</style>
