<template>
    <span>
        <p>{{ shrl.alias }}</p>
        woah
    </span>
</template>

<script>
export default {
    data() {
        return {
            shrl: { alias: "hello" }
        }
    },
}
</script>

<style scoped>
p {
    font-size: 18pt;
}
</style>