import Vue from "vue"

import Bookmarklet from "./components/bookmarklet.vue"

require('./bookmarklet.scss')
require('html2canvas');
require('copy-to-clipboard');

(() => {
    let bus_id = "data-shrls-modal"
    let event_bus = document.getElementById(bus_id)
    if (event_bus == undefined) {
        event_bus = document.createElement("span")
        event_bus.id = bus_id
        document.body.appendChild(event_bus);

        new Vue({
            el: event_bus,
            components: {Bookmarklet},
            template: `<div id=` + bus_id + `>
                <bookmarklet v-on:close="close"/>
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