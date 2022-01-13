<template>
    <tr>
        <td>
            <button v-on:click="edit" class="button is-small is-info">Edit</button>
            <button v-on:click="copyUrl" class="button is-small is-info">Copy</button>
        </td>
        <td>
            0
        </td>
        <td>
            <a v-bind:href="short_url">{{ shrl.Alias }}</a>
            <br>
            <small>
                <a v-bind:href="shrl.Location">{{ shrl.Location.slice(0, 50) }}</a>
            </small>
        </td>
        <shrl-edit
            v-on:save="save"
            v-on:remove="remove"
            v-on:close="closeEdit"
            v-if="edit"
            v-bind:shrl="shrl"
            v-bind:editing="editing"></shrl-edit>
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
        }
    },
    computed: {
        short_url: function() {
            return "http://localhost:8000/" + this.shrl.Alias;
        },
    },
    methods: {
        save: function() {
            let el = this;
            fetch("/api/shrl/" + el.shrl.ID, {
                method: "PUT",
                body: JSON.stringify(el.shrl),
            }).then(() => {
                bus.$emit("load-shrls")
                el.closeEdit();
            })
        },
        remove: function() {
            let el = this;
            fetch("/api/shrl/" + this.shrl.ID, {
                method: "DELETE",
                body: JSON.stringify(el.shrl),
            }).then(() => {
                bus.$emit("load-shrls")
                el.closeEdit();
            })
        },
        edit: function() {
            this.editing = true
        },
        closeEdit: function() {
            this.editing = false
        },
        copyUrl: function() {
            copy(document.location.protocol + "//" + document.location.host + "/" + this.shrl.Alias)
        }
    }
}
</script>