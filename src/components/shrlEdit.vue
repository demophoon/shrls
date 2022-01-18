<template>
    <div v-bind:class="{ 'is-active': editing }" class="modal">
        <div class="modal-background" v-on:click="$emit('close')"></div>
        <div class="modal-content">
            <button v-on:click="$emit('close')" class="modal-close is-large" aria-label="close"></button>
            <div class="box is-vcentered">

                <h1 class="title">
                    Edit
                </h1>
                <h2 class="subtitle">
                    {{ type }}
                </h2>

                <div class="field" v-if="shrl.type == 0">
                    <label class="label">Location</label>
                    <div class="control">
                        <input class="input" type="text" placeholder="https://example.com/"
                            v-model="shrl.location" 
                        />
                    </div>
                </div>

                <div class="field">
                    <label class="label">Alias</label>
                    <div class="control">
                        <input class="input" type="text" placeholder="short_url"
                            v-model="shrl.alias" 
                        />
                    </div>
                </div>

                <div class="field" v-if="shrl.type == 2">
                    <label class="label">Snippet Title</label>
                    <div class="control">
                        <input type="text" class="input" placeholder="Snippet Title" v-model="shrl.snippet_title"/>
                    </div>
                </div>

                <div class="field" v-if="shrl.type == 2">
                    <label class="label">Snippet</label>
                    <div class="control">
                        <textarea v-model="shrl.snippet" class="textarea" rows="6" placeholder="Snippet of text"></textarea>
                    </div>
                </div>

                <button v-on:click="$emit('save')" class="button is-small is-primary">Save</button>
                <button v-on:click="$emit('remove')" class="button is-small is-danger">Delete</button>
            </div>
        </div>
    </div>
</template>

<script>
import { ShrlEnum } from "../index.js"
export default {
    props: ["shrl", "editing"],
    computed: {
        type: function() {
            return ShrlEnum[this.shrl.type];
        },
    },
    mounted: function() {
        window.addEventListener("keydown", (e) => {
            if (e.keyCode === 27) {
                this.$emit("close");
            }
        })
    },
}
</script>