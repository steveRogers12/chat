import Vue from 'vue'
import App from './App.vue'
import router from './router'
import './css/app.css'

import Http from './js/http'
// import EventBus from './js/event-bus.js'
import User from './js/user.js'
import Utils from './js/utils.js'

Vue.config.productionTip = false
Vue.use(Http)
// Vue.use(EventBus)
Vue.use(User)
Vue.use(Utils)

new Vue({
  el: '#app',
  router,
  render: h => h(App),
}).$mount('#app')
