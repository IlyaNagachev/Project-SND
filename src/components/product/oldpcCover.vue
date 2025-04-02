<template>
<div class="oldpcCover">
	<h1>PC КАК НОВЫЕ</h1>
	<div class="oldpcCover__variants">
		<v-btn style="height: 40vh; width: 15vw; margin-right: 2vw">
			<v-img width="300" height="300" src="../../assets/images/oldpc/1.png"></v-img>
			<p>i3 rx580</p>
		</v-btn>
		<v-btn style="height: 40vh; width: 15vw; margin-right: 2vw">
			<v-img width="300" height="300" src="../../assets/images/oldpc/2.png"></v-img>
			<p>i5 rtx2060</p>
		</v-btn>
		<v-btn style="height: 40vh; width: 15vw; margin-right: 2vw">
			<v-img width="250" height="300" src="../../assets/images/oldpc/3.png"></v-img>
			<p>ryzen 5 gtx1060</p>
		</v-btn>
		<v-btn style="height: 40vh; width: 15vw">
			<v-img width="300" height="300" src="../../assets/images/oldpc/2.png"></v-img>
			<p>ryzen 5 rtx2060</p>
		</v-btn>
	</div>
	<v-btn style="height: 5vh; width: 13vw">
		<p>Посмотреть все варианты</p>
	</v-btn>

</div>
</template>


<script setup lang="ts">

import {mdiMagnify} from "@mdi/js";
</script>



<style lang="scss">
.oldpcCover{
	display: flex;
	flex-direction: column;
	align-items: center;
	gap: 4vw;
	button{
	}
	.v-btn__content{
		display:flex;
		flex-direction:column;
		gap: 1vh;
	}
}


</style>
