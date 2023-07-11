import Vue from "vue"
import SwaggerClient from 'swagger-client';

import ShrlList from "./components/shrlList.vue"
import ShrlItem from "./components/shrl.vue"
import ShrlEdit from "./components/shrlEdit.vue"
import ShrlNav from "./components/shrlNav.vue"
import ShrlOmnibar from "./components/omnibar.vue"
import ShrlSearch from "./components/shrlSearch.vue"
import ShrlFooter from "./components/footer.vue"

require('./styles.scss')
import '@fortawesome/fontawesome-free/css/all.css'

import swagger from '../../server/gen/shrls.swagger.json'

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
    shortenedURL: "LINK",
    uploadedFile: "UPLOAD",
    textSnippet: "SNIPPET",
}

const api = (await new SwaggerClient(swagger))

const app = new Vue({
    el: container,
    data: function() {
        return {
            api: api.apis.Shrls,
            shrls: [],
            count: 0,
            searchOpts: {
                search: "",
                page: 0,
                count: 15,
            },
        }
    },
    template: `
    <div>
        <div class="container">
            <shrl-nav></shrl-nav>
        </div>

        <div class="container mb-6 mt-4">

            
            <div class="columns is-centered">
                <div class="column is-two-thirds">
                    <shrl-omnibar v-bind:api='api' />
                </div>
            </div>

            <div class="columns is-centered">
                <div class="column is-two-thirds is-hidden-tablet">
                    <shrl-search />
                </div>
            </div>

            <div class="columns is-centered">
                <div class="column is-two-thirds">
                    <shrl-list v-bind:shrls='shrls' v-bind:count="count" v-bind:searchOpts='searchOpts' v-bind:api='api'></shrl-list>
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
            let limit = this.searchOpts.count || 15
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
            this.api.Shrls_ListShrls({
                search: this.searchOpts.search,
                page: this.searchOpts.page,
                count: this.searchOpts.count,
            }).then(res => {
                app.shrls = res.obj.shrls
                app.count = res.obj.totalShrls
            }).catch(err => { throw err });
        }, 1000),
    },
})

console.log(app.api)

export let bus = _bus;
