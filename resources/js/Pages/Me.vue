<template>
  <Layout class="text-center">
    <div class="container mx-auto px-6 md:px-8 lg:px-16 xl:px-32 mt-6">
      <div class="">
        <img :src="user.avatar" alt="" class="w-32 h-32 rounded-full mx-auto mb-8 shadow">
      </div>
      <h1 class="container mx-auto px-6 md:px-8 lg:px-16 xl:px-32 text-2xl mb-12">
        {{ user.name }}
      </h1>
      <div class="bg-gray-50 w-24 h-24 rounded-full mx-auto mb-8 shadow">
        <i class="las la-at text-6xl mt-5"></i>
      </div>
      <h5 class="text-sm">Email address</h5>
      <div class="mb-8">
        <h1 class="mb-2">
          <div v-if="!user.verified">
            <div class="text-lg">
              {{ account.email }}
            </div>
          </div>
        </h1>
      </div>
    </div>
  </Layout>
</template>

<script>
import Layout from '../Shared/Layout'

export default {
  name: "Me",
  components: {
    Layout
  },
  props: {
    user: Object,
    account: Object,
  },
  data() {
    return {
      editName: false,
      tempName: "",
      subscribed: false,
      showChangePassword: false,
      addressForm: false,
      address_has_error: false,
      address_error: "",
      form: {
        change_password: "",
        change_password_repeat: "",
      },
      address: {
        street: "",
        zip: "",
        city: ""
      }
    }
  },
  mounted() {
    this.tempName = this.user.name
  },
  methods: {
    storeName() {
      if (this.tempName.length > 2) {
        this.$inertia.post('/user/me/name/store', {'name': this.tempName}, {
          preserveScroll: true
        })
        this.editName = false
      }
    },
    submitAddress() {
      if (this.address.city != "" && this.address.city != "" && this.address.city != "") {
        this.address_has_error = false
        this.address_error = ""
      } else {
        this.address_has_error = true
        this.address_error = "Du må fylle inn alle adresse-feltene."
      }
      if (!this.address_has_error) {
        this.$inertia.post('/user/me/address/store', this.address, {
          preserveScroll: true
        })
        this.addressForm = false
        this.address_has_error = false
      }
    },
    ChangeSubscription(uuid) {
      this.$inertia.post('/user/me/subscribe/' + uuid, {}, {
        preserveScroll: true
      })
    },
    changePassword() {
      if (this.samePassword) {
        this.$inertia.post('/user/me/changepassword/', this.form)
        this.showChangePassword = false
      }
    },
    resendVerification() {
      this.$inertia.post('/user/me/verification', {}, {
        preserveScroll: true
      })
    },
    ChangeMarketingSubscription() {
      this.$inertia.post('/user/me/changemarketing/', {}, {
        preserveScroll: true
      })
    }
  },
  computed: {
    canSave() {
      return this.tempName.length > 2
    },
    samePassword() {
      return this.form.change_password == this.form.change_password_repeat
    }
  }
}
</script>

<style scoped>

</style>