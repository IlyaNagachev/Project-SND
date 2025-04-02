<script setup lang="ts">

import {mdiMagnify} from "@mdi/js";
</script>

<template>
	<div class="footer">
		<h1>© 2025  Brand  All Rights Reserved.</h1>
		<div class="footer__button">
			<v-col cols="auto">
				<a href=»tel:1234567890″ target="_blank"><v-btn icon="mdi-phone"></v-btn></a>
			</v-col>
			<v-col cols="auto">
				<a href="https://gmail.com" target="_blank"><v-btn icon="mdi-email-edit-outline"></v-btn></a>
			</v-col>
			<v-col cols="auto">
				<a href="https://t.me" target="_blank"><v-btn icon="mdi-send-outline" ></v-btn></a>
			</v-col>
		</div>
	</div>
</template>

<style scoped lang="scss">

.footer{
	display:flex;
	margin-top: 1vh;
	justify-content: space-between;
	align-content: center;
	height: 20vh;
	border-radius: 50px 50px 0 0;
	box-shadow: 0px 0px 30px 0px rgba(0, 0, 0, 0.12);
	background: white;
	opacity: 0.6;
	h1{
		margin-top: 7vh;
		margin-left: 15vw;
		font-size: 20px;
	}
	&__button{
		display: flex;
		flex-direction: row;
		margin-top: 5vh;
		margin-right: 15vh;
		button{
			border-radius: 0 !important;
		}
	}
}
</style>
