<template>
    <div class="box has-background-primary my-5">
        <span v-if="omnibarType == ShrlType.textSnippet">
            <div class="field">
                <label class="label has-text-light">Snippet Title</label>
            </div>

            <div class="field has-addons sticky">

                <div class="control is-expanded">
                    <input v-model="snippetTitle" class="input" type="text" placeholder="Untitled Snippet"/>
                </div>

                <div class="control">
                    <a class="button is-success"
                        v-on:click="uploadSnippet"
                    >
                        Upload
                    </a>
                </div>
            </div>
        </span>

        <div class="columns">
            <div class="column">
                <div class="field">
                    <label class="label has-text-light" v-if="omnibarType == ShrlType.textSnippet"></label>
                    <div class="control">
                        <textarea autocomplete="off" class="has-fixed-size" :placeholder="placeholder"
                            v-model:paste="omnibar"
                            v-on:paste="parsePaste"
                            v-on:keydown.enter="omnibarNewline"
                            v-bind:rows="omnibarType == ShrlType.shortenedURL ? 1 : Math.max(6, omnibar.split('\n').length)"
                            v-bind:type="omnibarType == ShrlType.shortenedURL ? 'text' : 'textarea'"
                            v-bind:class="omnibarType == ShrlType.shortenedURL ? 'input' : 'textarea'"
                        ></textarea>
                    </div>
                </div>
            </div>

            <div class="column is-narrow" v-if="omnibar.length == 0">
                <div class="file">
                    <label class="file-label">
                        <input class="file-input" type="file" name="resume" v-on:change="omnibarUpload">
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
            </div>
        </div>
    </div>
</template>

<script>
import { bus, ShrlType } from "../index.js"
import copy from "copy-to-clipboard"

function isValidHttpUrl(string) {
    string = string.toLowerCase()
    let url;
    if (string.indexOf("\n") > -1) {
        return false;
    }
    
    try {
        url = new URL(string);
    } catch (_) {
        return "http://".startsWith(string) || "https://".startsWith(string)
    }

        return url.protocol === "http:" || url.protocol === "https:";
}

export default {
    data: function() {
        return {
            ShrlType,
            snippetTitle: "",
            omnibar: "",
            notification: "",
        }
    },
    computed: {
        omnibarType() {
            if (this.omnibar.length == 0 || isValidHttpUrl(this.omnibar)) {
                return ShrlType.shortenedURL
            }
            return ShrlType.textSnippet
        },
        placeholder() {
            if (this.notification != "") {
                setTimeout(() => { this.notification = "" }, 5000)
                return this.notification
            }
            return "Paste URL / Snippet / File"
        }
    },
    methods: {
        resetOmnibar: function() {
            this.omnibar = ""
            this.snippetTitle = ""
        },
        parsePaste: function(event) {
            let el = this;
            let clipboard = (event.clipboardData || window.clipboardData)
            if (clipboard.files.length > 0) {
                for (let i of clipboard.files) {
                    el.createUpload(i)
                }
            } else {
                let paste = clipboard.getData("text")
                this.omnibar = paste
                this.parseOmnibar()
            }
        },
        omnibarNewline: function(e) {
            if (e.shiftKey) { return }
            if (this.omnibarType == ShrlType.shortenedURL) {
                this.parseOmnibar()
            }
        },
        parseOmnibar: function() {
            if (isValidHttpUrl(this.omnibar)) {
                this.createShrl(this.omnibar)
            }
        },
        uploadSnippet: function() {
            this.createSnippet(this.snippetTitle, this.omnibar)
        },
        omnibarUpload: function(e) {
            let file = e.target.files[0];
            this.createUpload(file)
        },
        copyAlias: function(alias) {
            let url = document.location.protocol + "//" + document.location.host + "/" + alias
            this.notification = "\"" + url + "\" copied to clipboard."
            copy(url)
        },
        // API posts
        createShrl: function(url) {
            fetch("/api/shrl", {
                method: "POST",
                body: JSON.stringify({
                    location: url
                })
            }).then((d) => {
                return d.json()
            }).then((d) => {
                this.copyAlias(d.shrl.alias)
                bus.$emit("load-shrls")
                this.resetOmnibar()
            })
        },
        createSnippet: function(title, paste) {
            fetch("/api/snippet", {
                method: "POST",
                body: JSON.stringify({
                    title: title,
                    body: paste,
                }),
            }).then((d) => {
                return d.json()
            }).then((d) => {
                this.copyAlias(d.shrl.alias)
                bus.$emit("load-shrls")
                this.resetOmnibar()
            })
        },
        createUpload: function(file) {
            let fd = new FormData()
            fd.append("file", file)
            fetch("/api/upload", {
                method: "POST",
                body: fd,
            }).then((d) => {
                return d.json()
            }).then((d) => {
                this.copyAlias(d.shrl.alias)
                bus.$emit("load-shrls")
                this.resetOmnibar()
            })
        },
    }
}
</script>

<style>
textarea {
    resize: none;
}
.sticky {
    position: sticky;
    z-index: 10;
    top: 0px;
}
</style>