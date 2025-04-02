<template>

	<catalogMenu v-if="!openMenu "/>
	<Catalog v-if="openMenu"/>
</template>

<script setup lang="ts">
import {provide, ref} from "vue";
import CatalogMenu from "./catalogMenu.vue";
import Catalog from "./catalog.vue";


const selectClass = ref<string>("")
provide("selectedClass", selectClass)

const openMenu = ref<boolean>(false)
provide("open", openMenu)
</script>

<style scoped lang="scss">

</style>
