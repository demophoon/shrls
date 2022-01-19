import Vue from "vue"
import ShrlList from "./components/shrlList.vue"
import ShrlItem from "./components/shrl.vue"
import ShrlEdit from "./components/shrlEdit.vue"
import ShrlNav from "./components/shrlNav.vue"
import ShrlOmnibar from "./components/omnibar.vue"
import ShrlSearch from "./components/shrlSearch.vue"
import ShrlFooter from "./components/footer.vue"

require('./styles.scss')
import '@fortawesome/fontawesome-free/css/all.css'

var _ = require('lodash');

Vue.component("shrl-list", ShrlList)
Vue.component("shrl-item", ShrlItem)
Vue.component("shrl-edit", ShrlEdit)
Vue.component("shrl-nav", ShrlNav)
Vue.component("shrl-omnibar", ShrlOmnibar)
Vue.component("shrl-search", ShrlSearch)
Vue.component("shrl-footer", ShrlFooter)

const container = document.createElement("div")

document.body.append(container)

const _bus = new Vue({});

export const ShrlEnum = [
    "Shortened URL",
    "Uploaded File",
    "Text Snippet",
]
export const ShrlType = {
    shortenedURL: 0,
    uploadedFile: 1,
    textSnippet: 2,
}

const app = new Vue({
    el: container,
    data: function() {
        return {
            shrls: [],
            count: 0,
            searchOpts: {
                search: "",
                page: 0,
                limit: 15,
            },
        }
    },
    template: `
    <div>
        <div class="container">
            <shrl-nav></shrl-nav>
        </div>

        <div class="container mb-6 mt-4">

            <shrl-omnibar />

            <div class="columns">
                <div class="column is-hidden-tablet">
                    <shrl-search />
                </div>
                <div class="column">
                    <shrl-list v-bind:shrls='shrls' v-bind:count="count" v-bind:searchOpts='searchOpts'></shrl-list>
                </div>
            </div>

        </div>
        <shrl-footer />
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

            if (e.search !== undefined) { this.searchOpts.search = e.search }
            if (e.page !== undefined) { this.searchOpts.page = e.page }

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
                    app.shrls = data.shrls.map((shrl) => {
                        shrl.tags = shrl.tags || []
                        return shrl
                    });
                    app.count = data.count;
                })
                .catch(err => { throw err });
        }, 1000),
    },
})

export let bus = _bus;