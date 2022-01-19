<template>
    <div>

        <div class="list has-visible-pointer-control has-overflow-ellipsis has-hoverable-list-items">
            <span v-for="shrl in shrls">
                <shrl-item v-bind:shrl="shrl"></shrl-item>
            </span>
        </div>

        <nav class="pagination">
            <a v-on:click="previousPage" v-bind:disabled="page <= 0" class="pagination-previous">Previous</a>
            <a v-on:click="nextPage" v-bind:disabled="page + 1 >= pageCount" class="pagination-next">Next Page</a>
            <ul class="pagination-list">
                <li v-if="page > numPages"><a v-on:click="setPage(0)" class="pagination-link"> 1 </a></li>
                <li v-if="page > numPages"><span class="pagination-ellipsis">&hellip;</span></li>

                <li v-for="p in pages">
                    <a v-on:click="setPage(p)" v-if="page == p" class="pagination-link is-current">
                        {{ p + 1 }}
                    </a>
                    <a v-on:click="setPage(p)" v-else class="pagination-link">
                        {{ p + 1 }}
                    </a>
                </li>

                <li v-if="page < pageCount - numPages - 1"><span class="pagination-ellipsis">&hellip;</span></li>
                <li v-if="page < pageCount - numPages - 1"><a v-on:click="setPage(pageCount - 1)" class="pagination-link"> {{ pageCount }} </a></li>
            </ul>
        </nav>

    </div>
</template>

<script>
import { bus } from "../index.js"

export default {
    data: function() {
        return {
            numPages: 2,
        }
    },
    computed: {
        page: function() {
            return this.searchOpts.page
        },
        pages: function() {
            let ps = []
            let startPage = Math.max(0, this.page - this.numPages)
            let endPage = Math.min(this.pageCount, this.page + this.numPages)
            if (this.page < this.numPages) {
                endPage += this.numPages - this.page
            }
            for (let p=startPage; p<Math.min(this.pageCount, endPage + 1); p++) {
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
        nextPage: function() {
            this.setPage(this.page + 1)
        },
        previousPage: function() {
            this.setPage(Math.max(0, this.page - 1))
        },
        setPage: function(p) {
            bus.$emit("setValue", {
                page: p,
            })
        },
    }
}
</script>