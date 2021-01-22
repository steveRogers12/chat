import Vue from 'vue'
import Router from 'vue-router'
import RouterConfig from './modules'
import CommonRoutes from './common'
import Utils from '../js/utils'

Vue.use(Router)

let router = new Router({
	mode: 'history',
	routes: RouterConfig.concat(CommonRoutes),
	scrollBehavior() {
		return { x: 0, y: 0 }
	}
})

router.beforeEach(async (to, from, next) => {
	if (to.meta.allowGuest) {
		return next()
	}
	if (!Utils.isLogin()) {
		return next({ path: '/login/login-in' });
	}
	return next();
})

export default router;