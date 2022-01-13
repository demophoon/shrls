<template>
    <span v-on:dragover.prevent v-on:drop.prevent>
        <div class="file is-boxed is-centered">
            <label class="file-label">
                <input v-on:change="uploadFile" v-on:drop="uploadFile" class="file-input" type="file">
                <span class="file-cta">
                    <span class="file-icon">
                        <i class="fas fa-upload"></i>
                    </span>
                    <span class="file-label">
                        Choose a file…
                    </span>
                </span>
            </label>
        </div>
    </span>
</template>

<script>
export default {
    methods: {
        uploadFile: function(e) {
            console.log(e)
        }
    },
}
</script>