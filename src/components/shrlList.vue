<template>
    <div>
        <table class="table is-fullwidth is-striped is-narrow">
            <thead>
                <tr>
                    <td>Actions</td>
                    <td>Views</td>
                    <td>URL</td>
                </tr>
            </thead>
            <tbody v-for="shrl in shrls">
                <shrl-item v-bind:shrl='shrl'></shrl-item>
            </tbody>
        </table>

        <div class="columns is-centered">
            <span class="column is-1">
                <button class="button" v-bind:disabled="page <= 0" v-on:click="previousPage">&lt;&lt;</button>
            </span>

            <span class="column is-1" v-for="p in pages">
                <button class="button is-dark" v-if="page == p">{{ p + 1 }}</button>
                <button class="button" v-else v-on:click="setPage(p)">{{ p + 1 }}</button>
            </span>

            <span class="column is-1">
                <button class="button" v-bind:disabled="page + 1 >= pageCount" v-on:click="nextPage">&gt;&gt;</button>
            </span>
        </div>

    </div>
</template>

<script>
import { bus } from "../index.js"

export default {
    data: function() {
        return {
            page: 0,
        }
    },
    computed: {
        pages: function() {
            let numPages = 2
            let ps = []
            let startPage = Math.max(0, this.page - numPages)
            let endPage = Math.min(this.pageCount, this.page + numPages)
            if (this.page < numPages) {
                endPage += numPages - this.page
            }
            for (let p=startPage; p<Math.min(this.pageCount, endPage); p++) {
                ps.push(p)
            }
            return ps.splice(0, (endPage - startPage) + 1)
        },
        pageCount: function() {
            return Math.ceil(this.count / this.searchOpts.limit)
        }
    },
    props: ["shrls", "count", "searchOpts"],
    methods: {
        searchShrls: function() {
            bus.$emit("setValue", {
                page: this.page,
            })
        },
        updateSearch: function() {
            this.page = 0;
            this.searchShrls();
        },
        nextPage: function() {
            this.page = this.page + 1;
            this.searchShrls();
        },
        previousPage: function() {
            this.page = Math.max(0, this.page - 1);
            this.searchShrls();
        },
        setPage: function(e) {
            this.page = e
            this.searchShrls();
        },
    }
}
</script>