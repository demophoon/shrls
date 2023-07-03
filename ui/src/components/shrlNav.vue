<template>
    <nav class="navbar my-3" role="navigation" aria-label="main navigation">
        <div class="navbar-brand">
            <a class="navbar-item" href="/admin/">
                <img src="../img/banner.png" width="112" height="28">
            </a>
        </div>

        <div class="navbar-end is-hidden-mobile">
            <div class="navbar-item">
                <shrl-search />
            </div>
            <a v-bind:href="bookmarklet" class="navbar-item">
                Shrls Bookmarklet
            </a>
        </div>
    </nav>
</template>

<script>
export default {
    computed: {
        bookmarklet() {
            let u = new URL(document.location.href)
            let shrlsServer = u.protocol + "//" + u.host
            return "javascript:(function()%7Bvar%20d=document,s=d.createElement('script');s.src='" + shrlsServer + "/admin/bookmarklet.js';void(d.body.appendChild(s));%7D)();"
        }
    }

}
</script>