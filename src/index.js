import Vue from "vue"
import ShrlList from "./components/shrlList.vue"
import ShrlItem from "./components/shrl.vue"
import ShrlNew from "./components/shrlNew.vue"
import ShrlEdit from "./components/shrlEdit.vue"
import ShrlNav from "./components/shrlNav.vue"
import FileUpload from "./components/FileUpload.vue"
import ShrlSnippet from "./components/ShrlSnippet.vue"

import 'bulma/css/bulma.css'
import '@fortawesome/fontawesome-free/css/all.css'

var _ = require('lodash');

Vue.component("shrl-list", ShrlList)
Vue.component("shrl-item", ShrlItem)
Vue.component("shrl-new", ShrlNew)
Vue.component("shrl-edit", ShrlEdit)
Vue.component("shrl-nav", ShrlNav)
Vue.component("file-upload", FileUpload)
Vue.component("shrl-snippet", ShrlSnippet)

const container = document.createElement("div")

document.body.append(container)

const _bus = new Vue({});

const app = new Vue({
    el: container,
    data: function() {
        return {
            shrls: [],
            count: 0,
            searchOpts: {
                search: "",
                page: 0,
                limit: 25,
            },
        }
    },
    template: `
    <div class="container">
        <shrl-nav></shrl-nav>
        <div class="box">
            <div class="columns">
                <div class="column is-two-thirds">
                    <shrl-new></shrl-new>

                    <shrl-snippet></shrl-snippet>
                </div>
                <div class="column is-vcentered">
                    <file-upload></file-upload>
                </div>
            </div>
        </div>

        <div class="columns">
            <div class="column">
                <shrl-list v-bind:shrls='shrls' v-bind:count="count" v-bind:searchOpts='searchOpts'></shrl-list>
            </div>
        </div>
    </div>
    `,
    created: function() {
        this.fetchShrls()
    },
    mounted: function() {
        _bus.$on("load-shrls", this.fetchShrls)
        _bus.$on("setValue", this.setValue)
    },
    computed: {
        skip: function() {
            let page = this.searchOpts.page || 0
            let limit = this.searchOpts.limit || 25
            return page * limit
        },
    },
    methods: {
        setValue: function(e) {
            let fetch = (
                this.searchOpts.search != e.search ||
                this.searchOpts.page != e.page
            )

            this.searchOpts.search = e.search;
            this.searchOpts.page = e.page

            if (fetch) { this.fetchShrls() }
        },
        fetchShrls: _.throttle(function() {
            let url = new URL(document.location.protocol + "//" + document.location.host + "/api/shrl")
            url.searchParams.append("search", this.searchOpts.search)
            url.searchParams.append("skip", this.skip)
            url.searchParams.append("limit", this.searchOpts.limit)
            fetch(url)
                .then(res => res.json())
                .then(data => {
                    app.shrls = data.shrls;
                    app.count = data.count;
                })
                .catch(err => { throw err });
        }, 1000),
    },
})

export let bus = _bus;