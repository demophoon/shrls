<template>
    <tr>
        <td>
            <button v-on:click="edit" class="button is-small is-info">Edit</button>
            <button v-on:click="copyUrl" class="button is-small is-info">Copy</button>
        </td>
        <td>
            {{ shrl.views }}
        </td>
        <td>
            <a v-bind:href="short_url">{{ shrl.alias }}</a>
            <span v-if="shrl.type == 0">
                <br>
                <small>
                    <a v-bind:href="shrl.location">{{ shrl.location.slice(0, 50) }}</a>
                </small>
            </span>
        </td>
        <shrl-edit
            v-on:save="save"
            v-on:remove="remove"
            v-on:close="closeEdit"
            v-if="edit"
            v-bind:shrl="shrl"
            v-bind:editing="editing"
            v-bind:params="params"></shrl-edit>
    </tr>
</template>

<script>
import { bus } from "../index.js"
import copy from "copy-to-clipboard"

export default {
    props: ["shrl"],
    data: function() {
        return {
            editing: false,
            params: {},
        }
    },
    computed: {
        short_url: function() {
            return "/" + this.shrl.alias;
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
            let self = this
            if (this.shrl.type == 2) {
                fetch("/api/snippet/" + this.shrl.id).then(d => {
                    return d.json()
                }).then((pl) => {
                    self.params.snippetTitle = pl.title
                    self.params.snippet = pl.body
                    self.editing = true
                })
            } else {
                this.editing = true
            }
        },
        closeEdit: function() {
            this.editing = false
            bus.$emit("load-shrls")
        },
        copyUrl: function() {
            copy(document.location.protocol + "//" + document.location.host + "/" + this.shrl.alias)
        }
    }
}
</script>