<template>
  <div>
    <header class="sticky top-0 transition duration-1000 ease-in-out bg-white"
            :class="{ 'scrolled': !view.atTopOfPage }">
      <div v-show="error" class="bg-red-50 text-center border-yellow-200 border-b-1">
        <div class="p-2">
          <div v-html="error"></div>
        </div>
      </div>
      <nav class="flex flex-wrap items-center justify-between px-2 py-3 navbar-expand-lg">
        <div class="container px-4 mx-auto flex flex-wrap items-center justify-between">
          <div class="w-full relative flex justify-between">
              <inertia-link aria-label="Logo"
                  class="text-sm font-bold leading-relaxed inline-block mr-4 py-2 whitespace-no-wrap uppercase text-lg serif z-10"
                  href="/">
                <h2>
                  LARAFIBER
                </h2>
              </inertia-link>

            <div class="flex z-10">
              <button
                  class=" cursor-pointer text-lg leading-none pl-4 py-1 border border-solid border-transparent rounded bg-transparent block outline-none focus:outline-none"
                  type="button" v-on:click="toggleNavbar()">
                MENU
                <i class="las la-bars" v-if="!showMenu"></i>
                <i class="las la-times" v-else></i>
              </button>
            </div>
          </div>
          <div v-bind:class="{'hidden': !showMenu, 'flex': showMenu}"
               class="text-md -mt-1 transition duration-1000 ease-in-out items-center absolute  bg-white top-0 bottom-0 left-0 right-0 h-screen">
            <ul class="text-center flex-col list-none mx-auto" v-if="!auth.uuid">
              <!-- Not logged in -->
              <li class=" nav-item">
                <inertia-link class="px-3 block py-2 uppercase font-bold leading-snug hover:opacity-75 text-sm"
                              href="/login">
                  Log in
                </inertia-link>
              </li>
              <li class=" nav-item">
                <inertia-link class="px-3 block py-2 uppercase font-bold leading-snug hover:opacity-75 text-sm"
                              href="/register">
                  Create account
                </inertia-link>
              </li>
            </ul>

            <ul class="flex flex-col text-center list-none mx-auto" v-else>
              <!-- Logged in -->
              <li class="">
                <inertia-link class="px-3 block py-2 uppercase font-bold leading-snug hover:opacity-75 text-sm"
                              href="/user/me">
                  My page
                </inertia-link>
              </li>

              <li v-for="artist in auth.artists">
                <inertia-link :href="'/artist/' + artist.uuid" class="px-3 block py-2 uppercase font-bold leading-snug hover:opacity-75 text-sm">{{artist.name}}</inertia-link>
              </li>
              <li v-if="auth.admin" class="">
                <inertia-link
                    class="px-3 block py-2 uppercase font-bold leading-snug hover:opacity-75 text-sm"
                    href="/admin">
                  Admin
                </inertia-link>
              </li>
              <li class="">
                <inertia-link class="px-3 block py-2 uppercase font-bold leading-snug hover:opacity-75 text-sm"
                              href="/logout">
                  Log out
                </inertia-link>
              </li>
            </ul>
          </div>
        </div>
      </nav>
    </header>
    <div v-if="error">
      {{ error }}
    </div>
    <div class="mb-4">
      <slot/>
    </div>
    <div class="text-center mb-2 mt-10">
      <p class="text-center text-gray-700 text-xs mb-4">
        &copy;{{ new Date().getFullYear() }} <inertia-link href="/" class="hover:underline">Larafiber</inertia-link>. A demo created by <a href="https://hakonleinan.com" class="text-gray-900 hover:underline">Håkon Leinan</a>.
      </p>
    </div>
  </div>
</template>


<script>
import { computed } from 'vue'
import { usePage } from '@inertiajs/inertia-vue3'

export default {
  setup() {
    const system_message = computed(() => usePage().props.value.system_message)
    const auth = computed(() => usePage().props.value.auth)
    const title = computed(() => usePage().props.value.title)
    const error = computed(() => usePage().props.value.error)
    return { system_message, auth, title, error }
  },

  data() {
    return {
      showMenu: false,
      view: {
        atTopOfPage: true
      }
    };
  },
  methods: {
    toggleNavbar: function () {
      this.showMenu = !this.showMenu;
    },
    handleScroll() {
      if (window.pageYOffset > 0) {
        if (this.view.atTopOfPage) this.view.atTopOfPage = false
      } else {
        if (!this.view.atTopOfPage) this.view.atTopOfPage = true
      }
    }
  },
  beforeMount() {
    window.addEventListener('scroll', this.handleScroll);
  },
  mounted() {
    // console.log("I am mounted: ", this.title)
    // document.querySelector('meta[name="description"]').setAttribute("content", this.title);
  },
  updated() {
    document.title = this.title
    // document.querySelector('meta[name="description"]').setAttribute("content", this.title);
  }
}
</script>

<style>
.st0 {
  fill: none;
}

.st1 {
  fill: #77D1F1;
}

header {
  z-index: 10
}

header.scrolled {
  @apply shadow;
  @apply: transition-shadow;
}

</style>