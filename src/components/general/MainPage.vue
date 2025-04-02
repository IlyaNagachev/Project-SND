<template>
	<Header/>
	<div>
		<section style="height: 83.5vh; margin-top: 1vh; border-radius: 35px; box-shadow: 0px 0px 30px 0px rgba(0, 0, 0, 0.12); background: white; opacity: 0.6;">
			<v-carousel style="height: 100%; border-radius: 35px"
				show-arrows="hover"
				cycle
				hide-delimiter-background
						:interval="30000"
			>
				<v-carousel-item
					v-for="(slide, i) in slides"
					:key="i"
				>
					<v-sheet
						height="100%"
					>
						<div v-if="i==0" class="d-flex fill-height justify-center align-center">
							<GlobalCatalog/>
						</div>
						<div v-if="i==1" class="d-flex fill-height justify-center align-center">
							<oldpc-cover/>
						</div>
						<div v-if="i==2" class="d-flex fill-height justify-center align-center">
							<trade-in-application/>
						</div>
					</v-sheet>
				</v-carousel-item>
			</v-carousel>
		</section>
	</div>
	<Footer/>

</template>

<script setup lang="ts">
import { mdiMagnify } from '@mdi/js';

import {onMounted, ref} from 'vue';
import {ProductStore} from "../../store/store";
import Header from "./Header.vue";
import {red} from "vuetify/util/colors";
import CatalogMenu from "../product/catalogMenu.vue";
import Footer from "./Footer.vue";
import OldpcCover from "../product/oldpcCover.vue";
import TradeInApplication from "../product/tradeInApplication.vue";
import GlobalCatalog from "../product/GlobalCatalog.vue";

const log = ref<string>("")
const passwd = ref<string>("")

const print = () => {
    console.log(log.value);
    console.log(passwd.value);
}


const slides = [
		0,
		1,
		2,
	]


const userStore = ProductStore()

onMounted(()=>{
	userStore.GetProducts()
})
</script>
