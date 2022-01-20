import Vue from "vue"

import Bookmarklet from "./components/bookmarklet.vue"

require('./bookmarklet.scss')
require('html2canvas');
require('copy-to-clipboard');

(() => {
    let u = new URL(document.currentScript.src)
    let shrlsServer = u.protocol + "//" + u.host
    let bus_id = "data-shrls-modal"
    let event_bus = document.getElementById(bus_id)
    if (event_bus == undefined) {
        event_bus = document.createElement("span")
        event_bus.id = bus_id
        document.body.appendChild(event_bus);

        new Vue({
            el: event_bus,
            data: function() {
                return {shrlsServer}
            },
            components: {Bookmarklet},
            template: `<div id=` + bus_id + `>
                <bookmarklet v-on:close="close" v-bind:shrlsServer="shrlsServer" />
            </div>`,
            methods: {
                close() {
                    this.$destroy();
                    this.$el.parentNode.removeChild(this.$el);
                },
            },
        })
    }

})();

document.currentScript.remove()