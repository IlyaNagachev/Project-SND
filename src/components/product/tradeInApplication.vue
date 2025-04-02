<template>
<div class="tradeInApplication">
	<div class="tradeInApplication__form">
		<h1>СДАЙ технику в TRADE-IN.</h1>
		<h2>Оставь заявку, вписав комплектующие или название компьютер с Вашим номером телефон, и мы с Вами свяжемся в ближайшее время.</h2>
		<v-form class="application">
			<v-text-field style="width: 30vw" label="Название техники" variant="solo"></v-text-field>
			<v-text-field style="width: 30vw" label="Номер" variant="solo"></v-text-field>
			<v-btn>
				<template v-slot:prepend>
					<p style="margin-left: 0.5vw">Оставить заявку</p>
				</template>
			</v-btn>
		</v-form>
	</div>

</div>
</template>


<script setup lang="ts">

import {mdiMagnify} from "@mdi/js";
</script>



<style lang="scss">
.tradeInApplication{
	background-image: url("../../assets/images/tradein/cover.png");
	background-size: cover;
	background-position: center;
	width: 100vw;
	height: 100vh;
	&__form{
		display: flex;
		flex-direction: column;
		align-items: center;
		margin-top: 20vh;
		h1{
			color: whitesmoke;
			margin-bottom: 1vh;
			font-size: 30px;
		}
		h2{
			color: whitesmoke;
			margin-bottom: 10vh;
		}
		.application{
			display: flex;
			flex-direction: column;
			align-items: center;
		}
	}
	.v-btn__content{
		display:flex;
		gap: 1vh;
	}
}


</style>
