<template>
	<div class="head" style="display:flex; justify-content: flex-start; align-content: center; height: 8vh; border-radius: 100px; box-shadow: 0px 0px 30px 0px rgba(0, 0, 0, 0.12); background: white; opacity: 0.6;">
		<h1 style="margin-top: 12px;margin-left: 20px">SND</h1>
		<v-text-field style="width: 30vw; height: 5vh; margin-top: 12px; margin-left: 20px" label="Поиск" variant="solo"></v-text-field>
		<v-btn style="width: 55px; height: 56px; background: rgb(0, 100, 210); margin-top: 12px; margin-right: 20px">
			<template v-slot:prepend>
				<svg-icon style="margin-left: 10px" type="mdi" :path="mdiMagnify"></svg-icon>
			</template>
		</v-btn>
		<v-col cols="auto">
			<v-btn icon="mdi-heart" size="large"></v-btn>
		</v-col>
		<v-col cols="auto">
			<v-btn icon="mdi-cart-outline" size="large"></v-btn>
		</v-col>
		<v-col cols="auto">
			<v-btn icon="mdi-menu" size="large"></v-btn>
		</v-col>
	</div>
</template>

<style scoped lang="scss">

</style>
<script setup lang="ts">
import {mdiMagnify} from "@mdi/js";
import SvgIcon from '@jamescoyle/vue-icon';
</script>
