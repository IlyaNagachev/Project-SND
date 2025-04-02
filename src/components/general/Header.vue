<template>
	<div class="head">
		<a href="/main" class="button-style" style="margin-top: 2vh; color: #2e3136; text-decoration: 0">
			<h1>SND</h1>
		</a>
		<v-text-field style="width: 30vw; height: 5vh; margin-top: 2vh; margin-left: 1vw" label="Поиск" variant="solo"></v-text-field>
		<v-btn style="width: 55px; height: 56px; background: rgb(0, 100, 210); margin-top: 2vh; margin-right: 1vw">
			<template v-slot:prepend>
				<svg-icon style="margin-left: 10px" type="mdi" :path="mdiMagnify"></svg-icon>
			</template>
		</v-btn>
		<div class="menuButton">
			<v-col cols="auto">
				<v-btn icon="mdi-heart" size="large"></v-btn>
			</v-col>
			<v-col cols="auto">
				<v-btn icon="mdi-cart-outline" size="large"></v-btn>
			</v-col>
			<v-menu>
				<template v-slot:activator="{ props }">
					<v-btn icon="mdi-menu" size="large"  v-bind="props"></v-btn>
				</template>
				<v-list>
					<v-list-item
						v-for="(item, index) in items"
						:key="index"
						:value="index"
					>
						<v-list-item-title>{{ item.title }}</v-list-item-title>
					</v-list-item>
				</v-list>
			</v-menu>
			<v-menu activator="#menu-activator">
				<v-list>
					<v-list-item
						v-for="(item, index) in items"
						:key="index"
						:value="index"
					>
						<v-list-item-title>{{ item.title }}</v-list-item-title>
					</v-list-item>
				</v-list>
			</v-menu>
		</div>

	</div>
</template>

<style scoped lang="scss">
.head{
	display:flex;
	justify-content: flex-start;
	align-content: center;
	height: 10vh;
	border-radius: 100px;
	box-shadow: 0px 0px 30px 0px rgba(0, 0, 0, 0.12);
	background: white;
	opacity: 0.6;
	h1{
		display: flex;
		align-items: center;
		margin-left: 20px
	}
}
.menuButton{
	display: flex;
	align-items: center;
	gap: 0.5vw;
	margin-right: 1vw;
}
</style>
<script setup lang="ts">
import {mdiMagnify} from "@mdi/js";
import SvgIcon from '@jamescoyle/vue-icon';

interface Item {
	title: string;
}

const items: Item[] = [
	{ title: 'Профиль' },
	{ title: 'Мои заказы' },
	{ title: 'Каталог' },
	{ title: 'Поддержка' },
];
</script>
