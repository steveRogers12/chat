/**
 * @content 公共方法
 */

import Vue from 'vue'

let utils = {
	getStorage(key) {
		if (window.sessionStorage) {
			let item = sessionStorage.getItem(key);
			return JSON.parse(item);
		}
		return null;
	},
	setStorage(key, value) {
		if (window.sessionStorage) {
			sessionStorage.setItem(key, JSON.stringify(value));
		}
	},
	removeStorage(key) {
		if (window.sessionStorage) {
			sessionStorage.removeItem(key);
		}
	},
	removeAllStorage() {
		if (window.sessionStorage) {
			sessionStorage.clear();
		}
	},
	removeStorages(keys) {
		if (window.sessionStorage) {
			keys && keys.length && keys.forEach(key => sessionStorage.removeItem(key));
		}
	},
	getAllStorageKeys() {
		let storageKeys = [];
		if (window.sessionStorage) {
			for (let i = 0; i < sessionStorage.length; i++) {
				let key = sessionStorage.key(i);
				storageKeys.push(key);
			}
		}
		return storageKeys;
	},
	// 是否登录
	isLogin() {
		return this.getStorage('currentUser') !== null;
	},
	// 获取当前用户信息
	getUser() {
		return this.getStorage('currentUser')
	}
}

utils.install = function (Vue) {
	Vue.prototype.$utils = Vue.utils = utils;
}

export default utils;
