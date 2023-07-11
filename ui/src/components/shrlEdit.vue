<template>
    <div v-bind:class="{ 'is-active': editing }" class="modal">
        <div class="modal-background" v-on:click="$emit('close')"></div>
        <div class="modal-content">
            <button v-on:click="$emit('close')" class="modal-close is-large" aria-label="close"></button>
            <div class="box is-vcentered">

                <h1 class="title">
                    Edit /{{ shrl.stub }}
                </h1>
                <h2 class="subtitle">
                    {{ type }}
                </h2>

                <div class="field">
                    <label class="label">Views</label>
                    <div class="control">
                        <input class="input" type="text" readonly
                            v-bind:value="shrl.views" 
                        />
                    </div>
                </div>

                <div class="field" v-if="shrl.type == ShrlType.shortenedURL">
                    <label class="label">Location</label>
                    <div class="control">
                        <input class="input" type="text" placeholder="https://example.com/"
                            v-model="shrl.content.url.url" 
                        />
                    </div>
                </div>

                <div class="field">
                    <label class="label">Alias</label>
                    <div class="control">
                        <input class="input" type="text" placeholder="short_url"
                            v-model="shrl.stub" 
                        />
                    </div>
                </div>

                <div class="field" v-if="shrl.type == ShrlType.textSnippet">
                    <label class="label">Snippet Title</label>
                    <div class="control">
                        <input type="text" class="input" placeholder="Snippet Title" v-model="shrl.content.snippet.title"/>
                    </div>
                </div>

                <div class="field" v-if="shrl.type == ShrlType.textSnippet">
                    <label class="label">Snippet</label>
                    <div class="control">
                        <textarea v-model="snippet" class="textarea" rows="6" placeholder="Snippet of text"></textarea>
                    </div>
                </div>

                <div class="field">
                    <label class="label">Tags</label>
                </div>
                <div class="field is-grouped is-grouped-multiline">
                    <div class="control" v-for="tag in shrl.tags">
                        <div class="tags has-addons are-large">
                            <span class="tag is-light is-primary">{{ tag }}</span>
                            <a v-on:click="removeTag(tag)" class="tag is-delete is-primary"></a>
                        </div>
                    </div>
                </div>
                <div class="field">
                    <div class="control">
                        <input type="text" class="input" placeholder="Add tag" v-model="currentTag" v-on:keypress.enter="addTag">
                    </div>
                </div>

                <button v-on:click="save" class="button is-small is-primary">Save</button>
                <button v-on:click="$emit('remove')" class="button is-small is-danger">Delete</button>
            </div>
        </div>
    </div>
</template>

<script>
import { ShrlEnum, ShrlType } from "../index.js"

export default {
    props: ["shrl", "editing"],
    data: function() {
        let d = {
            currentTag: "",
            snippet: "",
            ShrlType,
        }
        if (this.shrl.type == ShrlType.textSnippet) {
            d.snippet = atob(this.shrl.content.snippet.body);
        }
        return d
    },
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
    methods: {
        save() {
            if (this.type == ShrlType.textSnippet) {
                this.shrl.content.snippet.body = btoa(this.snippet)
            }
            this.$emit("save");
        },
        addTag() {
            this.shrl.tags.push(this.currentTag)
            this.currentTag = ""
        },
        removeTag(tag) {
            this.shrl.tags = this.shrl.tags.filter((t) => { return t != tag })
        }
    }
}
</script>