<template>
	<Header/>
	<div>
		<section style="height: 83.5vh; margin-top: 5px; border-radius: 35px; box-shadow: 0px 0px 30px 0px rgba(0, 0, 0, 0.12); background: white; opacity: 0.6;">
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
							<Gdfht/>
						</div>
					</v-sheet>
				</v-carousel-item>
			</v-carousel>
		</section>
	</div>
	<div class="border-xl !important" style="height: 100%; width: 100%; display: flex; align-items: center; justify-content: center;">
		<div style="border: 5px black solid; height: 50vh; width: 50vw; display: flex; align-items: center; justify-content: center;">
			<v-btn @click="userStore.Registration(log, passwd)" class="pr-10 !important" size="x-large" variant="outlined" width="200">Button</v-btn>
			<v-btn  @click="print" size="x-large" variant="outlined" width="200">Button</v-btn>
			<v-text-field
                v-model="log"
				name="name"
				label="label"
				id="id"
            />
            <v-text-field
                v-model="passwd"
                name="name"
                label="label"
                id="id"
            />
		</div>
	</div>


</template>

<script setup lang="ts">
import SvgIcon from '@jamescoyle/vue-icon';
import { mdiMagnify } from '@mdi/js';

import { ref } from 'vue';
import {UserStore} from "../store/store";
import Header from "./Header.vue";
import {red} from "vuetify/util/colors";
import Gdfht from "./Gdfht.vue";

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


const userStore = UserStore()
</script>
