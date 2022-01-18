<template>
    <input class="input" placeholder="Search" v-model="search" v-on:keyup="updateSearch">
</template>

<script>
import { bus } from "../index.js"

export default {
    data: function() {
        return {
            search: ""
        }
    },
    methods: {
        updateSearch() {
            bus.$emit("setValue", {
                search: this.search,
            })
        },
    }
}
</script>

<style>

</style>