<template>
    <div class="list-item">
        <div class="list-item-image">
            <span class="icon" v-if="shrl.type == ShrlType.shortenedURL">
                <img class="fas" v-if="shrl.content.url.favicon" :src="faviconHelper(shrl)"></img>
                <i class="fas fa-link" v-else></i>
            </span>
            <span class="icon" v-if="shrl.type == ShrlType.textSnippet">
                <i class="fas fa-code"></i>
            </span>
            <span class="icon" v-if="shrl.type == ShrlType.uploadedFile">
                <i class="fas fa-file"></i>
            </span>
        </div>

        <div class="list-item-content">
            <div class="list-item-title">
                <a target="_blank" v-bind:href="short_url">{{ shrl.stub }}</a>
            </div>

            <div class="list-item-description">
                <span class="is-size-7" v-if="shrl.type == ShrlType.shortenedURL">
                    <a target="_blank" v-bind:href="shrl.content.url.url">{{ domain }}</a>
                </span>
                <span class="is-size-7" v-if="shrl.type == ShrlType.textSnippet">
                    {{ shrl.content.snippet.title }}
                </span>
                <span class="is-size-7" v-if="shrl.type == ShrlType.uploadedFile">
                    Uploaded File
                </span>
                <div class="tags" v-if="shrl.tags.length > 0">
                    <span v-for="tag in shrl.tags" class="tag is-light is-primary">{{ tag }}</span>
                </div>
            </div>

            <div class="list-item-controls">
                <div class="buttons is-right">

                    <button  v-if="shrl.type == ShrlType.shortenedURL" v-on:click="copyQR" class="button">
                        <span class="icon is-small">
                            <i class="fas fa-qrcode"></i>
                        </span>
                        <span>QR Code</span>
                    </button>

                    <button v-on:click="edit" class="button">
                        <span class="icon is-small">
                            <i class="fas fa-edit"></i>
                        </span>
                        <span>Edit</span>
                    </button>

                    <button v-on:click="copyUrl" class="button">
                        <span class="icon is-small">
                            <i class="fas fa-copy"></i>
                        </span>
                        <span>Copy</span>
                    </button>

                </div>
            </div>

        </div>

        <shrl-edit
            v-on:save="save"
            v-on:remove="remove"
            v-on:close="closeEdit"
            v-if="edit"
            v-bind:shrl="shrl"
            v-bind:editing="editing"
        ></shrl-edit>
    </div>
</template>

<script>
import { bus, ShrlType } from "../index.js"
import copy from "copy-to-clipboard"

export default {
    props: ["shrl", "api"],
    data: function() {
        return {
            currentTag: '',
            editing: false,
            ShrlType,
        }
    },
    computed: {
        short_url: function() {
            return "/" + this.shrl.stub;
        },
        domain: function() {
            if (this.shrl.type == ShrlType.shortenedURL) {
                try {
                    return new URL(this.shrl.content.url.url).host
                } catch (_) {
                    return ""
                }
            }
            return ""
        },
    },
    methods: {
        save: function() {
            let el = this;
            this.api.Shrls.Shrls_PutShrl({
                "shrl.id": el.shrl.id,
                "body": {"shrl": el.shrl},
            }).then((r) => {
                console.log(r);
                el.closeEdit();
            }).catch(err => { throw err });
        },
        remove: function() {
            let el = this;
            this.api.Shrls.Shrls_DeleteShrl({
                "shrl.id": el.shrl.id,
            }).then((r) => {
                console.log(r);
                el.closeEdit();
            }).catch(err => { throw err });
        },
        edit: function() {
            this.editing = true
        },
        closeEdit: function() {
            this.editing = false
            bus.$emit("load-shrls")
        },
        copyUrl: function() {
            copy(document.location.protocol + "//" + document.location.host + "/" + this.shrl.stub)
        },
        copyQR: function() {
            copy(document.location.protocol + "//" + document.location.host + "/" + this.shrl.stub + ".qr")
        },
        faviconHelper: function(url) {
            return "data:image/png;base64, " + url.content.url.favicon;
        },
    }
}
</script>
