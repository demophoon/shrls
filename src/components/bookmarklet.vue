<template>
    <div v-if="visible" id="shrl-root">
        <div id="background" v-on:click="close"></div>

        <div id="shrl-ui" class="container">

            <article class="message is-primary">
                <div class="message-header">
                    <img src="../img/banner-white.png" width="112" height="28">
                    <button v-on:click="close" class="delete"></button>
                </div>
                <div class="message-body">
                    <div class="columns is-centered">
                        <div class="column is-half">
                            <figure class="image is-2by1 is-fullwidth">
                                <img id="screenshot" :src="screenshotDataUrl" />
                            </figure>
                        </div>
                    </div>

                    <div class="columns">
                        <div class="column">
                            <button class="button" v-on:click="shortenUrl">Shorten URL</button>
                        </div>
                        <div class="column">
                            <button class="button" v-on:click="saveScreenshot">Save Screenshot</button>
                        </div>
                        <div class="column">
                            <button class="button" v-on:click="saveSelection">Save Text Snippet</button>
                        </div>
                    </div>

                    <div v-if="currentUrl != ''" class="columns">
                        <div class="column">
                            "{{ currentUrl }}" Copied to clipboard
                        </div>
                    </div>
                </div>
            </article>

        </div>
    </div>
</template>

<script>
var html2canvas = require('html2canvas');
var copy = require('copy-to-clipboard');

const BookmarkletTypes = {
    Url: 'u',
    File: 'i',
    Snippet: 's',
}

export default {
    data() {
        return {
            visible: true,
            currentUrl: "",
            screenshot: undefined,
            selection: "",
            shrlsServer: this.shrlsServer,
        }
    },
    props: ["shrlsServer"],
    computed: {
        screenshotDataUrl() {
            return this.screenshot !== undefined ? this.screenshot.toDataURL() : "";
        },
        location() {
            let u = new URL(document.location.href)
            return u.pathname
        }
    },
    methods: {
        close() {
            this.visible = false
            this.$emit("close")
            this.$destroy();
            this.$el.parentNode.removeChild(this.$el);
        },
        shortenUrl() {
            this.upload(BookmarkletTypes.Url, document.location.href)
        },
        saveScreenshot() {
            this.upload(BookmarkletTypes.File, this.screenshotDataUrl.split(",")[1])
        },
        saveSelection() {
            console.log(this.selection)
            this.upload(BookmarkletTypes.Snippet, this.selection)
        },
        upload(type, data) {
            let s = document.createElement("script")
            let query = "?" + encodeURIComponent(type) + "=" + encodeURIComponent(data)
            s.src = this.shrlsServer + "/api/bookmarklet/new" + query
            s.addEventListener('shrls-response', this.handleResponse)
            this.$el.append(s)
        },
        handleResponse(e) {
            this.currentUrl = e.detail.shrl
            copy(this.currentUrl)
        }
    },
    beforeMount() {
        let el = this
        // Current Selection
        if (window.getSelection) {
            el.selection = window.getSelection().toString();
        } else if (document.selection && document.selection.type != "Control") {
            el.selection = document.selection.createRange().text
        }
        // Screenshot
        html2canvas(document.body).then((canvas) => {
            el.screenshot = canvas
        })
    },
}
</script>

<style scoped>
    #shrl-root, #background {
        display: inline-block;
        position: fixed;
        top: 0px;
        left: 0px;
        z-index: 99999;
        width: 100%;
        height: 100%;
    }
    #background {
        background: rgba(0, 0, 0, .5);
    }
    #screenshot {
        object-fit: cover;
    }
    #shrl-ui {
        z-index: 100000;
        padding-top: 6vh;
    }
</style>