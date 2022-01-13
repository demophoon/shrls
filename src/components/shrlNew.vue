<template>
    <span>
        <div class="field">
            <label class="label">Link</label>
            <div class="control">
                <input class="input" type="text" placeholder="https://example.com"
                    v-model="shrl.location" 
                    v-on:keydown.enter="create"
                ></input>
            </div>
        </div>
    </span>
</template>

<script>
import { bus } from "../index.js"
export default {
    data: function() {
        return {
            shrl: {
                location: null,
                alias: null,
                tags: null,
            },
        }
    },
    methods: {
        create: function() {
            let el = this;
            fetch("/api/shrl", {
                method: "POST",
                body: JSON.stringify(el.shrl)
            }).then(() => {
                el.shrl.location = null;
                el.shrl.alias = null;
                el.shrl.tags = null;

                bus.$emit("load-shrls")
            })
            return false;
        }
    }
}
</script>