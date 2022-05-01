<template>
  <Layout>
    <div class="container mx-auto flex justify-center">
      <div class="w-full max-w-sm mt-4 md:mt-16">
        <div class="px-8 pt-6 pb-8 mb-4">
          <h1 class="text-lg mb-2">Epost sendt</h1>
          <p class="mb-4">Da har vi sendt en epost til deg med instrukser hvordan sette nytt passord.</p>
          <p class="mb-8">Klikk <inertia-link class="text-pink-500 underline" href="/login">her</inertia-link> for å logge inn.</p>
        </div>
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
        data() {
            return {
            };
        },
    }
</script>