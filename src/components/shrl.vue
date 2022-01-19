<template>
    <tr>
        <td>
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
        </td>

        <td>
            <span class="icon" v-if="shrl.type == ShrlType.shortenedURL">
                <i class="fas fa-link"></i>
            </span>
            <span class="icon" v-if="shrl.type == ShrlType.uploadedFile">
                <i class="fas fa-file"></i>
            </span>
            <span class="icon" v-if="shrl.type == ShrlType.textSnippet">
                <i class="fas fa-code"></i>
            </span>
            <a target='_blank' v-bind:href="short_url">{{ shrl.alias }}</a>
        </td>

        <td>
            <span class='is-size-7' v-if="shrl.type == 0">
                <a target='_blank' v-bind:href="shrl.location">{{ domain }}</a>
            </span>
        </td>

        <td>
            <vue-tags-input v-bind:tags="tags" v-model="currentTag" v-on:tags-changed="updateTags"/>
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
import VueTagsInput from '@johmun/vue-tags-input'

export default {
    components: {
        VueTagsInput,
    },
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
        tags: function() {
            return this.shrl.tags.map((tag) => {
                return {
                    text: tag,
                    classes: "tag is-info is-light is-rounded"
                }
            })
        }
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
        updateTags: function(tags) {
            this.shrl.tags = tags.map((tag) => {
                return tag.text
            })
            this.save()
        },
    }
}
</script>