<template>
	<div>
		<h1>Followers</h1>
		<button @click="initFollowers">Init Followers</button>
		<button @click="followedStatus">FollowedStatus</button>
		<button @click="getFollowers">GetFollowers'List</button>

		<div class="dispAreaWrapper">
			<div class="news">
				<h3>New!</h3>
				<div v-for="newfollower in newfollowers">
					<img :src="newfollower.profile_image_url"><br/>
					<a v-bind:href="'https://twitter.com/' + newfollower.screen_name" target="_blank">{{ newfollower.name }}</a>
					<br/>
					<button class="">+</button>
					<button class="">x</button>
					<br/><br/>
				</div>
			</div>
			<div class="byes">
				<h3>Bye!</h3>
				<div v-for="(byefollower,bye) in byefollowers" :key="bye">
					<img :src="byefollower.profile_image_url"><br/>
					<a v-bind:href="'https://twitter.com/' + byefollower.screen_name" target="_blank">{{ byefollower.name }}</a>
					<br/>
					<button class="">+</button>
					<button class="">x</button>
					<br/><br/>
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
				})
				// 	.catch(error => {
				// 	// console.log(error)
				// })
		},
		getFollowers(){
			axios.get(base)
				.then(resp=>{
					this.followers = resp.data.users
				})
				// 	.catch(err=>{
				// 	// console.log(err)
				// })
		},
		initFollowers(){
			if (confirm("init for sure?")){
				axios.put(base+"/init")
					.then(resp=>{
						alert("success for initialize")
					}).catch(err=>{
						alert("failed to initialize")
					})
            }
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
	/* border:2px black solid; */
	color:lightgreen;
}
.byes{
	width:40%;
	margin-right:10%;
	float:right;
	/* border:2px black solid; */
	color:red;
}
	img{
        border-radius: 50%;
	}
</style>
