import Vue from "vue"
import Shrl from "./components/shrl.vue"

Vue.component("shrl-info", Shrl)

const container = document.createElement("div")

container.appendChild(document.createElement("shrl-info"))

document.body.append(container)

const app = new Vue({el: container})