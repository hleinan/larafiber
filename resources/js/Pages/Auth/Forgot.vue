<template>
  <Layout>
    <div class="container mx-auto flex justify-center">
      <div class="w-full max-w-sm mt-4 md:mt-16">
        <form class="px-8 pt-6 pb-8 mb-4" @submit.prevent="submit" autocomplete="off">
          <h1 class="text-lg mb-2">Glemt passordet ditt?</h1>
          <p class="mb-8">Legg inn epost-adressen din, så sender vi instrukser på mail.</p>
          <div class="mb-4">
            <label class="block text-gray-700 text-sm font-bold mb-2" for="email">
              Epost-adresse
            </label>
            <input v-model="form.email" class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline" id="email" type="text" placeholder="Epost-adresse">
          </div>
          <div  class="mb-4">
            &nbsp;<span class="text-red-700 text-sm" v-show="forgot_error">{{ forgot_error }}</span>
          </div>
          <div class="flex items-center justify-between mb-4">
            <button class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline" type="submit">
              Send epost
            </button>
            <inertia-link class="inline-block align-baseline font-bold text-sm text-pink-500 hover:text-pink-800" href="/login">
              Logg inn
            </inertia-link>
          </div>
          <div class="text-center mb-2">
          </div><div class="text-center">
            <inertia-link class="inline-block align-baseline font-bold text-sm text-pink-500 hover:text-pink-800" href="/register">
              Har du ikke konto? Opprett her.
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
    forgot_error: String
  },
  data() {
    return {
      form: {
        email: "",
      }
    };
  },
  methods: {
    submit() {
      this.$inertia.post('/forgot/', this.form)
    },
  },
}
</script>