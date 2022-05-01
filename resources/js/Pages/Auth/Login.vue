<template>
    <Layout>
      <div class="container mx-auto flex justify-center">
        <div class="w-full max-w-sm mt-4 md:mt-16">
          <form class="px-8 pt-6 pb-8 mb-4" @submit.prevent="submit">
            <h1 class="text-lg mb-2">Log in</h1>
            <div class="mb-4">
              <label class="block text-gray-700 text-sm font-bold mb-2" for="email">
                Email address
              </label>
              <input v-model="form.email" class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline" id="email" type="text" placeholder="Email address">
            </div>
            <div class="mb-2">
              <label class="block text-gray-700 text-sm font-bold mb-2" for="password">
                Password
              </label>
              <input v-model="form.password" class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 mb-3 leading-tight focus:outline-none focus:shadow-outline" id="password" type="password" placeholder="******************">
            </div>
            <div  class="mb-4">
              &nbsp;<span class="text-red-700 text-sm" v-show="login_error">{{ login_error }}</span>
            </div>
            <div class="flex items-center justify-between mb-4">
              <button class="bg-pink-500 hover:bg-pink-600 shadow hover:shadow-md text-white font-bold py-2 px-4 rounded-full focus:outline-none focus:shadow-outline" type="submit">
                Log in
              </button>
              <inertia-link class="inline-block align-baseline font-bold text-sm text-pink-500 hover:text-pink-800" href="/forgot">
                Forgot?
              </inertia-link>
            </div>
            <div class="text-center">
              <inertia-link class="inline-block align-baseline font-bold text-sm text-pink-500 hover:bg-pink-600" href="/register">
                No account yet. Create one here.
              </inertia-link>
            </div>
          </form>
        </div>
      </div>
    </Layout>
</template>

<script>
    import Layout from '../../Shared/Layout'

    export default {
        components: {
            Layout
        },
        props: {
          login_error: String
        },
        data() {
            return {
                form: {
                    email: "",
                    password: "",
                }
            };
        },
        methods: {
            submit() {
                this.$inertia.post(window.location.href, this.form)
            },
        },
    }
</script>