/**
 * @content 用户角色
 */
import http from './http'
import eventBus from './event-bus'

let user = eventBus.user;

function setCurrentUser(newUser) {
	for (let key in newUser) {
		eventBus.$set(eventBus.user, key, newUser[key])
	}
	window.sessionStorage.setItem('currentUser', JSON.stringify(eventBus.user));
}

user.setCurrentUser = setCurrentUser;
user.install = async function (vue) {
	let strUser = window.sessionStorage.getItem('currentUser');
	if (strUser) {
		setCurrentUser(JSON.parse(strUser));
	}
	vue.prototype.$user = vue.user = eventBus.user;
}

export default eventBus.user
export {setCurrentUser}