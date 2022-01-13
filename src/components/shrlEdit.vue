<template>
    <div v-on:keyup.esc="closeEdit" v-bind:class="{ 'is-active': editing }" class="modal">
        <div class="modal-background"></div>
        <div class="modal-content">
            <div class="box is-vcentered">

                <div class="field">
                    <label class="label">Location</label>
                    <div class="control">
                        <input class="input" type="text" placeholder="https://example.com/"
                            v-model="shrl.Location" 
                        ></input>
                    </div>
                </div>

                <div class="field">
                    <label class="label">Alias</label>
                    <div class="control">
                        <input class="input" type="text" placeholder="short_url"
                            v-model="shrl.Alias" 
                        ></input>
                    </div>
                </div>

                <button v-on:click="$emit('save')" class="button is-small is-primary">Save</button>
                <button v-on:click="$emit('remove')" class="button is-small is-danger">Delete</button>
            </div>
        </div>
        <button v-on:click="$emit('close')" class="modal-close is-large" aria-label="close"></button>
    </div>
</template>

<script>
export default {
    props: ["shrl", "editing"],
}
</script>