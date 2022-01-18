<template>
    <div class="box has-background-info">
        <span v-if="omnibarType == ShrlType.textSnippet">
            <div class="field">
                <label class="label has-text-light">Snippet Title</label>
            </div>

            <div class="field has-addons">

                <div class="control is-expanded">
                    <input v-model="snippetTitle" class="input" type="text" placeholder="Untitled Snippet"/>
                </div>

                <div class="control">
                    <a class="button is-success"
                        v-on:click="parseOmnibar"
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
                        <textarea class="has-fixed-size" placeholder="Paste URL / Snippet / File"
                            v-model:paste="omnibar"
                            v-on:paste="parsePaste"
                            v-on:keydown.enter="omnibarNewline"
                            v-bind:rows="omnibarType == ShrlType.shortenedURL ? 1 : 6"
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

function isValidHttpUrl(string) {
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
        }
    },
    computed: {
        omnibarType() {
            if (this.omnibar.length == 0 || isValidHttpUrl(this.omnibar)) {
                return ShrlType.shortenedURL
            }
            return ShrlType.textSnippet
        },
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
            let paste = this.omnibar
            if (isValidHttpUrl(paste)) {
                this.createShrl(paste)
            } else {
                this.createSnippet(this.snippetTitle, paste)
            }
        },
        omnibarUpload: function(e) {
            let file = e.target.files[0];
            this.createUpload(file)
        },
        // API posts
        createShrl: function(url) {
            this.resetOmnibar()
            fetch("/api/shrl", {
                method: "POST",
                body: JSON.stringify({
                    location: url
                })
            }).then(() => {
                bus.$emit("load-shrls")
            })
        },
        createSnippet: function(title, paste) {
            this.resetOmnibar()
            fetch("/api/snippet", {
                method: "POST",
                body: JSON.stringify({
                    title: title,
                    body: paste,
                }),
            }).then(() => {
                bus.$emit("load-shrls")
            })
        },
        createUpload: function(file) {
            this.resetOmnibar()
            let fd = new FormData()
            fd.append("file", file)
            fetch("/api/upload", {
                method: "POST",
                body: fd,
            }).then(() => {
                bus.$emit("load-shrls")
            })
        },
    }
}
</script>

<style>
textarea {
    resize: none;
}
</style>