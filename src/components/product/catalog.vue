<template>
	<Header/>
	<div>
		<section style="height: 120vh; margin-top: 1vh; border-radius: 35px; box-shadow: 0px 0px 30px 0px rgba(0, 0, 0, 0.12); background: white; opacity: 0.6">
			<div style="display:none">
				{{selectClass}}
			</div>
			<v-virtual-scroll
				:height="1200"
				:items="products"
			>
				<template class="cardList" v-slot:default="{ item }">
					<router-link to="./product" @click="productStore.SelectedProducts=item">
						<div class="card">
							<img :src="item.img" style="margin-top: 5vh; height: 20vh">
							<h2>{{item.name}}</h2>
							<h3>{{item.price}}</h3>
						</div>
					</router-link>

				</template>
			</v-virtual-scroll>
		</section>
	</div>
	<Footer/>

</template>

<script setup lang="ts">
import Header from "../general/Header.vue";
import Footer from "../general/Footer.vue";
import {inject, onMounted, ref} from "vue";
import {Product} from "../../models/requests";
import {ProductStore} from "../../store/store";

const selectClass = ref<string>(inject<string>("selectedClass")||"")


const products = ref<Product[]>([])
const productStore = ProductStore()

onMounted(()=> {
	console.log(productStore.Products)
	console.log(selectClass.value)
	productStore.Products.forEach((value) => {
		if (value.class == selectClass.value){
			products.value.push(value);
		}
	})
	console.log(products.value)
})




</script>


<style scoped lang="scss">
.cardList{

}
.card{
	color: #2e3136;
	display: flex;
	flex-direction: column;
	align-items: center;
	height: 30vh;
	gap: 2vh;
	margin-bottom: 10vh;
}
</style>
