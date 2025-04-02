<template>
<div class="catalogMenu">
	<router-link to="/catalog">
		<div class="button container">
			<v-btn @click="openMenu = true; classProduct = 'Материнская плата'" style="height: 25vh; width: 13vw; margin: 1vw">
				<v-img width="100" height="100" src="../../assets/images/catalog/materinskaja_plata.png"></v-img>
				<p>Материнская</p>
				<p>плата</p>
			</v-btn>
			<v-btn @click="openMenu = true; classProduct = 'Видеокарта'" style="height: 25vh; width: 13vw; margin: 1vw">
				<v-img width="150" height="150" src="../../assets/images/catalog/graphic-card.png"></v-img>
				<p>Видеокарта</p>
			</v-btn>
			<v-btn @click="openMenu = true; classProduct = 'Оперативная память'" style="height: 25vh; width: 13vw; margin: 1vw">
				<v-img width="150" height="100" src="../../assets/images/catalog/ram.png"></v-img>
				<p>Оперативная</p>
				<p>память</p>
			</v-btn>
			<v-btn @click="openMenu = true; classProduct = 'Накопители'" style="height: 25vh; width: 13vw; margin: 1vw">
				<v-img width="100" height="100" src="../../assets/images/catalog/hdd.png"></v-img>
				<p>Накопители</p>
			</v-btn>
			<v-btn @click="openMenu = true; classProduct = 'Блок питания'" style="height: 25vh; width: 13vw; margin: 1vw">
				<v-img width="100" height="100" src="../../assets/images/catalog/psu.png"></v-img>
				<p>Блок питания</p>
			</v-btn>
			<v-btn @click="openMenu = true; classProduct = 'Охлаждение'" style="height: 25vh; width: 13vw; margin: 1vw">
				<v-img width="100" height="100" src="../../assets/images/catalog/cooler.png"></v-img>
				<p>Охлаждение</p>
			</v-btn>
			<v-btn @click="openMenu = true; classProduct = 'Корпус'" style="height: 25vh; width: 13vw; margin: 1vw">
				<v-img width="100" height="100" src="../../assets/images/catalog/case.png"></v-img>
				<p>Корпус</p>
			</v-btn>
			<v-btn @click="openMenu = true; classProduct = 'Остальное'" style="height: 25vh; width: 13vw; margin: 1vw">
				<v-img width="100" height="100" src="../../assets/images/catalog/other.png"></v-img>
				<p>Остальное</p>
			</v-btn>
		</div>
	</router-link>
</div>
</template>


<script setup lang="ts">
import {inject, provide, ref} from "vue";
import {mdiMagnify} from "@mdi/js";

const openMenu = inject<boolean>("open")
const classProduct = inject<string>("selectedClass")
</script>



<style lang="scss">
.catalogMenu {
	display: flex;
	max-width: 70vw;
	flex-direction: column;
	align-items: center;
	justify-content: center;
	box-sizing: border-box;
	margin-left: 10vw;
}

.v-btn__content {
	color: #2e3136;
	display: flex;
	flex-direction: column;
	gap: 1vh;
}



</style>
