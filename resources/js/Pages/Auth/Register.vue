<template>
  <Layout>
    <div class="container mx-auto flex justify-center relative">
      <div class="w-full max-w-sm mt-4 md:mt-16">
        <form class="px-8 pt-6 pb-8 mb-4" @submit.prevent="submit" autocomplete="off">
          <h1 class="text-lg mb-2">Create account</h1>
          <p class="mb-8">Create account here.</p>
          <div class="mb-4">
            <label class="block text-gray-700 text-sm font-bold mb-2" for="name">
              Name
            </label>
            <input v-model="form.name"
                   class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
                   id="name" type="text" placeholder="Name" autocomplete="off">
          </div>
          <div class="mb-4">
            <label class="block text-gray-700 text-sm font-bold mb-2" for="username">
              Email address (username)
            </label>
            <input v-model="form.email"
                   class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
                   id="username" type="email" placeholder="Email" autocomplete="off">
          </div>
          <div class="mb-4">
            <label class="block text-gray-700 text-sm font-bold mb-2" for="password">
              Password (minimum 8 characters)
            </label>
            <input v-model="form.password"
                   class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 mb-3 leading-tight focus:outline-none focus:shadow-outline"
                   id="password" type="password" placeholder="******************" autocomplete="password">
          </div>
          <div class="mb-4">
            <label class="block text-gray-700 text-sm font-bold mb-2" for="password_repeat">
              Repeat password
            </label>
            <input v-model="form.password_repeat"
                   class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 mb-3 leading-tight focus:outline-none focus:shadow-outline"
                   id="password_repeat" type="password" placeholder="******************" autocomplete="new-password">
            <div class="mb-4">
              &nbsp;<span class="text-red-700 text-sm" v-show="register_error">{{ register_error }}</span>
            </div>
          </div>
          <div class="mb-6">
            <div class="relative inline-block w-10 mr-2 align-middle select-none transition duration-200 ease-in">
              <input v-model="form.agreed" type="checkbox" name="toggle" id="toggle"
                     class="toggle-checkbox absolute block w-6 h-6 rounded-full bg-white border-4 appearance-none cursor-pointer"/>
              <label for="toggle"
                     class="toggle-label block overflow-hidden h-6 rounded-full bg-gray-300 cursor-pointer"></label>
            </div>
            <label for="toggle" class="text-xs text-gray-800"><a target="_blank" class="text-brand-blue underline"
                                                                 href="#">Accept agreement</a></label>
          </div>
          <div class="mb-6">
            <div class="relative inline-block w-10 mr-2 align-middle select-none transition duration-200 ease-in">
              <input v-model="form.marketing" type="checkbox" name="toggle-3" id="toggle-3"
                     class="toggle-checkbox absolute block w-6 h-6 rounded-full bg-white border-4 appearance-none cursor-pointer"/>
              <label for="toggle-3"
                     class="toggle-label block overflow-hidden h-6 rounded-full bg-gray-300 cursor-pointer"></label>
            </div>
            <label for="toggle" class="text-xs text-gray-800">Subscribe to newsletter?</label>
          </div>
          <div class="flex items-center justify-between mb-4">
            <button
                class="bg-pink-500 hover:bg-pink-600 shadow hover:shadow-md text-white font-bold py-2 px-4 rounded-full focus:outline-none focus:shadow-outline"
                type="submit">
              Create account
            </button>
            <inertia-link class="inline-block align-baseline font-bold text-sm text-pink-500 hover:text-pink-700"
                          href="/login">
              Log in
            </inertia-link>
          </div>
        </form>
      </div>
    </div>


  </Layout>
</template>

<script>
import Layout from '../../Shared/Layout'

import { computed } from 'vue'
import { usePage } from '@inertiajs/inertia-vue3'

export default {
  components: {
    Layout
  },
  props: {
    register_error: String
  },
  data() {
    return {
      reg: /^(([^<>()\[\]\\.,;:\s@"]+(\.[^<>()\[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,24}))$/,
      form: {
        full_name: "",
        name: "",
        email: "",
        password: "",
        password_repeat: "",
        agreed: true,
        subscribe: true,
        marketing: true,
      },
      showAgreement: false,
    };
  },
  mounted(){
    let uri = window.location.search.substring(1);
    let params = new URLSearchParams(uri);
    this.form.email = params.get("email")
  },
  methods: {
    // this.$route.query.test
    submit() {
      if (this.form.password === this.form.password_repeat) {
        this.$inertia.post('/register', this.form)
      } else {
        this.register_error = "Passwords does not match"
      }
    },
  },
}
</script>