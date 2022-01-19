<template>
    <tr>
        <td>
            <span class="my-2">
                <button v-on:click="edit" class="button is-small">
                    <span class="icon">
                        <i class="fas fa-edit"></i>
                    </span>
                </button>
                <button v-on:click="copyUrl" class="button is-small">
                    <span class="icon">
                        <i class="fas fa-copy"></i>
                    </span>
                </button>
                <button  v-if="shrl.type == ShrlType.shortenedURL" v-on:click="copyQR" class="button is-small">
                    <span class="icon">
                        <i class="fas fa-qrcode"></i>
                    </span>
                </button>
            </span>
        </td>

        <td>
            <span class="my-2 mr-3 is-pulled-left">
                <span class="icon" v-if="shrl.type == ShrlType.shortenedURL">
                    <i class="fas fa-link"></i>
                </span>
                <span class="icon" v-if="shrl.type == ShrlType.textSnippet">
                    <i class="fas fa-code"></i>
                </span>
                <span class="icon" v-if="shrl.type == ShrlType.uploadedFile">
                    <i class="fas fa-file"></i>
                </span>
            </span>

            <div>
                <a target="_blank" v-bind:href="short_url">{{ shrl.alias }}</a>

                <br>

                <span class="is-size-7" v-if="shrl.type == ShrlType.shortenedURL">
                    <a target="_blank" v-bind:href="shrl.location">{{ domain }}</a>
                </span>
                <span class="is-size-7" v-if="shrl.type == ShrlType.textSnippet">
                    {{ shrl.snippet_title }}
                </span>
                <span class="is-size-7" v-if="shrl.type == ShrlType.uploadedFile">
                    Uploaded File
                </span>

            </div>
        </td>

        <td>
            <span class="tags">
                <span v-for="tag in shrl.tags" class="tag is-light is-primary">{{ tag }}</span>
            </span>
        </td>

        <shrl-edit
            v-on:save="save"
            v-on:remove="remove"
            v-on:close="closeEdit"
            v-if="edit"
            v-bind:shrl="shrl"
            v-bind:editing="editing"
        ></shrl-edit>
    </tr>
</template>

<script>
import { bus, ShrlType } from "../index.js"
import copy from "copy-to-clipboard"

export default {
    props: ["shrl"],
    data: function() {
        return {
            currentTag: '',
            editing: false,
            ShrlType,
        }
    },
    computed: {
        short_url: function() {
            return "/" + this.shrl.alias;
        },
        domain: function() {
            if (this.shrl.type == ShrlType.shortenedURL) {
                try {
                    return new URL(this.shrl.location).host
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
            fetch("/api/shrl/" + el.shrl.id, {
                method: "PUT",
                body: JSON.stringify(el.shrl),
            }).then(() => {
                el.closeEdit();
            })
        },
        remove: function() {
            let el = this;
            fetch("/api/shrl/" + this.shrl.id, {
                method: "DELETE",
                body: JSON.stringify(el.shrl),
            }).then(() => {
                el.closeEdit();
            })
        },
        edit: function() {
            this.editing = true
        },
        closeEdit: function() {
            this.editing = false
            bus.$emit("load-shrls")
        },
        copyUrl: function() {
            copy(document.location.protocol + "//" + document.location.host + "/" + this.shrl.alias)
        },
        copyQR: function() {
            copy(document.location.protocol + "//" + document.location.host + "/" + this.shrl.alias + ".qr")
        },
    }
}
</script>

<style>
.ti-input {
    border: none;
}
.vue-tags {
    max-width: 200px;
}
</style>