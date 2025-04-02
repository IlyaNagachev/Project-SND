<template>
	<Header/>
	<div class="productPage">
		<section class="productPage__photo">
			<v-carousel style="height: 95%; border-radius: 35px"
						show-arrows="hover"
						cycle
						hide-delimiter-background
						:interval="30000"
			>
				<v-carousel-item ><v-img style="width: 100%" :src="product?.img"></v-img></v-carousel-item>
				<v-carousel-item><v-img style="width: 100%" :src="product?.img"></v-img></v-carousel-item>
				<v-carousel-item><v-img style="width: 100%" :src="product?.img"></v-img></v-carousel-item>
			</v-carousel>
		</section>
		<div class="productPage__down">
			<div class="productPage__down__buy">
				<v-col cols="auto">
					<v-btn style="border-radius: 2vh; height: 6vh" icon="mdi-heart" size="large"></v-btn>
				</v-col>
				<h1>{{product?.price}}</h1>
				<v-btn style="border-radius: 2vh; height: 6vh" >Купить</v-btn>
			</div>
			<div class="productPage__down__info">
				<h1>{{product?.name}}</h1>
				<p>{{product?.description}}</p>
				<h2>{{product?.specs}}</h2>
			</div>
		</div>
	</div>
	<Footer/>
</template>

<style scoped lang="scss">
.productPage{
	margin-top: 1vh;
	border-radius: 35px;
	box-shadow: 0px 0px 30px 0px rgba(0, 0, 0, 0.12);
	background: white;
	&__photo{
		height: 83.5vh;
		opacity: 0.6;
	}
	&__down{
		display: flex;
		align-items: center;
		justify-content: center;
		flex-direction: column;
		gap: 10vh;
		margin-bottom: 5vh;
		&__buy{
			display: flex;
			align-items: center;
			gap: 10vw;

		}
		&__info{
			display: flex;
			flex-direction: column;
			align-items: center;
			gap: 7vh;
		}
	}

}
</style>

<script setup lang="ts" xmlns="http://www.w3.org/1999/html">

import Header from "../general/Header.vue";
import Footer from "../general/Footer.vue";
import CatalogMenu from "./catalogMenu.vue";
import OldpcCover from "./oldpcCover.vue";
import TradeInApplication from "./tradeInApplication.vue";
import {onMounted, onUnmounted, ref} from "vue";
import {Product} from "../../models/requests";
import {ProductStore} from "../../store/store";

const product = ref<Product>()
const userStore = ProductStore()

onMounted(()=>{
	product.value = userStore.SelectedProducts
	console.log(product.value)
})
</script>

