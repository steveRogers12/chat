/**
 * @content 封装axios
 */

import Axios from 'axios'
import Vue from 'vue'
import User from './user'
import Utils from './utils'

let axiosInst = Axios.create({
	timeout: 20000,
	withCredentials: true
})
let err;
axiosInst.install = function (Vue) {
	console.log('process.env', process.env.BASE_URL);
	let baseUrl = process.env.BASE_URL;
	Vue.http = Vue.prototype.$http = this;
	this.interceptors.request.use(function (config) {
		if (!config.url) {
			console.log('services请求地址出错', this, config)
		}
		if (config.url.indexOf('http') !== 0) {
			config.url = baseUrl + config.url;
			if (User.id) {
				config.headers.common['token'] = User.token;
			}
		}
		return config;
	})
	this.interceptors.response.use(function (res) {
		if (res.config.useCustomHandler) { 
			return res;
		}
		const resCode = parseInt(res.data.code);
		if (resCode !== 100100) {
			if (resCode === 100600) {
				if (err) {
					clearTimeout(err);
				}
				err = setTimeout(() => {
					// Modal.info({
					// 	title: '登录失效，请重新登录',
					// 	okText: '重新登录',
					// 	onOk() {
					// 		Utils.removeStorage('currentUser');
					// 		Utils.removeStorage('currentKey');
					// 		location.reload();
					// 	},
					// })
				}, 500)
				
			} else {
				// message.error(res.data.msg)
				return Promise.reject(res.data.msg);
			}
		}
		return res;
	})
}

export default axiosInst;

