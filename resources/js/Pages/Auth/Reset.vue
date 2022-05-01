<template>
  <Layout>
    <div class="container mx-auto flex justify-center">
      <div class="w-full max-w-sm mt-4 md:mt-16">
        <form class="px-8 pt-6 pb-8 mb-4" @submit.prevent="submit" autocomplete="off">
          <h1 class="text-lg mb-2">Lag nytt passord</h1>
          <p class="mb-8">Lag nytt passord til kontoen din. Når passordet er opprettet, kan du logge inn på nytt.</p>
          <div class="mb-4">
            <label class="block text-gray-700 text-sm font-bold mb-2" for="password">
              Passord (minimum 8 tegn)
            </label>
            <input v-model="form.password"
                   class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 mb-3 leading-tight focus:outline-none focus:shadow-outline"
                   id="password" type="password" placeholder="************"
            >
            <!--<p class="text-red-500 text-xs italic">Du må skrive inn ett passord.</p>-->
          </div>
          <div class="mb-6">
            <label class="block text-gray-700 text-sm font-bold mb-2" for="password_repeat">
              Gjenta passord
            </label>
            <input v-model="form.password_repeat"
                   class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 mb-3 leading-tight focus:outline-none focus:shadow-outline"
                   id="password_repeat" type="password" placeholder="************"
            >
            <div class="mb-4">
              &nbsp;<span class="text-red-700 text-sm" v-show="register_error">{{ register_error }}</span>
            </div>
            <!--<p class="text-red-500 text-xs italic">Du må skrive inn ett passord.</p>-->
          </div>
          <div class="flex items-center justify-between mb-4">
            <button
                class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline"
                type="submit">
              Endre passordet
            </button>
            <inertia-link class="inline-block align-baseline font-bold text-sm text-pink-500 hover:text-pink-800"
                          href="/login">
              Logg inn
            </inertia-link>
          </div>

          <div class="text-center">
            <inertia-link class="inline-block align-baseline font-bold text-sm text-pink-500 hover:text-pink-800"
                          href="/register">
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
    register_error: String,
    uuid: String
  },
  data() {
    return {
      form: {
        password: "",
        password_repeat: "",
      }
    };
  },
  methods: {
    submit() {
      if (this.form.password === this.form.password_repeat) {
        if (this.form.password.length < 8) {
          this.register_error = "Passordet må være minimum 8 tegn"
        } else {
          this.$inertia.post('/resetpassword/' + this.uuid, this.form)
        }
      } else {
        this.register_error = "Passordene er ikke like"
      }
    },
  },
}
</script>